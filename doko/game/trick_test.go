package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIncompleteTrick_AsCompleteTrick(t *testing.T) {
	it := IncompleteTrick{Player3, map[Player]Card{Player1: SpadesJ, Player2: Diamonds9, Player3: ClubsQ, Player4: Diamonds10}}
	ct := it.AsCompleteTrick(Player3)
	assert.Equal(t, ct.Forehand, it.Forehand)
	assert.ElementsMatch(t, ct.CardsOf, [NumPlayers]Card{ClubsQ, Diamonds10, SpadesJ, Diamonds9})
}
