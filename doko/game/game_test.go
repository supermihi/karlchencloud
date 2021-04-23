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
	assert.Equal(t, Player1, WinnerOfTrick([4]Card{KreuzA, KreuzA, Kreuz10, Kreuz9}, Player1, someNormalspiel))

	// Turnier-Spielregeln (TSR) des Deutschen Doppelkopf-Verbandes e. V., Stand 1. April 2019
	// https://www.doko-verband.de/Regeln__Ordnungen.html?file=tl_files/DDV/Docs/Downloads/Regeln%20und%20Ordnungen/&file=tl_files/DDV/Docs/Downloads/Regeln%20und%20Ordnungen/Turnier-Spielregeln%20Stand%2001.04.2019.pdf

	// Rule 2.3.2
	assert.Equal(t, Player4, WinnerOfTrick([4]Card{KreuzA, KreuzA, HerzA, Karo9}, Player1, someNormalspiel))

	// Rule 2.3.4
	AssertCardIsHigher(t, KreuzD, Herz10)
	AssertCardIsHigher(t, PikD, KreuzD)
	AssertCardIsHigher(t, HerzD, PikD)
	AssertCardIsHigher(t, KaroD, HerzD)
	AssertCardIsHigher(t, KreuzB, KaroD)
	AssertCardIsHigher(t, PikB, KreuzB)
	AssertCardIsHigher(t, HerzB, PikB)
	AssertCardIsHigher(t, KaroB, HerzB)
	AssertCardIsHigher(t, KaroA, KaroB)
	AssertCardIsHigher(t, Karo10, KaroA)
	AssertCardIsHigher(t, KaroK, Karo10)
	AssertCardIsHigher(t, Karo9, KaroK)

	AssertCardIsHigher(t, Kreuz10, KreuzA)
	AssertCardIsHigher(t, KreuzK, Kreuz10)
	AssertCardIsHigher(t, Kreuz9, KreuzK)

	AssertCardIsHigher(t, Pik10, PikA)
	AssertCardIsHigher(t, PikK, Pik10)
	AssertCardIsHigher(t, Pik9, PikK)

	AssertCardIsHigher(t, HerzK, HerzA)
	AssertCardIsHigher(t, Herz9, HerzK)
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
