package game

type HochzeitState int

const (
	WaitingForPartner HochzeitState = iota
	PartnerFound
	Solo
)
const MaxTricksWaitingForPartner = 3

type Hochzeit struct {
	announcer      Player
	partner        Player
	state          HochzeitState
	klaerungsstich int
}

func NewHochzeit(announcer Player) Hochzeit {
	return Hochzeit{announcer, NoPlayer, WaitingForPartner, 0}
}

func (h Hochzeit) OnCompletedTrick(t Trick, numTrick int) {
	if h.state != WaitingForPartner {
		return
	}
	winner := WinnerOfTrick(t.CardsOf, t.Forehand, h)
	if winner != h.announcer {
		h.partner = winner
		h.klaerungsstich = numTrick
	} else if numTrick >= MaxTricksWaitingForPartner-1 {
		h.state = Solo
		h.klaerungsstich = numTrick
	}
}

func (h Hochzeit) PartyOf(p Player) Party {
	if p == h.announcer {
		return ReParty
	}
	if !h.PartnerFound() {
		return NoParty
	}
	if p == h.partner {
		return ReParty
	}
	return ContraParty
}

func (h Hochzeit) PartnerFound() bool {
	return h.partner != NoPlayer
}

func (h Hochzeit) Partner() Player {
	return h.partner
}

// trick-taking rules equal those of normal mode
func (h Hochzeit) GameSuit(card Card) GameSuit {
	return NormalGameSuit(card)
}

func (h Hochzeit) Klaerungsstich() int {
	return h.klaerungsstich
}

func (h Hochzeit) Type() AnnouncedGameType {
	return HochzeitType
}
