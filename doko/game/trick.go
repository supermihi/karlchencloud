package game

type Trick struct {
	Forehand Player
	CardsOf  [NumPlayers]Card
	Winner   Player
}

func (t *Trick) NthCard(n int) Card {
	return t.CardsOf[t.Forehand.NthNext(n)]
}

func (t *Trick) Score() int {
	result := 0
	for _, card := range t.CardsOf {
		result += card.Score()
	}
	return result
}

type IncompleteTrick struct {
	Forehand     Player
	CardsInOrder []Card
}

func NewIncompleteTrick(forehand Player) *IncompleteTrick {
	return &IncompleteTrick{forehand, make(Hand, 0, NumPlayers)}
}

func (t *IncompleteTrick) NthCard(i int) Card {
	return t.CardsInOrder[i]
}

func (t *IncompleteTrick) IsComplete() bool {
	return len(t.CardsInOrder) == NumPlayers
}

func (t *IncompleteTrick) CardOf(player Player) Card {
	return t.CardsInOrder[t.Forehand.HopsTo(player)]
}

func (t *IncompleteTrick) TryCardOf(player Player) (c Card, ok bool) {
	if t.Forehand.HopsTo(player) >= len(t.CardsInOrder) {
		return Card{}, false
	}
	return t.CardOf(player), true
}
func (t *IncompleteTrick) CardsByPlayer() [NumPlayers]Card {
	return [...]Card{t.CardOf(Player1), t.CardOf(Player2), t.CardOf(Player3), t.CardOf(Player4)}
}

func (t *IncompleteTrick) AsCompleteTrick(winner Player) Trick {
	return Trick{t.Forehand, t.CardsByPlayer(), winner}
}

func (t *IncompleteTrick) WhoseTurn() Player {
	playedCards := len(t.CardsInOrder)
	return t.Forehand.NthNext(playedCards)
}
