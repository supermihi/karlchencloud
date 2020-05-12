package match

import (
	"github.com/stretchr/testify/assert"
	"github.com/supermihi/karlchencloud/doko/game"
	"testing"
)

func Test_winnerOfGame(t *testing.T) {
	assert.Equal(t, winnerOfGame(Re, Contra, 120, 17), game.ContraParty)
	assert.Equal(t, winnerOfGame(Re, NoBid, 121, 1), game.ReParty)
	assert.Equal(t, winnerOfGame(ReKeine90, ContraKeine90, 129, 1), game.NoParty)
}
