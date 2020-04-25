package match

import "github.com/supermihi/doppelgopf/pkg/game/core"

type GameEvaluation struct {
	Winner           core.Party
	TrickScoreRe     int
	TrickScoreContra int
	GameValue        int
}

func EvaluateGame(g *core.Game, b *Bids) GameEvaluation {
	if !g.IsFinished() {
		panic("cannot evaluate unfinished game")
	}
	reScore, contraScore := 0, 0
	reTricks, contraTricks := 0, 0
	for _, trick := range g.CompleteTricks {
		winner := g.WinnerOfTrick(trick)
		if g.Mode.PartyOf(winner) == core.Re {
			reScore += trick.Score()
			reTricks += 1
		} else {
			contraScore += trick.Score()
			contraTricks += 1
		}
	}
	winningParty := b.WinnerOfGame(reScore, reTricks)
	return GameEvaluation{
		winningParty,
		reScore,
		core.TotalScore - reScore,
		1,
	}
}
