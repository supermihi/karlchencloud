package game

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

var someNormalspiel = NormalspielMode{[...]Party{ReParty, ReParty, ContraParty, ContraParty}, false}

func TestWinnerOfTrick(t *testing.T) {
	winner := WinnerOfTrick([4]Card{KreuzA, KreuzA, Kreuz10, Kreuz9}, Player1, someNormalspiel)
	assert.Equal(t, winner, Player1)
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
