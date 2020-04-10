package doko

import "fmt"

type Suit int

const (
	Diamonds Suit = iota
	Hearts
	Spades
	Clubs
)

func (s Suit) String() string {
	return [...]string{"Diamonds", "Hearts", "Spades", "Clubs"}[s]
}

type Rank int

const (
	Nine Rank = iota
	Jack
	Queen
	King
	Ten
	Ace
)

func (s Rank) String() string {
	return [...]string{"Nine", "Jack", "Queen", "King", "Ten", "Ace"}[s]
}

func (s Rank) Value() int {
	return [...]int{0, 2, 3, 4, 10, 11}[s]
}

type Card struct {
	suit Suit
	rank Rank
}

func (c Card) IsDulle() bool {
	return c.suit == Hearts && c.rank == Ten
}

func (c Card) IsStandardTrump() bool {
	return c.suit == Diamonds || c.rank == Jack || c.rank == Queen || c.IsDulle()
}

func (c Card) String() string {
	return fmt.Sprintf("%s %s", c.suit, c.rank)
}

func (c Card) Value() int {
	return c.rank.Value()
}

func CreateDeck() []Card {
	var ans [48]Card
	pos := 0
	for suit := Diamonds; suit <= Clubs; suit++ {
		for rank := Nine; rank <= Ace; rank++ {
			ans[pos] = Card{suit, rank}
			ans[pos+1] = Card{suit, rank}
			pos += 2
		}
	}
	return ans[:]
}
