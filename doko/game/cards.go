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

func (h Hand) NumQueensOfClubs() int {
	return h.NumberOfCards(QueenOfClubs())
}

func CreateDeck() []Card {
	ans := make([]Card, DeckSize)
	pos := 0
	for suit := Diamonds; suit <= Clubs; suit++ {
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
	if c == TenOfHearts() {
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

var Diamonds9 = Card{Diamonds, Neun}
var DiamondsB = Card{Diamonds, Bube}
var DiamondsD = Card{Diamonds, Dame}
var DiamondsK = Card{Diamonds, Koenig}
var Diamonds10 = Card{Diamonds, Zehn}
var DiamondsA = Card{Diamonds, Ass}

var Hearts9 = Card{Hearts, Neun}
var HeartsB = Card{Hearts, Bube}
var HeartsD = Card{Hearts, Dame}
var HeartsK = Card{Hearts, Koenig}
var Hearts10 = Card{Hearts, Zehn}
var HeartsA = Card{Hearts, Ass}

var Spades9 = Card{Spades, Neun}
var SpadesB = Card{Spades, Bube}
var SpadesD = Card{Spades, Dame}
var SpadesK = Card{Spades, Koenig}
var Spades10 = Card{Spades, Zehn}
var SpadesA = Card{Spades, Ass}

var Clubs9 = Card{Clubs, Neun}
var ClubsB = Card{Clubs, Bube}
var ClubsD = Card{Clubs, Dame}
var ClubsK = Card{Clubs, Koenig}
var Clubs10 = Card{Clubs, Zehn}
var ClubsA = Card{Clubs, Ass}
