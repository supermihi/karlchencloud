package tables

import (
	"github.com/stretchr/testify/assert"
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/server/users"
	"testing"
)

func TestTable_Start(t *testing.T) {
	table := Initialize(users.Id(1), true, 0)
	_ = table.Join(users.Id(2))
	_ = table.Join(users.Id(3))
	_ = table.Join(users.Id(4))
	_ = table.Start()
	assert.Contains(t, table.CurrentMatch.Players, users.Id(1))
	assert.Contains(t, table.CurrentMatch.Players, users.Id(2))
	assert.Contains(t, table.CurrentMatch.Players, users.Id(3))
	assert.Contains(t, table.CurrentMatch.Players, users.Id(4))
	assert.Equal(t, api.TablePhase_PLAYING, table.Phase)
}
