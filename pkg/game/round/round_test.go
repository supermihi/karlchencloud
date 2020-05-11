package round

import (
	"github.com/stretchr/testify/assert"
	"github.com/supermihi/karlchencloud/pkg/game/core"
	"testing"
)

func TestPlayersInGame_4(t *testing.T) {
	for i := 0; i < core.NumPlayers; i++ {
		p := PlayersInGame(4, i)
		assert.ElementsMatch(t, [...]int{0, 1, 2, 3}, p)
	}
}

func TestPlayersInGame_5(t *testing.T) {
	p := PlayersInGame(5, 0)
	assert.ElementsMatch(t, [...]int{1, 2, 3, 4}, p)
	p = PlayersInGame(5, 1)
	assert.ElementsMatch(t, [...]int{0, 2, 3, 4}, p)
	p = PlayersInGame(5, 2)
	assert.ElementsMatch(t, [...]int{0, 1, 3, 4}, p)
	p = PlayersInGame(5, 3)
	assert.ElementsMatch(t, [...]int{0, 1, 2, 4}, p)
	p = PlayersInGame(5, 4)
	assert.ElementsMatch(t, [...]int{0, 1, 2, 3}, p)
}

func TestPlayersInGame_6(t *testing.T) {
	p := PlayersInGame(6, 0)
	assert.ElementsMatch(t, [...]int{0, 2, 3, 4}, p)
	p = PlayersInGame(6, 1)
	assert.ElementsMatch(t, [...]int{1, 3, 4, 5}, p)
	p = PlayersInGame(6, 2)
	assert.ElementsMatch(t, [...]int{0, 2, 4, 5}, p)
	p = PlayersInGame(6, 3)
	assert.ElementsMatch(t, [...]int{0, 1, 3, 5}, p)
	p = PlayersInGame(6, 4)
	assert.ElementsMatch(t, [...]int{0, 1, 2, 4}, p)
	p = PlayersInGame(6, 5)
	assert.ElementsMatch(t, [...]int{1, 2, 3, 5}, p)
}
