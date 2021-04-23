package game

import "fmt"

const NumPlayers = 4
const NumHandCards = 12
const NumTricks = NumHandCards
const DeckSize = NumPlayers * NumHandCards
const TotalScore = 240

type Suit int

const (
	Karo Suit = iota
	Herz
	Pik
	Kreuz
)

func (s Suit) String() string {
	return [...]string{"♦", "♥", "♠", "♣"}[s]
}

func (s Suit) ColorCode() int {
	return [...]int{33, 31, 32, 97}[s]
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
	// right padded to two runes
	return [...]string{"9 ", "B ", "D ", "K ", "10", "A "}[s]
}

type Card struct {
	Suit Suit
	Rank Rank
}

func TenOfHearts() Card {
	return Card{Herz, Zehn}
}

func QueenOfClubs() Card {
	return Card{Kreuz, Dame}
}

func (c Card) String() string {
	// ANSI escape sequence for color, Suit, Rank, ANSI escape sequence reset
	return fmt.Sprintf("\033[%dm" + "%s"+ "%s" + "\033[m", c.Suit.ColorCode(), c.Suit, c.Rank)
}

func (c Card) Score() int {
	return c.Rank.Score()
}

func (s Rank) Score() int {
	return [...]int{0, 2, 3, 4, 10, 11}[s]
}

type Party int

const (
	ReParty Party = iota
	ContraParty
	NoParty
)

func (p Party) Other() Party {
	switch p {
	case ReParty:
		return ContraParty
	case ContraParty:
		return ReParty
	default:
		panic("Other() not valid for NoParty")
	}
}

type GameSuit int

const (
	Trump GameSuit = iota
	KaroNonTrump
	HerzNonTrump
	PikNonTrump
	KreuzNonTrump
)

func (s Suit) AsNonTrump() GameSuit {
	switch s {
	case Karo:
		return KaroNonTrump
	case Herz:
		return HerzNonTrump
	case Pik:
		return PikNonTrump
	case Kreuz:
		return KreuzNonTrump
	default:
		panic("unexpected suit")
	}
}
