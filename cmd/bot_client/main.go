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
	botConfig, err := client.ReadBotConfig("bots.json")
	var bots []client.BotLogin
	if err == nil {
		bots = botConfig.Bots
	}
	client.StartBots(address, config.NumberOfBots, config.TableId, config.InviteCode, bots)
}
