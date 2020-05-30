package game

import "fmt"

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

func (t AnnouncedGameType) String() string {
	switch t {
	case NormalspielType:
		return "Normal Game"
	case HochzeitType:
		return "Marriage"
	case FleischlosType:
		return "Meatless Solo"
	case KaroSoloType:
		return "Diamonds Solo"
	case HerzSoloType:
		return "Hearts Solo"
	case PikSoloType:
		return "Spades Solo"
	case KreuzSoloType:
		return "Clubs Solo"
	}
	panic(fmt.Sprintf("unexpected game type: %d", int(t)))
}
