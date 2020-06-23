package server

import (
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
	"time"
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

type TableData struct {
	Id         string
	Owner      string
	InviteCode string
	Players    []string
	Phase      api.TablePhase
	Created    time.Time
}

func GetData(t *Table) *TableData {
	players := t.playersInOrder
	if t.playersInOrder == nil {
		players = t.players
	}
	return &TableData{Id: t.Id, Owner: t.Owner(), InviteCode: t.InviteCode, Players: players,
		Phase: t.Phase, Created: t.Created}
}

func (d *TableData) ContainsPlayer(id string) bool {
	for _, p := range d.Players {
		if p == id {
			return true
		}
	}
	return false
}
