package server

import (
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
	"strconv"
	"time"
)

type UserId int64

const InvalidUserId UserId = -1

func (t UserId) String() string {
	return strconv.FormatInt(int64(t), 10)
}

func ParseUserId(idStr string) (UserId, error) {
	idInt, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return InvalidUserId, err
	}
	return UserId(idInt), nil
}

type TableId int64

const InvalidTableId TableId = -1

func (t TableId) String() string {
	return strconv.FormatInt(int64(t), 10)
}
func ParseTableId(idStr string) (TableId, error) {
	idInt, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return InvalidTableId, err
	}
	return TableId(idInt), nil
}

type UserData struct {
	Id    UserId
	Name  string
	Email string
	Token string
}

func NewUserData(id UserId, name string, email string, token string) UserData {
	return UserData{id, name, email, token}
}

func (d UserData) String() string {
	return d.Name
}

type Users interface {
	Add(email string, password string, name string) (UserData, error)
	ListIds() ([]UserId, error)
	GetData(id UserId) (UserData, error)
	ChangeName(id UserId, newName string) error
	VerifyToken(token string) (UserData, error)
	Authenticate(email string, password string) (UserData, error)
}

type PlayerUserMap [game.NumPlayers]UserId

type TableMatch struct {
	Match   match.Match
	Players PlayerUserMap
}

func (pm PlayerUserMap) PlayerFor(user UserId) game.Player {
	for p, pId := range pm {
		if pId == user {
			return game.Player(p)
		}
	}
	return game.NoPlayer
}

func (pm PlayerUserMap) IdOf(p game.Player) UserId {
	if p == game.NoPlayer {
		return InvalidUserId
	}
	return pm[p]
}

func getActivePlayerIds(playersInOrder []UserId, pa match.PlayerAssignment) [game.NumPlayers]UserId {
	var ans [game.NumPlayers]UserId
	for inGamePlayerNumber, playerIndex := range pa.Playing() {
		ans[inGamePlayerNumber] = playersInOrder[playerIndex]
	}
	return ans
}

type TableData struct {
	Id         TableId
	Owner      UserId
	InviteCode string
	Players    []UserId
	Phase      api.TablePhase
	Created    time.Time
}

func GetData(t *Table) *TableData {
	players := t.playersInOrder
	if t.playersInOrder == nil {
		players = t.players
	}
	return &TableData{
		Id:         t.Id,
		Owner:      t.Owner(),
		InviteCode: t.InviteCode,
		Players:    players,
		Phase:      t.Phase,
		Created:    t.Created,
	}
}

func (d *TableData) ContainsPlayer(id UserId) bool {
	for _, p := range d.Players {
		if p == id {
			return true
		}
	}
	return false
}
