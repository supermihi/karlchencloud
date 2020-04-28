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
