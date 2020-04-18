package core

type Mode interface {
	GameSuit(Card) GameSuit
}
