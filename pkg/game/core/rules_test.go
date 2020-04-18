package core

import "testing"

func TestWinnerOfTrick(t *testing.T) {
	if WinnerOfTrick(Trick{Player1, [4]Card{
		{Kreuz, Ass},
		{Kreuz, Ass},
		{Kreuz, Zehn},
		{Kreuz, Neun},
	}}, Normalspiel) != Player1 {
		t.Error("player1 sollte gewinnen")
	}
}

func TestGame_IsValidMove(t *testing.T) {
	cards := DealCards(1487)
	game := CreateGame(cards, Player3, Normalspiel)
	for _, card := range game.HandCards[Player3] {
		if !game.IsValidMove(Player3, card) {
			t.Error("forehand player not allowed to play card but should")
		}
	}
	for _, player := range []Player{Player1, Player2, Player4} {
		for _, card := range game.HandCards[player] {
			if game.IsValidMove(player, card) {
				t.Errorf("non-forehand %v allowed to play %v but not her turn", player, card)
			}
		}
	}
}
