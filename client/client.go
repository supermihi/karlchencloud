package client

import (
	"context"
	"fmt"
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
	"github.com/supermihi/karlchencloud/server/pbconv"
	"log"
)

type Client struct {
	clientData LoginData
	client     *DokoClient
	handler    ClientHandler
	table      *TableView
}

func NewClient(c LoginData, handler ClientHandler) Client {
	return Client{clientData: c, handler: handler}
}

func (c *Client) Logf(format string, v ...interface{}) {
	log.Printf(c.clientData.Name+": "+format, v...)
}

func (c *Client) Table() *TableView {
	return c.table
}

func (c *Client) User() UserData {
	return c.client.user
}

func (c *Client) Start(ctx context.Context) {
	dokoclient, err := GetConnectedDokoClient(c.clientData, ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer dokoclient.CloseConnection()
	c.client = dokoclient
	c.handler.OnConnect()
	stream, err := c.client.Grpc.StartSession(c.client.Context, &api.Empty{})
	c.Logf("Listening for match events ...")
	if err != nil {
		c.Logf("error subscribing: %v", err)
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
		case *api.Event_NewTable:
			c.handler.OnNewTable(NewTableInfo(ev.NewTable))
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

func (c *Client) ListOpenTables() ([]TableInfo, error) {
	response, err := c.client.Grpc.ListTables(c.client.Context, &api.ListTablesRequest{})
	if err != nil {
		return nil, err
	}
	result := make([]TableInfo, len(response.Tables))
	for i, table := range response.Tables {
		result[i] = NewTableInfo(table)
	}
	return result, nil
}

func (c *Client) handleMemberEvent(ev *api.MemberEvent) {
	switch ev.Type {
	case api.MemberEventType_JOIN_TABLE:
		c.Logf("user %s joined table", ev.Name)
		c.handler.OnMemberJoin(ev.UserId, ev.Name)
	case api.MemberEventType_GO_ONLINE:
		c.Logf("user %s is now online", c.table.MemberNamesById[ev.UserId])
	case api.MemberEventType_GO_OFFLINE:
		c.Logf("user %s is now offline", c.table.MemberNamesById[ev.UserId])
	default:
		c.Logf("unexpected MemberEvent: %v", ev)
	}
}

func (c *Client) CreateTable(public bool) error {
	tableData, err := c.client.Grpc.CreateTable(c.client.Context, &api.CreateTableRequest{Public: public})
	if err != nil {
		return err
	}
	c.Logf("table %s created with invite code %s. Waiting for players ...", tableData.TableId, tableData.InviteCode)
	c.initView(&api.TableState{Data: tableData, Phase: api.TablePhase_NOT_STARTED})
	return nil
}

func (c *Client) JoinTable(invite string, tableId string) (err error) {
	request := &api.JoinTableRequest{}
	if invite != "" {
		request.TableDescription = &api.JoinTableRequest_InviteCode{InviteCode: invite}
	} else {
		request.TableDescription = &api.JoinTableRequest_TableId{TableId: tableId}
	}
	tableState, err := c.client.Grpc.JoinTable(c.client.Context, request)

	if err == nil {
		c.Logf("table %s joined", tableState.Data.TableId)
		c.initView(tableState)
	}
	return
}

func (c *Client) StartTable() error {
	matchState, err := c.client.Grpc.StartTable(c.client.Context, &api.StartTableRequest{TableId: c.table.Id})
	if err == nil {
		c.handleStart(matchState)
	}
	return err
}

func (c *Client) initView(state *api.TableState) {
	c.table = NewTableView(state)
}
func (c *Client) handleWelcome(us *api.UserState) {
	c.handler.OnWelcome(us)
	ts := us.CurrentTable
	if ts == nil {
		return
	}
	c.Logf("sitting at table %s", ts.Data.TableId)
	c.initView(ts)
	c.checkMyTurn()
}

func (c *Client) handleStart(s *api.MatchState) {
	c.table.Match = NewMatchView(s)
	c.handler.OnMatchStart()
	c.checkMyTurn()
}

func (c *Client) handleDeclare(d *api.Declaration) {
	c.Match().UpdateOnDeclare(d)
	if c.Match().Phase == match.InGame {
		c.Logf("game phase started (forehand is %s)", c.Table().MemberNamesById[c.Match().Trick.Forehand])
	}
	c.checkMyTurn()
}

func (c *Client) handlePlayedCard(card *api.PlayedCard) {
	c.table.Match.UpdateTrick(card)
	if card.Winner != nil {
		c.Match().Phase = match.MatchFinished
	}
	c.handler.OnPlayedCard(card)
	c.checkMyTurn()
}

func (c *Client) checkMyTurn() {
	matchView := c.Match()
	if matchView != nil && matchView.MyTurn && len(matchView.Cards) > 0 {
		switch matchView.Phase {
		case match.InAuction:
			c.handler.OnMyTurnAuction()
		case match.InGame:
			c.handler.OnMyTurnGame()
		default:
			panic(fmt.Sprintf("should not be here: handleMyTurn in neither auction nor game"))
		}
	}
}
func (c *Client) PlayCard(i int) error {
	card := c.Match().Cards[i]
	c.Logf("playing card: %v", card)
	result, err := c.client.Grpc.PlayCard(
		c.client.Context,
		&api.PlayCardRequest{Table: c.table.Id, Card: pbconv.ToPbCard(card)})
	if err == nil {
		c.Match().DrawCard(i)
		c.handlePlayedCard(result)
	}
	return err
}

func (c *Client) Declare(t game.AnnouncedGameType) error {
	result, err := c.client.Grpc.Declare(c.client.Context, &api.DeclareRequest{
		Table:       c.table.Id,
		Declaration: pbconv.ToPbGameType(t)})
	if err == nil {
		c.Logf("declared %s", t)
		c.table.Match.UpdateOnDeclare(result)
	}
	return err
}

func (c *Client) Api() api.DokoClient {
	return c.client.Grpc
}

func (c *Client) Match() *MatchView {
	return c.table.Match
}

func (c *Client) StartNextMatch() error {
	req := api.StartNextMatchRequest{TableId: c.table.Id}
	ans, err := c.client.Grpc.StartNextMatch(c.client.Context, &req)
	if err != nil {
		return err
	}
	c.handleStart(ans)
	return nil
}
