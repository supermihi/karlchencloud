package match

import "github.com/supermihi/karlchencloud/doko/game"

type ExtraPointType int

type ExtraPoint struct {
	Type   ExtraPointType
	Player game.Player
	Trick  int
}

func findExtraPoints(g *game.Game) []ExtraPoint {
	var ans []ExtraPoint
	if !game.IsNormalspiel(g.Mode) {
		return ans
	}
	ans = append(ans, doppelkoepfe(g)...)
	ans = append(ans, fuechse(g)...)
	ans = append(ans, karlchen(g)...)
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

func fuechse(g *game.Game) []ExtraPoint {
	var ans []ExtraPoint
	for i, trick := range g.CompleteTricks {
		for _, player := range game.Players() {
			if trick.CardsOf[player] == game.KaroA && g.Mode.PartyOf(player) != g.Mode.PartyOf(trick.Winner) {
				ans = append(ans, ExtraPoint{FuchsGefangen, trick.Winner, i})
			}
		}
	}
	return ans
}

func karlchen(g *game.Game) []ExtraPoint {
	lastTrick := g.CompleteTricks[game.NumTricks-1]
	if lastTrick.CardsOf[lastTrick.Winner] == game.KreuzB {
		ans := [1]ExtraPoint{{Karlchen, lastTrick.Winner, game.NumTricks - 1}}
		return ans[:]
	}
	return []ExtraPoint{}
}
