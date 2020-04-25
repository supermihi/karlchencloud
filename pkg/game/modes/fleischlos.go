package modes

import "github.com/supermihi/doppelgopf/pkg/game/core"

type FleischlosMode struct {
	Soloist core.Player
}

func (f FleischlosMode) GameSuit(card core.Card) core.GameSuit {
	return card.Suit.AsFehl()
}

func (FleischlosMode) Klaerungsstich() int {
	return -1
}

func (FleischlosMode) OnCompletedTrick(core.Trick, int) {
}

func (f FleischlosMode) PartyOf(p core.Player) core.Party {
	if p == f.Soloist {
		return core.ReParty
	}
	return core.ContraParty
}
