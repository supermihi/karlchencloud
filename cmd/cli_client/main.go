package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/client"
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
	"github.com/supermihi/karlchencloud/server"
	"log"
	"os"
	"strconv"
)

const address = "localhost:50501"

func main() {
	conn := client.ConnectData{
		DisplayName:    "client",
		ExistingUserId: nil,
		ExistingSecret: nil,
		Address:        address,
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	service, err := client.GetClientService(conn, ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer service.CloseConnection()

	cliHandler := NewCliHandler()
	cliHandler.Start(service)
	<-ctx.Done()
}

type CliHandler struct {
	client.TableClient
}

func NewCliHandler() CliHandler {
	return CliHandler{}
}

func (h *CliHandler) Start(service client.ClientService) {
	log.Printf("Choose: [_c_reate, _j_oin] table")
	action := UserInputRune()
	table := ""
	if action == 'c' {
		tableData, err := service.Api.CreateTable(service.Context, &api.Empty{})
		if err != nil {
			log.Fatalf("%s could not create table: %v", service.Name, err)
		}
		table = tableData.TableId
		service.Logf("table %s created with invite code %s", table, tableData.InviteCode)
	} else if action == 'j' {
		log.Printf("Input invite code:")
		invite := UserInputString()
		tableData, err := service.Api.JoinTable(service.Context, &api.JoinTableRequest{InviteCode: invite})
		if err != nil {
			service.Logf("could not join table: %v", err)
			return
		}
		table = tableData.Data.TableId
		service.Logf("table %s joined", table)
	} else {
		service.Logf("Invalid action %c", action)
		return
	}
	h.TableClient = client.NewTableClient(service, table, h)
	go h.TableClient.Start()
}

func (h *CliHandler) OnTableStateReceived(_ *api.TableState) {
	h.Logf("Table state received. Let the games begin!")
}

func (h *CliHandler) OnMemberEvent(ev *api.MemberEvent) {
	switch ev.Type {
	case api.MemberEventType_JOIN_TABLE:
		if ev.UserId == h.Service.UserId() {
			h.Logf("oh, I joined myself")
			return
		}
		h.Logf("user %s joined table", ev.Name)

		if len(h.View.MemberNamesById) >= 4 {
			matchState, err := h.Api().StartTable(h.Service.Context, &api.TableId{Value: h.TableId})
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

func (h *CliHandler) OnMatchStart(state *api.MatchState) {
	h.Logf("Game starts! Other players: %v", h.View.PlayerNames())
	h.Logf("Forehand: %s", h.View.MemberNamesById[state.Turn.UserId])
	h.Logf("my cards: %s", h.Match().Cards)
}

func (h *CliHandler) OnPlayedCard(ev *api.PlayedCard) {
	if ev.UserId != h.Service.UserId() {
		h.Logf("%v played %v", h.View.MemberNamesById[ev.UserId], server.ToCard(ev.Card))
	}
	if len(h.Match().Trick.Cards) == 0 {
		h.Logf("trick finished. Winner: %s", h.View.MemberNamesById[h.Match().Trick.Forehand])
	}
}

func (h *CliHandler) OnDeclaration(_ *api.Declaration) {
	if h.Match().Phase == match.InGame {
		h.Logf("now in game! Forehand: %s", h.View.MemberNamesById[h.Match().Trick.Forehand])
	}
}

func (h *CliHandler) OnMyTurn() {
	switch h.Match().Phase {
	case match.InAuction:
		h.declare()
	case match.InGame:
		h.playCard()
	default:
		panic(fmt.Sprintf("should not be here: handleMyTurn in neither auction nor game"))
	}

}
func (h *CliHandler) declare() {
	log.Printf("Choose: [_g_esund, _h_ochzeit]")
	char := UserInputRune()
	declaration := game.NormalGameType
	if char == 'h' {
		declaration = game.MarriageType
	}
	declareErr := h.Declare(declaration)
	if declareErr != nil {
		log.Fatalf("error declaring game: %v", declareErr)
	}
	h.Logf("successfully declared %s", declaration)
}

func (h *CliHandler) playCard() {

	log.Printf("your cards: %v", h.Match().Cards)
	log.Printf("Choose index to play: ")
	for {
		line := UserInputString()
		i, err := strconv.Atoi(line)
		if err != nil {
			log.Printf("could not read answer: %v. Please try again", err)
			continue
		}
		if i < 0 || i > len(h.Match().Cards)-1 {
			log.Printf("invalid card index %d", i)
			continue
		}
		log.Printf("playing card: %v", h.Match().Cards[i])
		err = h.PlayCard(h.Match().Cards[i])
		if err != nil {
			log.Printf("could not play card: %v. Try again", err)
			continue
		}
		h.Match().DrawCard(i)
		break
	}

}

func UserInputRune() rune {
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		log.Fatalf("error reading rune: %v", err)
	}
	return char
}

func UserInputString() string {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("error reading string: %v", err)
	}
	return line[:len(line)-1]
}
