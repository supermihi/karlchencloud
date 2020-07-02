package main

import (
	"flag"
	"github.com/supermihi/karlchencloud/client"
)

const (
	address = "localhost:50501"
)

func main() {
	var table, inviteCode string
	flag.StringVar(&table, "table", "", "table id")
	flag.StringVar(&inviteCode, "code", "", "table invite code")
	var numBots int
	flag.IntVar(&numBots, "num", 3, "number of bots to join")
	flag.Parse()
	client.StartBots(address, numBots, table, inviteCode)
}
