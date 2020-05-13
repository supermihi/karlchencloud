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
	cardByPlayer map[Player]Card
}

func NewIncompleteTrick(forehand Player) *IncompleteTrick {
	return &IncompleteTrick{forehand, make(map[Player]Card, NumPlayers)}
}

func (t *IncompleteTrick) NumCardsPlayed() int {
	return len(t.cardByPlayer)
}

func (t *IncompleteTrick) NthCard(i int) Card {
	return t.cardByPlayer[t.Forehand.NthNext(i)]
}

func (t *IncompleteTrick) IsComplete() bool {
	return len(t.cardByPlayer) == NumPlayers
}

func (t *IncompleteTrick) CardOf(player Player) (c Card, ok bool) {
	c, ok = t.cardByPlayer[player]
	return
}

func (t *IncompleteTrick) Play(player Player, card Card) {
	t.cardByPlayer[player] = card
}

func (t *IncompleteTrick) CardsByPlayer() [NumPlayers]Card {
	return [...]Card{t.cardByPlayer[Player1], t.cardByPlayer[Player2], t.cardByPlayer[Player3], t.cardByPlayer[Player4]}
}

func (t *IncompleteTrick) AsCompleteTrick(winner Player) Trick {
	return Trick{t.Forehand, t.CardsByPlayer(), winner}
}

func (t *IncompleteTrick) WhoseTurn() Player {
	playedCards := len(t.cardByPlayer)
	return t.Forehand.NthNext(playedCards)
}
