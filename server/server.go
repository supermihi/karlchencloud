package server

import (
	"context"
	grpcauth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"sync"
)

type dokoserver struct {
	api.UnimplementedDokoServer
	room    Room
	roomMtx sync.RWMutex
	auth    Auth
	streams clientStreams
}

func newDokoserver(room Room, auth Auth) *dokoserver {
	return &dokoserver{
		room: room, auth: auth, streams: newStreams(),
	}
}

func CreateServer(users Users, room *Room) *grpc.Server {
	if room == nil {
		room = NewRoom(users)
	}
	auth := NewAuth(users)
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpcauth.UnaryServerInterceptor(auth.Authenticate)),
		grpc.StreamInterceptor(grpcauth.StreamServerInterceptor(auth.Authenticate)))
	serv := newDokoserver(*room, auth)
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
	user, _ := GetAuthenticatedUser(ctx)
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
	ans := ToTableData(table, user.Id, s.createTableMembers(table))
	return ans, nil
}

func (s *dokoserver) StartTable(ctx context.Context, id *api.TableId) (*api.Empty, error) {
	user, _ := GetAuthenticatedUser(ctx)
	s.roomMtx.Lock()
	defer s.roomMtx.Unlock()
	log.Printf("starting table %s", id.Value)
	table, err := s.room.StartTable(id.Value, user.Id)
	if err != nil {
		return nil, toGrpcError(err)
	}
	matchData, err := s.room.GetMatchData(id.Value)
	if err != nil {
		return nil, toGrpcError(err)
	}
	for _, u := range table.Players {
		state := ToMatchState(matchData, u)
		ev := &api.Event{Event: &api.Event_Start{Start: state}}
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
	ans := &api.UserState{Name: user.Name}
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

func (s *dokoserver) StartSession(_ *api.Empty, srv api.Doko_StartSessionServer) error {
	user, _ := GetAuthenticatedUser(srv.Context())
	err := s.startEventSubscription(user, srv)
	if err != nil {
		return err
	}
	<-srv.Context().Done()
	return s.endEventSubscription(srv, user)
}

func (s *dokoserver) startEventSubscription(user UserData, srv api.Doko_StartSessionServer) error {
	s.roomMtx.RLock()
	defer s.roomMtx.RUnlock()

	userState := &api.UserState{}
	table := s.room.ActiveTableOf(user.Id)
	if table != nil {
		state, err := s.getTableState(table, user.Id)
		if err != nil {
			return err
		}
		userState.CurrentTable = state
		receivers := getRelatedUsers(user.Id, s.room)
		s.streams.sendToAll(receivers, api.NewMemberEvent(user.Id, user.Name, api.MemberEventType_GO_ONLINE))
	}
	s.streams.startNew(srv, user.Id, userState)
	log.Printf("user %s connected", user)
	return nil
}

func getRelatedUsers(userId string, room Room) []string {
	table := room.ActiveTableOf(userId)
	if table == nil {
		return []string{}
	}
	ans := make([]string, len(table.Players)-1)
	i := 0
	for _, p := range table.Players {
		if p != userId {
			ans[i] = p
			i++
		}
	}
	return ans
}

func (s *dokoserver) endEventSubscription(srv api.Doko_StartSessionServer, user UserData) error {

	log.Printf("user %s disconnected", user)
	s.roomMtx.RLock()
	defer s.roomMtx.RUnlock()
	table := s.room.ActiveTableOf(user.Id)
	if table != nil {
		receivers := getRelatedUsers(user.Id, s.room)
		s.streams.sendToAll(receivers, api.NewMemberEvent(user.Id, user.Name, api.MemberEventType_GO_OFFLINE))
	}
	return srv.Context().Err()
}

func (s *dokoserver) Play(ctx context.Context, req *api.PlayRequest) (*api.Empty, error) {
	user, _ := GetAuthenticatedUser(ctx)
	s.roomMtx.Lock()
	defer s.roomMtx.Unlock()
	table, err := s.room.GetTable(req.Table)
	if err != nil {
		return nil, toGrpcError(err)
	}
	result := &api.Event{}
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
		result = &api.Event{Event: &api.Event_Declared{Declared: declaration}}

	case *api.PlayRequest_Bid:
		_, err := s.room.PlaceBid(req.Table, user.Id, ToBid(action.Bid))
		if err != nil {
			return nil, toGrpcError(err)
		}
		bid := &api.Bid{UserId: user.Id, Bid: action.Bid}
		result.Event = &api.Event_PlacedBid{PlacedBid: bid}
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
		result.Event = &api.Event_PlayedCard{PlayedCard: card}
	}
	s.streams.sendToAll(table.Players, result)
	return &api.Empty{}, nil
}

func (s *dokoserver) getTableState(table *TableData, user string) (*api.TableState, error) {
	data := ToTableData(table, user, s.createTableMembers(table))
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