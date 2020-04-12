package game

import "fmt"

const NumPlayers = 4
const HandSize = 12
const DeckSize = NumPlayers * HandSize

type Suit int

const (
	Karo Suit = iota
	Herz
	Pik
	Kreuz
)

func (s Suit) String() string {
	return [...]string{"♦", "♠", "♠", "♣"}[s]
}

type Rank int

const (
	Neun Rank = iota
	Bube
	Dame
	Koenig
	Zehn
	Ass
)

func (s Rank) String() string {
	return [...]string{"9", "B", "D", "K", "10", "A"}[s]
}

type Card struct {
	Suit Suit
	Rank Rank
}

func (c Card) IsDulle() bool {
	return c.Suit == Herz && c.Rank == Zehn
}

func (c Card) String() string {
	return fmt.Sprintf("%s"+
		"%s", c.Suit, c.Rank)
}

func (c Card) Value() int {
	return c.Rank.Value()
}

func (s Rank) Value() int {
	return [...]int{0, 2, 3, 4, 10, 11}[s]
}
