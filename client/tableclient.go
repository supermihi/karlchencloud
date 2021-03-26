package client

import (
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/server"
	"log"
)

type ClientHandler interface {
	OnMyTurn()
	OnTableStateReceived(state *api.TableState)
	OnMemberEvent(ev *api.MemberEvent)
	OnMatchStart(s *api.MatchState)
	OnDeclaration(d *api.Declaration)
	OnPlayedCard(card *api.PlayedCard)
	OnMatchEnd(end *api.EndOfGame)
}
type TableClient struct {
	Service ClientService
	handler ClientHandler
	TableId string
	View    *TableView
}

func NewTableClient(client ClientService, tableId string, handler ClientHandler) TableClient {
	return TableClient{Service: client, handler: handler, TableId: tableId}
}

func (c *TableClient) Logf(format string, v ...interface{}) {
	c.Service.Logf(format, v...)
}
func (c *TableClient) Start() {
	stream, err := c.Service.Api.StartSession(c.Service.Context, &api.Empty{})
	c.Service.Logf("Listening for match events ...")
	if err != nil {
		c.Service.Logf("error subscribing: %v", err)
		return
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
				c.View.MemberNamesById[ev.Member.UserId] = ev.Member.Name
			}
			c.handler.OnMemberEvent(ev.Member)
		case *api.Event_Start:
			c.HandleStart(ev.Start)
		case *api.Event_Declared:
			c.handleDeclare(ev.Declared)
		case *api.Event_PlayedCard:
			c.handlePlayedCard(ev.PlayedCard)
		case *api.Event_Ended:
			c.handleMatchEnded(ev.Ended)
		default:
			log.Fatalf("unimplemented event occured: %v", msg)
		}
	}
}

func (c *TableClient) handleWelcome(us *api.UserState) {
	ts := us.CurrentTable
	if ts == nil {
		panic("table client expects active table")
	}
	c.View = NewTableView(ts)
	c.handler.OnTableStateReceived(ts)
	c.checkMyTurn()
}

func (c *TableClient) HandleStart(s *api.MatchState) {
	c.View.Match = NewMatchView(s)
	c.handler.OnMatchStart(s)
	c.checkMyTurn()
}

func (c *TableClient) handleDeclare(d *api.Declaration) {
	c.View.Match.UpdateOnDeclare(d)
	c.handler.OnDeclaration(d)
	c.checkMyTurn()
}

func (c *TableClient) handlePlayedCard(card *api.PlayedCard) {
	c.View.Match.UpdateTrick(card)
	c.handler.OnPlayedCard(card)
	c.checkMyTurn()
}

func (c *TableClient) handleMatchEnded(end *api.EndOfGame) {
	c.handler.OnMatchEnd(end)
}
func (c *TableClient) checkMyTurn() {
	if c.View.Match != nil && c.View.Match.MyTurn {
		c.handler.OnMyTurn()
	}
}
func (c *TableClient) PlayCard(card game.Card) error {
	result, err := c.Service.Api.PlayCard(
		c.Service.Context,
		&api.PlayCardRequest{Table: c.TableId, Card: server.ToApiCard(card)})
	if err == nil {
		c.View.Match.UpdateTrick(result)
	}
	return err
}

func (c *TableClient) Declare(t game.AnnouncedGameType) error {
	result, err := c.Service.Api.Declare(c.Service.Context, &api.DeclareRequest{
		Table:       c.TableId,
		Declaration: server.ToApiGameType(t)})
	if err == nil {
		c.View.Match.UpdateOnDeclare(result)
	}
	return err
}

func (c *TableClient) Api() api.DokoClient {
	return c.Service.Api
}

func (c *TableClient) Match() *MatchView {
	return c.View.Match
}
