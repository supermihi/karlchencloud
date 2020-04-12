package game

type Mode interface {
	GameSuit(Card) GameSuit
	TakesTrickFrom(neu Card, alt Card) bool
}

func WinnerOfTrick(t Trick, m Mode) Player {
	winner := t.Forehand
	for i := 1; i < NumPlayers; i++ {
		player := t.Forehand.NthNext(i)
		if m.TakesTrickFrom(t.CardsOf[player], t.CardsOf[winner]) {
			winner = player
		}
	}
	return winner
}
