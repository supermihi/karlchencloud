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

func Dulle() Card {
	return Card{Herz, Zehn}
}

func Alte() Card {
	return Card{Kreuz, Dame}
}

func (c Card) String() string {
	return fmt.Sprintf("%s"+
		"%s", c.Suit, c.Rank)
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
	Trumpf GameSuit = iota
	KaroFehl
	HerzFehl
	PikFehl
	KreuzFehl
)

func (s Suit) AsFehl() GameSuit {
	switch s {
	case Karo:
		return KaroFehl
	case Herz:
		return HerzFehl
	case Pik:
		return PikFehl
	case Kreuz:
		return KreuzFehl
	default:
		panic("unexpected suit")
	}
}
