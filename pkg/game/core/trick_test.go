package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIncompleteTrick_AsCompleteTrick(t *testing.T) {
	it := IncompleteTrick{Player3, []Card{PikB, Karo9, KreuzD, Karo10}}
	ct := it.AsCompleteTrick(Player3)
	assert.Equal(t, ct.Forehand, it.Forehand)
	assert.ElementsMatch(t, ct.CardsOf, [NumPlayers]Card{KreuzD, Karo10, PikB, Karo9})
}
