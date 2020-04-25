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

type NormalspielMode struct {
	Parties [NumPlayers]Party
}

func NewNormalspiel(dealtCards Cards) NormalspielMode {
	var parties [NumPlayers]Party
	for _, p := range Players() {
		if dealtCards[p].NumAlte() == 1 {
			parties[p] = ReParty
		} else {
			parties[p] = ContraParty
		}
	}
	return NormalspielMode{parties}
}

func (n NormalspielMode) OnCompletedTrick(Trick, int) {
}

func (NormalspielMode) Klaerungsstich() int {
	return -1
}

func (n NormalspielMode) GameSuit(k Card) GameSuit {
	return NormalGameSuit(k)
}

func (n NormalspielMode) PartyOf(p Player) Party {
	return n.Parties[p]
}
