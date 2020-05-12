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

func (f Farbsolo) GameSuit(card game.Card) game.GameSuit {
	if card.Suit == f.Trumpf || card.Rank == game.Bube || card.Rank == game.Dame || card == game.Dulle() {
		return game.Trumpf
	}
	return card.Suit.AsFehl()
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
	switch f.Trumpf {
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

func (v VorbehaltFarbsolo) CanAnnounceWith(handCards game.Hand) bool {
	return true
}

func (v VorbehaltFarbsolo) Identifier() ModeId {
	switch v.suit {
	case game.Karo:
		return "KARO_SOLO"
	case game.Herz:
		return "HERZ_SOLO"
	case game.Pik:
		return "PIK_SOLO"
	case game.Kreuz:
		return "KREUZ_SOLO"
	default:
		panic("unexpected suit")
	}

}

func (v VorbehaltFarbsolo) Priority() int {
	return VORBEHALT_PRIORITY_HOCHZEIT + 1
}

func (v VorbehaltFarbsolo) CreateMode(announcer game.Player) game.Mode {
	return NewFarbsolo(v.suit, announcer)
}

func (v VorbehaltFarbsolo) AnnouncerTakesForehand() bool {
	return false
}

func AllFarbsolos() []SonderspielMode {
	ans := make([]SonderspielMode, 4)
	ans[0] = VorbehaltFarbsolo{game.Karo}
	ans[1] = VorbehaltFarbsolo{game.Herz}
	ans[2] = VorbehaltFarbsolo{game.Pik}
	ans[3] = VorbehaltFarbsolo{game.Kreuz}
	return ans
}
