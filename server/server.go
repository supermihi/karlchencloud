package server

import (
	"context"
	pb "github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
	"github.com/supermihi/karlchencloud/room"
	"github.com/supermihi/karlchencloud/server/pbconv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"sync"
)

type dokoserver struct {
	pb.UnimplementedDokoServer
	room    room.IRoom
	roomMtx sync.RWMutex
	auth    Auth
	streams clientStreams
	config  ServerConfig
}

func CreateServer(users room.Users, room *room.Room, config ServerConfig) *grpc.Server {
	auth := NewAuth(users)
	grpcServer := CreateGrpcServerForAuth(auth)
	serv := &dokoserver{room: room, auth: auth, streams: newStreams(), config: config}
	pb.RegisterDokoServer(grpcServer, serv)
	return grpcServer
}

func (s *dokoserver) Register(_ context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	user, err := s.auth.Users.Add(req.Email, req.Password, req.Name)
	if err != nil {
		log.Printf("error registering: %v", err)
		return nil, status.Error(codes.Internal, "error registering")
	}
	log.Printf("Registered user %v with id %v", user.Name, user.Id)

	return &pb.RegisterReply{UserId: user.Id.String(), Token: user.Token}, nil
}

func (s *dokoserver) Login(_ context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	user, err := s.auth.Users.Authenticate(req.Email, req.Password)
	if err != nil {
		log.Printf("error logging in: %v", err)
		return nil, err
	}
	return &pb.LoginReply{Name: user.Name, UserId: user.Id.String(), Token: user.Token}, nil
}

func (s *dokoserver) CreateTable(ctx context.Context, request *pb.CreateTableRequest) (*pb.TableData, error) {
	user, _ := GetAuthenticatedUser(ctx)
	s.roomMtx.Lock()
	defer s.roomMtx.Unlock()

	table, err := s.room.CreateTable(
		user.Id,
		request.Public,
		s.config.Room.Seed(),
	)
	if err != nil {
		return nil, toGrpcError(err)
	}
	log.Printf("user %v created new table %s, code %s", user, table.Id, table.InviteCode)
	ans := pbconv.ToPbTableData(table, user.Id, s.createPbTableMember)
	return ans, nil
}

func (s *dokoserver) StartTable(ctx context.Context, req *pb.StartTableRequest) (*pb.MatchState, error) {
	user, _ := GetAuthenticatedUser(ctx)
	s.roomMtx.Lock()
	defer s.roomMtx.Unlock()
	id, err := room.ParseTableId(req.TableId)
	if err != nil {
		log.Printf("could not parse table id %s: %v", req.TableId, err)
		return nil, toGrpcError(err)
	}
	log.Printf("starting table %s", id)
	table, err := s.room.StartTable(id, user.Id)
	if err != nil {
		return nil, toGrpcError(err)
	}
	matchData, err := s.room.GetMatchData(id)
	if err != nil {
		return nil, toGrpcError(err)
	}
	for _, u := range getOtherPlayers(table, user.Id) {
		state := pbconv.ToPbMatchState(matchData, u)
		ev := &pb.Event{Event: &pb.Event_Start{Start: state}}
		s.streams.sendSingle(u, ev)
	}
	return pbconv.ToPbMatchState(matchData, user.Id), nil
}

func (s *dokoserver) JoinTable(ctx context.Context, req *pb.JoinTableRequest) (*pb.TableState, error) {
	user, _ := GetAuthenticatedUser(ctx)
	s.roomMtx.Lock()
	defer s.roomMtx.Unlock()
	var table *room.TableData
	var err error
	switch td := req.TableDescription.(type) {
	case *pb.JoinTableRequest_InviteCode:
		table, err = s.room.JoinTableByInviteCode(user.Id, td.InviteCode)
	case *pb.JoinTableRequest_TableId:
		var tableId room.TableId
		tableId, err = room.ParseTableId(td.TableId)
		if err != nil {
			return nil, err
		}
		table, err = s.room.JoinTableByTableId(user.Id, tableId)
	}

	if err != nil {
		log.Printf("user %v failed joining: %v", user, err)
		return nil, toGrpcError(err)
	}
	log.Printf("user %v joined table %v", user, table.Id)
	s.streams.send(getOtherPlayers(table, user.Id), pb.NewMemberEvent(user.Id.String(), user.Name, pb.MemberEventType_JOIN_TABLE))
	return s.getTableState(table, user.Id)
}

