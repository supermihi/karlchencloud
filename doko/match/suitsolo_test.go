package match

import (
	"github.com/stretchr/testify/assert"
	"github.com/supermihi/karlchencloud/doko/game"
	"testing"
)

func AssertCardIsHigher(t *testing.T, lower game.Card, higher game.Card, mode game.Mode) {
	assert.Equal(t, game.Player4, game.WinnerOfTrick([4]game.Card{lower, lower, lower, higher}, game.Player1, mode))
}

func TestWinnerOfTrickClubSolo(t *testing.T) {
	var clubSolo = NewSuitSolo(game.Clubs, game.Player1)

	// Turnier-Spielregeln (TSR) des Deutschen Doppelkopf-Verbandes e. V., Stand 1. April 2019
	// https://www.doko-verband.de/Regeln__Ordnungen.html?file=tl_files/DDV/Docs/Downloads/Regeln%20und%20Ordnungen/&file=tl_files/DDV/Docs/Downloads/Regeln%20und%20Ordnungen/Turnier-Spielregeln%20Stand%2001.04.2019.pdf

	// Rule 2.3.2
	assert.Equal(t,
		game.Player4,
		game.WinnerOfTrick([4]game.Card{game.DiamondsA, game.DiamondsA, game.Diamonds10, game.Clubs9}, game.Player1, clubSolo))

	// Rule 2.3.5
	AssertCardIsHigher(t, game.ClubsQ, game.Hearts10, clubSolo)
	AssertCardIsHigher(t, game.SpadesQ, game.ClubsQ, clubSolo)
	AssertCardIsHigher(t, game.HeartsQ, game.SpadesQ, clubSolo)
	AssertCardIsHigher(t, game.DiamondsQ, game.HeartsQ, clubSolo)
	AssertCardIsHigher(t, game.ClubsJ, game.DiamondsQ, clubSolo)
	AssertCardIsHigher(t, game.SpadesJ, game.ClubsJ, clubSolo)
	AssertCardIsHigher(t, game.HeartsJ, game.SpadesJ, clubSolo)
	AssertCardIsHigher(t, game.DiamondsJ, game.HeartsJ, clubSolo)
	AssertCardIsHigher(t, game.ClubsA, game.DiamondsJ, clubSolo)
	AssertCardIsHigher(t, game.Clubs10, game.ClubsA, clubSolo)
	AssertCardIsHigher(t, game.ClubsK, game.Clubs10, clubSolo)
	AssertCardIsHigher(t, game.Clubs9, game.ClubsK, clubSolo)

	AssertCardIsHigher(t, game.Spades10, game.SpadesA, clubSolo)
	AssertCardIsHigher(t, game.SpadesK, game.Spades10, clubSolo)
	AssertCardIsHigher(t, game.Spades9, game.SpadesK, clubSolo)

	AssertCardIsHigher(t, game.HeartsK, game.HeartsA, clubSolo)
	AssertCardIsHigher(t, game.Hearts9, game.HeartsK, clubSolo)

	AssertCardIsHigher(t, game.Diamonds10, game.DiamondsA, clubSolo)
	AssertCardIsHigher(t, game.DiamondsK, game.Diamonds10, clubSolo)
	AssertCardIsHigher(t, game.Diamonds9, game.DiamondsK, clubSolo)
}

