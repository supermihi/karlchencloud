package match

import "github.com/supermihi/doppelgopf/pkg/game/core"

func EvaluateGame(game *core.Game, bids *Bids) GameEvaluation {
	if !game.IsFinished() {
		panic("cannot evaluate unfinished game")
	}
	reScore, reTricks := countReScoreAndTricks(game)
	winningParty := WinnerOfGame(reScore, reTricks, bids)
	gamePoints := getGamePoints(bids, winningParty, reScore, reTricks)
	extraPoints := findExtraPoints(game)
	totalValue := len(extraPoints)
	for _, p := range gamePoints {
		totalValue += p.Value
	}
	return GameEvaluation{
		winningParty,
		reScore,
		core.TotalScore - reScore,
		gamePoints,
		extraPoints,
		totalValue,
	}
}

type GameEvaluation struct {
	Winner           core.Party
	TrickScoreRe     int
	TrickScoreContra int
	GamePoints       []GamePoint
	ExtraPoints      []ExtraPoint
	TotalValue       int
}

const (
	Karlchen ExtraPointType = iota
	Doppelkopf
	FuchsGefangen
)

func countReScoreAndTricks(game *core.Game) (int, int) {
	reScore, contraScore := 0, 0
	reTricks, contraTricks := 0, 0
	for _, trick := range game.CompleteTricks {
		if game.Mode.PartyOf(trick.Winner) == core.ReParty {
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
