package server

import (
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
)

type UserData struct {
	Id   string
	Name string
}

func (d UserData) String() string {
	return d.Name
}

type Users interface {
	Add(user string, name string, secret string) bool
	List() []string
	GetName(id string) (name string, ok bool)
	ChangeName(id string, name string) (ok bool)
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

func (pm PlayerUserMap) IdOf(p game.Player) string {
	if p == game.NoPlayer {
		return ""
	}
	return pm[p]
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

type TableData struct {
	Id         string
	Owner      string
	InviteCode string
	Players    []string
	InMatch    bool
}

func GetData(t *Table) *TableData {
	players := t.playersInOrder
	if t.playersInOrder == nil {
		players = t.players
	}
	return &TableData{Id: t.Id, Owner: t.Owner(), InviteCode: t.InviteCode, Players: players,
		InMatch: t.CurrentMatch != nil}
}

func (d *TableData) ContainsPlayer(id string) bool {
	for _, p := range d.Players {
		if p == id {
			return true
		}
	}
	return false
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

func (t *Tables) ActiveTableOf(user string) *TableData {
	for _, table := range t.ById {
		if table.ContainsPlayer(user) {
			return GetData(table)
		}
	}
	return nil
}
