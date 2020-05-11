package cloud

import (
	"context"
	"fmt"
	grpcauth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	uuid "github.com/satori/go.uuid"
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/pkg/game/core"
	"github.com/supermihi/karlchencloud/pkg/game/match"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

type grpcserver struct {
	api.UnimplementedKarlchencloudServer
	room Room
	auth Auth
}

func (s *grpcserver) Register(_ context.Context, req *api.RegisterRequest) (*api.RegisterReply, error) {
	id := uuid.NewV4().String()
	secret := RandomSecret()
	s.room.Users.Add(UserId(id), req.GetName(), secret)
	log.Printf("Registered user %v with id %v", req.GetName(), id)
	return &api.RegisterReply{Id: id, Secret: secret}, nil
}

func (s *grpcserver) CheckLogin(ctx context.Context, _ *api.EmptyRequest) (*api.OkOrNot, error) {
	user, ok := GetAuthenticatedUser(ctx)
	if !ok {
		log.Print("check login failed")
		return &api.OkOrNot{Value: false}, nil
	}
	log.Printf("user %v ok", user)
	return &api.OkOrNot{Value: true}, nil
}

func (s *grpcserver) CreateTable(ctx context.Context, _ *api.EmptyRequest) (*api.TableData, error) {
	user, _ := GetAuthenticatedUser(ctx)
	table := s.room.CreateTable(user)
	log.Printf("user %v created new table %v", s.room.Users.GetName(user), table)
	return toTableData(table, user), nil
}

func toTableData(table *Table, user UserId) *api.TableData {
	exposedInviteCode := ""
	if table.Owner() == user {
		exposedInviteCode = table.inviteCode
	}
	return &api.TableData{TableId: table.id, Owner: string(table.Owner()), InviteCode: exposedInviteCode}
}

func (s *grpcserver) ListTables(ctx context.Context, _ *api.EmptyRequest) (*api.TableList, error) {
	user, _ := GetAuthenticatedUser(ctx)
	tables := s.room.tables.List()
	result := make([]*api.TableData, len(tables))
	for i, table := range tables {
		result[i] = toTableData(table, user)
	}
	return &api.TableList{Tables: result}, nil
}

func (s *grpcserver) StartTable(ctx context.Context, id *api.TableId) (*api.EmptyReply, error) {
	user, _ := GetAuthenticatedUser(ctx)
	table, err := s.tryGetTable(id.Value)
	if err != nil {
		return nil, err
	}
	if table.Owner() != user {
		return nil, status.Error(codes.PermissionDenied, "you are not owner of the table")
	}
	err = table.Start()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &api.EmptyReply{}, nil

}

func (s *grpcserver) tryGetTable(id string) (*Table, error) {
	table, ok := s.room.tables.GetTable(id)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "table does not exist")
	}
	return table, nil
}

func (s *grpcserver) JoinTable(ctx context.Context, req *api.JoinTableRequest) (*api.EmptyReply, error) {
	user, _ := GetAuthenticatedUser(ctx)
	table, err := s.tryGetTable(req.TableId)
	if err != nil {
		return nil, err
	}
	if table.inviteCode != req.InviteCode {
		return nil, status.Error(codes.PermissionDenied, "invalid invite code")
	}
	err = table.Join(user)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &api.EmptyReply{}, nil
}

func StartServer(users Users, port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	room := NewRoom(users)
	auth := NewAuth(users)
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpcauth.UnaryServerInterceptor(auth.Authenticate)),
		grpc.StreamInterceptor(grpcauth.StreamServerInterceptor(auth.Authenticate)))

	serv := grpcserver{room: room, auth: auth}
	api.RegisterKarlchencloudServer(grpcServer, &serv)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func toPlayer(p core.Player, mapNoPlayerToPlayer1 bool) api.Player {
	switch p {
	case core.Player1:
		return api.Player_PLAYER_1
	case core.Player2:
		return api.Player_PLAYER_2
	case core.Player3:
		return api.Player_PLAYER_3
	case core.Player4:
		return api.Player_PLAYER_4
	default:
		if mapNoPlayerToPlayer1 {
			return api.Player_PLAYER_1
		}
		panic("unexpected NoPlayer in toPlayer")
	}
}

func toAuctionPhase(p match.AuctionPhase) api.AuctionPhase {
	switch p {
	case match.VorbehaltAbfrage:
		return api.AuctionPhase_DECLARATION
	case match.VorbehaltSpezifikation:
		return api.AuctionPhase_SPECIFICATION
	}
	panic("unsupported auction phase")
}

func toAuctionState(a *match.Auction) *api.AuctionState {
	declarations := make([]*api.Declaration, 0)
	for _, p := range core.Players() {
		decl := a.DeclarationOf(p)
		if decl != match.NotDeclared {
			apiDecl := &api.Declaration{Player: toPlayer(p, false), Vorbehalt: decl == match.Vorbehalt}
			declarations = append(declarations, apiDecl)
		}
	}
	return &api.AuctionState{Phase: toAuctionPhase(a.Phase()), Declarations: declarations}
}
func toSoloType()
func toApiMode(mode core.Mode) *api.Mode {
	var gameType api.GameType
	var soloInfo *api.SoloInfo
	switch mode.(type) {
	case core.NormalspielMode:
		gameType = api.GameType_NORMAL_GAME
	case core.Hochzeit:
		gameType = api.GameType_MARRIAGE
	default:
		gameType = api.GameType_VOLUNTARY_SOLO
		soloInfo = &api.SoloInfo{Soloist: toPlayer(core.Soloist(mode), false), SoloType: toSoloType()}
	}
	return &api.Mode{Type: gameType}
}
func toGameState(m *match.Match) *api.GameState {
	return &api.GameState{Mode: m.Mode().}
}
func toMatchState(m *match.Match) *api.MatchState {
	turn := toPlayer(m.WhoseTurn(), true)
	switch m.Phase() {
	case match.InAuction:
		auctionState := toAuctionState(m.Auction())
		return &api.MatchState{Phase: api.MatchPhase_AUCTION,
			Turn:    turn,
			Details: &api.MatchState_AuctionState{AuctionState: auctionState}}
	case match.InGame:
		gameState = toGameState(m.Game())

	}

}
func (s *grpcserver) GetMatchState(ctx context.Context, tableId *api.TableId) (*api.MatchState, error) {
	table, err := s.tryGetTable(tableId.Value)
	if err != nil {
		return nil, err
	}
	m := table.currentMatch
	if m == nil {
		return nil, status.Error(codes.Internal, "no active match at table")
	}

}

func toMatchPhase(phase match.Phase) api.MatchPhase {
	switch phase {
	case match.AuctionAbfragePhase:
		return api.MatchPhase_AUCTION_DECLARATION
	case match.AuctionSpezifikationPhase:
		return api.MatchPhase_AUCTION_SPECIFICATION
	case match.InGame:
		return api.MatchPhase_GAME
	case match.MatchFinished:
		return api.MatchPhase_FINISHED
	default:
		panic(fmt.Sprintf("unexpected phase %v in toMatchPhase", phase))
	}
}
