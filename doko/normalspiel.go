package doko

type Normalspiel struct{}

func (n Normalspiel) Valid(OpenTrick, Card) bool {
	panic("implement me")
}

func (Normalspiel) WinnerOfTrick(t Trick) Player {
	winner := t.Dealer
	cand := t.Dealer
	for range []int{1, 2, 3} {
		cand := cand.NextPlayer()
		if TakesTrick(t.CardsByPlayer[cand], t.CardsByPlayer[winner]) {
			winner = cand
		}
	}
	return winner
}

func TrumpValue(trumpCard Card) int {
	switch {
	case trumpCard.IsDulle():
		return 1000
	case trumpCard.rank == Queen:
		return 200 + int(trumpCard.suit)
	case trumpCard.rank == Jack:
		return 100 + int(trumpCard.suit)
	default:
		return trumpCard.Value()
	}
}
func TakesTrick(newCard Card, winner Card) bool {
	if winner.IsStandardTrump() {
		return newCard.IsStandardTrump() && TrumpValue(newCard) > TrumpValue(winner)
	}
	return newCard.IsStandardTrump() || newCard.Value() > winner.Value()
}
