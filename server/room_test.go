package server

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRoom_JoinTable(t *testing.T) {
	r := NewRoom(NewMemoryUserDb())
	assert.True(t, r.users.Add("owner", "owner", "secret"))
	table, err := r.CreateTable("owner")
	assert.Nil(t, err)
	_, err = r.JoinTable(table.Id, "player 2", "not the invite code")
	assert.NotNil(t, err)
	_, err = r.JoinTable(table.Id, "player 2", table.InviteCode)
	assert.Nil(t, err)
	_, err = r.JoinTable(table.Id, "player 2", table.InviteCode) // joining again
	assert.NotNil(t, err)
}
