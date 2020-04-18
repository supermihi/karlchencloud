package modes

import "github.com/supermihi/doppelgopf/pkg/core"

type Farbsolo struct {
	Trumpf core.Suit
}

func (f Farbsolo) GameSuit(card core.Card) core.GameSuit {
	if card.Suit == f.Trumpf || card.Rank == core.Bube || card.Rank == core.Dame || card == core.Dulle() {
		return core.Trumpf
	}
	return card.Suit.AsFehl()
}
