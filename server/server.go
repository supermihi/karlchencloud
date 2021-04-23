package server

import (
	"context"
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
	config  ServerConfig
}

func CreateServer(users Users, room *Room, config ServerConfig) *grpc.Server {
	auth := NewAuth(users)
	grpcServer := CreateGrpcServerForAuth(auth)
	serv := &dokoserver{room: *room, auth: auth, streams: newStreams(), config: config}
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

	table, err := s.room.CreateTable(
		user.Id,
		&s.config.Room.ConstantTableId,
		&s.config.Room.ConstantInviteCode,
		s.config.Room.Seed(),
	)
	if err != nil {
		return nil, toGrpcError(err)
	}
	log.Printf("user %v created new table %s, code %s", user, table.Id, table.InviteCode)
	ans := ToTableData(table, user.Id, s.createTableMembers(table))
	return ans, nil
}

func (s *dokoserver) StartTable(ctx context.Context, id *api.TableId) (*api.MatchState, error) {
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
	for _, u := range getOtherPlayers(table, user.Id) {
		state := ToMatchState(matchData, u)
		ev := &api.Event{Event: &api.Event_Start{Start: state}}
		s.streams.sendSingle(u, ev)
	}
	return ToMatchState(matchData, user.Id), nil
}

func (s *dokoserver) JoinTable(ctx context.Context, req *api.JoinTableRequest) (*api.TableState, error) {
	user, _ := GetAuthenticatedUser(ctx)
	s.roomMtx.Lock()
	defer s.roomMtx.Unlock()
	table, err := s.room.JoinTable(user.Id, req.InviteCode)
	if err != nil {
		log.Printf("user %v failed joining: %v", user, err)
		return nil, toGrpcError(err)
	}
	log.Printf("user %v joined table %v", user, table.Id)
	s.streams.send(getOtherPlayers(table, user.Id), api.NewMemberEvent(user.Id, user.Name, api.MemberEventType_JOIN_TABLE))
	return s.getTableState(table, user.Id)
}

func (s *dokoserver) getUserState(user UserData) (*api.UserState, error) {
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
	s.streams.mtx.Lock()
	s.roomMtx.RLock()
	defer s.roomMtx.RUnlock()
	defer s.streams.mtx.Unlock()
	s.streams.startNew(srv, user.Id)

	userState, err := s.getUserState(user)
	if err != nil {
		return err
	}
	s.streams.sendSingle(user.Id, &api.Event{Event: &api.Event_Welcome{Welcome: userState}})
	receivers := s.room.RelatedUsers(user.Id)
	go func() {
		// requires streams.mtx, hence the goroutine
		s.streams.send(receivers, api.NewMemberEvent(user.Id, user.Name, api.MemberEventType_GO_ONLINE))
	}()
	log.Printf("user %s connected", user)
	return nil
}

func (s *dokoserver) endEventSubscription(srv api.Doko_StartSessionServer, user UserData) error {

	log.Printf("user %s disconnected", user)
	s.roomMtx.RLock()
	defer s.roomMtx.RUnlock()
	table := s.room.ActiveTableOf(user.Id)
	if table != nil {
		receivers := s.room.RelatedUsers(user.Id)
		s.streams.send(receivers, api.NewMemberEvent(user.Id, user.Name, api.MemberEventType_GO_OFFLINE))
	}
	return srv.Context().Err()
}

func getOtherPlayers(table *TableData, player string) []string {
	return stringsExcept(table.Players, player)
}

