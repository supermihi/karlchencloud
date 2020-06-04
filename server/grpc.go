package server

import (
	"context"
	grpcauth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
	"sync"
)

type dokoserver struct {
	api.UnimplementedDokoServer
	room    Room
	roomMtx sync.RWMutex
	auth    Auth
	streams clientStreams
}

func NewServer(room Room, auth Auth) *dokoserver {
	return &dokoserver{
		room: room, auth: auth, streams: newStreams(),
	}
}

// https://rogchap.com/2019/07/26/in-process-grpc-web-proxy/
func WrapServer(grpcServer *grpc.Server) *http.Server {
	grpcWebServer := grpcweb.WrapServer(grpcServer)

	httpServer := &http.Server{
		Handler: h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ProtoMajor == 2 {
				grpcWebServer.ServeHTTP(w, r)
			} else {
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
				w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-User-Agent, X-Grpc-Web")
				w.Header().Set("grpc-status", "")
				w.Header().Set("grpc-message", "")
				if grpcWebServer.IsGrpcWebRequest(r) {
					grpcWebServer.ServeHTTP(w, r)
				}
			}
		}), &http2.Server{}),
	}
	return httpServer
}
func CreateServer(users Users) *grpc.Server {
	room := NewRoom(users)
	auth := NewAuth(users)
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpcauth.UnaryServerInterceptor(auth.Authenticate)),
		grpc.StreamInterceptor(grpcauth.StreamServerInterceptor(auth.Authenticate)))
	serv := NewServer(room, auth)
	api.RegisterDokoServer(grpcServer, serv)
	return grpcServer
}

func (s *dokoserver) Register(_ context.Context, req *api.UserName) (*api.RegisterReply, error) {
	id := RandomLetters(8)
	secret := RandomSecret()
	s.room.Users.Add(id, req.GetName(), secret)
	log.Printf("Registered user %v with id %v", req.GetName(), id)
	return &api.RegisterReply{Id: id, Secret: secret}, nil
}

func (s *dokoserver) CheckLogin(ctx context.Context, _ *api.Empty) (*api.UserName, error) {
	user, ok := GetAuthenticatedUser(ctx)
	if !ok {
		log.Print("check login failed")
		return nil, nil
	}
	log.Printf("user %v ok", user)
	return &api.UserName{Name: user.Name}, nil
}

func (s *dokoserver) CreateTable(ctx context.Context, _ *api.Empty) (*api.TableData, error) {
	user, _ := GetAuthenticatedUser(ctx)
	s.roomMtx.Lock()
	defer s.roomMtx.Unlock()
	table := s.room.CreateTable(user.Id)
	log.Printf("user %v created new table %v", user, table)
	return ToTableData(table, user.Id), nil
}

func (s *dokoserver) ListTables(ctx context.Context, _ *api.Empty) (*api.TableList, error) {
	user, _ := GetAuthenticatedUser(ctx)
	s.roomMtx.RLock()
	defer s.roomMtx.RUnlock()
	tables := s.room.Tables.List()
	result := make([]*api.TableData, len(tables))
	for i, table := range tables {
		result[i] = ToTableData(table, user.Id)
	}
	return &api.TableList{Tables: result}, nil
}

func (s *dokoserver) StartTable(ctx context.Context, id *api.TableId) (*api.Empty, error) {
	user, _ := GetAuthenticatedUser(ctx)
	s.roomMtx.Lock()
	defer s.roomMtx.Unlock()
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
		state := ToMatchState(table.CurrentMatch, u)
		ev := &api.MatchEvent{Event: &api.MatchEvent_Start{Start: state}}
		s.streams.send(u, ev)
	}
	return &api.Empty{}, nil
}

func (s *dokoserver) JoinTable(ctx context.Context, req *api.JoinTableRequest) (*api.Empty, error) {
	user, _ := GetAuthenticatedUser(ctx)
	s.roomMtx.Lock()
	defer s.roomMtx.Unlock()
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
	s.streams.sendToAll(table.Users(), api.NewMemberEvent(user.Id, user.Name, api.MemberEventType_JOIN_TABLE))
	return &api.Empty{}, nil
}

func (s *dokoserver) GetTableState(ctx context.Context, tableId *api.TableId) (*api.TableState, error) {
	user, _ := GetAuthenticatedUser(ctx)
	s.roomMtx.RLock()
	defer s.roomMtx.RUnlock()
	table, err := s.getTable(tableId.Value, user.Id, true, false)
	if err != nil {
		return nil, err
	}
	return s.getTableState(table, user.Id)
}

func (s *dokoserver) SubscribeMatchEvents(tableId *api.TableId, srv api.Doko_SubscribeMatchEventsServer) error {
	user, _ := GetAuthenticatedUser(srv.Context())
	err := s.startSubscribeMatchEvents(tableId.Value, user, srv)
	if err != nil {
		return err
	}

	<-srv.Context().Done()
	log.Printf("user %s disconnected from match events", user)
	s.roomMtx.RLock()
	defer s.roomMtx.RUnlock()
	table, err := s.getTable(tableId.Value, user.Id, true, false)
	if err != nil {
		return err
	}
	s.streams.sendToAll(table.Users(), api.NewMemberEvent(user.Id, user.Name, api.MemberEventType_GO_OFFLINE))
	return srv.Context().Err()
}

