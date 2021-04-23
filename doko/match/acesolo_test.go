package match

import (
	"github.com/stretchr/testify/assert"
	"github.com/supermihi/karlchencloud/doko/game"
	"testing"
)

func TestWinnerOfTrickAceSolo(t *testing.T) {
	var solo = AceSolo{Soloist: game.Player1}

	// Turnier-Spielregeln (TSR) des Deutschen Doppelkopf-Verbandes e. V., Stand 1. April 2019
	// https://www.doko-verband.de/Regeln__Ordnungen.html?file=tl_files/DDV/Docs/Downloads/Regeln%20und%20Ordnungen/&file=tl_files/DDV/Docs/Downloads/Regeln%20und%20Ordnungen/Turnier-Spielregeln%20Stand%2001.04.2019.pdf

	// Rule 2.3.1
	assert.Equal(t,
		game.Player1,
		game.WinnerOfTrick([4]game.Card{game.KaroA, game.HerzA, game.PikA, game.KreuzA}, game.Player1, solo))

	// Rule 2.3.7
	AssertCardIsHigher(t, game.Kreuz10, game.KreuzA, solo)
	AssertCardIsHigher(t, game.KreuzK, game.Kreuz10, solo)
	AssertCardIsHigher(t, game.KreuzD, game.KreuzK, solo)
	AssertCardIsHigher(t, game.KreuzB, game.KreuzD, solo)
	AssertCardIsHigher(t, game.Kreuz9, game.KreuzB, solo)

	AssertCardIsHigher(t, game.Pik10, game.PikA, solo)
	AssertCardIsHigher(t, game.PikK, game.Pik10, solo)
	AssertCardIsHigher(t, game.PikD, game.PikK, solo)
	AssertCardIsHigher(t, game.PikB, game.PikD, solo)
	AssertCardIsHigher(t, game.Pik9, game.PikB, solo)

	AssertCardIsHigher(t, game.Herz10, game.HerzA, solo)
	AssertCardIsHigher(t, game.HerzK, game.Herz10, solo)
	AssertCardIsHigher(t, game.HerzD, game.HerzK, solo)
	AssertCardIsHigher(t, game.HerzB, game.HerzD, solo)
	AssertCardIsHigher(t, game.Herz9, game.HerzB, solo)

	AssertCardIsHigher(t, game.Karo10, game.KaroA, solo)
	AssertCardIsHigher(t, game.KaroK, game.Karo10, solo)
	AssertCardIsHigher(t, game.KaroD, game.KaroK, solo)
	AssertCardIsHigher(t, game.KaroB, game.KaroD, solo)
	AssertCardIsHigher(t, game.Karo9, game.KaroB, solo)
}
