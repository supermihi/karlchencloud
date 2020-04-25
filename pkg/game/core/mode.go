package core

type Mode interface {
	GameSuit(Card) GameSuit
	PartyOf(Player) Party
	Klaerungsstich() int
	OnCompletedTrick(t Trick, numTrick int)
}
