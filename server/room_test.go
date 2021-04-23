package server

import (
	"github.com/stretchr/testify/assert"
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
	"math/rand"
	"testing"
)

func TestRoom_JoinTable(t *testing.T) {
	r := NewRoom(NewMemoryUserDb())
	assert.True(t, r.users.Add("owner", "owner", "secret"))
	table, err := r.CreateTable("owner", nil, nil, 0)
	assert.Nil(t, err)
	_, err = r.JoinTable("player 2", "not the invite code")
	assert.NotNil(t, err)
	_, err = r.JoinTable("player 2", table.InviteCode)
	assert.Nil(t, err)
	_, err = r.JoinTable("player 2", table.InviteCode) // joining again
	assert.NotNil(t, err)
}


func TestRoom_GetMatchData(t *testing.T) {
	rng := rand.New(rand.NewSource(123))
	cards := game.DealCards(rng)
	match := match.NewMatch(game.Player1, cards)
	for _, player := range game.PlayersFrom(game.Player1) {
		match.AnnounceGameType(player, game.NormalGameType)
	}
	curTrick := game.NewIncompleteTrick(game.Player1)
	curTrick.Play(game.Player1, game.Herz9)
	curTrick.Play(game.Player2, game.Herz10)
	match.Game = game.NewGame(cards, game.Player1, game.NewNormalGame(cards))
	match.Game.CurrentTrick = curTrick
	players := [game.NumPlayers]string {"1", "2", "3", "4"}
	tm := TableMatch{match, players}
	md := GetMatchData(&tm)
	assert.Equal(t, curTrick, md.CurrentTrick)
}