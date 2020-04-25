package core

import "testing"

var someNormalspiel = NormalspielMode{[...]Party{ReParty, ReParty, ContraParty, ContraParty}}

func TestWinnerOfTrick(t *testing.T) {
	if WinnerOfTrick(Trick{Player1, [4]Card{
		{Kreuz, Ass},
		{Kreuz, Ass},
		{Kreuz, Zehn},
		{Kreuz, Neun},
	}}, someNormalspiel) != Player1 {
		t.Error("player1 sollte gewinnen")
	}
}

func TestGame_IsValidMove(t *testing.T) {
	cards := DealCards(1487)
	game := NewGame(cards, Player3, someNormalspiel)
	for _, card := range game.HandCards[Player3] {
		if game.CanPlayCard(Player3, card) != CardPlayed {
			t.Error("forehand player not allowed to play card but should")
		}
	}
	for _, player := range []Player{Player1, Player2, Player4} {
		for _, card := range game.HandCards[player] {
			if game.CanPlayCard(player, card) != WrongPlayer {
				t.Errorf("non-forehand %v allowed to play %v but not her turn", player, card)
			}
		}
	}
}
