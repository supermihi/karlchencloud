package core

type StilleHochzeit struct {
	Soloist Player
}

func (StilleHochzeit) GameSuit(c Card) GameSuit {
	return NormalGameSuit(c)
}
