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

func StartBots(address string, numBots int, inviteCode string, initTable bool, logins []LoginData) {
	clients := make([]*KarlchenClient, numBots)
	ctx := context.Background()
	for i := 0; i < numBots; i++ {
		var login LoginData
		if len(logins) > i {
			login = logins[i]
		} else {
			login.RegisterIfEmptyUserId = true
			login.ServerAddress = address
			login.Name = fmt.Sprintf("Bot %v", i+1)
			login.Password = "123"
		}
		handler := NewBotHandler(i == 0 && initTable, inviteCode)
		log.Printf("starting bot %d", i)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		client := NewKarlchenClient(login, handler)
		clients[i] = &client
		go client.Start(ctx)
	}
	<-ctx.Done()
	// TODO use nontrivial contexts
	log.Printf("all bots finished")

}

type BotHandler struct {
	KarlchenClient
	context context.Context
	cancel  context.CancelFunc
	service *ClientService
	isOwner bool
	invite  string
}

func NewBotHandler(isOwner bool, invite string) *BotHandler {
	ctx, cancel := context.WithCancel(context.Background())
	return &BotHandler{
		context: ctx,
		cancel:  cancel,
		isOwner: isOwner,
		invite:  invite}
}

func (h *BotHandler) OnConnect(service *ClientService) {
	h.service = service
}
func (h *BotHandler) OnWelcome(client *KarlchenClient, us *api.UserState) {
	if us.CurrentTable != nil {
		log.Printf("already at table. Continuing ...")
	}
	if h.isOwner {
		tableData, err := h.service.Grpc.CreateTable(h.service.Context, &api.Empty{})
		if err != nil {
			log.Fatalf("%s could not create table: %v", h.service.Name, err)
		}
		h.service.Logf("table %s created with invite code %s", tableData.TableId, tableData.InviteCode)
		client.TableId = tableData.TableId
	} else {
		_, err := h.service.Grpc.JoinTable(h.Service.Context, &api.JoinTableRequest{InviteCode: h.invite})
		if err != nil {
			h.service.Logf("could not join table: %v", err)
			return
		}
		h.service.Logf("joined")
	}
}
func (h *BotHandler) OnMatchStart(_ *api.MatchState) {
	// pass
}

func (h *BotHandler) OnMyTurn() {
	switch h.Match().Phase {
	case match.InAuction:
		h.makeTurnAuction()
	case match.InGame:
		h.makeTurnGame()
	}
}

func (h *BotHandler) makeTurnAuction() {
	declaration := game.NormalGameType
	if h.Match().Cards.NumQueensOfClubs() == 2 {
		declaration = game.MarriageType
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
	h.Logf("playing card: %v", card)
	err := h.PlayCard(card)
	if err != nil {
		log.Fatalf("could not play card: %v", err)
	}
}

func (h *BotHandler) OnMemberEvent(ev *api.MemberEvent) {
	switch ev.Type {
	case api.MemberEventType_JOIN_TABLE:
		if ev.UserId == h.Service.UserId() {
			h.Logf("oh, I joined myself")
			return
		}
		h.Logf("user %s joined table", ev.Name)

		if len(h.View.MemberNamesById) >= 4 && h.isOwner {
			matchState, err := h.Api().StartTable(h.Service.Context, &api.StartTableRequest{TableId: h.TableId})
			if err != nil {
				log.Fatalf("error starting table: %v", err)
			}
			h.HandleStart(matchState)
		}
	case api.MemberEventType_GO_ONLINE:
		h.Logf("user %s is now online", h.View.MemberNamesById[ev.UserId])
	case api.MemberEventType_GO_OFFLINE:
		h.Logf("user %s is now offline", h.View.MemberNamesById[ev.UserId])
	default:
		h.Logf("unexpected MemberEvent: %v", ev)
	}
}

func (h *BotHandler) OnDeclaration(_ *api.Declaration) {
	// pass
}

func (h *BotHandler) OnPlayedCard(play *api.PlayedCard) {
	if play.Winner != nil {
		h.Match().Phase = match.MatchFinished
	}
	if h.Match().Phase == match.MatchFinished && h.isOwner {
		state, err := h.Service.Grpc.StartNextMatch(h.Service.Context, &api.StartNextMatchRequest{TableId: h.TableId})
		if err != nil {
			h.Logf("Failed to start next match!")
		}
		h.HandleStart(state)
	}
}

func (h *BotHandler) OnTableStateReceived(_ *api.TableState) {
	// pass
}
