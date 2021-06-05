package table

import (
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
	"github.com/supermihi/karlchencloud/server/users"
)

func getActivePlayerIds(playersInOrder []users.Id, pa match.PlayerAssignment) [game.NumPlayers]users.Id {
	var ans [game.NumPlayers]users.Id
	for inGamePlayerNumber, playerIndex := range pa.Playing() {
		ans[inGamePlayerNumber] = playersInOrder[playerIndex]
	}
	return ans
}
