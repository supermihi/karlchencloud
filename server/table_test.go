package server

import (
	"github.com/stretchr/testify/assert"
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/room"
	"testing"
)

func TestTable_Start(t *testing.T) {
	table := room.NewTable(room.UserId(1), room.InvalidTableId, nil, 0)
	_ = table.Join(room.UserId(2))
	_ = table.Join(room.UserId(3))
	_ = table.Join(room.UserId(4))
	_ = table.Start()
	assert.Contains(t, table.CurrentMatch.Players, room.UserId(1))
	assert.Contains(t, table.CurrentMatch.Players, room.UserId(2))
	assert.Contains(t, table.CurrentMatch.Players, room.UserId(3))
	assert.Contains(t, table.CurrentMatch.Players, room.UserId(4))
	assert.Equal(t, api.TablePhase_PLAYING, table.Phase)
}