func (s *dokoserver) Declare(ctx context.Context, d *api.DeclareRequest) (*api.Declaration, error) {
	user, _ := GetAuthenticatedUser(ctx)
	s.roomMtx.Lock()
	defer s.roomMtx.Unlock()
	gameType := ToGameType(d.Declaration)
	m, err := s.room.Declare(d.Table, user.Id, gameType)
	if err != nil {
		return nil, toGrpcError(err)
	}
	log.Printf("%v declares %v", user.Name, d.Declaration)
	declaration := &api.Declaration{UserId: user.Id, Vorbehalt: !game.IsNormalspiel(gameType)}
	if m.Phase == match.InGame {
		log.Printf("game has started at table %s", d.Table)
		declaration.DefinedGameMode = ToApiMode(m.Mode, m.Turn, m.Players)
	}
	event := &api.Event{Event: &api.Event_Declared{Declared: declaration}}
	s.streams.send(s.room.RelatedUsers(user.Id), event)
	return declaration, nil
}

func (s *dokoserver) PlaceBid(ctx context.Context, req *api.PlaceBidRequest) (*api.Bid, error) {
	user, _ := GetAuthenticatedUser(ctx)
	_, err := s.room.PlaceBid(req.Table, user.Id, ToBid(req.Bid))
	if err != nil {
		return nil, toGrpcError(err)
	}
	bid := &api.Bid{UserId: user.Id, Bid: req.Bid}
	event := &api.Event{Event: &api.Event_PlacedBid{PlacedBid: bid}}
	s.streams.send(s.room.RelatedUsers(user.Id), event)
	return bid, nil
}

func (s *dokoserver) PlayCard(ctx context.Context, r *api.PlayCardRequest) (*api.PlayedCard, error) {
	user, _ := GetAuthenticatedUser(ctx)
	s.roomMtx.Lock()
	defer s.roomMtx.Unlock()

	m, err := s.room.PlayCard(r.Table, user.Id, ToCard(r.Card))
	if err != nil {
		log.Printf("%s failed to play %s: %s", user.Name, ToCard(r.Card), err)
		return nil, toGrpcError(err)
	}
	log.Printf("%s plays %s", user.Name, ToCard(r.Card))
	table, err := s.room.GetTable(r.Table)
	if err != nil {
		return nil, toGrpcError(err)
	}
	card := &api.PlayedCard{UserId: user.Id, Card: r.Card}
	if m.CurrentTrick != nil && m.CurrentTrick.NumCardsPlayed() == 0 {
		card.TrickWinner = &api.PlayerValue{UserId: table.Players[int(m.PreviousTrick.Winner)]}
	}
	if m.Phase == match.MatchFinished {
		card.Winner = &api.PartyValue{}
		switch m.Evaluation.Winner {
		case game.ReParty:
			card.Winner.Party = api.Party_RE
		case game.ContraParty:
			card.Winner.Party = api.Party_CONTRA
		}
	}
	event := &api.Event{Event: &api.Event_PlayedCard{PlayedCard: card}}
	s.streams.send(getOtherPlayers(table, user.Id), event)
	return card, nil
}

func (s *dokoserver) StartNextMatch(ctx context.Context, tableId *api.TableId) (*api.MatchState, error) {
	user, _ := GetAuthenticatedUser(ctx)
	s.roomMtx.Lock()
	defer s.roomMtx.Unlock()

	matchData, err := s.room.StartNextMatch(tableId.Value, user.Id)
	if err != nil {
		return nil, toGrpcError(err)
	}
	matchState := ToMatchState(matchData, user.Id)
	event := &api.Event{Event: &api.Event_Start{Start: matchState}}
	table, err := s.room.GetTable(tableId.Value)
	if err != nil {
		return nil, toGrpcError(err)
	}
	s.streams.send(getOtherPlayers(table, user.Id), event)
	return matchState, nil
}

func (s *dokoserver) getTableState(table *TableData, user string) (*api.TableState, error) {
	data := ToTableData(table, user, s.createTableMembers(table))
	state := &api.TableState{Data: data, Phase: table.Phase}
	if table.Phase == api.TablePhase_PLAYING {
		matchData, err := s.room.GetMatchData(table.Id)
		if err != nil {
			return nil, err
		}
		state.CurrentMatch = ToMatchState(matchData, user)
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
