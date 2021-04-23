package match

import (
	"github.com/stretchr/testify/assert"
	"github.com/supermihi/karlchencloud/doko/game"
	"testing"
)

func AssertCardIsHigher(t *testing.T, lower game.Card, higher game.Card, mode game.Mode) {
	assert.Equal(t, game.Player4, game.WinnerOfTrick([4]game.Card{lower, lower, lower, higher}, game.Player1, mode))
}

func TestWinnerOfTrickKreuzSolo(t *testing.T) {
	var kreuzSolo = NewFarbsolo(game.Kreuz, game.Player1)

	// Turnier-Spielregeln (TSR) des Deutschen Doppelkopf-Verbandes e. V., Stand 1. April 2019
	// https://www.doko-verband.de/Regeln__Ordnungen.html?file=tl_files/DDV/Docs/Downloads/Regeln%20und%20Ordnungen/&file=tl_files/DDV/Docs/Downloads/Regeln%20und%20Ordnungen/Turnier-Spielregeln%20Stand%2001.04.2019.pdf

	// Rule 2.3.2
	assert.Equal(t,
		game.Player4,
		game.WinnerOfTrick([4]game.Card{game.KaroA, game.KaroA, game.Karo10, game.Kreuz9}, game.Player1, kreuzSolo))

	// Rule 2.3.5
	AssertCardIsHigher(t, game.KreuzD, game.Herz10, kreuzSolo)
	AssertCardIsHigher(t, game.PikD, game.KreuzD, kreuzSolo)
	AssertCardIsHigher(t, game.HerzD, game.PikD, kreuzSolo)
	AssertCardIsHigher(t, game.KaroD, game.HerzD, kreuzSolo)
	AssertCardIsHigher(t, game.KreuzB, game.KaroD, kreuzSolo)
	AssertCardIsHigher(t, game.PikB, game.KreuzB, kreuzSolo)
	AssertCardIsHigher(t, game.HerzB, game.PikB, kreuzSolo)
	AssertCardIsHigher(t, game.KaroB, game.HerzB, kreuzSolo)
	AssertCardIsHigher(t, game.KreuzA, game.KaroB, kreuzSolo)
	AssertCardIsHigher(t, game.Kreuz10, game.KreuzA, kreuzSolo)
	AssertCardIsHigher(t, game.KreuzK, game.Kreuz10, kreuzSolo)
	AssertCardIsHigher(t, game.Kreuz9, game.KreuzK, kreuzSolo)

	AssertCardIsHigher(t, game.Pik10, game.PikA, kreuzSolo)
	AssertCardIsHigher(t, game.PikK, game.Pik10, kreuzSolo)
	AssertCardIsHigher(t, game.Pik9, game.PikK, kreuzSolo)

	AssertCardIsHigher(t, game.HerzK, game.HerzA, kreuzSolo)
	AssertCardIsHigher(t, game.Herz9, game.HerzK, kreuzSolo)

	AssertCardIsHigher(t, game.Karo10, game.KaroA, kreuzSolo)
	AssertCardIsHigher(t, game.KaroK, game.Karo10, kreuzSolo)
	AssertCardIsHigher(t, game.Karo9, game.KaroK, kreuzSolo)
}

func TestWinnerOfTrickPikSolo(t *testing.T) {
	var pikSolo = NewFarbsolo(game.Pik, game.Player2)

	// Turnier-Spielregeln (TSR) des Deutschen Doppelkopf-Verbandes e. V., Stand 1. April 2019
	// https://www.doko-verband.de/Regeln__Ordnungen.html?file=tl_files/DDV/Docs/Downloads/Regeln%20und%20Ordnungen/&file=tl_files/DDV/Docs/Downloads/Regeln%20und%20Ordnungen/Turnier-Spielregeln%20Stand%2001.04.2019.pdf

	// Rule 2.3.2
	assert.Equal(t,
		game.Player4,
		game.WinnerOfTrick([4]game.Card{game.KaroA, game.KaroA, game.Karo10, game.Pik9}, game.Player1, pikSolo))

	// Rule 2.3.5
	AssertCardIsHigher(t, game.KreuzD, game.Herz10, pikSolo)
	AssertCardIsHigher(t, game.PikD, game.KreuzD, pikSolo)
	AssertCardIsHigher(t, game.HerzD, game.PikD, pikSolo)
	AssertCardIsHigher(t, game.KaroD, game.HerzD, pikSolo)
	AssertCardIsHigher(t, game.KreuzB, game.KaroD, pikSolo)
	AssertCardIsHigher(t, game.PikB, game.KreuzB, pikSolo)
	AssertCardIsHigher(t, game.HerzB, game.PikB, pikSolo)
	AssertCardIsHigher(t, game.KaroB, game.HerzB, pikSolo)
	AssertCardIsHigher(t, game.PikA, game.KaroB, pikSolo)
	AssertCardIsHigher(t, game.Pik10, game.PikA, pikSolo)
	AssertCardIsHigher(t, game.PikK, game.Pik10, pikSolo)
	AssertCardIsHigher(t, game.Pik9, game.PikK, pikSolo)

	AssertCardIsHigher(t, game.Kreuz10, game.KreuzA, pikSolo)
	AssertCardIsHigher(t, game.KreuzK, game.Kreuz10, pikSolo)
	AssertCardIsHigher(t, game.Kreuz9, game.KreuzK, pikSolo)

	AssertCardIsHigher(t, game.HerzK, game.HerzA, pikSolo)
	AssertCardIsHigher(t, game.Herz9, game.HerzK, pikSolo)

	AssertCardIsHigher(t, game.Karo10, game.KaroA, pikSolo)
	AssertCardIsHigher(t, game.KaroK, game.Karo10, pikSolo)
	AssertCardIsHigher(t, game.Karo9, game.KaroK, pikSolo)
}

