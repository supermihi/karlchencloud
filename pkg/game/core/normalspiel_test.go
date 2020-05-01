package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpielFarbeTrumpf(t *testing.T) {
	cards := []Card{Herz10, KreuzD, KaroA, HerzB}
	for _, card := range cards {
		assert.Equal(t, NormalGameSuit(card), Trumpf)
	}
}

func TestSpielFarbeKaro(t *testing.T) {
	for _, card := range CreateDeck() {
		assert.NotEqual(t, NormalGameSuit(card), KaroFehl)
	}
}

func TestSpielFarbeHerz(t *testing.T) {
	keinHerz := []Card{Herz10, HerzB, KreuzA}
	for _, card := range keinHerz {
		assert.NotEqual(t, NormalGameSuit(card), HerzFehl)
	}
	herz := []Card{Herz9, HerzK, HerzA}
	for _, card := range herz {
		assert.Equal(t, NormalGameSuit(card), HerzFehl)
	}
}

func TestTakesTrickFrom(t *testing.T) {
	sticht := []struct {
		neu Card
		alt Card
	}{
		{Herz10, KreuzD},
		{HerzA, HerzK},
		{Karo9, PikA},
		{Kreuz10, KreuzK},
	}
	for _, neuAlt := range sticht {
		assert.True(t, TakesTrickFrom(neuAlt.neu, neuAlt.alt, someNormalspiel))
	}
	stichtNicht := []struct {
		neu Card
		alt Card
	}{
		{Herz10, Herz10},
		{HerzA, HerzA},
		{Kreuz10, Karo9},
		{HerzA, Herz10},
	}
	for _, neuAlt := range stichtNicht {
		assert.False(t, TakesTrickFrom(neuAlt.neu, neuAlt.alt, someNormalspiel))
	}
}
