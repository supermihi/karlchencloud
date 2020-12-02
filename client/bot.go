package client

import (
	"context"
	"fmt"
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
	"log"
	"math/rand"
	"time"
)

func StartBots(address string, numBots int, table string, inviteCode string,
	logins []BotLogin) {
	clients := make([]*BotHandler, numBots)
	for i := 0; i < numBots; i++ {
		var user, secret *string
		if len(logins) > i {
			user = &logins[i].Id
			secret = &logins[i].Secret
		}
		connect := ConnectData{
			DisplayName:    fmt.Sprintf("Bot %v", i+1),
			ExistingUserId: user,
			ExistingSecret: secret,
			Address:        address,
		}
		clients[i] = NewBotHandler()
		log.Printf("starting bot %d", i)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		go clients[i].Run(connect, table, inviteCode)
	}
	for i := 0; i < numBots; i++ {
		<-clients[i].context.Done()
	}
	log.Printf("all bots finished")

}

type BotHandler struct {
	TableClient
	context context.Context
	cancel  context.CancelFunc
}

func NewBotHandler() *BotHandler {
	ctx, cancel := context.WithCancel(context.Background())
	return &BotHandler{
		context: ctx,
		cancel:  cancel}
}

func (h *BotHandler) Run(conn ConnectData, table string, invite string) {
	ctx := h.context
	service, err := GetClientService(conn, ctx)
	if err != nil {
		log.Print(err)
		return
	}
	c := service.Api
	defer h.cancel()
	defer service.CloseConnection()
	_, err = c.JoinTable(ctx, &api.JoinTableRequest{InviteCode: invite})
	if err != nil {
		service.Logf("could not join table: %v", err)
		return
	}
	service.Logf("joined")
	h.TableClient = NewTableClient(service, table, h)
	h.TableClient.Start()
}

func (h *BotHandler) OnMatchStart(_ *api.MatchState) {
	// pass
}

func (h *BotHandler) OnMyTurn() {
	time.Sleep(500 * time.Millisecond)
	switch h.Match().Phase {
	case match.InAuction:
		h.makeTurnAuction()
	case match.InGame:
		h.makeTurnGame()
	}
}

func (h *BotHandler) makeTurnAuction() {

	declaration := game.NormalspielType
	if h.Match().Cards.NumAlte() == 2 {
		declaration = game.HochzeitType
	}
	h.Logf("declaring '%s'...", declaration)
	err := h.Declare(declaration)
	if err != nil {
		log.Fatalf("could not make auction turn: %v", err)
	}

}

func (h *BotHandler) makeTurnGame() {
	cardIndex := -1
	m := h.Match()
	trick := m.Trick
	if len(trick.Cards) == 0 {
		// I am forehand
		cardIndex = 0
	} else {
		firstCard := trick.Cards[trick.Forehand]
		gs := m.GameSuit(firstCard)
		for i, card := range m.Cards {
			if m.GameSuit(card) == gs {
				cardIndex = i
				break
			}
		}
	}
	if cardIndex == -1 {
		cardIndex = 0 // no matchnig card -> can play anything
	}
	card := m.DrawCard(cardIndex)
	h.Logf("playing card '%s'...", card)
	err := h.PlayCard(card)
	if err != nil {
		log.Fatalf("could not play card: %v", err)
	}
}

func (h *BotHandler) OnMemberEvent(_ *api.MemberEvent) {
	// pass
}

func (h *BotHandler) OnDeclaration(_ *api.Declaration) {
	// pass
}

func (h *BotHandler) OnPlayedCard(_ *api.PlayedCard) {
	// pass
}

func (h *BotHandler) OnMatchEnd(_ *api.EndOfGame) {
	// pass
}

func (h *BotHandler) OnTableStateReceived(_ *api.TableState) {
	// pass
}
