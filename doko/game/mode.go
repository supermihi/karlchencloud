package game

type Mode interface {
	GameSuit(Card) GameSuit
	PartyOf(Player) Party
	Klaerungsstich() int
	Type() AnnouncedGameType
	OnCompletedTrick(t Trick, numTrick int)
}

func IsNormalspiel(t AnnouncedGameType) bool {
	switch t {
	case NormalspielType:
		return true
	case HochzeitType:
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
