package doko

type Player int

const (
	Player1 Player = iota
	Player2
	Player3
	Player4
)

func (p Player) NextPlayer() Player {
	switch p {
	case Player1:
		return Player2
	case Player2:
		return Player3
	case Player3:
		return Player4
	case Player4:
		return Player1
	default:
		panic("not a valid player")
	}
}
func (p Player) AdvanceBy(i int) Player {
	r := p
	for i > 0 {
		r = r.NextPlayer()
		i -= 1
	}
	return r
}

type Hand []Karte

type GameBasics struct {
	Handkarten [4]Hand
	Vorhand    Player
}

type Stich struct {
	Vorhand  Player
	KarteVon [4]Karte
}

func (t Stich) NthCard(n int) Karte {
	return t.KarteVon[t.Vorhand.AdvanceBy(n)]
}
func (t Stich) FirstCard() Karte {
	return t.NthCard(0)
}

type LaufenderStich struct {
	Vorhand      Player
	CardsInOrder []Karte
}

func (trick LaufenderStich) NthCard(i int) Karte {
	return trick.CardsInOrder[i]
}

type GameProgress struct {
	PlayedTricks []Stich
	CurrentTrick LaufenderStich
}

type GameState struct {
	Basics   GameBasics
	Progress GameProgress
}

type SpielFarbe int

const (
	Trumpf SpielFarbe = iota
	KaroFehl
	HerzFehl
	PikFehl
	KreuzFehl
)

func (s Farbe) AlsFehl() SpielFarbe {
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
		panic("unexpected farbe")
	}
}

type Spielmodus interface {
	SpielFarbe(Karte) SpielFarbe
	Sticht(neu Karte, alt Karte) bool
}

func StichGewinner(s Stich, m Spielmodus) Player {
	winner := s.Vorhand
	for i := range []int{1, 2, 3} {
		player := s.Vorhand.AdvanceBy(i)
		if m.Sticht(s.KarteVon[player], s.KarteVon[winner]) {
			winner = player
		}
	}
	return winner
}
