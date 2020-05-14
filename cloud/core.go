package cloud

import (
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
	"github.com/supermihi/karlchencloud/doko/round"
)

type UserId string

type Users interface {
	Add(user UserId, name string, secret string) bool
	List() []UserId
	GetName(id UserId) string
	ChangeName(id UserId, name string)
	Authenticate(id UserId, secret string) bool
}

type TableMatch struct {
	Match   *match.Match
	Players [game.NumPlayers]UserId
}

func getActivePlayerIds(playersInOrder []UserId, pa round.PlayerAssignment) [game.NumPlayers]UserId {
	var ans [game.NumPlayers]UserId
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

func (t *Tables) CreateTable(owner UserId) *Table {
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
