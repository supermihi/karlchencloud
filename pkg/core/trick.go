package core

type Trick struct {
	Forehand Player
	CardsOf  [NumPlayers]Card
}

func (t Trick) NthCard(n int) Card {
	return t.CardsOf[t.Forehand.NthNext(n)]
}
func (t Trick) FirstCard() Card {
	return t.NthCard(0)
}

type IncompleteTrick struct {
	Forehand     Player
	CardsInOrder []Card
}

func NewIncompleteTrick(forehand Player) IncompleteTrick {
	return IncompleteTrick{forehand, make(Hand, 0, NumPlayers)}
}

func (t IncompleteTrick) NthCard(i int) Card {
	return t.CardsInOrder[i]
}

func (t IncompleteTrick) IsComplete() bool {
	return len(t.CardsInOrder) == NumPlayers
}

func (t IncompleteTrick) cardsByPlayer() [NumPlayers]Card {
	return [NumPlayers]Card{t.NthCard(0), t.NthCard(1), t.NthCard(2), t.NthCard(3)}
}

func (t IncompleteTrick) AsCompleteTrick() Trick {
	return Trick{t.Forehand, t.cardsByPlayer()}
}

func (t IncompleteTrick) WhoseTurn() Player {
	playedCards := len(t.CardsInOrder)
	return t.Forehand.NthNext(playedCards)
}
