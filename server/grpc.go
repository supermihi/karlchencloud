package server

import (
	"context"
	grpcauth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/cloud"
	"github.com/supermihi/karlchencloud/common"
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"sync"
)

type grpcserver struct {
	api.UnimplementedKarlchencloudServer
	room                         cloud.Room
	auth                         Auth
	ClientTableStreams           map[string]chan *api.MatchEvent
	roomMutex, clientStreamMutex sync.RWMutex
}

func NewServer(room cloud.Room, auth Auth) *grpcserver {
	return &grpcserver{
		room: room, auth: auth,
		ClientTableStreams: make(map[string]chan *api.MatchEvent, 1000),
	}
}
func StartServer(users cloud.Users, port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	room := cloud.NewRoom(users)
	auth := NewAuth(users)
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpcauth.UnaryServerInterceptor(auth.Authenticate)),
		grpc.StreamInterceptor(grpcauth.StreamServerInterceptor(auth.Authenticate)))
	serv := NewServer(room, auth)
	api.RegisterKarlchencloudServer(grpcServer, serv)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *grpcserver) Register(_ context.Context, req *api.UserName) (*api.RegisterReply, error) {
	id := cloud.RandomLetters(8)
	secret := cloud.RandomSecret()
	s.room.Users.Add(id, req.GetName(), secret)
	log.Printf("Registered user %v with id %v", req.GetName(), id)
	return &api.RegisterReply{Id: id, Secret: secret}, nil
}

func (s *grpcserver) CheckLogin(ctx context.Context, _ *api.Empty) (*api.UserName, error) {
	user, ok := GetAuthenticatedUser(ctx)
	if !ok {
		log.Print("check login failed")
		return nil, nil
	}
	log.Printf("user %v ok", user)
	return &api.UserName{Name: user.Name}, nil
}

func (s *grpcserver) CreateTable(ctx context.Context, _ *api.Empty) (*api.TableData, error) {
	user, _ := GetAuthenticatedUser(ctx)
	s.roomMutex.Lock()
	table := s.room.CreateTable(user.Id)
	s.roomMutex.Unlock()
	log.Printf("user %v created new table %v", user, table)
	return common.ToTableData(table, user.Id), nil
}

func (s *grpcserver) ListTables(ctx context.Context, _ *api.Empty) (*api.TableList, error) {
	user, _ := GetAuthenticatedUser(ctx)
	s.roomMutex.RLock()
	tables := s.room.Tables.List()
	result := make([]*api.TableData, len(tables))
	for i, table := range tables {
		result[i] = common.ToTableData(table, user.Id)
	}
	s.roomMutex.RUnlock()
	return &api.TableList{Tables: result}, nil
}

func (s *grpcserver) StartTable(ctx context.Context, id *api.TableId) (*api.Empty, error) {
	user, _ := GetAuthenticatedUser(ctx)
	s.roomMutex.Lock()
	defer s.roomMutex.Unlock()
	table, err := s.getTable(id.Value, user.Id, true, true)
	if err != nil {
		return nil, err
	}
	log.Printf("starting table %s", table)
	err = table.Start()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	for _, u := range table.Users() {
		state, err := s.getMatchState(table.Id, u)
		if err != nil {
			return nil, err
		}
		ev := &api.MatchEvent{Event: &api.MatchEvent_Start{Start: state}}
		s.sendEventIfOnline(u, ev)
	}
	return &api.Empty{}, nil
}

func (s *grpcserver) getName(userId string) string {
	return s.auth.Users.GetName(userId)
}

