package server

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type ServerConfig struct {
	Room struct {
		ConstantTableId    string `yaml:"tableId" env:"CONSTANT_TABLE_ID" env-default:""`
		ConstantInviteCode string `yaml:"inviteCode" env:"CONSTANT_INVITE_CODE" env-default:""`
		Seed               int64  `yaml:"seed" env:"KC_DBG_SEED" env-default:""`
	} `yaml:"room"`
	NoProxy bool `yaml:"noProxy" env:"NO_PROXY" env-default:"0"`
}

func ReadConfig() (cfg ServerConfig, err error) {
	err = cleanenv.ReadEnv(&cfg)
	return
}
