package match

import "github.com/supermihi/doppelgopf/pkg/game/core"

type GameEvaluation struct {
	Winner           core.Party
	TrickScoreRe     int
	TrickScoreContra int
	GameValue        int
}

func CountReScoreAndTricks(game *core.Game) (int, int) {
	reScore, contraScore := 0, 0
	reTricks, contraTricks := 0, 0
	for _, trick := range game.CompleteTricks {
		winner := game.WinnerOfTrick(trick)
		if game.Mode.PartyOf(winner) == core.Re {
			reScore += trick.Score()
			reTricks += 1
		} else {
			contraScore += trick.Score()
			contraTricks += 1

		}
	}
	if reScore+contraScore != core.TotalScore {
		panic("total score != 240")
	}
	if reTricks+contraTricks != core.NumTricks {
		panic("total tricks != 12")
	}
	return reScore, reTricks
}

func EvaluateGame(game *core.Game, bids *Bids) GameEvaluation {
	if !game.IsFinished() {
		panic("cannot evaluate unfinished game")
	}
	reScore, reTricks := CountReScoreAndTricks(game)
	winningParty := WinnerOfGame(reScore, reTricks, bids)
	return GameEvaluation{
		winningParty,
		reScore,
		core.TotalScore - reScore,
		1,
	}
}
