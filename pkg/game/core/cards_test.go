package core

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestHand_CardIndex(t *testing.T) {
	var hand Hand = []Card{Karo9, Herz10, PikA, KreuzB}
	assert.Equal(t, -1, hand.CardIndex(Karo10))
	assert.Equal(t, 1, hand.CardIndex(Herz10))
}

func TestHand_RemoveCard(t *testing.T) {
	var hand Hand = []Card{Herz10, KaroK, PikD, Herz10}
	newHand := hand[:]
	newHand.RemoveCard(Herz10)
	var expected Hand = []Card{KaroK, PikD, Herz10}
	assert.ElementsMatch(t, newHand, expected)
}

//noinspection GoNilness
func TestDealCards(t *testing.T) {
	hands := DealCards(126795)
	var dealtDeck []Card
	for _, p := range Players() {
		dealtDeck = append(dealtDeck, hands[p]...)
	}
	sort.Sort(BySuitAndRank(dealtDeck))
	plainDeck := CreateDeck()
	sort.Sort(BySuitAndRank(plainDeck))
	assert.ElementsMatch(t, dealtDeck, plainDeck)
}
