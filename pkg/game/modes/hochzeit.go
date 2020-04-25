package modes

import (
	"github.com/supermihi/doppelgopf/pkg/game/auction"
	"github.com/supermihi/doppelgopf/pkg/game/core"
)

type HochzeitState int

const (
	WaitingForPartner HochzeitState = iota
	PartnerFound
	Solo
)
const MaxTricksWaitingForPartner = 3

type Hochzeit struct {
	announcer      core.Player
	partner        core.Player
	state          HochzeitState
	klaerungsstich int
}

func NewHochzeit(announcer core.Player) Hochzeit {
	return Hochzeit{announcer, core.NoPlayer, WaitingForPartner, 0}
}

func (h Hochzeit) OnCompletedTrick(t core.Trick, numTrick int) {
	if h.state != WaitingForPartner {
		return
	}
	winner := core.WinnerOfTrick(t, h)
	if winner != h.announcer {
		h.partner = winner
		h.klaerungsstich = numTrick
	} else if numTrick >= MaxTricksWaitingForPartner-1 {
		h.state = Solo
		h.klaerungsstich = numTrick
	}
}

func (h Hochzeit) PartyOf(p core.Player) core.Party {
	if p == h.announcer {
		return core.Re
	}
	if !h.PartnerFound() {
		return core.NoParty
	}
	if p == h.partner {
		return core.Re
	}
	return core.Contra
}

func (h Hochzeit) PartnerFound() bool {
	return h.partner != core.NoPlayer
}

// trick-taking rules equal those of normal mode
func (h Hochzeit) GameSuit(card core.Card) core.GameSuit {
	return core.NormalGameSuit(card)
}

func (h Hochzeit) Klaerungsstich() int {
	return h.klaerungsstich
}

type VorbehaltHochzeit struct {
}

func (v VorbehaltHochzeit) Identifier() auction.ModeId {
	return "HOCHZEIT"
}

func (v VorbehaltHochzeit) CanAnnounceWith(handCards core.Hand) bool {
	return handCards.NumAlte() == 2
}

func (v VorbehaltHochzeit) Priority() int {
	return 1
}

func (v VorbehaltHochzeit) CreateMode(announcer core.Player) core.Mode {
	return NewHochzeit(announcer)
}
