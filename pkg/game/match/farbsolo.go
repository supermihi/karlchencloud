package match

import (
	"github.com/supermihi/karlchencloud/pkg/game/core"
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

func (f Farbsolo) Type() core.AnnouncedGameType {
	switch f.Trumpf {
	case core.Karo:
		return core.KaroSoloType
	case core.Herz:
		return core.HerzSoloType
	case core.Pik:
		return core.PikSoloType
	case core.Kreuz:
		return core.KreuzSoloType
	}
	panic("unexpected solo suit")
}

type VorbehaltFarbsolo struct {
	suit core.Suit
}

func (v VorbehaltFarbsolo) CanAnnounceWith(handCards core.Hand) bool {
	return true
}

func (v VorbehaltFarbsolo) Identifier() ModeId {
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
	return VORBEHALT_PRIORITY_HOCHZEIT + 1
}

func (v VorbehaltFarbsolo) CreateMode(announcer core.Player) core.Mode {
	return NewFarbsolo(v.suit, announcer)
}

func (v VorbehaltFarbsolo) AnnouncerTakesForehand() bool {
	return false
}

func AllFarbsolos() []SonderspielMode {
	ans := make([]SonderspielMode, 4)
	ans[0] = VorbehaltFarbsolo{core.Karo}
	ans[1] = VorbehaltFarbsolo{core.Herz}
	ans[2] = VorbehaltFarbsolo{core.Pik}
	ans[3] = VorbehaltFarbsolo{core.Kreuz}
	return ans
}
