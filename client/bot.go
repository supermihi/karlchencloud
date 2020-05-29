package client

import (
	"context"
	"fmt"
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/doko/match"
	"log"
)

func StartBots(address string, numBots int, table string, inviteCode string) {
	clients := make([]*BotHandler, numBots)
	for i := 0; i < numBots; i++ {
		connect := ConnectData{
			DisplayName:    fmt.Sprintf("Bot %v", i+1),
			ExistingUserId: nil,
			ExistingSecret: nil,
			Address:        address,
		}
		clients[i] = NewBotHandler(table, inviteCode, connect)
		go clients[i].Run()
	}
	for i := 0; i < numBots; i++ {
		<-clients[i].context.Done()
	}
	log.Printf("all bots finished")

}

type BotHandler struct {
	inviteCode  string
	connectData ConnectData
	context     context.Context
	cancel      context.CancelFunc
	service     ClientService
	client      TableClient
	view        MatchView
}

func NewBotHandler(table string, inviteCode string, connect ConnectData) *BotHandler {
	ctx, cancel := context.WithCancel(context.Background())
	return &BotHandler{
		view:        NewMatchView(table),
		inviteCode:  inviteCode,
		connectData: connect,
		context:     ctx,
		cancel:      cancel,
		service:     ClientService{}}
}

func (h *BotHandler) HandleMatchStart(state *api.MatchState) {
	h.view.InitFromMatchState(state)
	h.checkMyTurn()
}

func (h *BotHandler) onMyTurn() {
	switch h.view.Phase {
	case match.InAuction:
		go h.makeTurnAuction()
	case match.InGame:
		go h.makeTurnGame()
	}
}

func (h *BotHandler) makeTurnAuction() {

	declaration := api.GameType_NORMAL_GAME
	if h.view.Cards.NumAlte() == 2 {
		declaration = api.GameType_MARRIAGE
	}
	h.service.Logf("declaring '%s'...", declaration)
	p := api.PlayRequest{Table: h.view.TableId, Request: &api.PlayRequest_Declaration{Declaration: declaration}}
	_, err := h.service.Api.Play(h.context, &p)
	if err != nil {
		log.Fatalf("could not make auction turn: %v", err)
	}

}

func (h *BotHandler) makeTurnGame() {
	cardIndex := -1
	trick := h.view.Trick
	if len(trick.Cards) == 0 {
		// I am forehand
		cardIndex = 0
	} else {
		firstCard := trick.Cards[trick.Forehand]
		gs := h.view.GameSuit(firstCard)
		for i, card := range h.view.Cards {
			if h.view.GameSuit(card) == gs {
				cardIndex = i
				break
			}
		}
	}
	if cardIndex == -1 {
		cardIndex = 0 // no matchnig card -> can play anything
	}
	card := h.view.DrawCard(cardIndex)
	err := h.service.Play(card, h.view.TableId)
	if err != nil {
		log.Fatalf("could not play card: %v", err)
	}

}

func (h *BotHandler) HandleMemberEvent(m *api.MemberEvent) {
	// pass
}

func (h *BotHandler) HandleDeclared(d *api.Declaration) {
	h.view.UpdateOnDeclare(d)
	h.checkMyTurn()
}

func (h *BotHandler) checkMyTurn() {
	if h.view.MyTurn {
		h.onMyTurn()
	}
}
func (h *BotHandler) HandlePlayedCard(c *api.PlayedCard) {
	h.view.UpdateTrick(c)
	h.checkMyTurn()
}

func (h *BotHandler) HandleEnd(ev *api.EndOfGame) {
	h.service.Logf("game ended with winner %s.", ev.Winner)
}

func (h *BotHandler) Run() {
	conn := h.connectData
	ctx := h.context
	service, err := GetClientService(conn, ctx)
	if err != nil {
		log.Print(err)
		return
	}
	h.service = service
	c := service.Api
	defer h.cancel()
	defer service.CloseConnection()
	_, err = c.JoinTable(ctx, &api.JoinTableRequest{InviteCode: h.inviteCode, TableId: h.view.TableId})
	if err != nil {
		h.service.Logf("could not join table: %v", err)
		return
	}
	h.service.Logf("joined")
	h.client = NewTableClient(h.context, h.service, h.view.TableId, h)
	h.client.Start()
}
