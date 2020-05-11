package core

type Mode interface {
	GameSuit(Card) GameSuit
	PartyOf(Player) Party
	Klaerungsstich() int
	Type() AnnouncedGameType
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

func Soloist(m Mode) Player {
	soloist := NoPlayer
	for _, player := range Players() {
		if m.PartyOf(player) == ReParty {
			if soloist != NoPlayer { // second Re player -> no solo
				return NoPlayer
			}
			soloist = player
		}
	}
	return soloist
}

func IsAnnouncedSolo(m Mode) bool {
	return m.Type() != NormalspielType && m.Type() != HochzeitType
}

func IsCountedSolo(m Mode) bool {
	switch u := m.(type) {
	case NormalspielMode:
		return u.stilleHochzeit
	case Hochzeit:
		return false
	default:
		return true
	}
}

type AnnouncedGameType int

const (
	NormalspielType AnnouncedGameType = iota
	HochzeitType
	FleischlosType
	KaroSoloType
	HerzSoloType
	PikSoloType
	KreuzSoloType
)
