package server

import (
	"github.com/ilyakaznacheev/cleanenv"
	"math/rand"
)

type TablesConfig struct {
	InputSeed int64 `yaml:"seed" env:"KC_DBG_SEED" env-default:"0"`
}

type ServerConfig struct {
	Tables  TablesConfig `yaml:"tables"`
	NoProxy bool         `yaml:"noProxy" env:"NO_PROXY" env-default:"0"`
}

func ReadConfig() (cfg ServerConfig, err error) {
	err = cleanenv.ReadEnv(&cfg)
	return
}

func (cfg *TablesConfig) Seed() int64 {
	if cfg.InputSeed == 0 {
		return rand.Int63()
	}
	return cfg.InputSeed
}
