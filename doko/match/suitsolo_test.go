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
	AssertCardIsHigher(t, game.ClubsD, game.Hearts10, clubSolo)
	AssertCardIsHigher(t, game.SpadesD, game.ClubsD, clubSolo)
	AssertCardIsHigher(t, game.HeartsD, game.SpadesD, clubSolo)
	AssertCardIsHigher(t, game.DiamondsD, game.HeartsD, clubSolo)
	AssertCardIsHigher(t, game.ClubsB, game.DiamondsD, clubSolo)
	AssertCardIsHigher(t, game.SpadesB, game.ClubsB, clubSolo)
	AssertCardIsHigher(t, game.HeartsB, game.SpadesB, clubSolo)
	AssertCardIsHigher(t, game.DiamondsB, game.HeartsB, clubSolo)
	AssertCardIsHigher(t, game.ClubsA, game.DiamondsB, clubSolo)
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
	AssertCardIsHigher(t, game.ClubsD, game.Hearts10, spadeSolo)
	AssertCardIsHigher(t, game.SpadesD, game.ClubsD, spadeSolo)
	AssertCardIsHigher(t, game.HeartsD, game.SpadesD, spadeSolo)
	AssertCardIsHigher(t, game.DiamondsD, game.HeartsD, spadeSolo)
	AssertCardIsHigher(t, game.ClubsB, game.DiamondsD, spadeSolo)
	AssertCardIsHigher(t, game.SpadesB, game.ClubsB, spadeSolo)
	AssertCardIsHigher(t, game.HeartsB, game.SpadesB, spadeSolo)
	AssertCardIsHigher(t, game.DiamondsB, game.HeartsB, spadeSolo)
	AssertCardIsHigher(t, game.SpadesA, game.DiamondsB, spadeSolo)
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
	AssertCardIsHigher(t, game.ClubsD, game.Hearts10, heartSolo)
	AssertCardIsHigher(t, game.SpadesD, game.ClubsD, heartSolo)
	AssertCardIsHigher(t, game.HeartsD, game.SpadesD, heartSolo)
	AssertCardIsHigher(t, game.DiamondsD, game.HeartsD, heartSolo)
	AssertCardIsHigher(t, game.ClubsB, game.DiamondsD, heartSolo)
	AssertCardIsHigher(t, game.SpadesB, game.ClubsB, heartSolo)
	AssertCardIsHigher(t, game.HeartsB, game.SpadesB, heartSolo)
	AssertCardIsHigher(t, game.DiamondsB, game.HeartsB, heartSolo)
	AssertCardIsHigher(t, game.HeartsA, game.DiamondsB, heartSolo)
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
	AssertCardIsHigher(t, game.ClubsD, game.Hearts10, diamondSolo)
	AssertCardIsHigher(t, game.SpadesD, game.ClubsD, diamondSolo)
	AssertCardIsHigher(t, game.HeartsD, game.SpadesD, diamondSolo)
	AssertCardIsHigher(t, game.DiamondsD, game.HeartsD, diamondSolo)
	AssertCardIsHigher(t, game.ClubsB, game.DiamondsD, diamondSolo)
	AssertCardIsHigher(t, game.SpadesB, game.ClubsB, diamondSolo)
	AssertCardIsHigher(t, game.HeartsB, game.SpadesB, diamondSolo)
	AssertCardIsHigher(t, game.DiamondsB, game.HeartsB, diamondSolo)
	AssertCardIsHigher(t, game.DiamondsA, game.DiamondsB, diamondSolo)
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
