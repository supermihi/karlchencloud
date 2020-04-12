package game

type Game struct {
	HandCards      [NumPlayers]Hand
	CompleteTricks []Trick
	CurrentTrick   IncompleteTrick
	Mode           Mode
}

func (g Game) WinnerOfTrick(t Trick) Player {
	return WinnerOfTrick(t, g.Mode)
}

func (g Game) WhoseTurn() Player {
	forehand := g.CurrentTrick.Forehand
	playedCards := len(g.CurrentTrick.CardsInOrder)
	return forehand.NthNext(playedCards)
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

func (g Game) PerformMove(player Player, card Card) bool {
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
