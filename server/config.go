package server

import (
	"math/rand"
)

type TablesConfig struct {
	InputSeed int64
}

type ServerConfig struct {
	Tables TablesConfig
}

func (cfg *TablesConfig) Seed() int64 {
	if cfg.InputSeed == 0 {
		return rand.Int63()
	}
	return cfg.InputSeed
}
