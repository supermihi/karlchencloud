package game

import (
	"fmt"
	"math/rand"
	"strings"
)

type Hand []Card
type Cards [NumPlayers]Hand

func (h Hand) CardIndex(card Card) int {
	for i, c := range h {
		if c == card {
			return i
		}
	}
	return -1
}

func (h Hand) String() string {
	cardStrings := make([]string, len(h))
	for i, card := range h {
		// two runes for index (left padded)
		cardStrings[i] = fmt.Sprintf("%2d:%s", i, card.String())
	}
	return strings.Join(cardStrings, " ")
}

func (h *Hand) RemoveCard(c Card) {
	ans := make([]Card, 0, len(*h)-1)
	removed := false
	for _, card := range *h {
		if !removed && card == c {
			removed = true
		} else {
			ans = append(ans, card)
		}
	}
	*h = ans
}

func AnyCard(cards []Card, predicate func(Card) bool) bool {
	for _, card := range cards {
		if predicate(card) {
			return true
		}
	}
	return false
}

func (h Hand) ContainsCard(c Card) bool {
	return AnyCard(h, func(card Card) bool { return c == card })
}

func (h Hand) NumberOfCards(c Card) int {
	ans := 0
	for _, card := range h {
		if card == c {
			ans += 1
		}
	}
	return ans
}

func (h Hand) NumPlayedCards() int {
	return NumHandCards - len(h)
}

func (h Hand) NumAlte() int {
	return h.NumberOfCards(Alte())
}

func CreateDeck() []Card {
	ans := make([]Card, DeckSize)
	pos := 0
	for suit := Karo; suit <= Kreuz; suit++ {
		for rank := Neun; rank <= Ass; rank++ {
			ans[pos] = Card{suit, rank}
			ans[pos+1] = Card{suit, rank}
			pos += 2
		}
	}
	return ans
}

func DealCards(rng *rand.Rand) Cards {
	deck := CreateDeck()
	rng.Shuffle(len(deck), func(i int, j int) { deck[i], deck[j] = deck[j], deck[i] })
	var ans [NumPlayers]Hand
	for i := 0; i < NumPlayers; i++ {
		ans[i] = deck[i*NumHandCards : (i+1)*NumHandCards]
	}
	return ans
}

type BySuitAndRank []Card

func (cards BySuitAndRank) Len() int {
	return len(cards)
}

func (cards BySuitAndRank) Less(i, j int) bool {
	a := cards[i]
	b := cards[j]
	return a.Suit < b.Suit || (a.Suit == b.Suit && a.Rank < b.Rank)
}

func (cards BySuitAndRank) Swap(i, j int) {
	cards[i], cards[j] = cards[j], cards[i]
}

func (c Card) NonTrumpTrickValue() int {
	return c.Rank.Score()
}

func (c Card) TrumpTrickValue() int {
	if c == Dulle() {
		return 100
	}
	switch c.Rank {
	case Dame:
		return 30 + int(c.Suit)
	case Bube:
		return 20 + int(c.Suit)
	default:
		return c.Rank.Score()
	}
}

var Karo9 = Card{Karo, Neun}
var KaroB = Card{Karo, Bube}
var KaroD = Card{Karo, Dame}
var KaroK = Card{Karo, Koenig}
var Karo10 = Card{Karo, Zehn}
var KaroA = Card{Karo, Ass}

var Herz9 = Card{Herz, Neun}
var HerzB = Card{Herz, Bube}
var HerzD = Card{Herz, Dame}
var HerzK = Card{Herz, Koenig}
var Herz10 = Card{Herz, Zehn}
var HerzA = Card{Herz, Ass}

var Pik9 = Card{Pik, Neun}
var PikB = Card{Pik, Bube}
var PikD = Card{Pik, Dame}
var PikK = Card{Pik, Koenig}
var Pik10 = Card{Pik, Zehn}
var PikA = Card{Pik, Ass}

var Kreuz9 = Card{Kreuz, Neun}
var KreuzB = Card{Kreuz, Bube}
var KreuzD = Card{Kreuz, Dame}
var KreuzK = Card{Kreuz, Koenig}
var Kreuz10 = Card{Kreuz, Zehn}
var KreuzA = Card{Kreuz, Ass}
