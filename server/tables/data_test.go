package tables

import (
	"github.com/stretchr/testify/assert"
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/server/users"
	"testing"
	"time"
)

func TestGetData_PlayersAreInGameOrder(t *testing.T) {
	players := []users.Id{users.Id(1), users.Id(2), users.Id(3), users.Id(4)}
	playersInOrder := []users.Id{users.Id(3), users.Id(2), users.Id(4), users.Id(1)}

	theTable := Table{
		Id:             TableId(123),
		Created:        time.Now(),
		InviteCode:     "123",
		Public:         true,
		Phase:          api.TablePhase_NOT_STARTED,
		Players:        players,
		PlayersInOrder: playersInOrder,
	}
	data := NewTableData(&theTable)
	assert.Equal(t, playersInOrder, data.Players)
}
