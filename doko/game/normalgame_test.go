package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGameSuitTrump(t *testing.T) {
	cards := []Card{Hearts10, ClubsD, DiamondsA, HeartsB}
	for _, card := range cards {
		assert.Equal(t, NormalGameSuit(card), Trump)
	}
}

func TestGameSuitDiamonds(t *testing.T) {
	for _, card := range CreateDeck() {
		assert.NotEqual(t, NormalGameSuit(card), DiamondsNonTrump)
	}
}

func TestGameSuitHearts(t *testing.T) {
	noHearts := []Card{Hearts10, HeartsB, ClubsA}
	for _, card := range noHearts {
		assert.NotEqual(t, NormalGameSuit(card), HeartsNonTrump)
	}
	hearts := []Card{Hearts9, HeartsK, HeartsA}
	for _, card := range hearts {
		assert.Equal(t, NormalGameSuit(card), HeartsNonTrump)
	}
}

func TestTakesTrickFrom(t *testing.T) {
	sticht := []struct {
		neu Card
		alt Card
	}{
		{Hearts10, ClubsD},
		{HeartsA, HeartsK},
		{Diamonds9, SpadesA},
		{Clubs10, ClubsK},
	}
	for _, neuAlt := range sticht {
		assert.True(t, TakesTrickFrom(neuAlt.neu, neuAlt.alt, someNormalspiel))
	}
	stichtNicht := []struct {
		neu Card
		alt Card
	}{
		{Hearts10, Hearts10},
		{HeartsA, HeartsA},
		{Clubs10, Diamonds9},
		{HeartsA, Hearts10},
	}
	for _, neuAlt := range stichtNicht {
		assert.False(t, TakesTrickFrom(neuAlt.neu, neuAlt.alt, someNormalspiel))
	}
}
