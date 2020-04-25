package match

import (
	"github.com/supermihi/doppelgopf/pkg/game/core"
	"testing"
)

func TestBids_winnerOfGame(t *testing.T) {
	if winnerOfGame(Re, Contra, 120, 17) != core.Contra {
		t.Error("Re should win with 120 if Re and Contra announced")
	}
	if winnerOfGame(Re, NoBid, 121, 1) != core.Re {
		t.Error("Re should win with 121 points if there are no Absagen")
	}
	if winnerOfGame(ReKeine90, ContraKeine90, 129, 1) != core.NoParty {
		t.Error("No party should win if both have unfulfilled Absagen")
	}
}
