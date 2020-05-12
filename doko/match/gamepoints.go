package match

import "github.com/supermihi/karlchencloud/doko/game"

type GamePointType int

const (
	Gewonnen GamePointType = iota
	GegenDieAlten
	Keine90Gespielt
	Keine60Gespielt
	Keine30Gespielt
	SchwarzGespielt
	ReAngesagt
	ContraAngesagt
	ReKeine90Abgesagt
	ReKeine60Abgesagt
	ReKeine30Abgesagt
	ReSchwarzAbgesagt
	ContraKeine90Abgesagt
	ContraKeine60Abgesagt
	ContraKeine30Abgesagt
	ContraSchwarzAbgesagt
	GegenKeine90Erreicht120
	GegenKeine60Erreicht90
	GegenKeine30Erreicht60
	GegenSchwarzErreicht30
)

func pointTypeFromBid(bid Bid) GamePointType {
	switch bid {
	case Re:
		return ReAngesagt
	case Contra:
		return ContraAngesagt
	case ReKeine90:
		return ReKeine90Abgesagt
	case ReKeine60:
		return ReKeine60Abgesagt
	case ReKeine30:
		return ReKeine30Abgesagt
	case ReSchwarz:
		return ReSchwarzAbgesagt
	case ContraKeine90:
		return ContraKeine90Abgesagt
	case ContraKeine60:
		return ContraKeine60Abgesagt
	case ContraKeine30:
		return ContraKeine30Abgesagt
	case ContraSchwarz:
		return ContraSchwarzAbgesagt
	default:
		panic("unexpected bid")
	}
}

type GamePoint struct {
	Type  GamePointType
	Value int
}

func NewGamePoint(typ GamePointType) GamePoint {
	if typ == ReAngesagt || typ == ContraAngesagt {
		return GamePoint{typ, 2}
	}
	return GamePoint{typ, 1}
}

func getGamePoints(bids *Bids, winner game.Party, reScore int, reTricks int) []GamePoint {
	ans := []GamePoint{NewGamePoint(Gewonnen)}
	addPoint := func(t GamePointType) {
		ans = append(ans, NewGamePoint(t))
	}
	if winner == game.NoParty {
		panic("getGamePoints not implemented correctly if no winner exists")
	}
	winnerScore := reScore
	winnerTricks := reTricks
	contraScore := game.TotalScore - reScore
	if winner != game.ReParty {
		addPoint(GegenDieAlten)
		winnerScore = game.TotalScore - reScore
		winnerTricks = game.NumTricks - winnerTricks
	}
	if winnerScore >= 150 {
		addPoint(Keine90Gespielt)
	}
	if winnerScore >= 180 {
		addPoint(Keine60Gespielt)
	}
	if winnerScore >= 210 {
		addPoint(Keine30Gespielt)
	}
	if winnerTricks == game.NumTricks {
		addPoint(SchwarzGespielt)
	}
	for _, b := range bids.AllBids() {
		addPoint(pointTypeFromBid(b))
	}
	if reScore >= 120 && bids.MaxPartyBid(game.ContraParty) >= ContraKeine90 {
		addPoint(GegenKeine90Erreicht120)
	}
	if reScore >= 90 && bids.MaxPartyBid(game.ContraParty) >= ContraKeine60 {
		addPoint(GegenKeine60Erreicht90)
	}
	if reScore >= 60 && bids.MaxPartyBid(game.ContraParty) >= ContraKeine30 {
		addPoint(GegenKeine30Erreicht60)
	}
	if reScore >= 30 && bids.MaxPartyBid(game.ContraParty) >= ContraSchwarz {
		addPoint(GegenSchwarzErreicht30)
	}
	if contraScore >= 120 && bids.MaxPartyBid(game.ReParty) >= ReKeine90 {
		addPoint(GegenKeine90Erreicht120)
	}
	if contraScore >= 90 && bids.MaxPartyBid(game.ReParty) >= ReKeine60 {
		addPoint(GegenKeine60Erreicht90)
	}
	if contraScore >= 60 && bids.MaxPartyBid(game.ReParty) >= ReKeine30 {
		addPoint(GegenKeine30Erreicht60)
	}
	if contraScore >= 30 && bids.MaxPartyBid(game.ReParty) >= ReSchwarz {
		addPoint(GegenSchwarzErreicht30)
	}
	return ans
}
