package match

import (
	"github.com/supermihi/karlchencloud/doko/game"
)

type FleischlosMode struct {
	Soloist game.Player
}

func (f FleischlosMode) GameSuit(card game.Card) game.GameSuit {
	return card.Suit.AsFehl()
}

func (FleischlosMode) Klaerungsstich() int {
	return -1
}

func (FleischlosMode) OnCompletedTrick(game.Trick, int) {
}

func (f FleischlosMode) PartyOf(p game.Player) game.Party {
	if p == f.Soloist {
		return game.ReParty
	}
	return game.ContraParty
}

func (FleischlosMode) Type() game.AnnouncedGameType {
	return game.FleischlosType
}

type VorbehaltFleischlos struct{}

func (VorbehaltFleischlos) CanAnnounceWith(handCards game.Hand) bool {
	return true
}

func (VorbehaltFleischlos) Identifier() ModeId {
	return "FLEISCHLOS"
}

func (VorbehaltFleischlos) Priority() int {
	return VORBEHALT_PRIORITY_HOCHZEIT + 1
}

func (VorbehaltFleischlos) CreateMode(announcer game.Player) game.Mode {
	return FleischlosMode{announcer}
}

func (VorbehaltFleischlos) AnnouncerTakesForehand() bool {
	return false
}
