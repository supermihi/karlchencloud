package core

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

// trick-taking rules equal those of normal mode
func (h Hochzeit) GameSuit(card Card) GameSuit {
	return NormalGameSuit(card)
}

func (h Hochzeit) Klaerungsstich() int {
	return h.klaerungsstich
}

type StilleHochzeit struct {
	Soloist Player
}

func (StilleHochzeit) OnCompletedTrick(Trick, int) {
}

func (h StilleHochzeit) PartyOf(p Player) Party {
	if p == h.Soloist {
		return ReParty
	}
	return ContraParty
}

func (StilleHochzeit) GameSuit(c Card) GameSuit {
	return NormalGameSuit(c)
}

func (StilleHochzeit) Klaerungsstich() int {
	return -1
}
