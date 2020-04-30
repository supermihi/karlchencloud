package match

import (
	"github.com/supermihi/karlchencloud/pkg/game/core"
	"testing"
)

func TestBids_winnerOfGame(t *testing.T) {
	if winnerOfGame(Re, Contra, 120, 17) != core.ContraParty {
		t.Error("ReParty should win with 120 if ReParty and ContraParty announced")
	}
	if winnerOfGame(Re, NoBid, 121, 1) != core.ReParty {
		t.Error("ReParty should win with 121 points if there are no Absagen")
	}
	if winnerOfGame(ReKeine90, ContraKeine90, 129, 1) != core.NoParty {
		t.Error("No party should win if both have unfulfilled Absagen")
	}
}
