package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/client"
	"github.com/supermihi/karlchencloud/doko/match"
	"log"
	"os"
	"strconv"
)

const address = "localhost:50051"

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

	cliHandler := NewCliHandler(service, ctx)
	cliHandler.Start()
	<-ctx.Done()
}

type CliHandler struct {
	numOtherPlayers int
	view            client.MatchView
	service         client.ClientService
	ctx             context.Context
	tableClient     client.TableClient
}

func NewCliHandler(s client.ClientService, ctx context.Context) CliHandler {
	return CliHandler{numOtherPlayers: 0, view: client.NewMatchView(""), service: s,
		ctx: ctx}
}

func (h *CliHandler) Start() {
	table, err := h.service.Api.CreateTable(h.ctx, &api.Empty{})
	if err != nil {
		log.Fatalf("%s could not create table: %v", h.service.Name, err)
	}
	h.view.TableId = table.TableId
	h.service.Logf("table %s created with invite code %s", h.view.TableId, table.InviteCode)
	/*err = ioutil.WriteFile("table.config", []byte(table.TableId+":"+table.InviteCode), 0644)
	if err != nil {
		log.Fatalf("error wrtiing table.config: %v", err)
	}
	log.Printf("wrote table.config")
	*/
	h.tableClient = client.NewTableClient(h.ctx, h.service, table.TableId, h)
	go h.tableClient.Start()
	go client.StartBots(address, 3, table.TableId, table.InviteCode)
}

func (h *CliHandler) HandleMemberEvent(ev *api.MemberEvent) {
	switch ev.Type {
	case api.MemberEventType_JOIN_TABLE:
		h.view.AddName(ev.UserId, ev.Name)
		if ev.UserId == h.service.UserId() {
			h.service.Logf("oh, I joined myself also")
			return
		}
		h.service.Logf("user %s joined table", ev.Name)
		h.numOtherPlayers += 1
		if h.numOtherPlayers >= 3 {
			h.service.Logf("all players there. Starting game ...")
			go func() {
				_, err := h.service.Api.StartTable(h.ctx, &api.TableId{Value: h.view.TableId})
				if err != nil {
					log.Fatalf("error starting table: %v", err)
				}
			}()
		}
	case api.MemberEventType_GO_ONLINE:
		h.service.Logf("user %s is now online", ev.UserId)
	case api.MemberEventType_GO_OFFLINE:
		h.service.Logf("user %s is now offline", ev.UserId)
	default:
		h.service.Logf("unexpected MemberEvent: %v", ev)
	}
}

func (h *CliHandler) HandleMatchStart(state *api.MatchState) {
	h.view.InitFromMatchState(state)
	h.service.Logf("Game starts! Other players: %v", h.view.PlayerNames())
	h.service.Logf("Forehand: %s", h.view.Names[state.Turn.UserId])
	h.service.Logf("my cards: %s", h.view.Cards)
	h.checkMyTurn()
}

func (h *CliHandler) HandlePlayedCard(ev *api.PlayedCard) {
	h.view.UpdateTrick(ev)
	h.checkMyTurn()
}

func (h *CliHandler) HandleEnd(ev *api.EndOfGame) {
	panic("implement me")
}

func (h *CliHandler) HandleDeclared(d *api.Declaration) {
	h.view.UpdateOnDeclare(d)
	h.checkMyTurn()
}

func (h *CliHandler) checkMyTurn() {
	if h.view.MyTurn {
		h.handleMyTurn()
	}
}
func (h *CliHandler) handleMyTurn() {
	switch h.view.Phase {
	case match.InAuction:
		go h.declare()
	case match.InGame:
		go h.playCard()
	default:
		panic(fmt.Sprintf("should not be here: handleMyTurn in neither auction nor game"))
	}

}
func (h *CliHandler) declare() {
	log.Printf("Choose: [_g_esund, _h_ochzeit]")
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		log.Fatalf("error reading rune: %v", err)
	}
	declaration := api.GameType_NORMAL_GAME
	if char == 'h' {
		declaration = api.GameType_MARRIAGE
	}
	_, err = h.service.Api.Play(h.ctx, &api.PlayRequest{Table: h.view.TableId,
		Request: &api.PlayRequest_Declaration{Declaration: declaration}})
	if err != nil {
		log.Fatalf("error declaring game: %v", err)
	}
	log.Printf("sucessfully declared %s", declaration)
}

func (h *CliHandler) playCard() {

	log.Printf("your cards: %v", h.view.Cards)
	log.Printf("Choose index to play: ")
	i := -1
	for {
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("error reading ans: %v", err)
		}
		i, err = strconv.Atoi(line[:len(line)-1])
		if err != nil {
			log.Printf("could not read answer: %v. Please try again", err)
		}
		break
	}
	err := h.service.Play(h.view.DrawCard(i), h.view.TableId)
	if err != nil {
		log.Fatalf("could not play card: %v", err)
	}
}
