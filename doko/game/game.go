package game

type Game struct {
	HandCards      Cards
	CompleteTricks []Trick
	CurrentTrick   *IncompleteTrick
	Mode           Mode
}

func NewGame(dealtCards Cards, forehand Player, mode Mode) *Game {
	game := Game{dealtCards, make([]Trick, 0, NumHandCards), NewIncompleteTrick(forehand), mode}
	return &game
}

func (g *Game) NumCompletedTricks() int {
	return len(g.CompleteTricks)
}

func (g *Game) PreviousTrick() *Trick {
	if g.NumCompletedTricks() > 0 {
		return &g.CompleteTricks[g.NumCompletedTricks()-1]
	}
	return nil
}

func (g *Game) IsFinished() bool {
	return g.NumCompletedTricks() == NumHandCards
}
func (g *Game) WinnerOfTrick(cardsOf [NumPlayers]Card, forehand Player) Player {
	return WinnerOfTrick(cardsOf, forehand, g.Mode)
}

func (g *Game) WhoseTurn() Player {
	if g.IsFinished() {
		return NoPlayer
	}
	return g.CurrentTrick.WhoseTurn()
}

func (g *Game) CanPlayCard(player Player, card Card) PlayCardResult {
	if g.IsFinished() {
		return GameFinished
	}
	if g.WhoseTurn() != player {
		return WrongPlayer

	}
	if !g.PlayerHasCard(player, card) {
		return PlayerDoesNotHaveCard
	}
	trick := g.CurrentTrick
	if trick.NumCardsPlayed() == 0 {
		return CardPlayed
	}
	trickSuit := g.Mode.GameSuit(trick.NthCard(0))
	if g.Mode.GameSuit(card) == trickSuit || !g.PlayerHasCardOfSuit(player, trickSuit) {
		return CardPlayed
	}
	return InvalidCard
}

func (g *Game) playCard(player Player, card Card) {
	g.HandCards[player].RemoveCard(card)
	g.CurrentTrick.Play(player, card)
}

type PlayCardResult int

const (
	CardPlayed PlayCardResult = iota
	PlayerDoesNotHaveCard
	WrongPlayer
	InvalidCard
	GameFinished
)

func (g *Game) TryPlayCard(player Player, card Card) PlayCardResult {
	canResult := g.CanPlayCard(player, card)
	if canResult == CardPlayed {
		g.playCard(player, card)
		g.finishOpenTrickIfComplete()
	}
	return canResult
}

func (g *Game) finishOpenTrickIfComplete() {
	currentTrick := g.CurrentTrick
	if !currentTrick.IsComplete() {
		return
	}
	winner := g.WinnerOfTrick(currentTrick.CardsByPlayer(), currentTrick.Forehand)
	finishedTrick := currentTrick.AsCompleteTrick(winner)
	g.Mode.OnCompletedTrick(finishedTrick, len(g.CompleteTricks))
	g.CompleteTricks = append(g.CompleteTricks, finishedTrick)
	g.CurrentTrick = NewIncompleteTrick(winner)
}

func (g *Game) PlayerHasCard(p Player, c Card) bool {
	return g.HandCards[p].ContainsCard(c)
}

func (g *Game) PlayerHasCardOfSuit(p Player, suit GameSuit) bool {
	return AnyCard(g.HandCards[p], func(c Card) bool { return g.Mode.GameSuit(c) == suit })
}

func WinnerOfTrick(cardsOf [NumPlayers]Card, forehand Player, m Mode) Player {
	winner := forehand
	for i := 1; i < NumPlayers; i++ {
		player := forehand.NthNext(i)
		if TakesTrickFrom(cardsOf[player], cardsOf[winner], m) {
			winner = player
		}
	}
	return winner
}

func TakesTrickFrom(new Card, old Card, m Mode) bool {
	newSuit := m.GameSuit(new)
	oldSuit := m.GameSuit(old)
	if newSuit == Trump && oldSuit != Trump {
		return true
	}
	if oldSuit == Trump && newSuit != Trump {
		return false
	}
	if newSuit == Trump {
		// both trump
		return new.TrumpTrickValue() > old.TrumpTrickValue()
	}
	// none trump
	return newSuit == oldSuit && new.NonTrumpTrickValue() > old.NonTrumpTrickValue()
}
