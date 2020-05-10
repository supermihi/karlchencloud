package modes

import (
	"github.com/supermihi/karlchencloud/pkg/game/core"
	"github.com/supermihi/karlchencloud/pkg/game/match"
)

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

type VorbehaltFarbsolo struct {
	suit core.Suit
}

func (v VorbehaltFarbsolo) CanAnnounceWith(handCards core.Hand) bool {
	return true
}

func (v VorbehaltFarbsolo) Identifier() match.ModeId {
	switch v.suit {
	case core.Karo:
		return "KARO_SOLO"
	case core.Herz:
		return "HERZ_SOLO"
	case core.Pik:
		return "PIK_SOLO"
	case core.Kreuz:
		return "KREUZ_SOLO"
	default:
		panic("unexpected suit")
	}

}

func (v VorbehaltFarbsolo) Priority() int {
	return match.VORBEHALT_PRIORITY_HOCHZEIT + 1
}

func (v VorbehaltFarbsolo) CreateMode(announcer core.Player) core.Mode {
	return NewFarbsolo(v.suit, announcer)
}

func (v VorbehaltFarbsolo) AnnouncerTakesForehand() bool {
	return false
}
