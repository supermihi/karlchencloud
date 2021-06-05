package tables

import (
	"github.com/stretchr/testify/assert"
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
	u "github.com/supermihi/karlchencloud/server/users"
	"math/rand"
	"testing"
)

func TestTables_JoinTable(t *testing.T) {
	r := NewTables()
	table, err := r.CreateTable(u.Id(1), true, 0)
	assert.Nil(t, err)
	assert.NotEqual(t, InvalidTableId, table.Id)
	_, err = r.JoinTableByInviteCode(2, "not the invite code")
	assert.NotNil(t, err)
	_, err = r.JoinTableByInviteCode(2, table.InviteCode)
	assert.Nil(t, err)
	_, err = r.JoinTableByInviteCode(2, table.InviteCode) // joining again
	assert.NotNil(t, err)
}

func TestTables_GetMatchData(tx *testing.T) {
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
	tm := TableMatch{Match: theMatch, Players: players}
	md := NewMatchData(&tm)
	assert.Equal(tx, curTrick, md.CurrentTrick)
}
