package match

import "github.com/supermihi/doppelgopf/pkg/game/core"

type ExtraPointType int

type ExtraPoint struct {
	Type   ExtraPointType
	Player core.Player
	Trick  int
}

func findExtraPoints(game *core.Game) []ExtraPoint {
	var ans []ExtraPoint
	if !core.IsNormalspiel(game.Mode) {
		return ans
	}
	ans = append(ans, doppelkoepfe(game)...)
	ans = append(ans, fuechse(game)...)
	ans = append(ans, karlchen(game)...)
	return ans
}

func doppelkoepfe(game *core.Game) []ExtraPoint {
	var ans []ExtraPoint
	for i, trick := range game.CompleteTricks {
		if trick.Score() >= 40 {
			ep := ExtraPoint{Doppelkopf, trick.Winner, i}
			ans = append(ans, ep)
		}
	}
	return ans
}

func fuechse(game *core.Game) []ExtraPoint {
	var ans []ExtraPoint
	for i, trick := range game.CompleteTricks {
		for _, player := range core.Players() {
			if trick.CardsOf[player] == core.KaroA && game.Mode.PartyOf(player) != game.Mode.PartyOf(trick.Winner) {
				ans = append(ans, ExtraPoint{FuchsGefangen, trick.Winner, i})
			}
		}
	}
	return ans
}

func karlchen(game *core.Game) []ExtraPoint {
	lastTrick := game.CompleteTricks[core.NumTricks-1]
	if lastTrick.CardsOf[lastTrick.Winner] == core.KreuzB {
		ans := [1]ExtraPoint{{Karlchen, lastTrick.Winner, core.NumTricks - 1}}
		return ans[:]
	}
	return []ExtraPoint{}
}
