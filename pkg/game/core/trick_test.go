package core

import "testing"

func TestIncompleteTrick_AsCompleteTrick(t *testing.T) {
	it := IncompleteTrick{Player3, []Card{PikB, Karo9, KreuzD, Karo10}}
	ct := it.AsCompleteTrick()
	if ct.Forehand != it.Forehand {
		t.Errorf("forehand %v in complete trick does not macth %v in incomplete", ct.Forehand, it.Forehand)
	}
	if ct.CardsOf != [NumPlayers]Card{KreuzD, Karo10, PikB, Karo9} {
		t.Error("unexpected cards in completed trick")
	}
}
