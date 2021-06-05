package implementations

import (
	"bufio"
	"github.com/eiannone/keyboard"
	pb "github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/client"
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/server/pbconv"
	"log"
	"os"
	"strconv"
)

type CliHandler struct {
	IsCreator bool
	*client.Client
}

func NewCliHandler(login client.LoginData, isCreator bool) CliHandler {
	cli := CliHandler{IsCreator: isCreator}
	embeddedClient := client.NewClient(login, &cli)
	cli.Client = &embeddedClient
	return cli
}

func (h *CliHandler) OnConnect() {
}

func (h *CliHandler) OnNewTable(_ client.TableInfo) {
}

func (h *CliHandler) OnWelcome(us *pb.UserState) {
	if us.CurrentTable != nil {
		return
	}
	for {
		log.Printf("Choose: [_c_reate, _j_oin] table")
		action := UserInputRune()
		if action == 'c' {
			h.IsCreator = true
			err := h.CreateTable(true)
			if err == nil {
				return
			}
			h.Logf("%s could not create table: %v", err)
		} else if action == 'j' {
			h.IsCreator = false
			log.Printf("Input invite code:")
			invite := UserInputString()
			err := h.JoinTable(invite, "")
			if err == nil {
				return
			}
			h.Logf("could not join table: %v", err)
		} else {
			h.Logf("Invalid action %c", action)
		}
	}
}
func (h *CliHandler) Fatalf(format string, v ...interface{}) {
	h.Logf(format, v...)
	os.Exit(1)
}
func (h *CliHandler) OnMemberJoin(_ string, _ string) {
	if len(h.Table().MemberNamesById) >= 4 && h.IsCreator {
		err := h.StartTable()
		if err != nil {
			h.Fatalf("error starting table: %v", err)
		}
	}
}

func (h *CliHandler) OnMatchStart() {
	h.Logf("Game starts! Other players: %v", h.Table().PlayerNames())
	h.Logf("my cards: %s", h.Match().Cards)
}

func (h *CliHandler) OnPlayedCard(ev *pb.PlayedCard) {
	if ev.UserId != h.User().Id {
		h.Logf("%v played %v", h.Table().MemberNamesById[ev.UserId], pbconv.ToCard(ev.Card))
	}
	if len(h.Match().Trick.Cards) == 0 {
		h.Logf("trick finished. Winner: %s", h.Table().MemberNamesById[h.Match().Trick.Forehand])
	}
}

func (h *CliHandler) OnMyTurnAuction() {
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
}

func (h *CliHandler) OnMyTurnGame() {
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
		err = h.PlayCard(i)
		if err != nil {
			log.Printf("could not play card: %v. Try again", err)
			continue
		}
		break
	}
}

func UserInputRune() rune {
	char, _, err := keyboard.GetSingleKey()
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
