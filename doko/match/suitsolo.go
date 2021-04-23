package match

import (
	"github.com/supermihi/karlchencloud/doko/game"
)

type SuitSolo struct {
	Trumpf  game.Suit
	Soloist game.Player
}

func NewSuitSolo(trump game.Suit, soloist game.Player) SuitSolo {
	return SuitSolo{trump, soloist}
}

func SuitSoloSuit(trump game.Suit, card game.Card) game.GameSuit {
	if card.Suit == trump || card.Rank == game.Bube || card.Rank == game.Dame || card == game.Dulle() {
		return game.Trumpf
	}
	return card.Suit.AsFehl()
}

func (f SuitSolo) GameSuit(card game.Card) game.GameSuit {
	return SuitSoloSuit(f.Trumpf, card)
}

func (f SuitSolo) PartyOf(p game.Player) game.Party {
	if p == f.Soloist {
		return game.ReParty
	}
	return game.ContraParty
}

func (SuitSolo) Klaerungsstich() int {
	return -1
}

func (SuitSolo) OnCompletedTrick(game.Trick, int) {
}

func (f SuitSolo) Type() game.AnnouncedGameType {
	return gameTypeForTrump(f.Trumpf)
}

func gameTypeForTrump(trump game.Suit) game.AnnouncedGameType {
	switch trump {
	case game.Karo:
		return game.DiamondSoloType
	case game.Herz:
		return game.HeartSoloType
	case game.Pik:
		return game.SpadeSoloType
	case game.Kreuz:
		return game.ClubSoloType
	}
	panic("unexpected solo suit")
}

type ReservationSuitSolo struct {
	suit game.Suit
}

func (v ReservationSuitSolo) CanAnnounceWith(_ game.Hand) bool {
	return true
}
func (v ReservationSuitSolo) Type() game.AnnouncedGameType {
	return gameTypeForTrump(v.suit)
}

func (v ReservationSuitSolo) Priority() int {
	return ReservationPriorityMarriage + 1
}

func (v ReservationSuitSolo) CreateMode(announcer game.Player) game.Mode {
	return NewSuitSolo(v.suit, announcer)
}

func (v ReservationSuitSolo) AnnouncerTakesForehand() bool {
	return false
}
