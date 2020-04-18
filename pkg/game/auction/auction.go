package auction

import (
	"fmt"
	"github.com/supermihi/doppelgopf/pkg/game/core"
)

type Auction struct {
	Sonderspiele        Sonderspiele
	Forehand            core.Player
	Cards               core.Cards
	Vorbehalte          map[core.Player]bool
	SpecifiedVorbehalte map[core.Player]SonderspielMode
}

func NewAuction(forehand core.Player, cards core.Cards, sonderspiele Sonderspiele) *Auction {
	return &Auction{
		sonderspiele,
		forehand,
		cards,
		make(map[core.Player]bool),
		make(map[core.Player]SonderspielMode),
	}
}

type Phase int

const (
	VorbehaltAbfrage Phase = iota
	VorbehaltSpezifikation
	Finished
)

func (a Auction) Phase() Phase {
	if len(a.Vorbehalte) < core.NumPlayers {
		return VorbehaltAbfrage
	}
	if len(a.SpecifiedVorbehalte) < a.NumVorbehalt() {
		return VorbehaltSpezifikation
	}
	return Finished
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
	case VorbehaltAbfrage:
		return a.Forehand.NthNext(len(a.Vorbehalte))
	case VorbehaltSpezifikation:
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

func (a Auction) Announce(player core.Player, vorbehalt bool) {
	if player != a.WhoseTurn() {
		panic(fmt.Sprintf("%v cannot anounce because it's not her turn", player))
	}
	a.Vorbehalte[player] = vorbehalt
}

type SpecifyVorbehaltResult int

const (
	Ok SpecifyVorbehaltResult = iota
	UnknownSonderspielId
	SonderspielRequirementsUnmet
)

func (a Auction) SpecifyVorbehalt(player core.Player, sonderspielId ModeId) SpecifyVorbehaltResult {
	if a.Phase() != VorbehaltSpezifikation || player != a.WhoseTurn() {
		panic("invalid call to SpecifyVorbehalt")
	}
	ans := a.Sonderspiele.FindSonderspiel(sonderspielId)
	if ans == nil {
		return UnknownSonderspielId
	}
	if !ans.CanAnnounceWith(a.Cards[player]) {
		return SonderspielRequirementsUnmet
	}
	a.SpecifiedVorbehalte[player] = ans
	return Ok
}

type Result struct {
	Sonderspiel   core.Mode
	IsSonderspiel bool
	Forehand      core.Player
}

func (a Auction) GetResult() Result {
	if a.Phase() != Finished {
		panic("trying to get result from unfinished auction")
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
		return Result{nil, false, a.Forehand}
	}
	sonderspiel := a.SpecifiedVorbehalte[winner]
	mode := sonderspiel.CreateMode(winner)
	forehand := a.Forehand
	if sonderspiel.AnnouncerTakesForehand() {
		forehand = winner
	}
	return Result{mode, true, forehand}
}