func (s *grpcserver) JoinTable(ctx context.Context, req *api.JoinTableRequest) (*api.Empty, error) {
	user, _ := GetAuthenticatedUser(ctx)
	s.roomMutex.Lock()
	defer s.roomMutex.Unlock()
	table, err := s.getTable(req.TableId, user.Id, false, false)
	if err != nil {
		return nil, err
	}
	if table.InviteCode != req.InviteCode {
		return nil, status.Error(codes.PermissionDenied, "invalid invite code")
	}
	err = table.Join(user.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	log.Printf("user %v joined table %v", user, table.Id)
	s.sendEventToAll(table.Users(), api.NewMemberEvent(user.Id, user.Name, api.MemberEventType_JOIN_TABLE))
	return &api.Empty{}, nil
}

func (s *grpcserver) GetTableState(ctx context.Context, tableId *api.TableId) (*api.TableState, error) {
	user, _ := GetAuthenticatedUser(ctx)
	s.roomMutex.Lock()
	defer s.roomMutex.Unlock()
	table, err := s.getTable(tableId.Value, user.Id, true, false)
	if err != nil {
		return nil, err
	}
	return s.getTableState(table, user.Id)
}

func (s *grpcserver) SubscribeMatchEvents(tableId *api.TableId, srv api.Karlchencloud_SubscribeMatchEventsServer) error {
	user, _ := GetAuthenticatedUser(srv.Context())
	err := s.startSubscribeMatchEvents(tableId.Value, user, srv)
	if err != nil {
		return err
	}

	<-srv.Context().Done()
	log.Printf("user %s disconnected from match events", user)
	s.roomMutex.RLock()
	defer s.roomMutex.RUnlock()
	table, err := s.getTable(tableId.Value, user.Id, true, false)
	if err != nil {
		return err
	}
	s.sendEventToAll(table.Users(), api.NewMemberEvent(user.Id, user.Name, api.MemberEventType_GO_OFFLINE))
	return srv.Context().Err()
}

func (s *grpcserver) startSubscribeMatchEvents(tableId string, user UserData, srv api.Karlchencloud_SubscribeMatchEventsServer) error {
	s.roomMutex.RLock()
	defer s.roomMutex.RUnlock()
	table, err := s.getTable(tableId, user.Id, true, false)
	if err != nil {
		return err
	}
	state, err := s.getTableState(table, user.Id)
	if err != nil {
		return err
	}
	s.sendTableBroadcasts(srv, user.Id, state)
	s.sendEventToAll(table.Users(), api.NewMemberEvent(user.Id, user.Name, api.MemberEventType_GO_ONLINE))
	log.Printf("user %s subscribed for match events", user)
	return nil
}

func (s *grpcserver) Play(ctx context.Context, req *api.PlayRequest) (*api.Empty, error) {
	user, _ := GetAuthenticatedUser(ctx)
	s.roomMutex.Lock()
	defer s.roomMutex.Unlock()
	table, err := s.getTable(req.Table, user.Id, true, false)
	if err != nil {
		return nil, err
	}
	players := table.CurrentMatch.Players
	if table.CurrentMatch == nil {
		return nil, status.Error(codes.InvalidArgument, "no current match")
	}
	player := players.PlayerFor(user.Id)
	if player == game.NoPlayer {
		return nil, status.Error(codes.InvalidArgument, "you are not playing in this match")
	}
	playerId := common.ToApiUserId(player, players)
	m := table.CurrentMatch.Match
	result := &api.MatchEvent{}
	switch action := req.Request.(type) {
	case *api.PlayRequest_Declaration:
		log.Printf("%v declares %v", user.Name, action.Declaration)
		gameType := common.ToGameType(action.Declaration)
		if !m.AnnounceGameType(player, gameType) {
			return nil, status.Error(codes.InvalidArgument, "cannot declare")
		}
		declaration := &api.Declaration{UserId: playerId, Vorbehalt: !game.IsNormalspiel(gameType)}
		if m.Phase() == match.InGame {
			log.Printf("game has started on table %s", table)
			declaration.DefinedGameMode = common.ToApiMode(m.Mode(), m.Game.WhoseTurn(), players)
		}
		result.Event = &api.MatchEvent_Declared{Declared: declaration}

	case *api.PlayRequest_Bid:
		if !m.PlaceBid(player, common.ToBid(action.Bid)) {
			return nil, status.Error(codes.InvalidArgument, "cannot place bid")
		}
		bid := &api.Bid{UserId: playerId, Bid: action.Bid}
		result.Event = &api.MatchEvent_PlacedBid{PlacedBid: bid}

	case *api.PlayRequest_Card:
		if !m.PlayCard(player, common.ToCard(action.Card)) {
			return nil, status.Error(codes.InvalidArgument, "cannot play card")
		}
		log.Printf("%s plays %s", user.Name, common.ToCard(action.Card))
		card := &api.PlayedCard{UserId: playerId, Card: action.Card}

		if m.Game.CurrentTrick.NumCardsPlayed() == 0 {
			card.TrickWinner = &api.PlayerValue{UserId: table.CurrentMatch.Players[int(m.Game.PreviousTrick().Winner)]}
		}
		result.Event = &api.MatchEvent_PlayedCard{PlayedCard: card}
	}
	s.sendEventToAll(table.Users(), result)
	return &api.Empty{}, nil
}

func (s *grpcserver) getTableState(table *cloud.Table, user string) (*api.TableState, error) {
	state := &api.TableState{Members: s.createTableMembers(table)}
	if table.CurrentMatch == nil {
		state.State = &api.TableState_NoMatch{NoMatch: &api.Empty{}}
	} else {
		matchState, err := s.getMatchState(table.Id, user)
		if err != nil {
			return nil, err
		}
		state.State = &api.TableState_InMatch{InMatch: matchState}
	}
	return state, nil
}

func (s *grpcserver) getMatchState(tableId string, user string) (*api.MatchState, error) {
	table, err := s.getTable(tableId, user, true, false)
	if err != nil {
		return nil, err
	}
	m := table.CurrentMatch
	if m == nil {
		return nil, status.Error(codes.Internal, "no active match at table")
	}
	return common.ToMatchState(m, user), nil
}

func (s *grpcserver) getTable(id string, user string, needUserAtTable bool, needUserOwnsTable bool) (table *cloud.Table, err error) {
	t, ok := s.room.Tables.ById[id]
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "table does not exist")
	}
	if needUserAtTable && !t.ContainsPlayer(user) {
		return nil, status.Error(codes.PermissionDenied, "user not at table")
	}
	if needUserOwnsTable && t.Owner() != user {
		return nil, status.Error(codes.PermissionDenied, "not your table")
	}
	return t, nil
}

