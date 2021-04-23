package match

import (
	"github.com/supermihi/karlchencloud/doko/game"
)

type AceSolo struct {
	Soloist game.Player
}

func AceSoloSuite(card game.Card) game.GameSuit {
	return card.Suit.AsNonTrump()
}
func (f AceSolo) GameSuit(card game.Card) game.GameSuit {
	return AceSoloSuite(card)
}

func (AceSolo) ClarificationTrick() int {
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

type ReservationAceSolo struct{}

func (ReservationAceSolo) CanAnnounceWith(_ game.Hand) bool {
	return true
}
func (ReservationAceSolo) Type() game.AnnouncedGameType {
	return game.AceSoloType
}
func (ReservationAceSolo) Priority() int {
	return ReservationPriorityMarriage + 1
}

func (ReservationAceSolo) CreateMode(announcer game.Player) game.Mode {
	return AceSolo{announcer}
}

func (ReservationAceSolo) AnnouncerTakesForehand() bool {
	return false
}
