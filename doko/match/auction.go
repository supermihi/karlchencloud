package match

import (
	"github.com/supermihi/karlchencloud/doko/game"
)

type Declaration struct {
	Gesund    bool
	Vorbehalt Vorbehalt
}

func DeclarationGesund() Declaration {
	return Declaration{true, nil}
}

func DeclarationVorbehalt(v Vorbehalt) Declaration {
	return Declaration{false, v}
}

type Auction struct {
	forehand     game.Player
	cards        game.Cards
	Declarations map[game.Player]Declaration
}

func NewAuction(forehand game.Player, cards game.Cards) *Auction {
	return &Auction{
		forehand,
		cards,
		make(map[game.Player]Declaration),
	}
}

func (a *Auction) IsFinished() bool {
	return len(a.Declarations) == game.NumPlayers
}

func (a *Auction) WhoseTurn() game.Player {
	if a.IsFinished() {
		return game.NoPlayer
	}
	return a.forehand.NthNext(len(a.Declarations))
}

type DeclarationResult int

func (a *Auction) Declare(player game.Player, t game.AnnouncedGameType) bool {
	if player != a.WhoseTurn() {
		panic("invalid call to declare")
	}
	v := GetVorbehalt(t)
	if v == nil {
		// gesund
		a.Declarations[player] = DeclarationGesund()
	} else {
		if !v.CanAnnounceWith(a.cards[player]) {
			return false
		}
		a.Declarations[player] = DeclarationVorbehalt(v)
	}
	return true
}

type Result struct {
	Mode     game.Mode
	Forehand game.Player
}

func (a Auction) GetResult() Result {
	if !a.IsFinished() {
		panic("trying to get result from unfinished auction")
	}
	winner := game.NoPlayer
	maxPrio := -1
	for _, player := range game.PlayersFrom(a.forehand) {
		d := a.Declarations[player]
		if !d.Gesund && d.Vorbehalt.Priority() > maxPrio {
			winner = player
			maxPrio = d.Vorbehalt.Priority()
		}
	}
	if winner == game.NoPlayer {
		return Result{game.NewNormalspiel(a.cards), a.forehand}
	}
	vorbehalt := a.Declarations[winner].Vorbehalt
	forehand := a.forehand
	if vorbehalt.AnnouncerTakesForehand() {
		forehand = winner
	}
	return Result{vorbehalt.CreateMode(winner), forehand}
}
