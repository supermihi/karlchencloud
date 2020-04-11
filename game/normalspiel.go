package game

func NormalTakesTrick(neu Card, alt Card) bool {
	if IsNormalTrump(alt) {
		return NormalGameSuit(neu) == Trumpf && NormalTrickValue(neu) > NormalTrickValue(alt)
	}
	if IsNormalTrump(neu) {
		return true
	}
	if NormalGameSuit(neu) != NormalGameSuit(alt) {
		return false
	}
	return NormalTrickValue(neu) > NormalTrickValue(alt)
}

func IsNormalTrump(k Card) bool {
	return k.Suit == Karo || k.Rank == Bube || k.Rank == Dame || k.IsDulle()
}

func NormalGameSuit(k Card) GameSuit {
	if IsNormalTrump(k) {
		return Trumpf
	}
	return k.Suit.AsFehl()
}

func NormalTrickValue(card Card) int {
	switch {
	case card.IsDulle():
		return 1000
	case card.Rank == Dame:
		return 200 + int(card.Suit)
	case card.Rank == Bube:
		return 100 + int(card.Suit)
	default:
		return card.Value()
	}
}

type NormalspielMode struct{}

func (n NormalspielMode) GameSuit(k Card) GameSuit {
	return NormalGameSuit(k)
}

func (n NormalspielMode) TakesTrickFrom(neu Card, alt Card) bool {
	return NormalTakesTrick(neu, alt)
}

var Normalspiel = NormalspielMode{}
