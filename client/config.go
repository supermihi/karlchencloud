package client

import (
	"encoding/json"
	"github.com/ilyakaznacheev/cleanenv"
	"io/ioutil"
	"os"
)

type ClientConfig struct {
	TableId      string `env:"TABLE_ID" env-default:""`
	InviteCode   string `env:"INVITE_CODE" env-default:""`
	NumberOfBots int    `env:"NUM_BOTS" env-default:"3"`
}

func ReadConfig() (cfg ClientConfig, err error) {
	err = cleanenv.ReadEnv(&cfg)
	return
}


type BotLogin struct {
	Id string
	Name string
	Secret string
}

type BotConfig struct {
	Bots []BotLogin
}


func ReadBotConfig(filename string) (*BotConfig, error) {
	if _, statErr := os.Stat(filename); statErr != nil {
		return nil, statErr
	}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var config BotConfig
	err = json.Unmarshal(data, &config)
	return &config, nil
}