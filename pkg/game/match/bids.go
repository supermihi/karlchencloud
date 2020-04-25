package match

import "github.com/supermihi/doppelgopf/pkg/game/core"

type Bid int

const (
	NoBid Bid = iota
	Re
	ReKeine90
	ReKeine60
	ReKeine30
	ReSchwarz
	Contra
	ContraKeine90
	ContraKeine60
	ContraKeine30
	ContraSchwarz
)

func (bid Bid) Party() core.Party {
	if bid <= ReSchwarz {
		return core.Re
	}
	return core.Contra
}

func (bid Bid) MaxPlayedCardsAfterKlaerung() int {
	switch bid {
	case Re, Contra:
		return 1
	case ReKeine90, ContraKeine90:
		return 2
	case ReKeine60, ContraKeine60:
		return 3
	case ReKeine30, ContraKeine30:
		return 4
	case ReSchwarz, ContraSchwarz:
		return 5
	default:
		panic("unexpected bid")
	}
}

type Bids struct {
	bids      [core.NumPlayers][]Bid
	partyBids map[core.Party][]Bid
}

func (bids *Bids) placeBid(player core.Player, party core.Party, bid Bid) {
	bids.bids[player] = append(bids.bids[player], bid)
	bids.partyBids[party] = append(bids.partyBids[party], bid)
}

func (bid Bid) WinsOwnAbsage(score int, oppenentHasTricks bool) bool {
	if (bid == ReKeine90 || bid == ContraKeine90) && score >= 151 {
		return true
	}
	if (bid == ReKeine60 || bid == ContraKeine60) && score >= 181 {
		return true
	}
	if (bid == ReKeine30 || bid == ContraKeine30) && score >= 211 {
		return true
	}
	if (bid == ReSchwarz || bid == ContraSchwarz) && !oppenentHasTricks {
		return true
	}
	return false
}

// See: Turnier-Spielordnung, https://www.doko-verband.de/Regeln__Ordnungen.html, 7.1.2 und 7.1.3
func (bids *Bids) WinnerOfGame(reScore int, reTricks int) core.Party {
	maxContraBid := bids.MaxPartyBid(core.Contra)
	maxReBid := bids.MaxPartyBid(core.Re)
	return winnerOfGame(maxReBid, maxContraBid, reScore, reTricks)
}
func winnerOfGame(maxReBid Bid, maxContraBid Bid, reScore int, reTricks int) core.Party {
	contraScore := core.TotalScore - reScore

	if maxContraBid == Contra && maxReBid == Re {
		// 7.1.2, 1-4 und 7.1.3, 1-4
		if reScore >= 121 {
			return core.Re
		}
		// 7.1.3 1, 2 und 3
		if contraScore >= 121 {
			return core.Contra
		}
		if maxContraBid == Contra && maxReBid == NoBid {
			return core.Re
		}
		return core.Contra
	}
	// now: at least one "keine 90"
	// 7.1.2, 5-6
	if maxReBid > Re && maxReBid.WinsOwnAbsage(reScore, reTricks == core.NumTricks) {
		return core.Re
	}
	// 7.1.3, 5-6
	if maxContraBid > Contra && maxContraBid.WinsOwnAbsage(contraScore, reTricks > 0) {
		return core.Contra
	}
	// 7.1.2, 7-8
	if maxReBid <= Re && maxContraBid > Contra && !maxContraBid.WinsOwnAbsage(contraScore, reTricks > 0) {
		return core.Re
	}
	// 7.1.3: 7-8
	if maxContraBid <= Contra && maxReBid > Re && !maxReBid.WinsOwnAbsage(reScore, reTricks == core.NumTricks) {
		return core.Contra
	}
	return core.NoParty

}

func (bids *Bids) MaxPartyBid(p core.Party) Bid {
	max := NoBid
	for _, bid := range bids.partyBids[p] {
		if bid > max {
			max = bid
		}
	}
	return max
}

func playedCardsSinceKlaerung(player core.Player, game *core.Game) int {
	playedCards := game.HandCards[player].NumPlayedCards()
	klearungsstich := game.Mode.Klaerungsstich()
	return playedCards - klearungsstich - 1
}
func CanPlaceBid(player core.Player, bid Bid, bids *Bids, g *core.Game) bool {
	if g.Mode.PartyOf(player) != bid.Party() {
		return false
	}
	if player == core.NoPlayer || bid == NoBid {
		return false
	}
	if bid != bids.MaxPartyBid(g.Mode.PartyOf(player))+1 {
		return false
	}
	if playedCardsSinceKlaerung(player, g) > bid.MaxPlayedCardsAfterKlaerung() {
		return false
	}
	return true

}

func TryPlaceBid(player core.Player, bid Bid, bids *Bids, g *core.Game) bool {
	if CanPlaceBid(player, bid, bids, g) {
		return true
	}
	return false
}
