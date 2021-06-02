package room

import (
	"github.com/stretchr/testify/assert"
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
	users2 "github.com/supermihi/karlchencloud/room/users"
	"math/rand"
	"testing"
)

func TestRoom_JoinTable(t *testing.T) {
	r := NewRoom(users2.NewMemoryUserDb())
	id, err := r.users.Add("owner@example.com", "secret", "owner", false)
	assert.NotNil(t, err)
	assert.NotEqual(t, InvalidUserId, id)
	table, err := r.CreateTable(id, InvalidTableId, nil, 0)
	assert.Nil(t, err)
	assert.NotEqual(t, InvalidTableId, table.Id)
	_, err = r.JoinTableByInviteCode(2, "not the invite code")
	assert.NotNil(t, err)
	_, err = r.JoinTableByInviteCode(2, table.InviteCode)
	assert.Nil(t, err)
	_, err = r.JoinTableByInviteCode(2, table.InviteCode) // joining again
	assert.NotNil(t, err)
}

func TestRoom_GetMatchData(t *testing.T) {
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
	players := [game.NumPlayers]UserId{UserId(1), UserId(2), UserId(3), UserId(4)}
	tm := TableMatch{theMatch, players}
	md := GetMatchData(&tm)
	assert.Equal(t, curTrick, md.CurrentTrick)
}
