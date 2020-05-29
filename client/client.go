package client

import (
	"context"
	"github.com/supermihi/karlchencloud/api"
	"log"
)

type ClientHandler interface {
	HandleMatchStart(state *api.MatchState)
	HandleMemberEvent(ev *api.MemberEvent)
	HandleDeclared(ev *api.Declaration)
	HandlePlayedCard(ev *api.PlayedCard)
	HandleEnd(ev *api.EndOfGame)
	OnInitialState(state *api.TableState)
}
type TableClient struct {
	ctx     context.Context
	service ClientService
	handler ClientHandler
	tableId string
}

func NewTableClient(ctx context.Context, client ClientService, tableId string, handler ClientHandler) TableClient {
	return TableClient{ctx: ctx, service: client, handler: handler, tableId: tableId}
}

func (c *TableClient) Start() {
	stream, err := c.service.Api.SubscribeMatchEvents(c.ctx, &api.TableId{Value: c.tableId})
	c.service.Logf("Listening for match events ...")
	if err != nil {
		c.service.Logf("error subscribing: %v", err)
		return
	}
	go func() {
		state, err := c.service.Api.GetTableState(c.ctx, &api.TableId{Value: c.tableId})
		if err != nil {
			c.handler.OnInitialState(state)
		}
	}()
	handler := c.handler
	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Fatalf("error receiving from stream: %v", err)
		}
		switch ev := msg.Event.(type) {
		case *api.MatchEventStream_Member:
			handler.HandleMemberEvent(ev.Member)
		case *api.MatchEventStream_Start:
			handler.HandleMatchStart(ev.Start)
		case *api.MatchEventStream_Declared:
			handler.HandleDeclared(ev.Declared)
		case *api.MatchEventStream_PlayedCard:
			handler.HandlePlayedCard(ev.PlayedCard)
		case *api.MatchEventStream_Ended:
			handler.HandleEnd(ev.Ended)
		default:
			log.Fatalf("unimplemented event occured: %v", msg)
		}
	}
}
