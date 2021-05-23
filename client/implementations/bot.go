package implementations

import (
	"context"
	"fmt"
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/client"
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
	"log"
	"math/rand"
	"os"
	"time"
)

const superSecretBotPassword = "123"

func CreateBotLogin(num int, address string) client.LoginData {
	name := fmt.Sprintf("bot_%v", num+1)
	return client.LoginData{
		Name:               name,
		Email:              fmt.Sprintf("%s@example.com", name),
		Password:           superSecretBotPassword,
		RegisterOnAuthFail: true,
		ServerAddress:      address,
	}
}

func StartBots(address string, numBots int, inviteCode string, initTable bool, logins []client.LoginData) {
	clients := make([]*client.ClientImplementation, numBots)
	ctx := context.Background()
	for i := 0; i < numBots; i++ {
		var login client.LoginData
		if len(logins) > i {
			login = logins[i]
		} else {
			login = CreateBotLogin(i+1, address)
		}
		handler := NewBotHandler(i == 0 && initTable, inviteCode)
		log.Printf("starting bot %d", i)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		kc := client.NewClientImplementation(login, handler)
		clients[i] = &kc
		go kc.Start(ctx)
	}
	<-ctx.Done()
	// TODO use nontrivial contexts
	log.Printf("all bots finished")

}

type BotHandler struct {
	isOwner bool
	invite  string
}

func NewBotHandler(isOwner bool, invite string) *BotHandler {
	return &BotHandler{
		isOwner: isOwner,
		invite:  invite}
}

func (h *BotHandler) OnConnect(_ client.ClientApi) {
	// pass
}

func Fatalf(client client.ClientApi, format string, v ...interface{}) {
	client.Logf(format, v...)
	os.Exit(1)
}

func (h *BotHandler) OnWelcome(client client.ClientApi, us *api.UserState) {
	if us.CurrentTable != nil {
		client.Logf("already at table. Continuing ...")
	}
	if h.isOwner {
		err := client.CreateTable()
		if err != nil {
			Fatalf(client, "%s could not create table: %v", err)
		}
	} else {
		err := client.JoinTable(h.invite)
		if err != nil {
			Fatalf(client, "could not join table: %v", err)
			return
		}
	}
}
func (h *BotHandler) OnMatchStart(_ client.ClientApi) {
	// pass
}

func (h *BotHandler) OnMyTurnAuction(client client.ClientApi) {
	declaration := game.NormalGameType
	if client.Match().Cards.NumQueensOfClubs() == 2 {
		declaration = game.MarriageType
	}
	err := client.Declare(declaration)
	if err != nil {
		Fatalf(client, "could not make auction turn: %v", err)
	}
}

func (h *BotHandler) OnMyTurnGame(client client.ClientApi) {
	cardIndex := -1
	m := client.Match()
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
	err := client.PlayCard(cardIndex)
	if err != nil {
		Fatalf(client, "could not play card: %v", err)
	}
}

func (h *BotHandler) OnMemberJoin(client client.ClientApi, _ string, _ string) {
	if len(client.Table().MemberNamesById) >= 4 && h.isOwner {
		err := client.StartTable()
		if err != nil {
			log.Fatalf("error starting table: %v", err)
		}
	}
}

func (h *BotHandler) OnPlayedCard(client client.ClientApi, _ *api.PlayedCard) {
	if client.Match().Phase == match.MatchFinished && h.isOwner {
		err := client.StartNextMatch()
		if err != nil {
			client.Logf("Failed to start next match!")
		}
	}
}
