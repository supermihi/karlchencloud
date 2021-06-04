package room

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/doko/match"
)

func TestGetData_PlayersAreInGameOrder(t *testing.T) {
	rng := rand.New(rand.NewSource(123))
	players := []UserId{UserId(1), UserId(2), UserId(3), UserId(4)}
	playersInOrder := []UserId{UserId(3), UserId(2), UserId(4), UserId(1)}
	table := &Table{TableId(123), time.Now(), "123", api.TablePhase_NOT_STARTED, players, playersInOrder,
		match.NewRound(4, rng), nil, rng}
	data := getData(table)
	assert.Equal(t, playersInOrder, data.Players)
}
