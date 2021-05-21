package match

import "github.com/supermihi/karlchencloud/doko/game"

func getParties(mode game.Mode) [game.NumPlayers]game.Party {
	var result [game.NumPlayers]game.Party
	for i, p := range game.Players() {
		result[i] = mode.PartyOf(p)
	}
	return result
}

func EvaluateGame(g *game.Game, bids *Bids) GameEvaluation {
	if !g.IsFinished() {
		panic("cannot evaluate unfinished game")
	}
	reScore, reTricks := countReScoreAndTricks(g)
	winningParty := WinnerOfGame(reScore, reTricks, bids)
	gamePoints := getGamePoints(bids, winningParty, reScore, reTricks)
	extraPoints := findExtraPoints(g)
	totalValue := len(extraPoints)
	for _, p := range gamePoints {
		totalValue += p.Value
	}
	parties := getParties(g.Mode)
	return GameEvaluation{
		winningParty,
		reScore,
		game.TotalScore - reScore,
		gamePoints,
		extraPoints,
		totalValue,
		parties,
		IsCountedSolo(g.Mode),
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
		if IsCountedSolo(mode) && mode.PartyOf(p) == game.ReParty {
			// soloist
			value = value * 3
		}
		ans[p] = value
	}
	return ans
}

const (
	Charlie ExtraPointType = iota
	Doppelkopf
	FoxCaught
)

func countReScoreAndTricks(g *game.Game) (int, int) {
	reScore, contraScore := 0, 0
	reTricks, contraTricks := 0, 0
	for _, trick := range g.CompleteTricks {
		if g.Mode.PartyOf(trick.Winner) == game.ReParty {
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

func IsCountedSolo(m game.Mode) bool {
	switch u := m.(type) {
	case game.NormalGameMode:
		return u.IsSilentSolo()
	case game.Marriage:
		return false
	default:
		return true
	}
}
