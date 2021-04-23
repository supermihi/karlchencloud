package match

import (
	"github.com/stretchr/testify/assert"
	"github.com/supermihi/karlchencloud/doko/game"
	"testing"
)

func TestWinnerOfTrickJacksSolo(t *testing.T) {
	var jacksSolo = NewRankSolo(game.Jack, game.Player1)

	// Turnier-Spielregeln (TSR) des Deutschen Doppelkopf-Verbandes e. V., Stand 1. April 2019
	// https://www.doko-verband.de/Regeln__Ordnungen.html?file=tl_files/DDV/Docs/Downloads/Regeln%20und%20Ordnungen/&file=tl_files/DDV/Docs/Downloads/Regeln%20und%20Ordnungen/Turnier-Spielregeln%20Stand%2001.04.2019.pdf

	// Rule 2.3.2
	assert.Equal(t,
		game.Player4,
		game.WinnerOfTrick([4]game.Card{game.HeartsA, game.Hearts10, game.ClubsQ, game.DiamondsJ}, game.Player1, jacksSolo))

	// Rule 2.3.6
	AssertCardIsHigher(t, game.SpadesJ, game.ClubsJ, jacksSolo)
	AssertCardIsHigher(t, game.HeartsJ, game.SpadesJ, jacksSolo)
	AssertCardIsHigher(t, game.DiamondsJ, game.HeartsJ, jacksSolo)

	AssertCardIsHigher(t, game.Clubs10, game.ClubsA, jacksSolo)
	AssertCardIsHigher(t, game.ClubsK, game.Clubs10, jacksSolo)
	AssertCardIsHigher(t, game.ClubsQ, game.ClubsK, jacksSolo)
	AssertCardIsHigher(t, game.Clubs9, game.ClubsQ, jacksSolo)

	AssertCardIsHigher(t, game.Spades10, game.SpadesA, jacksSolo)
	AssertCardIsHigher(t, game.SpadesK, game.Spades10, jacksSolo)
	AssertCardIsHigher(t, game.SpadesQ, game.SpadesK, jacksSolo)
	AssertCardIsHigher(t, game.Spades9, game.SpadesQ, jacksSolo)

	AssertCardIsHigher(t, game.Hearts10, game.HeartsA, jacksSolo)
	AssertCardIsHigher(t, game.HeartsK, game.Hearts10, jacksSolo)
	AssertCardIsHigher(t, game.HeartsQ, game.HeartsK, jacksSolo)
	AssertCardIsHigher(t, game.Hearts9, game.HeartsQ, jacksSolo)

	AssertCardIsHigher(t, game.Diamonds10, game.DiamondsA, jacksSolo)
	AssertCardIsHigher(t, game.DiamondsK, game.Diamonds10, jacksSolo)
	AssertCardIsHigher(t, game.DiamondsQ, game.DiamondsK, jacksSolo)
	AssertCardIsHigher(t, game.Diamonds9, game.DiamondsQ, jacksSolo)
}

func TestWinnerOfTrickQueensSolo(t *testing.T) {
	var queensSolo = NewRankSolo(game.Queen, game.Player2)

	// Turnier-Spielregeln (TSR) des Deutschen Doppelkopf-Verbandes e. V., Stand 1. April 2019
	// https://www.doko-verband.de/Regeln__Ordnungen.html?file=tl_files/DDV/Docs/Downloads/Regeln%20und%20Ordnungen/&file=tl_files/DDV/Docs/Downloads/Regeln%20und%20Ordnungen/Turnier-Spielregeln%20Stand%2001.04.2019.pdf

	// Rule 2.3.2
	assert.Equal(t,
		game.Player4,
		game.WinnerOfTrick([4]game.Card{game.HeartsA, game.Hearts10, game.ClubsJ, game.DiamondsQ}, game.Player1, queensSolo))

	// Rule 2.3.6
	AssertCardIsHigher(t, game.SpadesQ, game.ClubsQ, queensSolo)
	AssertCardIsHigher(t, game.HeartsQ, game.SpadesQ, queensSolo)
	AssertCardIsHigher(t, game.DiamondsQ, game.HeartsQ, queensSolo)

	AssertCardIsHigher(t, game.Clubs10, game.ClubsA, queensSolo)
	AssertCardIsHigher(t, game.ClubsK, game.Clubs10, queensSolo)
	AssertCardIsHigher(t, game.ClubsJ, game.ClubsK, queensSolo)
	AssertCardIsHigher(t, game.Clubs9, game.ClubsJ, queensSolo)

	AssertCardIsHigher(t, game.Spades10, game.SpadesA, queensSolo)
	AssertCardIsHigher(t, game.SpadesK, game.Spades10, queensSolo)
	AssertCardIsHigher(t, game.SpadesJ, game.SpadesK, queensSolo)
	AssertCardIsHigher(t, game.Spades9, game.SpadesJ, queensSolo)

	AssertCardIsHigher(t, game.Hearts10, game.HeartsA, queensSolo)
	AssertCardIsHigher(t, game.HeartsK, game.Hearts10, queensSolo)
	AssertCardIsHigher(t, game.HeartsJ, game.HeartsK, queensSolo)
	AssertCardIsHigher(t, game.Hearts9, game.HeartsJ, queensSolo)

	AssertCardIsHigher(t, game.Diamonds10, game.DiamondsA, queensSolo)
	AssertCardIsHigher(t, game.DiamondsK, game.Diamonds10, queensSolo)
	AssertCardIsHigher(t, game.DiamondsJ, game.DiamondsK, queensSolo)
	AssertCardIsHigher(t, game.Diamonds9, game.DiamondsJ, queensSolo)
}
