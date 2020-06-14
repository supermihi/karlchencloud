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

func toGrpcError(err error) error {
	if _, ok := status.FromError(err); ok {
		return err // already a GRPC error
	}
	if cloudErr, ok := err.(CloudError); ok {
		return status.Error(codes.Internal, cloudErr.Error())
	}
	return status.Error(codes.Unknown, err.Error())
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
func CreateServer(users Users, room *Room) *grpc.Server {
	if room == nil {
		room = NewRoom(users)
	}
	auth := NewAuth(users)
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpcauth.UnaryServerInterceptor(auth.Authenticate)),
		grpc.StreamInterceptor(grpcauth.StreamServerInterceptor(auth.Authenticate)))
	serv := NewServer(*room, auth)
	api.RegisterDokoServer(grpcServer, serv)
	return grpcServer
}

func (s *dokoserver) Register(_ context.Context, req *api.UserName) (*api.RegisterReply, error) {
	id := RandomLetters(8)
	secret := RandomSecret()
	ok := s.room.AddUser(id, req.GetName(), secret)
	if !ok {
		return nil, status.Error(codes.Internal, "user ID clash")
	}
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
	table, err := s.room.CreateTable(user.Id)
	if err != nil {
		return nil, toGrpcError(err)
	}
	log.Printf("user %v created new table %v", user, table.Id)
	return s.createTableData(table, user.Id), nil
}

func (s *dokoserver) createTableData(table *TableData, user string) *api.TableData {
	ans := &api.TableData{TableId: table.Id, Owner: table.Owner, Members: s.createTableMembers(table)}
	if table.Owner == user {
		ans.InviteCode = table.InviteCode
	}
	return ans
}

