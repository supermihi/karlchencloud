package match

import (
	"github.com/stretchr/testify/assert"
	"github.com/supermihi/karlchencloud/doko/game"
	"testing"
)

func Test_getGamePoints(t *testing.T) {
	bids := NewBids()
	bids.placeBid(game.Player1, game.ReParty, Re)
	bids.placeBid(game.Player2, game.ReParty, ReKeine90)
	points := getGamePoints(&bids, game.ReParty, 151, 10)
	expected := []GamePoint{NewGamePoint(Gewonnen), NewGamePoint(ReAngesagt), NewGamePoint(ReKeine90Abgesagt), NewGamePoint(Keine90Gespielt)}
	assert.ElementsMatch(t, expected, points)
}
