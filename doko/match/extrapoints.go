package match

import "github.com/supermihi/karlchencloud/doko/game"

type ExtraPointType int

type ExtraPoint struct {
	Type   ExtraPointType
	Player game.Player
	Trick  int
}

func findExtraPoints(game *game.Game) []ExtraPoint {
	var ans []ExtraPoint
	if !game.IsNormalspiel(game.Mode) {
		return ans
	}
	ans = append(ans, doppelkoepfe(game)...)
	ans = append(ans, fuechse(game)...)
	ans = append(ans, karlchen(game)...)
	return ans
}

func doppelkoepfe(game *game.Game) []ExtraPoint {
	var ans []ExtraPoint
	for i, trick := range game.CompleteTricks {
		if trick.Score() >= 40 {
			ep := ExtraPoint{Doppelkopf, trick.Winner, i}
			ans = append(ans, ep)
		}
	}
	return ans
}

func fuechse(game *game.Game) []ExtraPoint {
	var ans []ExtraPoint
	for i, trick := range game.CompleteTricks {
		for _, player := range game.Players() {
			if trick.CardsOf[player] == game.KaroA && game.Mode.PartyOf(player) != game.Mode.PartyOf(trick.Winner) {
				ans = append(ans, ExtraPoint{FuchsGefangen, trick.Winner, i})
			}
		}
	}
	return ans
}

func karlchen(game *game.Game) []ExtraPoint {
	lastTrick := game.CompleteTricks[game.NumTricks-1]
	if lastTrick.CardsOf[lastTrick.Winner] == game.KreuzB {
		ans := [1]ExtraPoint{{Karlchen, lastTrick.Winner, game.NumTricks - 1}}
		return ans[:]
	}
	return []ExtraPoint{}
}
