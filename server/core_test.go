package server

import (
	"github.com/stretchr/testify/assert"
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/doko/match"
	"testing"
	"time"
)

func TestGetData_PlayersAreInGameOrder(t *testing.T) {
	table := &Table{"123", time.Now(), "123", api.TablePhase_NOT_STARTED,
		[]string{"p1", "p2", "p3", "p4"}, []string{"p3", "p2", "p4", "p1"},
		match.NewRound(4, 123), nil}
	data := GetData(table)
	assert.Equal(t, []string{"p3", "p2", "p4", "p1"}, data.Players)
}
