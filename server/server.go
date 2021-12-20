package server

import (
	"context"
	pb "github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
	"github.com/supermihi/karlchencloud/server/errors"
	"github.com/supermihi/karlchencloud/server/pbconv"
	t "github.com/supermihi/karlchencloud/server/tables"
	u "github.com/supermihi/karlchencloud/server/users"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type dokoserver struct {
	pb.UnimplementedDokoServer
	tables  *t.Tables
	users   u.Users
	streams ClientStreams
	config  ServerConfig
}

func CreateServer(users u.Users, tables *t.Tables, config ServerConfig) *grpc.Server {
	grpcServer := CreateAuthenticatingGrpcServer(users)
	serv := &dokoserver{tables: tables, users: users, streams: NewClientStreams(), config: config}
	pb.RegisterDokoServer(grpcServer, serv)
	return grpcServer
}

func (srv *dokoserver) Register(_ context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	user, err := srv.users.Add(req.Email, req.Password, req.Name)
	if err != nil {
		log.Printf("error registering: %v", err)
		return nil, status.Error(codes.Internal, "error registering")
	}
	log.Printf("Registered user %v with id %v", user.Name, user.Id)

	return &pb.RegisterReply{UserId: user.Id.String(), Token: user.Token}, nil
}

func (srv *dokoserver) Login(_ context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	user, err := srv.users.Authenticate(req.Email, req.Password)
	if err != nil {
		log.Printf("error logging in: %v", err)
		return nil, toGrpcError(err)
	}
	log.Printf("%s logged in", user.Name)
	return &pb.LoginReply{Name: user.Name, UserId: user.Id.String(), Token: user.Token}, nil
}

func (srv *dokoserver) StartSession(_ *pb.Empty, pbSrv pb.Doko_StartSessionServer) error {
  user, _ := GetAuthenticatedUser(pbSrv.Context())
  kicked, err := srv.startEventSubscription(user, pbSrv)
  if err != nil {
    return toGrpcError(err)
  }
  select {
  case <-pbSrv.Context().Done():
    log.Printf("%s: session ended by client", user.Name)
  case <-kicked:
    log.Printf("%s started a new session", user.Name)
    return errors.NewCloudError(errors.StartedAnotherSession)
  }
  return srv.endEventSubscription(pbSrv, user)
}

func (srv *dokoserver) CreateTable(ctx context.Context, request *pb.CreateTableRequest) (*pb.TableData, error) {
	user, _ := GetAuthenticatedUser(ctx)
	table, err := srv.tables.CreateTable(
		user.Id,
		request.Public,
		srv.config.Tables.Seed(),
	)
	if err != nil {
		return nil, toGrpcError(err)
	}
	log.Printf("user %v created new table %s, code %s", user, table.Id, table.InviteCode)
	tableData := pbconv.ToPbTableData(table, user.Id, srv.createPbTableMember)
	usersInLobby, err := srv.UsersInLobby()
	if err != nil {
		return nil, err
	}
	srv.streams.Send(usersInLobby, &pb.Event{Event: &pb.Event_NewTable{NewTable: tableData}})
	return tableData, nil
}

func (srv *dokoserver) UsersInLobby() ([]u.Id, error) {
	users, err := srv.users.ListIds()
	if err != nil {
		return nil, err
	}
	usersInLobby := make([]u.Id, 0)
	for _, uid := range users {
		if !srv.tables.IsAtAnyTable(uid) {
			usersInLobby = append(usersInLobby, uid)
		}
	}
	return usersInLobby, nil
}

func (srv *dokoserver) StartTable(ctx context.Context, req *pb.StartTableRequest) (*pb.MatchState, error) {
	user, _ := GetAuthenticatedUser(ctx)
	id, err := t.ParseTableId(req.TableId)
	if err != nil {
		log.Printf("could not parse table id %s: %v", req.TableId, err)
		return nil, toGrpcError(err)
	}
	log.Printf("starting table %s", id)
	table, err := srv.tables.StartTable(id, user.Id)
	if err != nil {
		return nil, toGrpcError(err)
	}
	matchData, err := srv.tables.GetMatchData(id)
	if err != nil {
		return nil, toGrpcError(err)
	}
	for _, player := range getOtherPlayers(table, user.Id) {
		state := pbconv.ToPbMatchState(matchData, player)
		ev := &pb.Event{Event: &pb.Event_Start{Start: state}}
		srv.streams.SendSingle(player, ev)
	}
	return pbconv.ToPbMatchState(matchData, user.Id), nil
}

func (srv *dokoserver) JoinTable(ctx context.Context, req *pb.JoinTableRequest) (*pb.TableState, error) {
	user, _ := GetAuthenticatedUser(ctx)
	var table *t.TableData
	var err error
	switch td := req.TableDescription.(type) {
	case *pb.JoinTableRequest_InviteCode:
		table, err = srv.tables.JoinTableByInviteCode(user.Id, td.InviteCode)
	case *pb.JoinTableRequest_TableId:
		var tableId t.TableId
		tableId, err = t.ParseTableId(td.TableId)
		if err != nil {
			return nil, err
		}
		table, err = srv.tables.JoinTableByTableId(user.Id, tableId)
	}

	if err != nil {
		log.Printf("user %v failed joining: %v", user, err)
		return nil, toGrpcError(err)
	}
	log.Printf("user %v joined table %v", user, table.Id)
	srv.streams.Send(getOtherPlayers(table, user.Id), pb.NewMemberEvent(user.Id.String(), user.Name, pb.MemberEventType_JOIN_TABLE))
	return srv.getTableState(table, user.Id)
}

func (srv *dokoserver) getUserState(user u.AccountData) (*pb.UserState, error) {
	ans := &pb.UserState{}
	activeTable := srv.tables.ActiveTableOf(user.Id)
	if activeTable != nil {
		tableState, err := srv.getTableState(activeTable, user.Id)
		if err != nil {
			return nil, toGrpcError(err)
		}
		ans.CurrentTable = tableState
	}
	return ans, nil
}



func (srv *dokoserver) startEventSubscription(user u.AccountData, pbSrv pb.Doko_StartSessionServer) (kicked chan int, err error) {
	kicked = srv.streams.StartNew(pbSrv, user.Id)
	userState, err := srv.getUserState(user)
	if err != nil {
		return nil, err
	}
	srv.streams.SendSingle(user.Id, &pb.Event{Event: &pb.Event_Welcome{Welcome: userState}})
	receivers := srv.tables.UsersAtSameTable(user.Id)
	if receivers != nil {
		srv.streams.Send(receivers, pb.NewMemberEvent(user.Id.String(), user.Name, pb.MemberEventType_GO_ONLINE))
	}
	log.Printf("user %s connected", user)
	return
}

func (srv *dokoserver) endEventSubscription(pbSrv pb.Doko_StartSessionServer, user u.AccountData) error {
	log.Printf("user %s disconnected", user)
	table := srv.tables.ActiveTableOf(user.Id)
	if table != nil {
		receivers := srv.tables.UsersAtSameTable(user.Id)
		if receivers != nil {
			srv.streams.Send(receivers, pb.NewMemberEvent(user.Id.String(), user.Name, pb.MemberEventType_GO_OFFLINE))
		}
	}
	return pbSrv.Context().Err()
}

func getOtherPlayers(table *t.TableData, player u.Id) []u.Id {
	return u.IdsExcept(table.Players, player)
}

func (srv *dokoserver) Declare(ctx context.Context, d *pb.DeclareRequest) (*pb.Declaration, error) {
	user, _ := GetAuthenticatedUser(ctx)
	tableId, err := t.ParseTableId(d.Table)
	if err != nil {
		log.Printf("declare: could not parse table id %s: %v\n", d.Table, err)
		return nil, toGrpcError(err)
	}
	gameType := pbconv.ToGameType(d.Declaration)
	m, err := srv.tables.Declare(tableId, user.Id, gameType)
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
	srv.streams.Send(srv.tables.UsersAtSameTable(user.Id), event)
	return declaration, nil
}

func (srv *dokoserver) PlaceBid(ctx context.Context, req *pb.PlaceBidRequest) (*pb.Bid, error) {
	user, _ := GetAuthenticatedUser(ctx)
	tableId, err := t.ParseTableId(req.Table)
	if err != nil {
		log.Printf("placeBid: could not parse table id %s: %v\n", req.Table, err)
		return nil, toGrpcError(err)
	}
	_, err = srv.tables.PlaceBid(tableId, user.Id, pbconv.ToBid(req.Bid))
	if err != nil {
		return nil, toGrpcError(err)
	}
	bid := &pb.Bid{UserId: user.Id.String(), Bid: req.Bid}
	event := &pb.Event{Event: &pb.Event_PlacedBid{PlacedBid: bid}}
	srv.streams.Send(srv.tables.UsersAtSameTable(user.Id), event)
	return bid, nil
}

func (srv *dokoserver) PlayCard(ctx context.Context, req *pb.PlayCardRequest) (*pb.PlayedCard, error) {
	user, _ := GetAuthenticatedUser(ctx)
	tableId, err := t.ParseTableId(req.Table)
	if err != nil {
		log.Printf("playCard: could not parse table id %s: %v\n", req.Table, err)
		return nil, toGrpcError(err)
	}
	m, err := srv.tables.PlayCard(tableId, user.Id, pbconv.ToCard(req.Card))
	if err != nil {
		log.Printf("%s failed to play %s: %s", user.Name, pbconv.ToCard(req.Card), err)
		return nil, toGrpcError(err)
	}
	log.Printf("%s plays %s", user.Name, pbconv.ToCard(req.Card))
	table, err := srv.tables.GetTable(tableId)
	if err != nil {
		return nil, toGrpcError(err)
	}
	card := &pb.PlayedCard{UserId: user.Id.String(), Card: req.Card}
	if m.CurrentTrick != nil && m.CurrentTrick.NumCardsPlayed() == 0 {
		card.TrickWinner = &pb.PlayerValue{UserId: table.Players[m.PreviousTrick.Winner].String()}
		log.Printf("trick finished. winner: %s", table.Players[m.PreviousTrick.Winner].String())
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
	srv.streams.Send(getOtherPlayers(table, user.Id), event)
	return card, nil
}

func (srv *dokoserver) StartNextMatch(ctx context.Context, req *pb.StartNextMatchRequest) (*pb.MatchState, error) {
	user, _ := GetAuthenticatedUser(ctx)
	tableId, err := t.ParseTableId(req.TableId)
	if err != nil {
		log.Printf("could not parse table id %s: %v\n", req.TableId, err)
		return nil, toGrpcError(err)
	}
	matchData, err := srv.tables.StartNextMatch(tableId, user.Id)
	if err != nil {
		return nil, toGrpcError(err)
	}
	table, err := srv.tables.GetTable(tableId)
	if err != nil {
		return nil, toGrpcError(err)
	}
	for _, player := range getOtherPlayers(table, user.Id) {
		state := pbconv.ToPbMatchState(matchData, player)
		ev := &pb.Event{Event: &pb.Event_Start{Start: state}}
		srv.streams.SendSingle(player, ev)
	}
	return pbconv.ToPbMatchState(matchData, user.Id), nil
}

func (srv *dokoserver) ListTables(ctx context.Context, _ *pb.ListTablesRequest) (*pb.ListTablesResult, error) {
	user, _ := GetAuthenticatedUser(ctx)
	tables := srv.tables.GetOpenTables(user.Id)
	pbResult := pbconv.ToPbTables(tables, user.Id, srv.createPbTableMember)
	return &pb.ListTablesResult{Tables: pbResult}, nil
}

func (srv *dokoserver) getTableState(table *t.TableData, user u.Id) (*pb.TableState, error) {
	data := pbconv.ToPbTableData(table, user, srv.createPbTableMember)
	state := &pb.TableState{Data: data, Phase: table.Phase}
	if table.Phase == pb.TablePhase_PLAYING {
		matchData, err := srv.tables.GetMatchData(table.Id)
		if err != nil {
			return nil, err
		}
		state.CurrentMatch = pbconv.ToPbMatchState(matchData, user)
	}
	return state, nil
}

func (srv *dokoserver) createPbTableMember(id u.Id) *pb.TableMember {
	data, err := srv.users.GetData(id)
	if err != nil {
		panic("not existingt user at table - should not be here!")
	}
	return &pb.TableMember{UserId: id.String(), Name: data.Name, Online: srv.streams.IsOnline(id)}
}
