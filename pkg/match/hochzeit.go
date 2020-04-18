package match

import (
	"github.com/supermihi/doppelgopf/pkg/core"
	"github.com/supermihi/doppelgopf/pkg/modes"
)

type VorbehaltHochzeit struct {
}

func (v VorbehaltHochzeit) IsValid(handCards core.Hand) bool {
	return handCards.NumAlte() == 2
}

func (v VorbehaltHochzeit) Priority() int {
	return 1
}

func (v VorbehaltHochzeit) CreateMode(announcer core.Player) core.Mode {
	return modes.NewHochzeit(announcer)
}
