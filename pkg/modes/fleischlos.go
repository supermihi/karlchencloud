package modes

import "github.com/supermihi/doppelgopf/pkg/core"

type FleischlosMode struct {
}

func (f FleischlosMode) GameSuit(card core.Card) core.GameSuit {
	return card.Suit.AsFehl()
}
