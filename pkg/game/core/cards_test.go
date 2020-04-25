package core

import (
	"sort"
	"testing"
)

func SameHands(a Hand, b Hand) bool {
	if len(a) != len(b) {
		return false
	}
	for i, aCard := range a {
		if aCard != b[i] {
			return false
		}
	}
	return true
}

func TestHand_RemoveCard(t *testing.T) {
	var hand Hand = []Card{
		{Herz, Zehn},
		{Karo, Koenig},
		{Pik, Dame},
		{Herz, Zehn},
	}
	newHand := hand[:]
	newHand.RemoveCard(Card{Herz, Zehn})
	var expected Hand = []Card{
		{Karo, Koenig},
		{Pik, Dame},
		{Herz, Zehn},
	}
	if !SameHands(newHand, expected) {
		t.Error("unexpected result")
	}

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
	if len(plainDeck) != len(dealtDeck) {
		t.Error("something strange happened")
	}
	for i, plain := range plainDeck {
		if dealtDeck[i] != plain {
			t.Errorf("dealt deck wrong: %v != %v", dealtDeck[i], plain)
		}
	}

}
