package main

import (
	"flag"
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
	client.StartBots(address, numBots, table, inviteCode)
}
