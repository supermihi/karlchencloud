package match

import (
	"github.com/supermihi/karlchencloud/doko/game"
)

type AceSolo struct {
	Soloist game.Player
}

func AceSoloSuite(card game.Card) game.GameSuit {
	return card.Suit.AsFehl()
}
func (f AceSolo) GameSuit(card game.Card) game.GameSuit {
	return AceSoloSuite(card)
}

func (AceSolo) Klaerungsstich() int {
	return -1
}

func (AceSolo) OnCompletedTrick(game.Trick, int) {
}

func (f AceSolo) PartyOf(p game.Player) game.Party {
	if p == f.Soloist {
		return game.ReParty
	}
	return game.ContraParty
}

func (AceSolo) Type() game.AnnouncedGameType {
	return game.AceSoloType
}

type VorbehaltAceSolo struct{}

func (VorbehaltAceSolo) CanAnnounceWith(_ game.Hand) bool {
	return true
}
func (VorbehaltAceSolo) Type() game.AnnouncedGameType {
	return game.AceSoloType
}
func (VorbehaltAceSolo) Priority() int {
	return VorbehaltPriorityMarriage + 1
}

func (VorbehaltAceSolo) CreateMode(announcer game.Player) game.Mode {
	return AceSolo{announcer}
}

func (VorbehaltAceSolo) AnnouncerTakesForehand() bool {
	return false
}
