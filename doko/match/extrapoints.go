package match

import "github.com/supermihi/karlchencloud/doko/game"

type ExtraPointType int

type ExtraPoint struct {
	Type   ExtraPointType
	Player game.Player
	Trick  int
}

func extraPointsApply(t game.AnnouncedGameType) bool {
	return t == game.NormalGameType || t == game.MarriageType
}

func findExtraPoints(g *game.Game) (ans []ExtraPoint) {
	if !extraPointsApply(g.Mode.Type()) {
		return
	}
	ans = append(ans, doppelkopfs(g)...)
	ans = append(ans, foxesCaught(g)...)
	ans = append(ans, charlie(g)...)
	return
}

func EventsOccured(g *game.Game, trick *game.Trick) (ans []ExtraPointType) {
	if !extraPointsApply(g.Mode.Type()) {
		return
	}
	if scoresDoppelkopf(trick) {
		ans = append(ans, Doppelkopf)
	}
	for _, player := range game.Players() {
		if maybeFoxCaught(player, trick) {
			ans = append(ans, FoxCaught)
		}
	}
	if charlieWinLastTrick(g) {
		ans = append(ans, Charlie)
	}
	return ans
}

func scoresDoppelkopf(trick *game.Trick) bool {
	return trick.Score() >= 40
}

func doppelkopfs(game *game.Game) []ExtraPoint {
	var ans []ExtraPoint
	for i, trick := range game.CompleteTricks {
		if scoresDoppelkopf(&trick) {
			ep := ExtraPoint{Doppelkopf, trick.Winner, i}
			ans = append(ans, ep)
		}
	}
	return ans
}

//logic without party information
func maybeFoxCaught(player game.Player, trick *game.Trick) bool {
	return trick.CardsOf[player] == game.DiamondsA && player != trick.Winner
}

func foxesCaught(g *game.Game) []ExtraPoint {
	var ans []ExtraPoint
	for i, trick := range g.CompleteTricks {
		for _, player := range game.Players() {
			if trick.CardsOf[player] == game.DiamondsA && g.Mode.PartyOf(player) != g.Mode.PartyOf(trick.Winner) {
				ans = append(ans, ExtraPoint{FoxCaught, trick.Winner, i})
			}
		}
	}
	return ans
}

func charlieWinLastTrick(g *game.Game) bool {
	if !g.IsFinished() {
		return false
	}
	lastTrick := g.CompleteTricks[game.NumTricks-1]
	return lastTrick.CardsOf[lastTrick.Winner] == game.ClubsJ
}

func charlie(g *game.Game) []ExtraPoint {
	lastTrick := g.CompleteTricks[game.NumTricks-1]
	if charlieWinLastTrick(g) {
		ans := [1]ExtraPoint{{Charlie, lastTrick.Winner, game.NumTricks - 1}}
		return ans[:]
	}
	return []ExtraPoint{}
}