func TestWinnerOfTrickSpadeSolo(t *testing.T) {
	var spadeSolo = NewSuitSolo(game.Spades, game.Player2)

	// Turnier-Spielregeln (TSR) des Deutschen Doppelkopf-Verbandes e. V., Stand 1. April 2019
	// https://www.doko-verband.de/Regeln__Ordnungen.html?file=tl_files/DDV/Docs/Downloads/Regeln%20und%20Ordnungen/&file=tl_files/DDV/Docs/Downloads/Regeln%20und%20Ordnungen/Turnier-Spielregeln%20Stand%2001.04.2019.pdf

	// Rule 2.3.2
	assert.Equal(t,
		game.Player4,
		game.WinnerOfTrick([4]game.Card{game.DiamondsA, game.DiamondsA, game.Diamonds10, game.Spades9}, game.Player1, spadeSolo))

	// Rule 2.3.5
	AssertCardIsHigher(t, game.ClubsQ, game.Hearts10, spadeSolo)
	AssertCardIsHigher(t, game.SpadesQ, game.ClubsQ, spadeSolo)
	AssertCardIsHigher(t, game.HeartsQ, game.SpadesQ, spadeSolo)
	AssertCardIsHigher(t, game.DiamondsQ, game.HeartsQ, spadeSolo)
	AssertCardIsHigher(t, game.ClubsJ, game.DiamondsQ, spadeSolo)
	AssertCardIsHigher(t, game.SpadesJ, game.ClubsJ, spadeSolo)
	AssertCardIsHigher(t, game.HeartsJ, game.SpadesJ, spadeSolo)
	AssertCardIsHigher(t, game.DiamondsJ, game.HeartsJ, spadeSolo)
	AssertCardIsHigher(t, game.SpadesA, game.DiamondsJ, spadeSolo)
	AssertCardIsHigher(t, game.Spades10, game.SpadesA, spadeSolo)
	AssertCardIsHigher(t, game.SpadesK, game.Spades10, spadeSolo)
	AssertCardIsHigher(t, game.Spades9, game.SpadesK, spadeSolo)

	AssertCardIsHigher(t, game.Clubs10, game.ClubsA, spadeSolo)
	AssertCardIsHigher(t, game.ClubsK, game.Clubs10, spadeSolo)
	AssertCardIsHigher(t, game.Clubs9, game.ClubsK, spadeSolo)

	AssertCardIsHigher(t, game.HeartsK, game.HeartsA, spadeSolo)
	AssertCardIsHigher(t, game.Hearts9, game.HeartsK, spadeSolo)

	AssertCardIsHigher(t, game.Diamonds10, game.DiamondsA, spadeSolo)
	AssertCardIsHigher(t, game.DiamondsK, game.Diamonds10, spadeSolo)
	AssertCardIsHigher(t, game.Diamonds9, game.DiamondsK, spadeSolo)
}

func TestWinnerOfTrickHeartSolo(t *testing.T) {
	var heartSolo = NewSuitSolo(game.Hearts, game.Player3)

	// Turnier-Spielregeln (TSR) des Deutschen Doppelkopf-Verbandes e. V., Stand 1. April 2019
	// https://www.doko-verband.de/Regeln__Ordnungen.html?file=tl_files/DDV/Docs/Downloads/Regeln%20und%20Ordnungen/&file=tl_files/DDV/Docs/Downloads/Regeln%20und%20Ordnungen/Turnier-Spielregeln%20Stand%2001.04.2019.pdf

	// Rule 2.3.2
	assert.Equal(t,
		game.Player4,
		game.WinnerOfTrick([4]game.Card{game.DiamondsA, game.DiamondsA, game.Diamonds10, game.Hearts9}, game.Player1, heartSolo))

	// Rule 2.3.5
	AssertCardIsHigher(t, game.ClubsQ, game.Hearts10, heartSolo)
	AssertCardIsHigher(t, game.SpadesQ, game.ClubsQ, heartSolo)
	AssertCardIsHigher(t, game.HeartsQ, game.SpadesQ, heartSolo)
	AssertCardIsHigher(t, game.DiamondsQ, game.HeartsQ, heartSolo)
	AssertCardIsHigher(t, game.ClubsJ, game.DiamondsQ, heartSolo)
	AssertCardIsHigher(t, game.SpadesJ, game.ClubsJ, heartSolo)
	AssertCardIsHigher(t, game.HeartsJ, game.SpadesJ, heartSolo)
	AssertCardIsHigher(t, game.DiamondsJ, game.HeartsJ, heartSolo)
	AssertCardIsHigher(t, game.HeartsA, game.DiamondsJ, heartSolo)
	AssertCardIsHigher(t, game.HeartsK, game.HeartsA, heartSolo)
	AssertCardIsHigher(t, game.Hearts9, game.HeartsK, heartSolo)

	AssertCardIsHigher(t, game.Clubs10, game.ClubsA, heartSolo)
	AssertCardIsHigher(t, game.ClubsK, game.Clubs10, heartSolo)
	AssertCardIsHigher(t, game.Clubs9, game.ClubsK, heartSolo)

	AssertCardIsHigher(t, game.Spades10, game.SpadesA, heartSolo)
	AssertCardIsHigher(t, game.SpadesK, game.Spades10, heartSolo)
	AssertCardIsHigher(t, game.Spades9, game.SpadesK, heartSolo)

	AssertCardIsHigher(t, game.Diamonds10, game.DiamondsA, heartSolo)
	AssertCardIsHigher(t, game.DiamondsK, game.Diamonds10, heartSolo)
	AssertCardIsHigher(t, game.Diamonds9, game.DiamondsK, heartSolo)
}

