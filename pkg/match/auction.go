package match

import (
	"github.com/supermihi/doppelgopf/pkg/core"
)

type ModeId string

type Sonderspiel interface {
	IsValid(handCards core.Hand) bool
	Identifier() ModeId
	Priority() int
	CreateMode(announcer core.Player) core.Mode
}

type RuleSet struct {
	Sonderspiele map[ModeId]Sonderspiel
	Normalspiel  core.Mode
}

type Auction struct {
	Forehand            core.Player
	Cards               [core.NumPlayers]core.Hand
	Vorbehalte          map[core.Player]bool
	Rules               RuleSet
	SpecifiedVorbehalte map[core.Player]Sonderspiel
}
type AuctionPhase int

const (
	Announcements AuctionPhase = iota
	VorbehaltSpecification
	AuctionFinished
)

func (a Auction) Phase() AuctionPhase {
	if len(a.Vorbehalte) < core.NumPlayers {
		return Announcements
	}
	if len(a.SpecifiedVorbehalte) < a.NumVorbehalt() {
		return VorbehaltSpecification
	}
	return AuctionFinished
}

func (a Auction) NumVorbehalt() int {
	ans := 0
	for _, hasVorbehalt := range a.Vorbehalte {
		if hasVorbehalt {
			ans += 1
		}
	}
	return ans
}

func (a Auction) WhoseTurn() core.Player {
	switch a.Phase() {
	case Announcements:
		return a.Forehand.NthNext(len(a.Vorbehalte))
	case VorbehaltSpecification:
		for _, candidate := range core.PlayersFrom(a.Forehand) {
			if a.Vorbehalte[candidate] {
				_, ok := a.SpecifiedVorbehalte[candidate]
				if !ok {
					return candidate
				}
			}
		}
		panic("should not happen")
	default:
		return core.NoPlayer
	}
}

func (a Auction) Announce(player core.Player, vorbehalt bool) bool {
	if player != a.WhoseTurn() {
		return false
	}
	a.Vorbehalte[player] = vorbehalt
	return true
}

func (a Auction) SpecifyVorbehalt(player core.Player, sonderspielId ModeId) bool {
	if a.Phase() != VorbehaltSpecification || player != a.WhoseTurn() {
		return false
	}
	ans := a.Rules.FindSonderspiel(sonderspielId)
	if ans == nil {
		return false
	}
	if !ans.IsValid(a.Cards[player]) {
		return false
	}
	a.SpecifiedVorbehalte[player] = ans
	return true
}

func (r *RuleSet) FindSonderspiel(id ModeId) Sonderspiel {
	var result, ok = r.Sonderspiele[id]
	if !ok {
		return nil
	}
	return result
}

func (a Auction) WinningMode() (core.Mode, bool) {
	if a.Phase() != AuctionFinished {
		return nil, false
	}
	winner := core.NoPlayer
	maxPrio := -1
	for _, player := range core.PlayersFrom(a.Forehand) {
		vorbehalt, hasVorbehalt := a.SpecifiedVorbehalte[player]
		if hasVorbehalt && vorbehalt.Priority() > maxPrio {
			winner = player
			maxPrio = vorbehalt.Priority()
		}
	}
	if winner == core.NoPlayer {
		return a.Rules.Normalspiel, true
	}
	return a.SpecifiedVorbehalte[winner].CreateMode(winner), true
}
