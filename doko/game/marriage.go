package game

type MarriageState int

const (
	WaitingForPartner MarriageState = iota
	PartnerFound
	Solo
)
const MaxTricksWaitingForPartner = 3

type Marriage struct {
	announcer      Player
	partner        Player
	state          MarriageState
	klaerungsstich int
}

func NewMarriage(announcer Player) Marriage {
	return Marriage{announcer, NoPlayer, WaitingForPartner, 0}
}

func (h Marriage) OnCompletedTrick(t Trick, numTrick int) {
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

func (h Marriage) PartyOf(p Player) Party {
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

func (h Marriage) PartnerFound() bool {
	return h.partner != NoPlayer
}

func (h Marriage) Partner() Player {
	return h.partner
}

// trick-taking rules equal those of normal mode
func (h Marriage) GameSuit(card Card) GameSuit {
	return NormalGameSuit(card)
}

func (h Marriage) Klaerungsstich() int {
	return h.klaerungsstich
}

func (h Marriage) Type() AnnouncedGameType {
	return MarriageType
}
