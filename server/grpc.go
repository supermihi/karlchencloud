package server

import (
	"context"
	grpcauth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/cloud"
	"github.com/supermihi/karlchencloud/common"
	"github.com/supermihi/karlchencloud/doko/game"
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
	ClientTableStreams           map[cloud.UserId]chan *api.MatchEventStream
	roomMutex, clientStreamMutex sync.RWMutex
}

func NewServer(room cloud.Room, auth Auth) *grpcserver {
	return &grpcserver{
		room: room, auth: auth,
		ClientTableStreams: make(map[cloud.UserId]chan *api.MatchEventStream, 1000),
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

func (s *grpcserver) Register(_ context.Context, req *api.RegisterRequest) (*api.RegisterReply, error) {
	id := cloud.RandomLetters(8)
	secret := cloud.RandomSecret()
	s.room.Users.Add(cloud.UserId(id), req.GetName(), secret)
	log.Printf("Registered user %v with id %v", req.GetName(), id)
	return &api.RegisterReply{Id: id, Secret: secret}, nil
}

func (s *grpcserver) CheckLogin(ctx context.Context, _ *api.Empty) (*api.OkOrNot, error) {
	user, ok := GetAuthenticatedUser(ctx)
	if !ok {
		log.Print("check login failed")
		return &api.OkOrNot{Value: false}, nil
	}
	log.Printf("user %v ok", user)
	return &api.OkOrNot{Value: true}, nil
}

func (s *grpcserver) CreateTable(ctx context.Context, _ *api.Empty) (*api.TableData, error) {
	user, _ := GetAuthenticatedUser(ctx)
	s.roomMutex.Lock()
	table := s.room.CreateTable(user)
	s.roomMutex.Unlock()
	log.Printf("user %v created new table %v", s.room.Users.GetName(user), table)
	return common.ToTableData(table, user), nil
}

func (s *grpcserver) ListTables(ctx context.Context, _ *api.Empty) (*api.TableList, error) {
	user, _ := GetAuthenticatedUser(ctx)
	s.roomMutex.RLock()
	tables := s.room.Tables.List()
	result := make([]*api.TableData, len(tables))
	for i, table := range tables {
		result[i] = common.ToTableData(table, user)
	}
	s.roomMutex.RUnlock()
	return &api.TableList{Tables: result}, nil
}

func (s *grpcserver) StartTable(ctx context.Context, id *api.TableId) (*api.Empty, error) {
	user, _ := GetAuthenticatedUser(ctx)
	s.roomMutex.Lock()
	defer s.roomMutex.Unlock()
	table, err := s.getTable(id.Value, user, true, true)
	if err != nil {
		return nil, err
	}
	err = table.Start()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	for _, u := range table.Users() {
		state, err := s.getMatchState(table.Id, u)
		if err != nil {
			return nil, err
		}
		ev := &api.MatchEventStream{Event: &api.MatchEventStream_Start{Start: state}}
		s.sendEventIfOnline(u, ev)
	}
	return &api.Empty{}, nil
}

func (s *grpcserver) JoinTable(ctx context.Context, req *api.JoinTableRequest) (*api.TableState, error) {
	user, _ := GetAuthenticatedUser(ctx)
	s.roomMutex.Lock()
	defer s.roomMutex.Unlock()
	table, err := s.getTable(req.TableId, user, false, false)
	if err != nil {
		return nil, err
	}
	if table.InviteCode != req.InviteCode {
		return nil, status.Error(codes.PermissionDenied, "invalid invite code")
	}
	err = table.Join(user)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	log.Printf("user %v joined table %v", user, table.Id)
	s.sendEventToAll(table.Users(), api.NewMemberEvent(user, api.MemberEventType_JOIN_TABLE))
	return s.getTableState(table, user)
}

func (s *grpcserver) SubscribeMatchEvents(tableId *api.TableId, srv api.Karlchencloud_SubscribeMatchEventsServer) error {
	user, _ := GetAuthenticatedUser(srv.Context())
	s.roomMutex.RLock()
	table, err := s.getTable(tableId.Value, user, true, false)
	if err != nil {
		s.roomMutex.RUnlock()
		return err
	}
	go s.sendTableBroadcasts(srv, user)
	s.sendEventToAll(table.Users(), api.NewMemberEvent(user, api.MemberEventType_GO_ONLINE))
	s.roomMutex.RUnlock()
	log.Printf("user %s subscribed for match events", user)
	<-srv.Context().Done()
	log.Printf("user %s disconnected from match events", user)
	s.roomMutex.RLock()
	table, err = s.getTable(tableId.Value, user, true, false)
	if err != nil {
		s.roomMutex.RUnlock()
		return err
	}
	s.sendEventToAll(table.Users(), api.NewMemberEvent(user, api.MemberEventType_GO_OFFLINE))
	s.roomMutex.RUnlock()

	return srv.Context().Err()
}

func (s *grpcserver) Play(ctx context.Context, req *api.PlayRequest) (*api.Empty, error) {
	user, _ := GetAuthenticatedUser(ctx)
	s.roomMutex.Lock()
	table, err := s.getTable(req.Table, user, true, false)
	log.Printf("incoming play request from %s", user)
	if err != nil {
		s.roomMutex.Unlock()
		return nil, err
	}
	if table.CurrentMatch == nil {
		s.roomMutex.Unlock()
		return nil, status.Error(codes.InvalidArgument, "no current match")
	}
	player := table.CurrentMatch.PlayerFor(user)
	if player == game.NoPlayer {
		s.roomMutex.Unlock()
		return nil, status.Error(codes.InvalidArgument, "you are not playing in this match")
	}
	p := common.ToApiPlayer(player, false)
	m := table.CurrentMatch.Match
	result := &api.MatchEventStream{}
	switch action := req.Request.(type) {
	case *api.PlayRequest_Declaration:
		gameType := common.ToGameType(action.Declaration)
		if !m.AnnounceGameType(player, gameType) {
			return nil, status.Error(codes.InvalidArgument, "cannot declare")
		}
		declaration := &api.Declaration{Player: p, Vorbehalt: !game.IsNormalspiel(gameType)}
		result.Event = &api.MatchEventStream_Declared{Declared: declaration}

	case *api.PlayRequest_Bid:
		if !m.PlaceBid(player, common.ToBid(action.Bid)) {
			return nil, status.Error(codes.InvalidArgument, "cannot place bid")
		}
		bid := &api.Bid{Player: p, Bid: action.Bid}
		result.Event = &api.MatchEventStream_PlacedBid{PlacedBid: bid}

	case *api.PlayRequest_Card:
		if !m.PlayCard(player, common.ToCard(action.Card)) {
			return nil, status.Error(codes.InvalidArgument, "cannot play card")
		}
		card := &api.PlayedCard{Player: p, Card: action.Card}
		result.Event = &api.MatchEventStream_PlayedCard{PlayedCard: card}
	}
	s.sendEventToAll(table.Users(), result)
	s.roomMutex.Unlock()
	return &api.Empty{}, nil
}

func (s *grpcserver) getTableState(table *cloud.Table, user cloud.UserId) (*api.TableState, error) {
	state := &api.TableState{Users: s.createTableMembers(table)}
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

func (s *grpcserver) getMatchState(tableId string, user cloud.UserId) (*api.MatchState, error) {
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

func (s *grpcserver) getTable(id string, user cloud.UserId, needUserAtTable bool, needUserOwnsTable bool) (table *cloud.Table, err error) {
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

func (s *grpcserver) sendEventIfOnline(user cloud.UserId, event *api.MatchEventStream) {
	s.clientStreamMutex.RLock()
	defer s.clientStreamMutex.RUnlock()
	stream, ok := s.ClientTableStreams[user]
	if ok {
		stream <- event
	}
}

func (s *grpcserver) sendEventToAll(users []cloud.UserId, event *api.MatchEventStream) {
	for _, user := range users {
		s.sendEventIfOnline(user, event)
	}
}

func (s *grpcserver) sendTableBroadcasts(srv api.Karlchencloud_SubscribeMatchEventsServer, user cloud.UserId) {
	stream := s.openTableStream(user)
	defer s.closeTableStream(user)
	log.Printf("waiting for messages to %s", user)
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
}
func (s *grpcserver) openTableStream(user cloud.UserId) (stream chan *api.MatchEventStream) {
	stream = make(chan *api.MatchEventStream, 10)
	s.clientStreamMutex.Lock()
	s.ClientTableStreams[user] = stream
	s.clientStreamMutex.Unlock()

	log.Printf("created table stream for %s", user)

	return
}

func (s *grpcserver) closeTableStream(user cloud.UserId) {
	s.clientStreamMutex.Lock()

	if stream, ok := s.ClientTableStreams[user]; ok {
		delete(s.ClientTableStreams, user)
		close(stream)
	}
	log.Printf("closed table stream for %s", user)
	s.clientStreamMutex.Unlock()
}

func (s *grpcserver) isOnline(user cloud.UserId) bool {
	s.clientStreamMutex.RLock()
	_, ok := s.ClientTableStreams[user]
	s.clientStreamMutex.RUnlock()
	return ok
}

func (s *grpcserver) createTableMembers(table *cloud.Table) []*api.TableMember {
	ans := make([]*api.TableMember, len(table.Users()))
	for i, id := range table.Users() {
		ans[i] = &api.TableMember{Id: string(id), Name: s.room.Users.GetName(id),
			Online: s.isOnline(id)}
	}
	return ans
}
