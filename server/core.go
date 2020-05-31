package server

import (
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
)

type Users interface {
	Add(user string, name string, secret string) bool
	List() []string
	GetName(id string) string
	ChangeName(id string, name string)
	Authenticate(id string, secret string) bool
}

type PlayerUserMap [game.NumPlayers]string
type TableMatch struct {
	Match   match.Match
	Players PlayerUserMap
}

func (pm PlayerUserMap) PlayerFor(user string) game.Player {
	for p, pId := range pm {
		if pId == user {
			return game.Player(p)
		}
	}
	return game.NoPlayer
}

func getActivePlayerIds(playersInOrder []string, pa match.PlayerAssignment) [game.NumPlayers]string {
	var ans [game.NumPlayers]string
	for inGamePlayerNumber, playerIndex := range pa.Playing() {
		ans[inGamePlayerNumber] = playersInOrder[playerIndex]
	}
	return ans
}

type Tables struct {
	ById map[string]*Table
}

func NewTables() *Tables {
	result := Tables{make(map[string]*Table)}
	return &result
}

func (t *Tables) CreateTable(owner string) *Table {
	table := NewTable(owner)
	t.ById[table.Id] = table
	return table
}

func (t *Tables) List() []*Table {
	ids := make([]*Table, 0, len(t.ById))
	for _, v := range t.ById {
		ids = append(ids, v)
	}
	return ids
}
