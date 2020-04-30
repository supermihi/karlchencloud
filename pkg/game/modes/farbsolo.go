package modes

import "github.com/supermihi/karlchencloud/pkg/game/core"

type Farbsolo struct {
	Trumpf  core.Suit
	Soloist core.Player
}

func NewFarbsolo(trump core.Suit, soloist core.Player) Farbsolo {
	return Farbsolo{trump, soloist}
}

func (f Farbsolo) GameSuit(card core.Card) core.GameSuit {
	if card.Suit == f.Trumpf || card.Rank == core.Bube || card.Rank == core.Dame || card == core.Dulle() {
		return core.Trumpf
	}
	return card.Suit.AsFehl()
}

func (f Farbsolo) PartyOf(p core.Player) core.Party {
	if p == f.Soloist {
		return core.ReParty
	}
	return core.ContraParty
}

func (Farbsolo) Klaerungsstich() int {
	return -1
}

func (Farbsolo) OnCompletedTrick(core.Trick, int) {
}
