package core

type Game struct {
	HandCards      [NumPlayers]Hand
	CompleteTricks []Trick
	CurrentTrick   IncompleteTrick
	Mode           Mode
}

func CreateGame(handCards [NumPlayers]Hand, forehand Player, mode Mode) Game {
	return Game{handCards, make([]Trick, 0, NumHandCards), NewIncompleteTrick(forehand), mode}
}

func (g Game) IsFinished() bool {
	return len(g.CompleteTricks) == NumHandCards
}
func (g Game) WinnerOfTrick(t Trick) Player {
	return WinnerOfTrick(t, g.Mode)
}

func (g Game) WhoseTurn() Player {
	if g.IsFinished() {
		return NoPlayer
	}
	return g.CurrentTrick.WhoseTurn()
}

func (g Game) IsValidMove(player Player, card Card) bool {
	if g.WhoseTurn() != player || !g.PlayerHasCard(player, card) {
		return false
	}
	trick := g.CurrentTrick
	if len(trick.CardsInOrder) == 0 {
		return true
	}
	trickSuit := g.Mode.GameSuit(trick.NthCard(0))
	return g.Mode.GameSuit(card) == trickSuit || !g.PlayerHasCardOfSuit(player, trickSuit)
}

func (g Game) playCard(player Player, card Card) {
	g.HandCards[player].RemoveCard(card)
	g.CurrentTrick.CardsInOrder = append(g.CurrentTrick.CardsInOrder, card)
}

func (g Game) TryPlayCard(player Player, card Card) bool {
	if !g.IsValidMove(player, card) {
		return false
	}
	g.playCard(player, card)
	g.finishOpenTrickIfComplete()
	return true
}

func (g Game) finishOpenTrickIfComplete() {
	currentTrick := g.CurrentTrick
	if !currentTrick.IsComplete() {
		return
	}
	finishedTrick := currentTrick.AsCompleteTrick()
	g.CompleteTricks = append(g.CompleteTricks, finishedTrick)
	g.CurrentTrick = NewIncompleteTrick(g.WinnerOfTrick(finishedTrick))
}

func (g Game) PlayerHasCard(p Player, c Card) bool {
	return g.HandCards[p].ContainsCard(c)
}

func (g Game) PlayerHasCardOfSuit(p Player, suit GameSuit) bool {
	return AnyCard(g.HandCards[p], func(c Card) bool { return g.Mode.GameSuit(c) == suit })
}

func WinnerOfTrick(t Trick, m Mode) Player {
	winner := t.Forehand
	for i := 1; i < NumPlayers; i++ {
		player := t.Forehand.NthNext(i)
		if TakesTrickFrom(t.CardsOf[player], t.CardsOf[winner], m) {
			winner = player
		}
	}
	return winner
}

func TakesTrickFrom(new Card, old Card, m Mode) bool {
	newSuit := m.GameSuit(new)
	oldSuit := m.GameSuit(old)
	if newSuit == Trumpf && oldSuit != Trumpf {
		return true
	}
	if oldSuit == Trumpf && newSuit != Trumpf {
		return false
	}
	if newSuit == Trumpf {
		// both trumpf
		return new.TrumpfTrickValue() > old.TrumpfTrickValue()
	}
	// none trumpf
	return newSuit == oldSuit && new.FehlTrickValue() > old.FehlTrickValue()
}
