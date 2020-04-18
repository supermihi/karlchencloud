package modes

import "github.com/supermihi/doppelgopf/pkg/core"

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
