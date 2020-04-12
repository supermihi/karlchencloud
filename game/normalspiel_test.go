package game

import "testing"

func TestSpielFarbeTrumpf(t *testing.T) {
	truempfe := []Card{
		{Herz, Zehn},
		{Kreuz, Dame},
		{Karo, Ass},
		{Herz, Bube},
	}
	for _, trumpf := range truempfe {
		if NormalGameSuit(trumpf) != Trumpf {
			t.Errorf("not trumpf as expected: %v", trumpf)
		}
	}
}

func TestSpielFarbeKaro(t *testing.T) {
	for _, card := range CreateDeck() {
		if NormalGameSuit(card) == KaroFehl {
			t.Error("karo fehl gibt es im NormalspielMode nicht")
		}
	}
}

func TestSpielFarbeHerz(t *testing.T) {
	keinHerz := []Card{{Herz, Zehn}, {Herz, Bube}, {Kreuz, Ass}}
	for _, card := range keinHerz {
		if NormalGameSuit(card) == HerzFehl {
			t.Errorf("kein Herz fehl: %v", card)
		}
	}
	herz := []Card{{Herz, Neun}, {Herz, Koenig}, {Herz, Ass}}
	for _, card := range herz {
		if NormalGameSuit(card) != HerzFehl {
			t.Errorf("ist Herz fehl: %v", card)
		}
	}
}

func TestSticht(t *testing.T) {
	sticht := []struct {
		neu Card
		alt Card
	}{
		{Card{Herz, Zehn}, Card{Kreuz, Dame}},
		{Card{Herz, Ass}, Card{Herz, Koenig}},
		{Card{Karo, Neun}, Card{Pik, Ass}},
		{Card{Kreuz, Zehn}, Card{Kreuz, Koenig}},
	}
	for _, neuAlt := range sticht {
		if !NormalTakesTrick(neuAlt.neu, neuAlt.alt) {
			t.Errorf("%v sticht %v", neuAlt.neu, neuAlt.alt)
		}
	}
}
