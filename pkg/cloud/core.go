package cloud

import (
	uuid "github.com/satori/go.uuid"
	"github.com/supermihi/karlchencloud/pkg/game/match"
)

type UserId string

type Users interface {
	Add(user UserId, name string, secret string) bool
	List() []UserId
	GetName(id UserId) string
	ChangeName(id UserId, name string)
	Authenticate(id UserId, secret string) bool
}

type Scores []map[UserId]int
type TableId string
type Table struct {
	id           TableId
	players      []UserId
	scores       Scores
	currentMatch *match.Match
}

func (t *Table) Owner() UserId {
	return t.players[0]
}

func NewTable(owner UserId) *Table {
	id := TableId(uuid.NewV4().String())
	table := Table{id, []UserId{owner}, make([]map[UserId]int, 0), nil}
	return &table
}

type Tables struct {
	tables map[TableId]*Table
}

func NewTables() *Tables {
	result := Tables{make(map[TableId]*Table)}
	return &result
}

func (t *Tables) CreateTable(owner UserId) TableId {
	table := NewTable(owner)
	t.tables[table.id] = table
	return table.id
}

func (t *Tables) List() []TableId {
	ids := make([]TableId, 0, len(t.tables))
	for k := range t.tables {
		ids = append(ids, k)
	}
	return ids
}
