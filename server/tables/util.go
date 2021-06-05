package tables

import (
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
	"github.com/supermihi/karlchencloud/server/errors"
	"github.com/supermihi/karlchencloud/server/users"
	"github.com/supermihi/karlchencloud/utils"
)

func getActivePlayerIds(playersInOrder []users.Id, pa match.PlayerAssignment) [game.NumPlayers]users.Id {
	var ans [game.NumPlayers]users.Id
	for inGamePlayerNumber, playerIndex := range pa.Playing() {
		ans[inGamePlayerNumber] = playersInOrder[playerIndex]
	}
	return ans
}

func ensureUserIsOwnerUnsafe(table *Table, userId users.Id) error {
	if !table.ContainsPlayer(userId) {
		return errors.NewCloudError(errors.UserNotAtTable)
	}
	if table.Owner() != userId {
		return errors.NewCloudError(errors.NotOwnerOfTable)
	}
	return nil
}

func getMatchAndPlayerUnsafe(table *Table, user users.Id) (match *TableMatch, p game.Player, err error) {
	if table.CurrentMatch == nil {
		err = errors.NewCloudError(errors.NoCurrentMatch)
		return
	}
	players := table.CurrentMatch.Players
	p = players.PlayerFor(user)
	if p == game.NoPlayer {
		err = errors.NewCloudError(errors.UserNotPlayingInMatch)
		return
	}
	return table.CurrentMatch, p, nil
}

func randomInviteCode() string {
	return utils.RandomLetters(12)
}
