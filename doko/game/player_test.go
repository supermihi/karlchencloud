package game

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

func TestPlayer_NthNext(t *testing.T) {
	assert.Equal(t, Player2, Player1.NthNext(1))
	assert.Equal(t, Player3, Player1.NthNext(2))
	assert.Equal(t, Player4, Player1.NthNext(3))
	assert.Equal(t, Player2, Player3.NthNext(3))
}