func (s *dokoserver) getUserState(user room.UserData) (*pb.UserState, error) {
	ans := &pb.UserState{Name: user.Name}
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

func (s *dokoserver) StartSession(_ *pb.Empty, srv pb.Doko_StartSessionServer) error {
	user, _ := GetAuthenticatedUser(srv.Context())
	err := s.startEventSubscription(user, srv)
	if err != nil {
		return err
	}
	<-srv.Context().Done()
	return s.endEventSubscription(srv, user)
}

func (s *dokoserver) startEventSubscription(user room.UserData, srv pb.Doko_StartSessionServer) error {
	s.streams.mtx.Lock()
	s.roomMtx.RLock()
	defer s.roomMtx.RUnlock()
	defer s.streams.mtx.Unlock()
	s.streams.startNew(srv, user.Id)

	userState, err := s.getUserState(user)
	if err != nil {
		return err
	}
	s.streams.sendSingle(user.Id, &pb.Event{Event: &pb.Event_Welcome{Welcome: userState}})
	receivers := s.room.RelatedUsers(user.Id)
	go func() {
		// requires streams.mtx, hence the goroutine
		s.streams.send(receivers, pb.NewMemberEvent(user.Id.String(), user.Name, pb.MemberEventType_GO_ONLINE))
	}()
	log.Printf("user %s connected", user)
	return nil
}

func (s *dokoserver) endEventSubscription(srv pb.Doko_StartSessionServer, user room.UserData) error {
	log.Printf("user %s disconnected", user)
	s.roomMtx.RLock()
	defer s.roomMtx.RUnlock()
	table := s.room.ActiveTableOf(user.Id)
	if table != nil {
		receivers := s.room.RelatedUsers(user.Id)
		s.streams.send(receivers, pb.NewMemberEvent(user.Id.String(), user.Name, pb.MemberEventType_GO_OFFLINE))
	}
	return srv.Context().Err()
}

func getOtherPlayers(table *room.TableData, player room.UserId) []room.UserId {
	return room.UsersExcept(table.Players, player)
}

func (s *dokoserver) Declare(ctx context.Context, d *pb.DeclareRequest) (*pb.Declaration, error) {
	user, _ := GetAuthenticatedUser(ctx)
	tableId, err := room.ParseTableId(d.Table)
	if err != nil {
		log.Printf("declare: could not parse table id %s: %v\n", d.Table, err)
		return nil, toGrpcError(err)
	}
	s.roomMtx.Lock()
	defer s.roomMtx.Unlock()
	gameType := pbconv.ToGameType(d.Declaration)
	m, err := s.room.Declare(tableId, user.Id, gameType)
	if err != nil {
		return nil, toGrpcError(err)
	}
	log.Printf("%v declares %v", user.Name, d.Declaration)
	declaration := &pb.Declaration{UserId: user.Id.String(), Vorbehalt: !game.IsNormalspiel(gameType)}
	if m.Phase == match.InGame {
		log.Printf("game has started at table %s", d.Table)
		declaration.DefinedGameMode = pbconv.ToPbMode(m.Mode, m.Turn, m.Players)
	}
	event := &pb.Event{Event: &pb.Event_Declared{Declared: declaration}}
	s.streams.send(s.room.RelatedUsers(user.Id), event)
	return declaration, nil
}

func (s *dokoserver) PlaceBid(ctx context.Context, req *pb.PlaceBidRequest) (*pb.Bid, error) {
	user, _ := GetAuthenticatedUser(ctx)
	tableId, err := room.ParseTableId(req.Table)
	if err != nil {
		log.Printf("placeBid: could not parse table id %s: %v\n", req.Table, err)
		return nil, toGrpcError(err)
	}
	_, err = s.room.PlaceBid(tableId, user.Id, pbconv.ToBid(req.Bid))
	if err != nil {
		return nil, toGrpcError(err)
	}
	bid := &pb.Bid{UserId: user.Id.String(), Bid: req.Bid}
	event := &pb.Event{Event: &pb.Event_PlacedBid{PlacedBid: bid}}
	s.streams.send(s.room.RelatedUsers(user.Id), event)
	return bid, nil
}

func (s *dokoserver) PlayCard(ctx context.Context, req *pb.PlayCardRequest) (*pb.PlayedCard, error) {
	user, _ := GetAuthenticatedUser(ctx)
	tableId, err := room.ParseTableId(req.Table)
	if err != nil {
		log.Printf("playCard: could not parse table id %s: %v\n", req.Table, err)
		return nil, toGrpcError(err)
	}
	s.roomMtx.Lock()
	defer s.roomMtx.Unlock()
	m, err := s.room.PlayCard(tableId, user.Id, pbconv.ToCard(req.Card))
	if err != nil {
		log.Printf("%s failed to play %s: %s", user.Name, pbconv.ToCard(req.Card), err)
		return nil, toGrpcError(err)
	}
	log.Printf("%s plays %s", user.Name, pbconv.ToCard(req.Card))
	table, err := s.room.GetTable(tableId)
	if err != nil {
		return nil, toGrpcError(err)
	}
	card := &pb.PlayedCard{UserId: user.Id.String(), Card: req.Card}
	if m.CurrentTrick != nil && m.CurrentTrick.NumCardsPlayed() == 0 {
		card.TrickWinner = &pb.PlayerValue{UserId: table.Players[m.PreviousTrick.Winner].String()}
	}
	if m.Phase == match.MatchFinished {
		card.Winner = &pb.PartyValue{}
		switch m.Evaluation.Winner {
		case game.ReParty:
			card.Winner.Party = pb.Party_RE
		case game.ContraParty:
			card.Winner.Party = pb.Party_CONTRA
		}
	}
	event := &pb.Event{Event: &pb.Event_PlayedCard{PlayedCard: card}}
	s.streams.send(getOtherPlayers(table, user.Id), event)
	return card, nil
}

func (s *dokoserver) StartNextMatch(ctx context.Context, req *pb.StartNextMatchRequest) (*pb.MatchState, error) {
	user, _ := GetAuthenticatedUser(ctx)
	s.roomMtx.Lock()
	defer s.roomMtx.Unlock()
	tableId, err := room.ParseTableId(req.TableId)
	if err != nil {
		log.Printf("could not parse table id %s: %v\n", req.TableId, err)
		return nil, toGrpcError(err)
	}
	matchData, err := s.room.StartNextMatch(tableId, user.Id)
	if err != nil {
		return nil, toGrpcError(err)
	}
	table, err := s.room.GetTable(tableId)
	if err != nil {
		return nil, toGrpcError(err)
	}
	for _, u := range getOtherPlayers(table, user.Id) {
		state := pbconv.ToPbMatchState(matchData, u)
		ev := &pb.Event{Event: &pb.Event_Start{Start: state}}
		s.streams.sendSingle(u, ev)
	}
	return pbconv.ToPbMatchState(matchData, user.Id), nil
}

func (s *dokoserver) ListTables(ctx context.Context, req *pb.ListTablesRequest) (*pb.ListTablesResult, error) {
	user, _ := GetAuthenticatedUser(ctx)
	s.roomMtx.Lock()
	defer s.roomMtx.Unlock()
	tables := s.room.GetOpenTables(user.Id)
	pbResult := pbconv.ToPbTables(tables, user.Id, s.createPbTableMember)
	return &pb.ListTablesResult{Tables: pbResult}, nil
}

func (s *dokoserver) getTableState(table *room.TableData, user room.UserId) (*pb.TableState, error) {
	data := pbconv.ToPbTableData(table, user, s.createPbTableMember)
	state := &pb.TableState{Data: data, Phase: table.Phase}
	if table.Phase == pb.TablePhase_PLAYING {
		matchData, err := s.room.GetMatchData(table.Id)
		if err != nil {
			return nil, err
		}
		state.CurrentMatch = pbconv.ToPbMatchState(matchData, user)
	}
	return state, nil
}

func (s *dokoserver) createPbTableMember(id room.UserId) *pb.TableMember {
	data, err := s.auth.Users.GetData(id)
	if err != nil {
		panic("not existingt user at table - should not be here!")
	}
	return &pb.TableMember{UserId: id.String(), Name: data.Name, Online: s.streams.isOnline(id)}
}
