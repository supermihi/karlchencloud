package doko

import "testing"

func TestStichGewinner(t *testing.T) {
	if StichGewinner(Stich{Player1, [4]Karte{
		{Kreuz, Ass},
		{Kreuz, Ass},
		{Kreuz, Zehn},
		{Kreuz, Neun},
	}}, Normalspiel) != Player1 {
		t.Error("player1 sollte gewinnen")
	}
}
