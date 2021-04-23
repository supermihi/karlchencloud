package server

import (
	"github.com/ilyakaznacheev/cleanenv"
	"math/rand"
)

type RoomConfig struct {
	ConstantTableId    string `yaml:"tableId" env:"CONSTANT_TABLE_ID" env-default:""`
	ConstantInviteCode string `yaml:"inviteCode" env:"CONSTANT_INVITE_CODE" env-default:""`
	InputSeed          int64  `yaml:"seed" env:"KC_DBG_SEED" env-default:"0"`
}

type ServerConfig struct {
	Room RoomConfig `yaml:"room"`
	NoProxy bool `yaml:"noProxy" env:"NO_PROXY" env-default:"0"`
}

func ReadConfig() (cfg ServerConfig, err error) {
	err = cleanenv.ReadEnv(&cfg)
	return
}

func (cfg *RoomConfig) Seed() int64 {
	if cfg.InputSeed == 0 {
		return rand.Int63()
	}
	return cfg.InputSeed
}
