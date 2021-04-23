package game

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

var someNormalspiel = NormalGameMode{[...]Party{ReParty, ReParty, ContraParty, ContraParty}, false}

func AssertCardIsHigher(t *testing.T, lower Card, higher Card) {
	assert.Equal(t, Player4, WinnerOfTrick([4]Card{lower, lower, lower, higher}, Player1, someNormalspiel))
}

func TestWinnerOfTrick(t *testing.T) {
	assert.Equal(t, Player1, WinnerOfTrick([4]Card{ClubsA, ClubsA, Clubs10, Clubs9}, Player1, someNormalspiel))

	// Turnier-Spielregeln (TSR) des Deutschen Doppelkopf-Verbandes e. V., Stand 1. April 2019
	// https://www.doko-verband.de/Regeln__Ordnungen.html?file=tl_files/DDV/Docs/Downloads/Regeln%20und%20Ordnungen/&file=tl_files/DDV/Docs/Downloads/Regeln%20und%20Ordnungen/Turnier-Spielregeln%20Stand%2001.04.2019.pdf

	// Rule 2.3.2
	assert.Equal(t, Player4, WinnerOfTrick([4]Card{ClubsA, ClubsA, HeartsA, Diamonds9}, Player1, someNormalspiel))

	// Rule 2.3.4
	AssertCardIsHigher(t, ClubsD, Hearts10)
	AssertCardIsHigher(t, SpadesD, ClubsD)
	AssertCardIsHigher(t, HeartsD, SpadesD)
	AssertCardIsHigher(t, DiamondsD, HeartsD)
	AssertCardIsHigher(t, ClubsB, DiamondsD)
	AssertCardIsHigher(t, SpadesB, ClubsB)
	AssertCardIsHigher(t, HeartsB, SpadesB)
	AssertCardIsHigher(t, DiamondsB, HeartsB)
	AssertCardIsHigher(t, DiamondsA, DiamondsB)
	AssertCardIsHigher(t, Diamonds10, DiamondsA)
	AssertCardIsHigher(t, DiamondsK, Diamonds10)
	AssertCardIsHigher(t, Diamonds9, DiamondsK)

	AssertCardIsHigher(t, Clubs10, ClubsA)
	AssertCardIsHigher(t, ClubsK, Clubs10)
	AssertCardIsHigher(t, Clubs9, ClubsK)

	AssertCardIsHigher(t, Spades10, SpadesA)
	AssertCardIsHigher(t, SpadesK, Spades10)
	AssertCardIsHigher(t, Spades9, SpadesK)

	AssertCardIsHigher(t, HeartsK, HeartsA)
	AssertCardIsHigher(t, Hearts9, HeartsK)
}

func TestGame_IsValidMove(t *testing.T) {
	rng := rand.New(rand.NewSource(1487))
	cards := DealCards(rng)
	game := NewGame(cards, Player3, someNormalspiel)
	for _, card := range game.HandCards[Player3] {
		assert.Equal(t, game.CanPlayCard(Player3, card), CardPlayed)
	}
	for _, player := range []Player{Player1, Player2, Player4} {
		for _, card := range game.HandCards[player] {
			assert.Equal(t, game.CanPlayCard(player, card), WrongPlayer)
		}
	}
}
