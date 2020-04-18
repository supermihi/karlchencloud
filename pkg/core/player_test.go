package core

import "testing"

func TestPlayersFrom(t *testing.T) {
	result := PlayersFrom(Player2)
	if result != [...]Player{Player2, Player3, Player4, Player1} {
		t.Error("unexpected players in order")
	}
}

func TestPlayersFrom2(t *testing.T) {
	result := PlayersFrom(Player1)
	if result != [...]Player{Player1, Player2, Player3, Player4} {
		t.Error("unexpected players in order")
	}
}

func TestPlayersFromFails(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("PlayersFrom did not panic on NoPlayer input")
		}
	}()
	PlayersFrom(NoPlayer)
}
