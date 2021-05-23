package server

import (
	"github.com/stretchr/testify/assert"
	"github.com/supermihi/karlchencloud/api"
	"testing"
)

func TestTable_Start(t *testing.T) {
	table := NewTable(UserId(1), InvalidTableId, nil, 0)
	_ = table.Join(UserId(2))
	_ = table.Join(UserId(3))
	_ = table.Join(UserId(4))
	_ = table.Start()
	assert.Contains(t, table.CurrentMatch.Players, UserId(1))
	assert.Contains(t, table.CurrentMatch.Players, UserId(2))
	assert.Contains(t, table.CurrentMatch.Players, UserId(3))
	assert.Contains(t, table.CurrentMatch.Players, UserId(4))
	assert.Equal(t, api.TablePhase_PLAYING, table.Phase)
}
