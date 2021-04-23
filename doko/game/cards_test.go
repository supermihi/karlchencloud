package game

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sort"
	"testing"
)

func TestHand_CardIndex(t *testing.T) {
	var hand Hand = []Card{Diamonds9, Hearts10, SpadesA, ClubsJ}
	assert.Equal(t, -1, hand.CardIndex(Diamonds10))
	assert.Equal(t, 1, hand.CardIndex(Hearts10))
}

func TestHand_RemoveCard(t *testing.T) {
	var hand Hand = []Card{Hearts10, DiamondsK, SpadesQ, Hearts10}
	newHand := hand[:]
	newHand.RemoveCard(Hearts10)
	var expected Hand = []Card{DiamondsK, SpadesQ, Hearts10}
	assert.ElementsMatch(t, newHand, expected)
}

//noinspection GoNilness
func TestDealCards(t *testing.T) {
	rng := rand.New(rand.NewSource(126795))
	hands := DealCards(rng)
	var dealtDeck []Card
	for _, p := range Players() {
		dealtDeck = append(dealtDeck, hands[p]...)
	}
	sort.Sort(BySuitAndRank(dealtDeck))
	plainDeck := CreateDeck()
	sort.Sort(BySuitAndRank(plainDeck))
	assert.ElementsMatch(t, dealtDeck, plainDeck)
}
