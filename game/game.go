package game

type Mode interface {
	GameSuit(Card) GameSuit
	TakesTrickFrom(neu Card, alt Card) bool
}

type Game struct {
	HandCards      [NumPlayers]Hand
	CompleteTricks []Trick
	CurrentTrick   IncompleteTrick
	Mode           Mode
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

func (game Game) WinnerOfTrick(t Trick) Player {
	return WinnerOfTrick(t, game.Mode)
}

func (game Game) WhoseTurn() Player {
	forehand := game.CurrentTrick.Forehand
	playedCards := len(game.CurrentTrick.CardsInOrder)
	return forehand.NthNext(playedCards)
}

func IsValidMove(g Game, player Player, card Card) bool {
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

func (game Game) PlayCard(player Player, card Card) {
	game.HandCards[player].RemoveCard(card)
	game.CurrentTrick.CardsInOrder = append(game.CurrentTrick.CardsInOrder, card)
}

func (game Game) PerformMove(player Player, card Card) bool {
	if !IsValidMove(game, player, card) {
		return false
	}
	game.PlayCard(player, card)
	game.FinishOpenTrickIfComplete()
	return true
}

func (game Game) FinishOpenTrickIfComplete() {
	currentTrick := game.CurrentTrick
	if !currentTrick.IsComplete() {
		return
	}
	finishedTrick := currentTrick.AsCompleteTrick()
	game.CompleteTricks = append(game.CompleteTricks, finishedTrick)
	game.CurrentTrick = NewIncompleteTrick(game.WinnerOfTrick(finishedTrick))
}
