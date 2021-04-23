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
		game.WinnerOfTrick([4]game.Card{game.DiamondsA, game.HeartsA, game.SpadesA, game.ClubsA}, game.Player1, solo))

	// Rule 2.3.7
	AssertCardIsHigher(t, game.Clubs10, game.ClubsA, solo)
	AssertCardIsHigher(t, game.ClubsK, game.Clubs10, solo)
	AssertCardIsHigher(t, game.ClubsQ, game.ClubsK, solo)
	AssertCardIsHigher(t, game.ClubsJ, game.ClubsQ, solo)
	AssertCardIsHigher(t, game.Clubs9, game.ClubsJ, solo)

	AssertCardIsHigher(t, game.Spades10, game.SpadesA, solo)
	AssertCardIsHigher(t, game.SpadesK, game.Spades10, solo)
	AssertCardIsHigher(t, game.SpadesQ, game.SpadesK, solo)
	AssertCardIsHigher(t, game.SpadesJ, game.SpadesQ, solo)
	AssertCardIsHigher(t, game.Spades9, game.SpadesJ, solo)

	AssertCardIsHigher(t, game.Hearts10, game.HeartsA, solo)
	AssertCardIsHigher(t, game.HeartsK, game.Hearts10, solo)
	AssertCardIsHigher(t, game.HeartsQ, game.HeartsK, solo)
	AssertCardIsHigher(t, game.HeartsJ, game.HeartsQ, solo)
	AssertCardIsHigher(t, game.Hearts9, game.HeartsJ, solo)

	AssertCardIsHigher(t, game.Diamonds10, game.DiamondsA, solo)
	AssertCardIsHigher(t, game.DiamondsK, game.Diamonds10, solo)
	AssertCardIsHigher(t, game.DiamondsQ, game.DiamondsK, solo)
	AssertCardIsHigher(t, game.DiamondsJ, game.DiamondsQ, solo)
	AssertCardIsHigher(t, game.Diamonds9, game.DiamondsJ, solo)
}
