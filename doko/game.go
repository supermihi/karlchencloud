package doko

type Player int

const (
	Player1 Player = iota
	Player2
	Player3
	Player4
)

func (p Player) NextPlayer() Player {
	switch p {
	case Player1:
		return Player2
	case Player2:
		return Player3
	case Player3:
		return Player4
	case Player4:
		return Player1
	default:
		panic("not a valid player")
	}
}
func (p Player) AdvanceBy(i int) Player {
	r := p
	for i > 0 {
		r = r.NextPlayer()
		i -= 1
	}
	return r
}

type Hand [12]Card

type GameBasics struct {
	Hands    [4]Hand
	Forehand Player
}

type Trick struct {
	Dealer        Player
	CardsByPlayer [4]Card
}

func (t Trick) NthCard(n int) Card {
	return t.CardsByPlayer[t.Dealer.AdvanceBy(n)]
}
func (t Trick) FirstCard() Card {
	return t.NthCard(0)
}

type OpenTrick struct {
	Dealer       Player
	CardsInOrder []Card
}

type GameProgress struct {
	PlayedTricks []Trick
	CurrentTrick OpenTrick
}

type GameState struct {
	Basics   GameBasics
	Progress GameProgress
}

type Rules interface {
	Valid(OpenTrick, Card) bool
	WinnerOfTrick(Trick) Player
}

type Game struct {
	State GameState
	Rules Rules
}
