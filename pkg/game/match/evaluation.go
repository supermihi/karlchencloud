package match

import "github.com/supermihi/doppelgopf/pkg/game/core"

type ExtraPointType int

const (
	Karlchen ExtraPointType = iota
	Doppelkopf
	FuchsGefangen
)

type ExtraPoint struct {
	Type   ExtraPointType
	Player core.Player
	Trick  int
}
type GameEvaluation struct {
	Winner           core.Party
	TrickScoreRe     int
	TrickScoreContra int
	GameValue        int
	ExtraPoints      []ExtraPoint
}

func CountReScoreAndTricks(game *core.Game) (int, int) {
	reScore, contraScore := 0, 0
	reTricks, contraTricks := 0, 0
	for _, trick := range game.CompleteTricks {
		winner := game.WinnerOfTrick(trick)
		if game.Mode.PartyOf(winner) == core.ReParty {
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
