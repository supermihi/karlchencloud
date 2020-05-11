package match

import (
	"github.com/supermihi/karlchencloud/pkg/game/core"
)

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

func (FleischlosMode) Type() core.AnnouncedGameType {
	return core.FleischlosType
}

type VorbehaltFleischlos struct{}

func (VorbehaltFleischlos) CanAnnounceWith(handCards core.Hand) bool {
	return true
}

func (VorbehaltFleischlos) Identifier() ModeId {
	return "FLEISCHLOS"
}

func (VorbehaltFleischlos) Priority() int {
	return VORBEHALT_PRIORITY_HOCHZEIT + 1
}

func (VorbehaltFleischlos) CreateMode(announcer core.Player) core.Mode {
	return FleischlosMode{announcer}
}

func (VorbehaltFleischlos) AnnouncerTakesForehand() bool {
	return false
}