func TestWinnerOfTrickDiamondSolo(t *testing.T) {
	var diamondSolo = NewSuitSolo(game.Diamonds, game.Player4)

	// Turnier-Spielregeln (TSR) des Deutschen Doppelkopf-Verbandes e. V., Stand 1. April 2019
	// https://www.doko-verband.de/Regeln__Ordnungen.html?file=tl_files/DDV/Docs/Downloads/Regeln%20und%20Ordnungen/&file=tl_files/DDV/Docs/Downloads/Regeln%20und%20Ordnungen/Turnier-Spielregeln%20Stand%2001.04.2019.pdf

	// same as in Normalspiel

	// Rule 2.3.2
	assert.Equal(t,
		game.Player4,
		game.WinnerOfTrick([4]game.Card{game.ClubsA, game.ClubsA, game.Clubs10, game.Diamonds9}, game.Player1, diamondSolo))

	// Rule 2.3.5
	AssertCardIsHigher(t, game.ClubsQ, game.Hearts10, diamondSolo)
	AssertCardIsHigher(t, game.SpadesQ, game.ClubsQ, diamondSolo)
	AssertCardIsHigher(t, game.HeartsQ, game.SpadesQ, diamondSolo)
	AssertCardIsHigher(t, game.DiamondsQ, game.HeartsQ, diamondSolo)
	AssertCardIsHigher(t, game.ClubsJ, game.DiamondsQ, diamondSolo)
	AssertCardIsHigher(t, game.SpadesJ, game.ClubsJ, diamondSolo)
	AssertCardIsHigher(t, game.HeartsJ, game.SpadesJ, diamondSolo)
	AssertCardIsHigher(t, game.DiamondsJ, game.HeartsJ, diamondSolo)
	AssertCardIsHigher(t, game.DiamondsA, game.DiamondsJ, diamondSolo)
	AssertCardIsHigher(t, game.Diamonds10, game.DiamondsA, diamondSolo)
	AssertCardIsHigher(t, game.DiamondsK, game.Diamonds10, diamondSolo)
	AssertCardIsHigher(t, game.Diamonds9, game.DiamondsK, diamondSolo)
	
	AssertCardIsHigher(t, game.Clubs10, game.ClubsA, diamondSolo)
	AssertCardIsHigher(t, game.ClubsK, game.Clubs10, diamondSolo)
	AssertCardIsHigher(t, game.Clubs9, game.ClubsK, diamondSolo)

	AssertCardIsHigher(t, game.Spades10, game.SpadesA, diamondSolo)
	AssertCardIsHigher(t, game.SpadesK, game.Spades10, diamondSolo)
	AssertCardIsHigher(t, game.Spades9, game.SpadesK, diamondSolo)

	AssertCardIsHigher(t, game.HeartsK, game.HeartsA, diamondSolo)
	AssertCardIsHigher(t, game.Hearts9, game.HeartsK, diamondSolo)
}
