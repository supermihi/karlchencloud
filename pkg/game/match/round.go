package match

import (
	"github.com/supermihi/karlchencloud/pkg/game/core"
)

type PlayerAssignment struct {
	playing [core.NumPlayers]int // core.Player->round-player index
	dealer  int
}

func NewPlayerAssignment(numPlayers int, round int) PlayerAssignment {
	dealer := Dealer(numPlayers, round)
	playing := PlayersInGame(numPlayers, round)
	return PlayerAssignment{playing, dealer}
}

func (p *PlayerAssignment) Playing() [core.NumPlayers]int {
	return p.playing
}

type MatchStats struct {
	evaluation *GameEvaluation
	players    *PlayerAssignment
}

func NewMatchStats(evaluation *GameEvaluation, assignment *PlayerAssignment) MatchStats {
	return MatchStats{evaluation, assignment}
}

type Round struct {
	numPlayers int
	scores     []MatchStats
	rules      Sonderspiele
	cardSeed   int64
}

func NewRound(numPlayers int, rules Sonderspiele, cardSeed int64) *Round {
	return &Round{numPlayers, make([]MatchStats, 0), rules, cardSeed}
}
func (r *Round) CurrentPlayerAssignment() PlayerAssignment {
	return NewPlayerAssignment(r.numPlayers, r.NumFinishedGames())
}

func (r *Round) NextMatch() *Match {
	index := r.NumFinishedGames()
	assignment := NewPlayerAssignment(r.numPlayers, index)
	ans := NewMatch(assignment.Forehand(r.numPlayers), r.rules, core.DealCards(r.cardSeed+int64(index)))
	return &ans
}

func (r *Round) NumFinishedGames() int {
	return len(r.scores)
}

func Dealer(numPlayers int, round int) int {
	return round % numPlayers
}

func PlayersInGame(numPlayers int, round int) [core.NumPlayers]int {
	dealer := Dealer(numPlayers, round)
	var ans [core.NumPlayers]int
	inGame := func(i int) bool {
		switch numPlayers {
		case 4:
			return true
		case 5:
			return i != dealer
		case 6:
			return (i+1)%numPlayers != dealer && (i+numPlayers-1)%numPlayers != dealer
		default:
			panic("unsupported number of players")
		}
	}
	i := 0
	for player := 0; player < numPlayers; player++ {
		if inGame(player) {
			ans[i] = player
			i += 1
		}
	}
	return ans
}

func (a *PlayerAssignment) Forehand(numPlayers int) core.Player {
	for i := (a.dealer + 1) % numPlayers; ; i = (i + 1) % numPlayers {
		for _, j := range a.playing {
			if j == (i % numPlayers) {
				return core.Player(j)
			}
		}
	}
}
