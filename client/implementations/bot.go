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
	"strings"
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
	clients := make([]*BotClient, numBots)
	ctx := context.Background()
	for i := 0; i < numBots; i++ {
		var login client.LoginData
		if len(logins) > i {
			login = logins[i]
		} else {
			login = CreateBotLogin(i+1, address)
		}
		log.Printf("starting bot %d", i)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		clients[i] = NewBotClient(login, i == 0 && initTable, inviteCode)
		go clients[i].Start(ctx)
	}
	<-ctx.Done()
	// TODO use nontrivial contexts
	log.Printf("all bots finished")

}

type BotClient struct {
	*client.Client
	isOwner bool
	invite  string
}

func NewBotClient(login client.LoginData, isOwner bool, invite string) *BotClient {
	bot := &BotClient{isOwner: isOwner, invite: invite}
	embeddedClient := client.NewClient(login, bot)
	bot.Client = &embeddedClient
	return bot
}

func (h *BotClient) OnConnect() {
	// pass
}

func (h *BotClient) Fatalf(format string, v ...interface{}) {
	h.Logf(format, v...)
	os.Exit(1)
}

func (h *BotClient) OnWelcome(us *api.UserState) {
	if us.CurrentTable != nil {
		h.Logf("already at table. Continuing ...")
		return
	}
	if h.isOwner {
		err := h.CreateTable(true)
		if err != nil {
			h.Fatalf("%s could not create table: %v", err)
		}
	} else if h.invite != "" {
		err := h.JoinTable(h.invite, "")
		if err != nil {
			h.Fatalf("could not join table: %v", err)
			return
		}
	} else {
		tables, err := h.ListOpenTables()
		if err != nil {
			h.Fatalf("could not list open tables: %v", err)
			return
		}
		for _, table := range tables {
			memberNames := make([]string, len(table.MemberNamesById))
			i := 0
			for _, name := range table.MemberNamesById {
				memberNames[i] = name
				i += 1
			}
			fmt.Printf("Seeing table %s [%s]", table.Id, strings.Join(memberNames, ", "))

		}
		if len(tables) > 0 {
			err := h.JoinTable("", tables[0].Id)
			if err != nil {
				h.Fatalf("could not join table: %v", err)
				return
			}
		}

	}
}
func (h *BotClient) OnMatchStart() {
	// pass
}

func (h *BotClient) OnMyTurnAuction() {
	declaration := game.NormalGameType
	if h.Match().Cards.NumQueensOfClubs() == 2 {
		declaration = game.MarriageType
	}
	err := h.Declare(declaration)
	if err != nil {
		h.Fatalf("could not make auction turn: %v", err)
	}
}

func (h *BotClient) OnMyTurnGame() {
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
	err := h.PlayCard(cardIndex)
	if err != nil {
		h.Fatalf("could not play card: %v", err)
	}
}

func (h *BotClient) OnMemberJoin(_ string, _ string) {
	if len(h.Table().MemberNamesById) >= 4 && h.isOwner {
		err := h.StartTable()
		if err != nil {
			log.Fatalf("error starting table: %v", err)
		}
	}
}

func (h *BotClient) OnPlayedCard(_ *api.PlayedCard) {
	if h.Match().Phase == match.MatchFinished && h.isOwner {
		err := h.StartNextMatch()
		if err != nil {
			h.Logf("Failed to start next match!")
		}
	}
}