func (s *dokoserver) StartTable(ctx context.Context, id *api.TableId) (*api.Empty, error) {
	user, _ := GetAuthenticatedUser(ctx)
	s.roomMtx.Lock()
	defer s.roomMtx.Unlock()
	table, err := s.getTable(id.Value, user.Id, true, true)
	if err != nil {
		return nil, err
	}
	log.Printf("starting table %s", table.Id)
	err = s.room.StartTable(table.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	matchData, err := s.room.GetMatchData(table.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	for _, u := range table.Players {
		state := ToMatchState(matchData, u)
		ev := &api.MatchEvent{Event: &api.MatchEvent_Start{Start: state}}
		s.streams.send(u, ev)
	}
	return &api.Empty{}, nil
}

func (s *dokoserver) JoinTable(ctx context.Context, req *api.JoinTableRequest) (*api.Empty, error) {
	user, _ := GetAuthenticatedUser(ctx)
	s.roomMtx.Lock()
	defer s.roomMtx.Unlock()
	table, err := s.room.JoinTable(req.TableId, user.Id, req.InviteCode)
	if err != nil {
		return nil, toGrpcError(err)
	}
	log.Printf("user %v joined table %v", user, req.TableId)
	s.streams.sendToAll(table.Players, api.NewMemberEvent(user.Id, user.Name, api.MemberEventType_JOIN_TABLE))
	return &api.Empty{}, nil
}
func (s *dokoserver) GetUserState(ctx context.Context, _ *api.Empty) (*api.UserState, error) {
	user, _ := GetAuthenticatedUser(ctx)
	s.roomMtx.Lock()
	defer s.roomMtx.Unlock()
	ans := &api.UserState{}
	activeTable := s.room.ActiveTableOf(user.Id)
	if activeTable != nil {
		tableState, err := s.getTableState(activeTable, user.Id)
		if err != nil {
			return nil, toGrpcError(err)
		}
		ans.CurrentTable = tableState
	}
	return ans, nil
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
	s.streams.sendToAll(table.Players, api.NewMemberEvent(user.Id, user.Name, api.MemberEventType_GO_OFFLINE))
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
	s.streams.sendToAll(table.Players, api.NewMemberEvent(user.Id, user.Name, api.MemberEventType_GO_ONLINE))
	log.Printf("user %s subscribed for match events", user)
	return nil
}

func (s *dokoserver) Play(ctx context.Context, req *api.PlayRequest) (*api.Empty, error) {
	user, _ := GetAuthenticatedUser(ctx)
	s.roomMtx.Lock()
	defer s.roomMtx.Unlock()
	table, err := s.room.GetTable(req.Table)
	if err != nil {
		return nil, toGrpcError(err)
	}
	result := &api.MatchEvent{}
	switch action := req.Request.(type) {
	case *api.PlayRequest_Declaration:
		gameType := ToGameType(action.Declaration)
		m, err := s.room.Declare(req.Table, user.Id, gameType)
		if err != nil {
			return nil, toGrpcError(err)
		}
		log.Printf("%v declares %v", user.Name, action.Declaration)
		declaration := &api.Declaration{UserId: user.Id, Vorbehalt: !game.IsNormalspiel(gameType)}
		if m.Phase == match.InGame {
			log.Printf("game has started on table %s", table.Id)
			declaration.DefinedGameMode = ToApiMode(m.Mode, m.Turn, m.Players)
		}
		result = &api.MatchEvent{Event: &api.MatchEvent_Declared{Declared: declaration}}

	case *api.PlayRequest_Bid:
		_, err := s.room.PlaceBid(req.Table, user.Id, ToBid(action.Bid))
		if err != nil {
			return nil, toGrpcError(err)
		}
		bid := &api.Bid{UserId: user.Id, Bid: action.Bid}
		result.Event = &api.MatchEvent_PlacedBid{PlacedBid: bid}
	case *api.PlayRequest_Card:
		m, err := s.room.PlayCard(req.Table, user.Id, ToCard(action.Card))
		if err != nil {
			return nil, toGrpcError(err)
		}
		log.Printf("%s plays %s", user.Name, ToCard(action.Card))
		card := &api.PlayedCard{UserId: user.Id, Card: action.Card}

		if m.CurrentTrick != nil && m.CurrentTrick.NumCardsPlayed() == 0 {
			card.TrickWinner = &api.PlayerValue{UserId: table.Players[int(m.PreviousTrick.Winner)]}
		}
		result.Event = &api.MatchEvent_PlayedCard{PlayedCard: card}
	}
	s.streams.sendToAll(table.Players, result)
	return &api.Empty{}, nil
}

func (s *dokoserver) getTableState(table *TableData, user string) (*api.TableState, error) {
	data := s.createTableData(table, user)
	state := &api.TableState{Data: data}
	if !table.InMatch {
		state.State = &api.TableState_NoMatch{NoMatch: &api.Empty{}}
	} else {
		matchData, err := s.room.GetMatchData(table.Id)
		if err != nil {
			return nil, err
		}
		matchState := ToMatchState(matchData, user)
		state.State = &api.TableState_InMatch{InMatch: matchState}
	}
	return state, nil
}

func (s *dokoserver) getTable(id string, user string, needUserAtTable bool, needUserOwnsTable bool) (table *TableData, err error) {
	tableData, e := s.room.GetTable(id)
	if e != nil {
		return nil, e
	}
	if needUserAtTable && !tableData.ContainsPlayer(user) {
		return nil, NewCloudError(UserNotAtTable)
	}
	if needUserOwnsTable && tableData.Owner != user {
		return nil, NewCloudError(NotOwnerOfTable)
	}
	return tableData, nil
}

func (s *dokoserver) createTableMembers(table *TableData) []*api.TableMember {
	ans := make([]*api.TableMember, len(table.Players))
	for i, id := range table.Players {
		data, err := s.room.GetUserData(id)
		if err != nil {
			panic("not existingt user at table - should not be here!")
		}
		ans[i] = &api.TableMember{UserId: id, Name: data.Name, Online: s.streams.isOnline(id)}
	}
	return ans
}
