package match

import "github.com/supermihi/karlchencloud/doko/game"

func getParties(mode game.Mode) [game.NumPlayers]game.Party {
	var result [game.NumPlayers]game.Party
	for i, p := range game.Players() {
		result[i] = mode.PartyOf(p)
	}
	return result
}

func EvaluateGame(game *game.Game, bids *Bids) GameEvaluation {
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
		game.TotalScore - reScore,
		gamePoints,
		extraPoints,
		totalValue,
		parties,
		game.IsCountedSolo(game.Mode),
	}
}

type GameEvaluation struct {
	Winner           game.Party
	TrickScoreRe     int
	TrickScoreContra int
	GamePoints       []GamePoint
	ExtraPoints      []ExtraPoint
	TotalValue       int
	Parties          [game.NumPlayers]game.Party
	SoloGame         bool
}

func PointsByPlayer(eval *GameEvaluation, mode game.Mode) [game.NumPlayers]int {
	var ans [game.NumPlayers]int
	for _, p := range game.Players() {
		value := eval.TotalValue
		if mode.PartyOf(p) != eval.Winner {
			value = -value
		}
		if game.IsCountedSolo(mode) && mode.PartyOf(p) == game.ReParty {
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

func countReScoreAndTricks(game *game.Game) (int, int) {
	reScore, contraScore := 0, 0
	reTricks, contraTricks := 0, 0
	for _, trick := range game.CompleteTricks {
		if game.Mode.PartyOf(trick.Winner) == game.ReParty {
			reScore += trick.Score()
			reTricks += 1
		} else {
			contraScore += trick.Score()
			contraTricks += 1

		}
	}
	if reScore+contraScore != game.TotalScore {
		panic("total score != 240")
	}
	if reTricks+contraTricks != game.NumTricks {
		panic("total tricks != 12")
	}
	return reScore, reTricks
}
