package doko

import "testing"

func TestSpielFarbeTrumpf(t *testing.T) {
	truempfe := []Karte{
		{Herz, Zehn},
		{Kreuz, Dame},
		{Karo, Ass},
		{Herz, Bube},
	}
	for _, trumpf := range truempfe {
		if Normalspiel.SpielFarbe(trumpf) != Trumpf {
			t.Errorf("not trumpf as expected: %v", trumpf)
		}
	}
}

func TestSpielFarbeKaro(t *testing.T) {
	for _, card := range Deck() {
		if Normalspiel.SpielFarbe(card) == KaroFehl {
			t.Error("karo fehl gibt es im NormalspielModus nicht")
		}
	}
}

func TestSpielFarbeHerz(t *testing.T) {
	keinHerz := []Karte{{Herz, Zehn}, {Herz, Bube}, {Kreuz, Ass}}
	for _, card := range keinHerz {
		if Normalspiel.SpielFarbe(card) == HerzFehl {
			t.Errorf("kein Herz fehl: %v", card)
		}
	}
	herz := []Karte{{Herz, Neun}, {Herz, Koenig}, {Herz, Ass}}
	for _, card := range herz {
		if Normalspiel.SpielFarbe(card) != HerzFehl {
			t.Errorf("ist Herz fehl: %v", card)
		}
	}
}

func TestSticht(t *testing.T) {
	sticht := []struct {
		neu Karte
		alt Karte
	}{
		{Karte{Herz, Zehn}, Karte{Kreuz, Dame}},
		{Karte{Herz, Ass}, Karte{Herz, Koenig}},
		{Karte{Karo, Neun}, Karte{Pik, Ass}},
		{Karte{Kreuz, Zehn}, Karte{Kreuz, Koenig}},
	}
	for _, neuAlt := range sticht {
		if !NormalspielSticht(neuAlt.neu, neuAlt.alt) {
			t.Errorf("%v sticht %v", neuAlt.neu, neuAlt.alt)
		}
	}
}
