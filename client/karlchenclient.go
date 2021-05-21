package client

import (
	"context"
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/server"
	"log"
)

type KarlchenClient struct {
	clientData LoginData
	Service    *ClientService
	handler    ClientHandler
	TableId    string
	View       *TableView
}

func NewKarlchenClient(c LoginData, handler ClientHandler) KarlchenClient {
	return KarlchenClient{clientData: c, handler: handler}
}

func (c *KarlchenClient) Logf(format string, v ...interface{}) {
	c.Service.Logf(format, v...)
}
func (c *KarlchenClient) Start(ctx context.Context) {
	service, err := GetConnectedClientService(c.clientData, ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer service.CloseConnection()
	c.Service = service
	c.handler.OnConnect(c.Service)
	stream, err := c.Service.Grpc.StartSession(c.Service.Context, &api.Empty{})
	c.Service.Logf("Listening for match events ...")
	if err != nil {
		c.Service.Logf("error subscribing: %v", err)
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
				c.View.MemberNamesById[ev.Member.UserId] = ev.Member.Name
			}
			c.handler.OnMemberEvent(ev.Member)
		case *api.Event_Start:
			c.HandleStart(ev.Start)
		case *api.Event_Declared:
			c.handleDeclare(ev.Declared)
		case *api.Event_PlayedCard:
			c.handlePlayedCard(ev.PlayedCard)
		default:
			log.Fatalf("unimplemented event occured: %v", msg)
		}
	}
}

func (c *KarlchenClient) handleWelcome(us *api.UserState) {
	c.handler.OnWelcome(c, us)
	ts := us.CurrentTable
	if ts != nil {
		c.TableId = ts.Data.TableId
		c.Logf("sitting at table %s", c.TableId)
	}
	c.View = NewTableView(ts)
	c.handler.OnTableStateReceived(ts)
	c.checkMyTurn()
}

func (c *KarlchenClient) HandleStart(s *api.MatchState) {
	c.View.Match = NewMatchView(s)
	c.handler.OnMatchStart(s)
	c.checkMyTurn()
}

func (c *KarlchenClient) handleDeclare(d *api.Declaration) {
	c.View.Match.UpdateOnDeclare(d)
	c.handler.OnDeclaration(d)
	c.checkMyTurn()
}

func (c *KarlchenClient) handlePlayedCard(card *api.PlayedCard) {
	c.View.Match.UpdateTrick(card)
	c.handler.OnPlayedCard(card)
	c.checkMyTurn()
}

func (c *KarlchenClient) checkMyTurn() {
	if c.View.Match != nil && c.View.Match.MyTurn && len(c.View.Match.Cards) > 0 {
		c.handler.OnMyTurn()
	}
}
func (c *KarlchenClient) PlayCard(card game.Card) error {
	result, err := c.Service.Grpc.PlayCard(
		c.Service.Context,
		&api.PlayCardRequest{Table: c.TableId, Card: server.ToApiCard(card)})
	if err == nil {
		c.handlePlayedCard(result)
	}
	return err
}

func (c *KarlchenClient) Declare(t game.AnnouncedGameType) error {
	result, err := c.Service.Grpc.Declare(c.Service.Context, &api.DeclareRequest{
		Table:       c.TableId,
		Declaration: server.ToApiGameType(t)})
	if err == nil {
		c.View.Match.UpdateOnDeclare(result)
	}
	return err
}

func (c *KarlchenClient) Api() api.DokoClient {
	return c.Service.Grpc
}

func (c *KarlchenClient) Match() *MatchView {
	return c.View.Match
}
