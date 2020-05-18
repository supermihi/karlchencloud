package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/client"
	"io/ioutil"
	"log"
	"strings"
)

const (
	address = "localhost:50051"
)

func main() {
	var table, inviteCode string
	flag.StringVar(&table, "table", "", "table id")
	flag.StringVar(&inviteCode, "code", "", "table invite code")
	var numBots int
	flag.IntVar(&numBots, "num", 3, "number of bots to join")
	flag.Parse()
	tc, err := ioutil.ReadFile("table.config")
	if err != nil {
		log.Fatal(err)
	}
	tableIdAndInvite := strings.Split(string(tc), ":")
	if len(tableIdAndInvite) != 2 {
		log.Fatalf("unexpected format %v in table.config", string(tc))
	}
	table = tableIdAndInvite[0]
	inviteCode = tableIdAndInvite[1]
	clients := make([]*GoBotClient, numBots)
	for i := 0; i < numBots; i++ {
		connect := client.ConnectData{
			DisplayName:    fmt.Sprintf("Bot %v", i+1),
			ExistingUserId: nil,
			ExistingSecret: nil,
			Address:        address,
		}
		clients[i] = NewGoBotClient(table, inviteCode, connect)
		go clients[i].Run()
	}
	for i := 0; i < numBots; i++ {
		<-clients[i].context.Done()
	}
	log.Printf("all bots finished")

}

type GoBotClient struct {
	table       string
	inviteCode  string
	connectData client.ConnectData
	context     context.Context
	cancel      context.CancelFunc
}

func NewGoBotClient(table string, inviteCode string, connect client.ConnectData) *GoBotClient {
	ctx, cancel := context.WithCancel(context.Background())
	return &GoBotClient{table, inviteCode, connect, ctx, cancel}
}

func (gbc *GoBotClient) Run() {
	conn := gbc.connectData
	ctx := gbc.context
	clientObj, err := client.GetConnectedService(conn, ctx)
	if err != nil {
		log.Print(err)
		return
	}
	c := clientObj.Kc
	defer gbc.cancel()
	defer clientObj.Close()
	_, err = c.JoinTable(ctx, &api.JoinTableRequest{InviteCode: gbc.inviteCode, TableId: gbc.table})
	if err != nil {
		log.Printf("%s could not join table: %v", gbc.connectData.DisplayName, err)
		return
	}
	log.Printf("%s joined", gbc.connectData.DisplayName)
	serv, err := c.SubscribeMatchEvents(ctx, &api.TableId{Value: gbc.table})
	if err != nil {
		log.Printf("%s could not subscribe match events: %v", gbc.connectData.DisplayName, err)
		return
	}
	for {
		msg, err := serv.Recv()
		if err != nil {
			log.Printf("%s could not receive event: %v", gbc.connectData.DisplayName, err)
			return
		}
		log.Printf("%s incoming message: %s", gbc.connectData.DisplayName, client.MatchEventString(msg))

	}
}
