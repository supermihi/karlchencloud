package match

import (
	"github.com/supermihi/karlchencloud/doko/game"
)

type Farbsolo struct {
	Trumpf  game.Suit
	Soloist game.Player
}

func NewFarbsolo(trump game.Suit, soloist game.Player) Farbsolo {
	return Farbsolo{trump, soloist}
}

func FarbsoloSuit(trump game.Suit, card game.Card) game.GameSuit {
	if card.Suit == trump || card.Rank == game.Bube || card.Rank == game.Dame || card == game.Dulle() {
		return game.Trumpf
	}
	return card.Suit.AsFehl()
}

func (f Farbsolo) GameSuit(card game.Card) game.GameSuit {
	return FarbsoloSuit(f.Trumpf, card)
}

func (f Farbsolo) PartyOf(p game.Player) game.Party {
	if p == f.Soloist {
		return game.ReParty
	}
	return game.ContraParty
}

func (Farbsolo) Klaerungsstich() int {
	return -1
}

func (Farbsolo) OnCompletedTrick(game.Trick, int) {
}

func (f Farbsolo) Type() game.AnnouncedGameType {
	return gameTypeForTrump(f.Trumpf)
}

func gameTypeForTrump(trump game.Suit) game.AnnouncedGameType {
	switch trump {
	case game.Karo:
		return game.KaroSoloType
	case game.Herz:
		return game.HerzSoloType
	case game.Pik:
		return game.PikSoloType
	case game.Kreuz:
		return game.KreuzSoloType
	}
	panic("unexpected solo suit")
}

type VorbehaltFarbsolo struct {
	suit game.Suit
}

func (v VorbehaltFarbsolo) CanAnnounceWith(_ game.Hand) bool {
	return true
}
func (v VorbehaltFarbsolo) Type() game.AnnouncedGameType {
	return gameTypeForTrump(v.suit)
}

func (v VorbehaltFarbsolo) Priority() int {
	return VorbehaltPriorityHochzeit + 1
}

func (v VorbehaltFarbsolo) CreateMode(announcer game.Player) game.Mode {
	return NewFarbsolo(v.suit, announcer)
}

func (v VorbehaltFarbsolo) AnnouncerTakesForehand() bool {
	return false
}
