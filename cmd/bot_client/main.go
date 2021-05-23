package main

import (
	"github.com/supermihi/karlchencloud/client"
	"github.com/supermihi/karlchencloud/client/implementations"
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
	/*botConfig, err := client.ReadBotConfig("bots.json")
	var bots []client.BotLogin
	if err == nil {
		bots = botConfig.Bots
	}*/
	implementations.StartBots(address, config.NumberOfBots, config.InviteCode, config.InitTable, []client.LoginData{})
}
