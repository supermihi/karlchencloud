package game

import "testing"

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
	newHand := hand.RemoveCard(Card{Herz, Zehn})
	var expected Hand = []Card{
		{Karo, Koenig},
		{Pik, Dame},
		{Herz, Zehn},
	}
	if !SameHands(newHand, expected) {
		t.Error("unexpected result")
	}

}
