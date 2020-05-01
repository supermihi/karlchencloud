package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlayersFrom(t *testing.T) {
	result := PlayersFrom(Player2)
	expected := [...]Player{Player2, Player3, Player4, Player1}
	assert.ElementsMatch(t, result, expected)
}

func TestPlayersFrom2(t *testing.T) {
	result := PlayersFrom(Player1)
	expected := [...]Player{Player1, Player2, Player3, Player4}
	assert.ElementsMatch(t, result, expected)
}

func TestPlayersFromFails(t *testing.T) {
	assert.Panics(t, func() { PlayersFrom(NoPlayer) })
}
