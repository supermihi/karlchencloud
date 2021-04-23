package match

import (
	"github.com/stretchr/testify/assert"
	"github.com/supermihi/karlchencloud/doko/game"
	"testing"
)

func TestWinnerOfTrickJacksSolo(t *testing.T) {
	var jacksSolo = NewRankSolo(game.Bube, game.Player1)

	// Turnier-Spielregeln (TSR) des Deutschen Doppelkopf-Verbandes e. V., Stand 1. April 2019
	// https://www.doko-verband.de/Regeln__Ordnungen.html?file=tl_files/DDV/Docs/Downloads/Regeln%20und%20Ordnungen/&file=tl_files/DDV/Docs/Downloads/Regeln%20und%20Ordnungen/Turnier-Spielregeln%20Stand%2001.04.2019.pdf

	// Rule 2.3.2
	assert.Equal(t,
		game.Player4,
		game.WinnerOfTrick([4]game.Card{game.HerzA, game.Herz10, game.KreuzD, game.KaroB}, game.Player1, jacksSolo))

	// Rule 2.3.6
	AssertCardIsHigher(t, game.PikB, game.KreuzB, jacksSolo)
	AssertCardIsHigher(t, game.HerzB, game.PikB, jacksSolo)
	AssertCardIsHigher(t, game.KaroB, game.HerzB, jacksSolo)

	AssertCardIsHigher(t, game.Kreuz10, game.KreuzA, jacksSolo)
	AssertCardIsHigher(t, game.KreuzK, game.Kreuz10, jacksSolo)
	AssertCardIsHigher(t, game.KreuzD, game.KreuzK, jacksSolo)
	AssertCardIsHigher(t, game.Kreuz9, game.KreuzD, jacksSolo)

	AssertCardIsHigher(t, game.Pik10, game.PikA, jacksSolo)
	AssertCardIsHigher(t, game.PikK, game.Pik10, jacksSolo)
	AssertCardIsHigher(t, game.PikD, game.PikK, jacksSolo)
	AssertCardIsHigher(t, game.Pik9, game.PikD, jacksSolo)

	AssertCardIsHigher(t, game.Herz10, game.HerzA, jacksSolo)
	AssertCardIsHigher(t, game.HerzK, game.Herz10, jacksSolo)
	AssertCardIsHigher(t, game.HerzD, game.HerzK, jacksSolo)
	AssertCardIsHigher(t, game.Herz9, game.HerzD, jacksSolo)

	AssertCardIsHigher(t, game.Karo10, game.KaroA, jacksSolo)
	AssertCardIsHigher(t, game.KaroK, game.Karo10, jacksSolo)
	AssertCardIsHigher(t, game.KaroD, game.KaroK, jacksSolo)
	AssertCardIsHigher(t, game.Karo9, game.KaroD, jacksSolo)
}

func TestWinnerOfTrickQueensSolo(t *testing.T) {
	var queensSolo = NewRankSolo(game.Dame, game.Player2)

	// Turnier-Spielregeln (TSR) des Deutschen Doppelkopf-Verbandes e. V., Stand 1. April 2019
	// https://www.doko-verband.de/Regeln__Ordnungen.html?file=tl_files/DDV/Docs/Downloads/Regeln%20und%20Ordnungen/&file=tl_files/DDV/Docs/Downloads/Regeln%20und%20Ordnungen/Turnier-Spielregeln%20Stand%2001.04.2019.pdf

	// Rule 2.3.2
	assert.Equal(t,
		game.Player4,
		game.WinnerOfTrick([4]game.Card{game.HerzA, game.Herz10, game.KreuzB, game.KaroD}, game.Player1, queensSolo))

	// Rule 2.3.6
	AssertCardIsHigher(t, game.PikD, game.KreuzD, queensSolo)
	AssertCardIsHigher(t, game.HerzD, game.PikD, queensSolo)
	AssertCardIsHigher(t, game.KaroD, game.HerzD, queensSolo)

	AssertCardIsHigher(t, game.Kreuz10, game.KreuzA, queensSolo)
	AssertCardIsHigher(t, game.KreuzK, game.Kreuz10, queensSolo)
	AssertCardIsHigher(t, game.KreuzB, game.KreuzK, queensSolo)
	AssertCardIsHigher(t, game.Kreuz9, game.KreuzB, queensSolo)

	AssertCardIsHigher(t, game.Pik10, game.PikA, queensSolo)
	AssertCardIsHigher(t, game.PikK, game.Pik10, queensSolo)
	AssertCardIsHigher(t, game.PikB, game.PikK, queensSolo)
	AssertCardIsHigher(t, game.Pik9, game.PikB, queensSolo)

	AssertCardIsHigher(t, game.Herz10, game.HerzA, queensSolo)
	AssertCardIsHigher(t, game.HerzK, game.Herz10, queensSolo)
	AssertCardIsHigher(t, game.HerzB, game.HerzK, queensSolo)
	AssertCardIsHigher(t, game.Herz9, game.HerzB, queensSolo)

	AssertCardIsHigher(t, game.Karo10, game.KaroA, queensSolo)
	AssertCardIsHigher(t, game.KaroK, game.Karo10, queensSolo)
	AssertCardIsHigher(t, game.KaroB, game.KaroK, queensSolo)
	AssertCardIsHigher(t, game.Karo9, game.KaroB, queensSolo)
}
