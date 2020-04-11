package game

type GameSuit int

const (
	Trumpf GameSuit = iota
	KaroFehl
	HerzFehl
	PikFehl
	KreuzFehl
)

func (s Suit) AsFehl() GameSuit {
	switch s {
	case Karo:
		return KaroFehl
	case Herz:
		return HerzFehl
	case Pik:
		return PikFehl
	case Kreuz:
		return KreuzFehl
	default:
		panic("unexpected suit")
	}
}
