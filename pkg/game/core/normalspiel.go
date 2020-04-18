package core

func IsNormalTrump(k Card) bool {
	return k.Suit == Karo || k.Rank == Bube || k.Rank == Dame || k == Dulle()
}

func NormalGameSuit(c Card) GameSuit {
	if IsNormalTrump(c) {
		return Trumpf
	}
	return c.Suit.AsFehl()
}

type NormalspielMode struct{}

func (n NormalspielMode) GameSuit(k Card) GameSuit {
	return NormalGameSuit(k)
}

var Normalspiel = NormalspielMode{}
