package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIncompleteTrick_AsCompleteTrick(t *testing.T) {
	it := IncompleteTrick{Player3, map[Player]Card{Player1: SpadesB, Player2: Diamonds9, Player3: ClubsD, Player4: Diamonds10}}
	ct := it.AsCompleteTrick(Player3)
	assert.Equal(t, ct.Forehand, it.Forehand)
	assert.ElementsMatch(t, ct.CardsOf, [NumPlayers]Card{ClubsD, Diamonds10, SpadesB, Diamonds9})
}
