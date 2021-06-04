package implementations

import (
	"fmt"
	pb "github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/client"
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
	"log"
	"os"
	"strings"
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

func (h *BotClient) OnWelcome(us *pb.UserState) {
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
			fmt.Printf("Open table %s\n", formatTable(table))
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

func formatTable(table client.TableInfo) string {
	memberNames := make([]string, len(table.MemberNamesById))
	i := 0
	for _, name := range table.MemberNamesById {
		memberNames[i] = name
		i += 1
	}
	return fmt.Sprintf("%s [%s]", table.Id, strings.Join(memberNames, ", "))
}

func (h *BotClient) OnMatchStart() {
	// pass
}

func (h *BotClient) OnNewTable(table client.TableInfo) {
	if h.Table() == nil && table.Public {
		h.Logf("joining new public table %s", formatTable(table))
		err := h.JoinTable("", table.Id)
		if err != nil {
			h.Fatalf("error joining new table: %v", err)
		}
	}
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

func (h *BotClient) OnPlayedCard(_ *pb.PlayedCard) {
	if h.Match().Phase == match.MatchFinished && h.isOwner {
		err := h.StartNextMatch()
		if err != nil {
			h.Logf("Failed to start next match!")
		}
	}
}
