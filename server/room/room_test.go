package room

import (
	"github.com/stretchr/testify/assert"
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
	t "github.com/supermihi/karlchencloud/server/table"
	u "github.com/supermihi/karlchencloud/server/users"
	"math/rand"
	"testing"
)

func TestRoom_JoinTable(te *testing.T) {
	r := NewRoom()
	table, err := r.CreateTable(u.Id(1), true, 0)
	assert.Nil(te, err)
	assert.NotEqual(te, t.InvalidId, table.Id)
	_, err = r.JoinTableByInviteCode(2, "not the invite code")
	assert.NotNil(te, err)
	_, err = r.JoinTableByInviteCode(2, table.InviteCode)
	assert.Nil(te, err)
	_, err = r.JoinTableByInviteCode(2, table.InviteCode) // joining again
	assert.NotNil(te, err)
}

func TestRoom_GetMatchData(tx *testing.T) {
	rng := rand.New(rand.NewSource(123))
	cards := game.DealCards(rng)
	theMatch := match.NewMatch(game.Player1, cards)
	for _, player := range game.PlayersFrom(game.Player1) {
		theMatch.AnnounceGameType(player, game.NormalGameType)
	}
	curTrick := game.NewIncompleteTrick(game.Player1)
	curTrick.Play(game.Player1, game.Hearts9)
	curTrick.Play(game.Player2, game.Hearts10)
	theMatch.Game = game.NewGame(cards, game.Player1, game.NewNormalGame(cards))
	theMatch.Game.CurrentTrick = curTrick
	players := [game.NumPlayers]u.Id{u.Id(1), u.Id(2), u.Id(3), u.Id(4)}
	tm := t.TableMatch{Match: theMatch, Players: players}
	md := NewMatchData(&tm)
	assert.Equal(tx, curTrick, md.CurrentTrick)
}
