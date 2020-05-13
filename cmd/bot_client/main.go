package main

import (
	"flag"
	"fmt"
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/client"
	"log"
	"time"
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
	for i := 0; i < numBots; i++ {
		go RunBot(table, inviteCode, fmt.Sprintf("RunBot %v", i+1))
	}

}

func RunBot(table string, inviteCode string, name string) {
	clientObj := client.GetConnectedService(address, "bot", 10*time.Second)
	ctx := clientObj.Ctx
	c := clientObj.Kc
	defer clientObj.Cancel()
	defer func() { _ = clientObj.Connection.Close() }()
	_, err := c.JoinTable(ctx, &api.JoinTableRequest{InviteCode: inviteCode, TableId: table})
	if err != nil {
		log.Fatalf("%s could not join table: %v", name, err)
	}
	log.Printf("%s joined", name)
}
