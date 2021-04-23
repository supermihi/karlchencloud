package server

import (
	"github.com/stretchr/testify/assert"
	"github.com/supermihi/karlchencloud/api"
	"testing"
)

func TestTable_Start(t *testing.T) {
	table := NewTable("user1", nil, nil, 0)
	_ = table.Join("user2")
	_ = table.Join("user3")
	_ = table.Join("user4")
	_ = table.Start()
	assert.Contains(t, table.CurrentMatch.Players, "user1")
	assert.Contains(t, table.CurrentMatch.Players, "user2")
	assert.Contains(t, table.CurrentMatch.Players, "user3")
	assert.Contains(t, table.CurrentMatch.Players, "user4")
	assert.Equal(t, api.TablePhase_PLAYING, table.Phase)
}
