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
		game.WinnerOfTrick([4]game.Card{game.HeartsA, game.Hearts10, game.ClubsD, game.DiamondsB}, game.Player1, jacksSolo))

	// Rule 2.3.6
	AssertCardIsHigher(t, game.SpadesB, game.ClubsB, jacksSolo)
	AssertCardIsHigher(t, game.HeartsB, game.SpadesB, jacksSolo)
	AssertCardIsHigher(t, game.DiamondsB, game.HeartsB, jacksSolo)

	AssertCardIsHigher(t, game.Clubs10, game.ClubsA, jacksSolo)
	AssertCardIsHigher(t, game.ClubsK, game.Clubs10, jacksSolo)
	AssertCardIsHigher(t, game.ClubsD, game.ClubsK, jacksSolo)
	AssertCardIsHigher(t, game.Clubs9, game.ClubsD, jacksSolo)

	AssertCardIsHigher(t, game.Spades10, game.SpadesA, jacksSolo)
	AssertCardIsHigher(t, game.SpadesK, game.Spades10, jacksSolo)
	AssertCardIsHigher(t, game.SpadesD, game.SpadesK, jacksSolo)
	AssertCardIsHigher(t, game.Spades9, game.SpadesD, jacksSolo)

	AssertCardIsHigher(t, game.Hearts10, game.HeartsA, jacksSolo)
	AssertCardIsHigher(t, game.HeartsK, game.Hearts10, jacksSolo)
	AssertCardIsHigher(t, game.HeartsD, game.HeartsK, jacksSolo)
	AssertCardIsHigher(t, game.Hearts9, game.HeartsD, jacksSolo)

	AssertCardIsHigher(t, game.Diamonds10, game.DiamondsA, jacksSolo)
	AssertCardIsHigher(t, game.DiamondsK, game.Diamonds10, jacksSolo)
	AssertCardIsHigher(t, game.DiamondsD, game.DiamondsK, jacksSolo)
	AssertCardIsHigher(t, game.Diamonds9, game.DiamondsD, jacksSolo)
}

func TestWinnerOfTrickQueensSolo(t *testing.T) {
	var queensSolo = NewRankSolo(game.Dame, game.Player2)

	// Turnier-Spielregeln (TSR) des Deutschen Doppelkopf-Verbandes e. V., Stand 1. April 2019
	// https://www.doko-verband.de/Regeln__Ordnungen.html?file=tl_files/DDV/Docs/Downloads/Regeln%20und%20Ordnungen/&file=tl_files/DDV/Docs/Downloads/Regeln%20und%20Ordnungen/Turnier-Spielregeln%20Stand%2001.04.2019.pdf

	// Rule 2.3.2
	assert.Equal(t,
		game.Player4,
		game.WinnerOfTrick([4]game.Card{game.HeartsA, game.Hearts10, game.ClubsB, game.DiamondsD}, game.Player1, queensSolo))

	// Rule 2.3.6
	AssertCardIsHigher(t, game.SpadesD, game.ClubsD, queensSolo)
	AssertCardIsHigher(t, game.HeartsD, game.SpadesD, queensSolo)
	AssertCardIsHigher(t, game.DiamondsD, game.HeartsD, queensSolo)

	AssertCardIsHigher(t, game.Clubs10, game.ClubsA, queensSolo)
	AssertCardIsHigher(t, game.ClubsK, game.Clubs10, queensSolo)
	AssertCardIsHigher(t, game.ClubsB, game.ClubsK, queensSolo)
	AssertCardIsHigher(t, game.Clubs9, game.ClubsB, queensSolo)

	AssertCardIsHigher(t, game.Spades10, game.SpadesA, queensSolo)
	AssertCardIsHigher(t, game.SpadesK, game.Spades10, queensSolo)
	AssertCardIsHigher(t, game.SpadesB, game.SpadesK, queensSolo)
	AssertCardIsHigher(t, game.Spades9, game.SpadesB, queensSolo)

	AssertCardIsHigher(t, game.Hearts10, game.HeartsA, queensSolo)
	AssertCardIsHigher(t, game.HeartsK, game.Hearts10, queensSolo)
	AssertCardIsHigher(t, game.HeartsB, game.HeartsK, queensSolo)
	AssertCardIsHigher(t, game.Hearts9, game.HeartsB, queensSolo)

	AssertCardIsHigher(t, game.Diamonds10, game.DiamondsA, queensSolo)
	AssertCardIsHigher(t, game.DiamondsK, game.Diamonds10, queensSolo)
	AssertCardIsHigher(t, game.DiamondsB, game.DiamondsK, queensSolo)
	AssertCardIsHigher(t, game.Diamonds9, game.DiamondsB, queensSolo)
}
