package match

import "github.com/supermihi/karlchencloud/pkg/game/core"

func getParties(mode core.Mode) [core.NumPlayers]core.Party {
	var result [core.NumPlayers]core.Party
	for i, p := range core.Players() {
		result[i] = mode.PartyOf(p)
	}
	return result
}

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
	parties := getParties(game.Mode)
	return GameEvaluation{
		winningParty,
		reScore,
		core.TotalScore - reScore,
		gamePoints,
		extraPoints,
		totalValue,
		parties,
		core.IsSolo(game.Mode),
	}
}

type GameEvaluation struct {
	Winner           core.Party
	TrickScoreRe     int
	TrickScoreContra int
	GamePoints       []GamePoint
	ExtraPoints      []ExtraPoint
	TotalValue       int
	Parties          [core.NumPlayers]core.Party
	SoloGame         bool
}

func PointsByPlayer(eval *GameEvaluation, mode core.Mode) [core.NumPlayers]int {
	var ans [core.NumPlayers]int
	for _, p := range core.Players() {
		value := eval.TotalValue
		if mode.PartyOf(p) != eval.Winner {
			value = -value
		}
		if core.IsSolo(mode) && mode.PartyOf(p) == core.ReParty {
			// soloist
			value = value * 3
		}
		ans[p] = value
	}
	return ans
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
