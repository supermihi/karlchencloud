package match

import (
	"fmt"
	"github.com/supermihi/karlchencloud/doko/game"
)

type Auction struct {
	sonderspiele   Sonderspiele
	forehand       game.Player
	cards          game.Cards
	vorbehalte     map[game.Player]bool
	specifications map[game.Player]SonderspielMode
}

func NewAuction(forehand game.Player, cards game.Cards, sonderspiele Sonderspiele) *Auction {
	return &Auction{
		sonderspiele,
		forehand,
		cards,
		make(map[game.Player]bool),
		make(map[game.Player]SonderspielMode),
	}
}

type AuctionPhase int

const (
	VorbehaltAbfrage AuctionPhase = iota
	VorbehaltSpezifikation
	AuctionFinished
)

func (a *Auction) Phase() AuctionPhase {
	if len(a.vorbehalte) < game.NumPlayers {
		return VorbehaltAbfrage
	}
	if len(a.specifications) < a.NumVorbehalt() {
		return VorbehaltSpezifikation
	}
	return AuctionFinished
}

type Declaration int

const (
	NotDeclared Declaration = iota
	Gesund
	Vorbehalt
)

func (a *Auction) DeclarationOf(p game.Player) Declaration {
	vorbehalt, ok := a.vorbehalte[p]
	if !ok {
		return NotDeclared
	}
	if vorbehalt {
		return Vorbehalt
	}
	return Gesund
}

func (a *Auction) NumVorbehalt() int {
	ans := 0
	for _, hasVorbehalt := range a.vorbehalte {
		if hasVorbehalt {
			ans += 1
		}
	}
	return ans
}

func (a *Auction) WhoseTurn() game.Player {
	switch a.Phase() {
	case VorbehaltAbfrage:
		return a.forehand.NthNext(len(a.vorbehalte))
	case VorbehaltSpezifikation:
		for _, candidate := range game.PlayersFrom(a.forehand) {
			if a.vorbehalte[candidate] {
				_, ok := a.specifications[candidate]
				if !ok {
					return candidate
				}
			}
		}
		panic("should not happen")
	default:
		return game.NoPlayer
	}
}

func (a *Auction) Announce(player game.Player, vorbehalt bool) {
	if player != a.WhoseTurn() {
		panic(fmt.Sprintf("%v cannot anounce because it's not her turn", player))
	}
	a.vorbehalte[player] = vorbehalt
}

type SpecifyVorbehaltResult int

const (
	Ok SpecifyVorbehaltResult = iota
	UnknownSonderspielId
	SonderspielRequirementsUnmet
)

func (a *Auction) SpecifyVorbehalt(player game.Player, sonderspielId ModeId) SpecifyVorbehaltResult {
	if a.Phase() != VorbehaltSpezifikation || player != a.WhoseTurn() {
		panic("invalid call to SpecifyVorbehalt")
	}
	ans := a.sonderspiele.FindSonderspiel(sonderspielId)
	if ans == nil {
		return UnknownSonderspielId
	}
	if !ans.CanAnnounceWith(a.cards[player]) {
		return SonderspielRequirementsUnmet
	}
	a.specifications[player] = ans
	return Ok
}

type Result struct {
	Sonderspiel   game.Mode
	IsSonderspiel bool
	Forehand      game.Player
}

func (a Auction) GetResult() Result {
	if a.Phase() != AuctionFinished {
		panic("trying to get result from unfinished auction")
	}
	winner := game.NoPlayer
	maxPrio := -1
	for _, player := range game.PlayersFrom(a.forehand) {
		vorbehalt, hasVorbehalt := a.specifications[player]
		if hasVorbehalt && vorbehalt.Priority() > maxPrio {
			winner = player
			maxPrio = vorbehalt.Priority()
		}
	}
	if winner == game.NoPlayer {
		return Result{nil, false, a.forehand}
	}
	sonderspiel := a.specifications[winner]
	mode := sonderspiel.CreateMode(winner)
	forehand := a.forehand
	if sonderspiel.AnnouncerTakesForehand() {
		forehand = winner
	}
	return Result{mode, true, forehand}
}
