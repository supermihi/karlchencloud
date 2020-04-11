package game

import "fmt"

type Suit int

const (
	Karo Suit = iota
	Herz
	Pik
	Kreuz
)

func (s Suit) String() string {
	return [...]string{"Karo", "Herz", "Pik", "Kreuz"}[s]
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
	return [...]string{"Neun", "Bube", "Dame", "Koenig", "Zehn", "Ass"}[s]
}

func (s Rank) Value() int {
	return [...]int{0, 2, 3, 4, 10, 11}[s]
}

type Card struct {
	Suit Suit
	Rank Rank
}

func (c Card) IsDulle() bool {
	return c.Suit == Herz && c.Rank == Zehn
}

func (c Card) String() string {
	return fmt.Sprintf("%s %s", c.Suit, c.Rank)
}

func (c Card) Value() int {
	return c.Rank.Value()
}

func Deck() []Card {
	var ans [48]Card
	pos := 0
	for suit := Karo; suit <= Kreuz; suit++ {
		for rank := Neun; rank <= Ass; rank++ {
			ans[pos] = Card{suit, rank}
			ans[pos+1] = Card{suit, rank}
			pos += 2
		}
	}
	return ans[:]
}

type Hand []Card

func (h Hand) CardIndex(card Card) int {
	for i, c := range h {
		if c == card {
			return i
		}
	}
	return -1
}

func (h Hand) RemoveCard(c Card) Hand {
	ans := make([]Card, 0, len(h)-1)
	removed := false
	for _, card := range h {
		if !removed && card == c {
			removed = true
		} else {
			ans = append(ans, card)
		}
	}
	return ans
}

type CardPredicate func(Card) bool

func AnyCard(cards []Card, predicate CardPredicate) bool {
	for _, card := range cards {
		if predicate(card) {
			return true
		}
	}
	return false
}

func (game Game) PlayerHasCard(p Player, c Card) bool {
	return AnyCard(game.HandCards[p], func(card Card) bool { return c == card })
}

func (game Game) PlayerHasCardOfSuit(p Player, suit GameSuit) bool {
	return AnyCard(game.HandCards[p], func(c Card) bool { return game.Mode.GameSuit(c) == suit })
}
