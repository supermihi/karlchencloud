package main

import (
	"bufio"
	"context"
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/client"
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
	"io/ioutil"
	"log"
	"os"
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
	c, err := client.GetConnectedService(conn, ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	table, err := c.Kc.CreateTable(ctx, &api.Empty{})
	if err != nil {
		log.Fatalf("%s could not create table: %v", conn.DisplayName, err)
	}
	log.Printf("table %s created with invite code %s", table.TableId, table.InviteCode)
	err = ioutil.WriteFile("table.config", []byte(table.TableId+":"+table.InviteCode), 0644)
	if err != nil {
		log.Fatalf("error wrtiing table.config: %v", err)
	}
	log.Printf("wrote table.config")

	cliClient := NewCliClient(ctx, *c, table.TableId)
	cliClient.Start(table.TableId)
	<-cliClient.ctx.Done()
}

type CliClient struct {
	ctx             context.Context
	client          client.Client
	stream          api.Karlchencloud_SubscribeMatchEventsClient
	numOtherPlayers int
	cards           game.Hand
	players         *api.Players
	phase           match.Phase
	tableId         string
}

func NewCliClient(ctx context.Context, client client.Client, tableId string) CliClient {
	return CliClient{ctx: ctx, client: client, numOtherPlayers: 0, phase: match.MatchFinished,
		tableId: tableId}
}

func (c *CliClient) UserId() string {
	return c.client.Creds.UserId()
}

func (c *CliClient) Start(tableId string) {
	stream, err := c.client.Kc.SubscribeMatchEvents(c.ctx, &api.TableId{Value: tableId})
	log.Printf("listening for match events ...")
	if err != nil {
		log.Printf("error subscribing: %v", err)
		return
	}
	c.stream = stream
	go func() {
		for {
			msg, err := stream.Recv()
			if err != nil {
				log.Fatalf("error receiving from stream: %v", err)
			}
			c.HandleEvent(msg)
		}
	}()
}
func (c *CliClient) HandleEvent(msg *api.MatchEventStream) {
	switch ev := msg.Event.(type) {
	case *api.MatchEventStream_Member:
		c.HandleMemberEvent(ev.Member)
	case *api.MatchEventStream_Start:
		c.HandleMatchStart(ev.Start)
	default:
		log.Fatalf("unimplemented event occured: %v", msg)
	}
}

func (c *CliClient) HandleMemberEvent(ev *api.MemberEvent) {
	switch ev.Type {
	case api.MemberEventType_JOIN_TABLE:
		if ev.UserId == c.UserId() {
			log.Printf("oh, I joined myself also")
			return
		}
		log.Printf("user %s joined table", ev.UserId)
		c.numOtherPlayers += 1
		if c.numOtherPlayers >= 3 {
			_, err := c.client.Kc.StartTable(c.ctx, &api.TableId{Value: c.tableId})
			if err != nil {
				log.Fatalf("error starting table: %v", err)
			}
		}
	case api.MemberEventType_GO_ONLINE:
		log.Printf("user %s is now online", ev.UserId)
	default:
		log.Printf("unexpected MemberEvent: %v", ev)
	}
}

func (c *CliClient) HandleMatchStart(state *api.MatchState) {
	switch r := state.Role.(type) {
	case *api.MatchState_Spectator:
		log.Fatalf("unexpected role spectator")
	case *api.MatchState_OwnCards:
		c.cards = client.ToHand(r.OwnCards.Cards)
		log.Printf("own cards received: %s", c.cards.String())
	}
	c.players = state.Players
	c.phase = match.InAuction
	if state.Turn.UserId == c.UserId() {
		c.Declare()
	}
}

func (c *CliClient) Declare() {
	log.Printf("looks like I'm the first player. Say something! [_g_esund, _h_ochzeit]")
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		log.Fatalf("error reading rune: %v", err)
	}
	declaration := api.GameType_NORMAL_GAME
	if char == 'h' {
		declaration = api.GameType_MARRIAGE
	}
	_, err = c.client.Kc.Play(c.ctx, &api.PlayRequest{Table: c.tableId,
		Request: &api.PlayRequest_Declaration{Declaration: declaration}})
	if err != nil {
		log.Fatalf("error declaring game: %v", err)
	}
	log.Printf("sucessfully declared %s", declaration)
}
