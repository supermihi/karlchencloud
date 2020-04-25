package core

type StilleHochzeit struct {
	Soloist Player
}

func (StilleHochzeit) OnCompletedTrick(Trick, int) {
}

func (h StilleHochzeit) PartyOf(p Player) Party {
	if p == h.Soloist {
		return ReParty
	}
	return ContraParty
}

func (StilleHochzeit) GameSuit(c Card) GameSuit {
	return NormalGameSuit(c)
}

func (StilleHochzeit) Klaerungsstich() int {
	return -1
}
