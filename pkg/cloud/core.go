package cloud

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/supermihi/karlchencloud/pkg/game/core"
	"github.com/supermihi/karlchencloud/pkg/game/match"
	"github.com/supermihi/karlchencloud/pkg/game/modes"
	"math/rand"
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
	match   *match.Match
	players [core.NumPlayers]UserId
}

type TablePhase int

const (
	BeforeFirstGame TablePhase = iota
	Playing
	WaitingForNextGame
	AllGamesEnded
)

type Table struct {
	id             string
	phase          TablePhase
	players        []UserId
	playersInOrder []UserId
	round          *match.Round
	currentMatch   *TableMatch
}

func (t *Table) Owner() UserId {
	return t.players[0]
}

func NewTable(owner UserId) *Table {
	id := uuid.NewV4().String()
	table := Table{id, BeforeFirstGame, []UserId{owner}, nil, nil, nil}
	return &table
}

func (t *Table) Start() error {
	if t.phase != BeforeFirstGame {
		return Error("table already started")
	}
	if len(t.players) <= core.NumPlayers || len(t.players) >= 7 {
		return Error(fmt.Sprintf("Cannot start table with %v players", len(t.players)))
	}
	t.playersInOrder = make([]UserId, len(t.players))
	rand.Shuffle(len(t.players), func(i int, j int) {
		t.playersInOrder[i], t.playersInOrder[j] = t.playersInOrder[j], t.playersInOrder[i]
	})
	t.round = match.NewRound(len(t.players), modes.StandardSonderspiele(), rand.Int63())
	t.phase = WaitingForNextGame
	return t.StartMatch()
}
func (t *Table) StartMatch() error {
	if t.phase != WaitingForNextGame {
		return Error("can only start match in phase WaitingForNextGame")
	}
	nextMatch := t.round.NextMatch()
	pa := t.round.CurrentPlayerAssignment()

	t.currentMatch = &TableMatch{nextMatch, playerIds(t.playersInOrder, pa)}
	return nil
}

func playerIds(playersInOrder []UserId, pa match.PlayerAssignment) [core.NumPlayers]UserId {
	var ans [core.NumPlayers]UserId
	for inGamePlayerNumber, playerIndex := range pa.Playing() {
		ans[inGamePlayerNumber] = playersInOrder[playerIndex]
	}
	return ans
}

type Tables struct {
	tables map[string]*Table
}

func NewTables() *Tables {
	result := Tables{make(map[string]*Table)}
	return &result
}

func (t *Tables) CreateTable(owner UserId) string {
	table := NewTable(owner)
	t.tables[table.id] = table
	return table.id
}

func (t *Tables) List() []string {
	ids := make([]string, 0, len(t.tables))
	for k := range t.tables {
		ids = append(ids, k)
	}
	return ids
}

func (t *Tables) GetTable(id string) (table *Table, ok bool) {
	table, ok = t.tables[id]
	return
}
