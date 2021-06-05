package room

import (
	"github.com/supermihi/karlchencloud/server/table"
	"github.com/supermihi/karlchencloud/server/users"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/supermihi/karlchencloud/api"
)

func TestGetData_PlayersAreInGameOrder(t *testing.T) {
	players := []users.Id{users.Id(1), users.Id(2), users.Id(3), users.Id(4)}
	playersInOrder := []users.Id{users.Id(3), users.Id(2), users.Id(4), users.Id(1)}

	theTable := table.Table{
		Id:             table.Id(123),
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
