package main

import (
	"github.com/supermihi/karlchencloud/client"
	"log"
)

const (
	address = "localhost:50501"
)

func main() {
	config, err := client.ReadConfig()
	if err != nil {
		log.Fatalf("Error reading client config: %v", err)
	}
	client.StartBots(address, config.NumberOfBots, config.TableId, config.InviteCode)
}
