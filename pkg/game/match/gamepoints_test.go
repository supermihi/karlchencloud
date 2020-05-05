package match

import (
	"github.com/stretchr/testify/assert"
	"github.com/supermihi/karlchencloud/pkg/game/core"
	"testing"
)

func Test_getGamePoints(t *testing.T) {
	bids := NewBids()
	bids.placeBid(core.Player1, core.ReParty, Re)
	bids.placeBid(core.Player2, core.ReParty, ReKeine90)
	points := getGamePoints(&bids, core.ReParty, 151, 10)
	expected := []GamePoint{NewGamePoint(Gewonnen), NewGamePoint(ReAngesagt), NewGamePoint(ReKeine90Abgesagt), NewGamePoint(Keine90Gespielt)}
	assert.ElementsMatch(t, expected, points)
}