func (s *grpcserver) sendEventIfOnline(user string, event *api.MatchEvent) {
	stream, ok := s.ClientTableStreams[user]
	if ok {
		stream <- event
	}
}

func (s *grpcserver) sendEventToAll(users []string, event *api.MatchEvent) {
	s.clientStreamMutex.RLock()
	for _, user := range users {
		s.sendEventIfOnline(user, event)
	}
	defer s.clientStreamMutex.RUnlock()
}

func (s *grpcserver) sendTableBroadcasts(srv api.Karlchencloud_SubscribeMatchEventsServer, user string,
	state *api.TableState) {
	stream := s.openTableStream(user)
	stream <- &api.MatchEvent{Event: &api.MatchEvent_Welcome{Welcome: state}}
	go func() {
		defer s.closeTableStream(user)
		for {
			select {
			case <-srv.Context().Done():
				log.Printf("no longer waiting for messages to %s", user)
				return
			case event := <-stream:
				if s, ok := status.FromError(srv.Send(event)); ok {
					switch s.Code() {
					case codes.OK:
						// pass
					case codes.Unavailable, codes.Canceled, codes.DeadlineExceeded:
						log.Printf("client %s terminated connection", user)
						return

					default:
						log.Printf("failed to send to client %s: %v", user, s.Err())
					}
				}
			}
		}
	}()
}
func (s *grpcserver) openTableStream(user string) (stream chan *api.MatchEvent) {
	stream = make(chan *api.MatchEvent, 10)
	s.clientStreamMutex.Lock()
	defer s.clientStreamMutex.Unlock()
	s.ClientTableStreams[user] = stream
	return
}

func (s *grpcserver) closeTableStream(user string) {
	s.clientStreamMutex.Lock()

	if stream, ok := s.ClientTableStreams[user]; ok {
		delete(s.ClientTableStreams, user)
		close(stream)
	}
	log.Printf("closed table stream for %s", user)
	s.clientStreamMutex.Unlock()
}

func (s *grpcserver) isOnline(user string) bool {
	s.clientStreamMutex.RLock()
	_, ok := s.ClientTableStreams[user]
	s.clientStreamMutex.RUnlock()
	return ok
}

func (s *grpcserver) createTableMembers(table *cloud.Table) []*api.TableMember {
	ans := make([]*api.TableMember, len(table.Users()))
	for i, id := range table.Users() {
		ans[i] = &api.TableMember{UserId: string(id), Name: s.room.Users.GetName(id),
			Online: s.isOnline(id)}
	}
	return ans
}