func (s *dokoserver) startSubscribeMatchEvents(tableId string, user UserData, srv api.Doko_SubscribeMatchEventsServer) error {
	s.roomMtx.RLock()
	defer s.roomMtx.RUnlock()
	table, err := s.getTable(tableId, user.Id, true, false)
	if err != nil {
		return err
	}
	state, err := s.getTableState(table, user.Id)
	if err != nil {
		return err
	}
	s.streams.startNew(srv, user.Id, state)
	s.streams.sendToAll(table.Users(), api.NewMemberEvent(user.Id, user.Name, api.MemberEventType_GO_ONLINE))
	log.Printf("user %s subscribed for match events", user)
	return nil
}

func (s *dokoserver) getTableWithMatchAndActivePlayer(user string, tableId string) (table *Table, p game.Player, err error) {
	table, err = s.getTable(tableId, user, true, false)
	if err != nil {
		return
	}
	if table.CurrentMatch == nil {
		err = status.Error(codes.InvalidArgument, "no current match")
		return
	}
	players := table.CurrentMatch.Players
	p = players.PlayerFor(user)
	if p == game.NoPlayer {
		err = status.Error(codes.InvalidArgument, "you are not playing in this match")
		return
	}
	return
}
func (s *dokoserver) Play(ctx context.Context, req *api.PlayRequest) (*api.Empty, error) {
	user, _ := GetAuthenticatedUser(ctx)
	s.roomMtx.Lock()
	defer s.roomMtx.Unlock()
	table, player, err := s.getTableWithMatchAndActivePlayer(user.Id, req.Table)
	if err != nil {
		return nil, err
	}
	m := table.CurrentMatch.Match
	result := &api.MatchEvent{}
	switch action := req.Request.(type) {
	case *api.PlayRequest_Declaration:
		log.Printf("%v declares %v", user.Name, action.Declaration)
		gameType := ToGameType(action.Declaration)
		if !m.AnnounceGameType(player, gameType) {
			return nil, status.Error(codes.InvalidArgument, "cannot declare")
		}
		declaration := &api.Declaration{UserId: user.Id, Vorbehalt: !game.IsNormalspiel(gameType)}
		if m.Phase() == match.InGame {
			log.Printf("game has started on table %s", table.Id)
			declaration.DefinedGameMode = ToApiMode(m.Mode(), m.Game.WhoseTurn(), table.CurrentMatch.Players)
		}
		result = &api.MatchEvent{Event: &api.MatchEvent_Declared{Declared: declaration}}

	case *api.PlayRequest_Bid:
		if !m.PlaceBid(player, ToBid(action.Bid)) {
			return nil, status.Error(codes.InvalidArgument, "cannot place bid")
		}
		bid := &api.Bid{UserId: user.Id, Bid: action.Bid}
		result.Event = &api.MatchEvent_PlacedBid{PlacedBid: bid}

	case *api.PlayRequest_Card:
		if !m.PlayCard(player, ToCard(action.Card)) {
			return nil, status.Error(codes.InvalidArgument, "cannot play card")
		}
		log.Printf("%s plays %s", user.Name, ToCard(action.Card))
		card := &api.PlayedCard{UserId: user.Id, Card: action.Card}

		if m.Game.CurrentTrick.NumCardsPlayed() == 0 {
			card.TrickWinner = &api.PlayerValue{UserId: table.CurrentMatch.Players[int(m.Game.PreviousTrick().Winner)]}
		}
		result.Event = &api.MatchEvent_PlayedCard{PlayedCard: card}
	}
	s.streams.sendToAll(table.Users(), result)
	return &api.Empty{}, nil
}

func (s *dokoserver) getTableState(table *Table, user string) (*api.TableState, error) {
	state := &api.TableState{Members: s.createTableMembers(table)}
	if table.CurrentMatch == nil {
		state.State = &api.TableState_NoMatch{NoMatch: &api.Empty{}}
	} else {
		matchState := ToMatchState(table.CurrentMatch, user)
		state.State = &api.TableState_InMatch{InMatch: matchState}
	}
	return state, nil
}

func (s *dokoserver) getTable(id string, user string, needUserAtTable bool, needUserOwnsTable bool) (table *Table, err error) {
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

func (s *dokoserver) createTableMembers(table *Table) []*api.TableMember {
	ans := make([]*api.TableMember, len(table.Users()))
	for i, id := range table.Users() {
		ans[i] = &api.TableMember{UserId: id, Name: s.room.Users.GetName(id),
			Online: s.streams.isOnline(id)}
	}
	return ans
}
