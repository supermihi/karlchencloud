package match

import (
	"github.com/stretchr/testify/assert"
	"github.com/supermihi/karlchencloud/pkg/game/core"
	"testing"
)

func Test_winnerOfGame(t *testing.T) {
	assert.Equal(t, winnerOfGame(Re, Contra, 120, 17), core.ContraParty)
	assert.Equal(t, winnerOfGame(Re, NoBid, 121, 1), core.ReParty)
	assert.Equal(t, winnerOfGame(ReKeine90, ContraKeine90, 129, 1), core.NoParty)
}
