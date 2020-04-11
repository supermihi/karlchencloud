package game

import "testing"

func TestStichGewinner(t *testing.T) {
	if WinnerOfTrick(Trick{Player1, [4]Card{
		{Kreuz, Ass},
		{Kreuz, Ass},
		{Kreuz, Zehn},
		{Kreuz, Neun},
	}}, Normalspiel) != Player1 {
		t.Error("player1 sollte gewinnen")
	}
}
