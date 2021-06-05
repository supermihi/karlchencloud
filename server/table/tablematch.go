package table

import (
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
	"github.com/supermihi/karlchencloud/server/users"
)

type PlayerUserMap [game.NumPlayers]users.Id

type TableMatch struct {
	Match   match.Match
	Players PlayerUserMap
}

func (pm PlayerUserMap) PlayerFor(user users.Id) game.Player {
	for p, pId := range pm {
		if pId == user {
			return game.Player(p)
		}
	}
	return game.NoPlayer
}

func (pm PlayerUserMap) IdOf(p game.Player) users.Id {
	if p == game.NoPlayer {
		return users.InvalidId
	}
	return pm[p]
}
