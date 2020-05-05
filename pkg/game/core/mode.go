package core

type Mode interface {
	GameSuit(Card) GameSuit
	PartyOf(Player) Party
	Klaerungsstich() int
	OnCompletedTrick(t Trick, numTrick int)
}

func IsNormalspiel(m Mode) bool {
	switch m.(type) {
	case NormalspielMode:
		return true
	case Hochzeit:
		return true
	default:
		return false
	}
}

func IsSolo(m Mode) bool {
	numContra := 0
	for _, player := range Players() {
		if m.PartyOf(player) == ContraParty {
			numContra += 1
		}
	}
	return numContra != 2
}
