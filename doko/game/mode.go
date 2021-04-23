package game

import "fmt"

type Mode interface {
	GameSuit(Card) GameSuit
	PartyOf(Player) Party
	ClarificationTrick() int
	Type() AnnouncedGameType
	OnCompletedTrick(t Trick, numTrick int)
}

func IsNormalspiel(t AnnouncedGameType) bool {
	switch t {
	case NormalGameType:
		return true
	case MarriageType:
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
	NormalGameType AnnouncedGameType = iota
	MarriageType
	AceSoloType
	DiamondSoloType
	HeartSoloType
	SpadeSoloType
	ClubSoloType
	JackSoloType
	QueenSoloType
)

func (t AnnouncedGameType) String() string {
	switch t {
	case NormalGameType:
		return "Normal Game"
	case MarriageType:
		return "Marriage"
	case AceSoloType:
		return "Ace Solo"
	case DiamondSoloType:
		return "Diamond Solo"
	case HeartSoloType:
		return "Heart Solo"
	case SpadeSoloType:
		return "Spade Solo"
	case ClubSoloType:
		return "Club Solo"
	case JackSoloType:
		return "Jack Solo"
	case QueenSoloType:
		return "Queen Solo"
	}
	panic(fmt.Sprintf("unexpected game type: %d", int(t)))
}
