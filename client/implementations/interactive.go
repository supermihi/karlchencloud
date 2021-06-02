package implementations

import (
	"bufio"
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/client"
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/server/pbconv"
	"log"
	"os"
	"strconv"
)

type CliHandler struct {
	IsCreator bool
}

func (h *CliHandler) OnConnect(_ client.ClientApi) {
	// pass
}

func (h *CliHandler) OnWelcome(client client.ClientApi, us *api.UserState) {
	if us.CurrentTable != nil {
		return
	}
	for {
		log.Printf("Choose: [_c_reate, _j_oin] table")
		action := UserInputRune()
		if action == 'c' {
			h.IsCreator = true
			err := client.CreateTable(true)
			if err == nil {
				return
			}
			client.Logf("%s could not create table: %v", err)
		} else if action == 'j' {
			h.IsCreator = false
			log.Printf("Input invite code:")
			invite := UserInputString()
			err := client.JoinTable(invite, "")
			if err == nil {
				return
			}
			client.Logf("could not join table: %v", err)
		} else {
			client.Logf("Invalid action %c", action)
		}
	}
}

func (h *CliHandler) OnMemberJoin(client client.ClientApi, id string, name string) {
	if len(client.Table().MemberNamesById) >= 4 && h.IsCreator {
		err := client.StartTable()
		if err != nil {
			Fatalf(client, "error starting table: %v", err)
		}
	}
}

func (h *CliHandler) OnMatchStart(client client.ClientApi) {
	client.Logf("Game starts! Other players: %v", client.Table().PlayerNames())
	client.Logf("Forehand: %s", client.Table().MemberNamesById[client.Match().Trick.Forehand])
	client.Logf("my cards: %s", client.Match().Cards)
}

func (h *CliHandler) OnPlayedCard(client client.ClientApi, ev *api.PlayedCard) {
	if ev.UserId != client.User().Id {
		client.Logf("%v played %v", client.Table().MemberNamesById[ev.UserId], pbconv.ToCard(ev.Card))
	}
	if len(client.Match().Trick.Cards) == 0 {
		client.Logf("trick finished. Winner: %s", client.Table().MemberNamesById[client.Match().Trick.Forehand])
	}
}

func (h *CliHandler) OnMyTurnAuction(client client.ClientApi) {
	log.Printf("Choose: [_g_esund, _h_ochzeit]")
	char := UserInputRune()
	declaration := game.NormalGameType
	if char == 'h' {
		declaration = game.MarriageType
	}
	declareErr := client.Declare(declaration)
	if declareErr != nil {
		log.Fatalf("error declaring game: %v", declareErr)
	}
}

func (h *CliHandler) OnMyTurnGame(client client.ClientApi) {
	log.Printf("your cards: %v", client.Match().Cards)
	log.Printf("Choose index to play: ")
	for {
		line := UserInputString()
		i, err := strconv.Atoi(line)
		if err != nil {
			log.Printf("could not read answer: %v. Please try again", err)
			continue
		}
		if i < 0 || i > len(client.Match().Cards)-1 {
			log.Printf("invalid card index %d", i)
			continue
		}
		err = client.PlayCard(i)
		if err != nil {
			log.Printf("could not play card: %v. Try again", err)
			continue
		}
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