func TestWinnerOfTrickHerzSolo(t *testing.T) {
	var herzSolo = NewFarbsolo(game.Herz, game.Player3)

	// Turnier-Spielregeln (TSR) des Deutschen Doppelkopf-Verbandes e. V., Stand 1. April 2019
	// https://www.doko-verband.de/Regeln__Ordnungen.html?file=tl_files/DDV/Docs/Downloads/Regeln%20und%20Ordnungen/&file=tl_files/DDV/Docs/Downloads/Regeln%20und%20Ordnungen/Turnier-Spielregeln%20Stand%2001.04.2019.pdf

	// Rule 2.3.2
	assert.Equal(t,
		game.Player4,
		game.WinnerOfTrick([4]game.Card{game.KaroA, game.KaroA, game.Karo10, game.Herz9}, game.Player1, herzSolo))

	// Rule 2.3.5
	AssertCardIsHigher(t, game.KreuzD, game.Herz10, herzSolo)
	AssertCardIsHigher(t, game.PikD, game.KreuzD, herzSolo)
	AssertCardIsHigher(t, game.HerzD, game.PikD, herzSolo)
	AssertCardIsHigher(t, game.KaroD, game.HerzD, herzSolo)
	AssertCardIsHigher(t, game.KreuzB, game.KaroD, herzSolo)
	AssertCardIsHigher(t, game.PikB, game.KreuzB, herzSolo)
	AssertCardIsHigher(t, game.HerzB, game.PikB, herzSolo)
	AssertCardIsHigher(t, game.KaroB, game.HerzB, herzSolo)
	AssertCardIsHigher(t, game.HerzA, game.KaroB, herzSolo)
	AssertCardIsHigher(t, game.HerzK, game.HerzA, herzSolo)
	AssertCardIsHigher(t, game.Herz9, game.HerzK, herzSolo)

	AssertCardIsHigher(t, game.Kreuz10, game.KreuzA, herzSolo)
	AssertCardIsHigher(t, game.KreuzK, game.Kreuz10, herzSolo)
	AssertCardIsHigher(t, game.Kreuz9, game.KreuzK, herzSolo)

	AssertCardIsHigher(t, game.Pik10, game.PikA, herzSolo)
	AssertCardIsHigher(t, game.PikK, game.Pik10, herzSolo)
	AssertCardIsHigher(t, game.Pik9, game.PikK, herzSolo)

	AssertCardIsHigher(t, game.Karo10, game.KaroA, herzSolo)
	AssertCardIsHigher(t, game.KaroK, game.Karo10, herzSolo)
	AssertCardIsHigher(t, game.Karo9, game.KaroK, herzSolo)
}

func TestWinnerOfTrickKaroSolo(t *testing.T) {
	var karoSolo = NewFarbsolo(game.Karo, game.Player4)

	// Turnier-Spielregeln (TSR) des Deutschen Doppelkopf-Verbandes e. V., Stand 1. April 2019
	// https://www.doko-verband.de/Regeln__Ordnungen.html?file=tl_files/DDV/Docs/Downloads/Regeln%20und%20Ordnungen/&file=tl_files/DDV/Docs/Downloads/Regeln%20und%20Ordnungen/Turnier-Spielregeln%20Stand%2001.04.2019.pdf

	// same as in Normalspiel

	// Rule 2.3.2
	assert.Equal(t,
		game.Player4,
		game.WinnerOfTrick([4]game.Card{game.KreuzA, game.KreuzA, game.Kreuz10, game.Karo9}, game.Player1, karoSolo))

	// Rule 2.3.5
	AssertCardIsHigher(t, game.KreuzD, game.Herz10, karoSolo)
	AssertCardIsHigher(t, game.PikD, game.KreuzD, karoSolo)
	AssertCardIsHigher(t, game.HerzD, game.PikD, karoSolo)
	AssertCardIsHigher(t, game.KaroD, game.HerzD, karoSolo)
	AssertCardIsHigher(t, game.KreuzB, game.KaroD, karoSolo)
	AssertCardIsHigher(t, game.PikB, game.KreuzB, karoSolo)
	AssertCardIsHigher(t, game.HerzB, game.PikB, karoSolo)
	AssertCardIsHigher(t, game.KaroB, game.HerzB, karoSolo)
	AssertCardIsHigher(t, game.KaroA, game.KaroB, karoSolo)
	AssertCardIsHigher(t, game.Karo10, game.KaroA, karoSolo)
	AssertCardIsHigher(t, game.KaroK, game.Karo10, karoSolo)
	AssertCardIsHigher(t, game.Karo9, game.KaroK, karoSolo)
	
	AssertCardIsHigher(t, game.Kreuz10, game.KreuzA, karoSolo)
	AssertCardIsHigher(t, game.KreuzK, game.Kreuz10, karoSolo)
	AssertCardIsHigher(t, game.Kreuz9, game.KreuzK, karoSolo)

	AssertCardIsHigher(t, game.Pik10, game.PikA, karoSolo)
	AssertCardIsHigher(t, game.PikK, game.Pik10, karoSolo)
	AssertCardIsHigher(t, game.Pik9, game.PikK, karoSolo)

	AssertCardIsHigher(t, game.HerzK, game.HerzA, karoSolo)
	AssertCardIsHigher(t, game.Herz9, game.HerzK, karoSolo)
}
