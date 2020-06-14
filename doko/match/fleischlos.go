package match

import (
	"github.com/supermihi/karlchencloud/doko/game"
)

type FleischlosMode struct {
	Soloist game.Player
}

func FleischlosSuite(card game.Card) game.GameSuit {
	return card.Suit.AsFehl()
}
func (f FleischlosMode) GameSuit(card game.Card) game.GameSuit {
	return FleischlosSuite(card)
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

func (VorbehaltFleischlos) CanAnnounceWith(_ game.Hand) bool {
	return true
}
func (VorbehaltFleischlos) Type() game.AnnouncedGameType {
	return game.FleischlosType
}
func (VorbehaltFleischlos) Priority() int {
	return VorbehaltPriorityHochzeit + 1
}

func (VorbehaltFleischlos) CreateMode(announcer game.Player) game.Mode {
	return FleischlosMode{announcer}
}

func (VorbehaltFleischlos) AnnouncerTakesForehand() bool {
	return false
}
