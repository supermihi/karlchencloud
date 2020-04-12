package game

import "math/rand"

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

func DealCards(seed int64) [NumPlayers]Hand {
	deck := CreateDeck()
	random := rand.New(rand.NewSource(seed))
	random.Shuffle(len(deck), func(i int, j int) { deck[i], deck[j] = deck[j], deck[i] })
	var ans [NumPlayers]Hand
	for i := 0; i < NumPlayers; i++ {
		ans[i] = deck[i*HandSize : (i+1)*HandSize]
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
