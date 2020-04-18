package modes

import (
	"github.com/supermihi/doppelgopf/pkg/game/auction"
	"github.com/supermihi/doppelgopf/pkg/game/core"
)

type Hochzeit struct {
	Announcer core.Player
	Partner   core.Player
}

func NewHochzeit(announcer core.Player) Hochzeit {
	return Hochzeit{announcer, core.NoPlayer}
}

func (h Hochzeit) PartnerFound() bool {
	return h.Partner != core.NoPlayer
}

// trick-taking rules equal those of normal mode
func (h Hochzeit) GameSuit(card core.Card) core.GameSuit {
	return core.NormalGameSuit(card)
}

type VorbehaltHochzeit struct {
}

func (v VorbehaltHochzeit) Identifier() auction.ModeId {
	return "HOCHZEIT"
}

func (v VorbehaltHochzeit) CanAnnounceWith(handCards core.Hand) bool {
	return handCards.NumAlte() == 2
}

func (v VorbehaltHochzeit) Priority() int {
	return 1
}

func (v VorbehaltHochzeit) CreateMode(announcer core.Player) core.Mode {
	return NewHochzeit(announcer)
}
