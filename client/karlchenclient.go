package client

import (
	"context"
	"fmt"
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
	"github.com/supermihi/karlchencloud/server"
	"log"
)

type KarlchenClient struct {
	clientData LoginData
	service    *ClientService
	handler    ClientHandler
	table      *TableView
}

func NewKarlchenClient(c LoginData, handler ClientHandler) KarlchenClient {
	return KarlchenClient{clientData: c, handler: handler}
}

func (c *KarlchenClient) Logf(format string, v ...interface{}) {
	c.service.Logf(format, v...)
}

func (c *KarlchenClient) UserId() string {
	return c.service.UserId()
}

func (c *KarlchenClient) Table() *TableView {
	return c.table
}

func (c *KarlchenClient) Start(ctx context.Context) {
	service, err := GetConnectedClientService(c.clientData, ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer service.CloseConnection()
	c.service = service
	c.handler.OnConnect(c)
	stream, err := c.service.Grpc.StartSession(c.service.Context, &api.Empty{})
	c.service.Logf("Listening for match events ...")
	if err != nil {
		c.service.Logf("error subscribing: %v", err)
		log.Fatal(err)
	}
	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Fatalf("error receiving from stream: %v", err)
		}
		switch ev := msg.Event.(type) {
		case *api.Event_Welcome:
			c.handleWelcome(ev.Welcome)
		case *api.Event_Member:
			if ev.Member.Type == api.MemberEventType_JOIN_TABLE {
				c.table.MemberNamesById[ev.Member.UserId] = ev.Member.Name
			}
			c.handleMemberEvent(ev.Member)
		case *api.Event_Start:
			c.handleStart(ev.Start)
		case *api.Event_Declared:
			c.handleDeclare(ev.Declared)
		case *api.Event_PlayedCard:
			c.handlePlayedCard(ev.PlayedCard)
		default:
			log.Fatalf("unimplemented event occured: %v", msg)
		}
	}
}

func (c *KarlchenClient) handleMemberEvent(ev *api.MemberEvent) {
	switch ev.Type {
	case api.MemberEventType_JOIN_TABLE:
		c.Logf("user %s joined table", ev.Name)
		c.handler.OnMemberJoin(c, ev.UserId, ev.Name)
	case api.MemberEventType_GO_ONLINE:
		c.Logf("user %s is now online", c.table.MemberNamesById[ev.UserId])
	case api.MemberEventType_GO_OFFLINE:
		c.Logf("user %s is now offline", c.table.MemberNamesById[ev.UserId])
	default:
		c.Logf("unexpected MemberEvent: %v", ev)
	}
}

func (c *KarlchenClient) CreateTable() error {
	tableData, err := c.service.Grpc.CreateTable(c.service.Context, &api.Empty{})
	if err != nil {
		return err
	}
	c.Logf("table %s created with invite code %s", tableData.TableId, tableData.InviteCode)
	c.initView(&api.TableState{Data: tableData, Phase: api.TablePhase_NOT_STARTED})
	return nil
}

func (c *KarlchenClient) JoinTable(invite string) (err error) {
	tableState, err := c.service.Grpc.JoinTable(c.service.Context, &api.JoinTableRequest{InviteCode: invite})
	if err == nil {
		c.Logf("table %s joined", tableState.Data.TableId)
		c.initView(tableState)
	}
	return
}

func (c *KarlchenClient) StartTable() error {
	matchState, err := c.service.Grpc.StartTable(c.service.Context, &api.StartTableRequest{TableId: c.table.Id})
	if err == nil {
		c.handleStart(matchState)
	}
	return err
}

func (c *KarlchenClient) initView(state *api.TableState) {
	c.table = NewTableView(state)
}
func (c *KarlchenClient) handleWelcome(us *api.UserState) {
	c.handler.OnWelcome(c, us)
	ts := us.CurrentTable
	if ts != nil {
		c.Logf("sitting at table %s", ts.Data.TableId)
	}
	c.initView(ts)
	c.checkMyTurn()
}

func (c *KarlchenClient) handleStart(s *api.MatchState) {
	c.table.Match = NewMatchView(s)
	c.handler.OnMatchStart(c)
	c.checkMyTurn()
}

func (c *KarlchenClient) handleDeclare(d *api.Declaration) {
	c.table.Match.UpdateOnDeclare(d)
	c.handler.OnDeclaration(c, d)
	c.checkMyTurn()
}

func (c *KarlchenClient) handlePlayedCard(card *api.PlayedCard) {
	c.table.Match.UpdateTrick(card)
	c.handler.OnPlayedCard(c, card)
	c.checkMyTurn()
}

func (c *KarlchenClient) checkMyTurn() {
	matchView := c.Match()
	if matchView != nil && matchView.MyTurn && len(matchView.Cards) > 0 {
		switch matchView.Phase {
		case match.InAuction:
			c.handler.OnMyTurnAuction(c)
		case match.InGame:
			c.handler.OnMyTurnGame(c)
		default:
			panic(fmt.Sprintf("should not be here: handleMyTurn in neither auction nor game"))
		}
	}
}
func (c *KarlchenClient) PlayCard(i int) error {
	card := c.Match().Cards[i]
	log.Printf("playing card: %v", card)
	result, err := c.service.Grpc.PlayCard(
		c.service.Context,
		&api.PlayCardRequest{Table: c.table.Id, Card: server.ToApiCard(card)})
	if err == nil {
		c.Match().DrawCard(i)
		c.handlePlayedCard(result)
	}
	return err
}

func (c *KarlchenClient) Declare(t game.AnnouncedGameType) error {
	result, err := c.service.Grpc.Declare(c.service.Context, &api.DeclareRequest{
		Table:       c.table.Id,
		Declaration: server.ToApiGameType(t)})
	if err == nil {
		c.Logf("successfully declared %s", t)
		c.table.Match.UpdateOnDeclare(result)
	}
	return err
}

func (c *KarlchenClient) Api() api.DokoClient {
	return c.service.Grpc
}

func (c *KarlchenClient) Match() *MatchView {
	return c.table.Match
}

func (c *KarlchenClient) StartNextMatch() error {
	req := api.StartNextMatchRequest{TableId: c.table.Id}
	ans, err := c.service.Grpc.StartNextMatch(c.service.Context, &req)
	if err != nil {
		return err
	}
	c.handleStart(ans)
	return nil
}

type ClientApi interface {
	Logf(format string, v ...interface{})
	UserId() string
	Table() *TableView
	Match() *MatchView
	CreateTable() error
	JoinTable(invite string) (err error)
	StartTable() error
	PlayCard(i int) error
	Declare(t game.AnnouncedGameType) error
	StartNextMatch() error
}
