package client

import (
	"github.com/ilyakaznacheev/cleanenv"
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
