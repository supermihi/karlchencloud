package match

import (
	"github.com/supermihi/doppelgopf/pkg/game/core"
)

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
func WinnerOfGame(reScore int, reTricks int, bids *Bids) core.Party {
	maxContraBid := bids.MaxPartyBid(core.ContraParty)
	maxReBid := bids.MaxPartyBid(core.ReParty)
	return winnerOfGame(maxReBid, maxContraBid, reScore, reTricks)
}
func winnerOfGame(maxReBid Bid, maxContraBid Bid, reScore int, reTricks int) core.Party {
	contraScore := core.TotalScore - reScore

	if maxContraBid <= Contra && maxReBid <= Re {
		// 7.1.2, 1-4 und 7.1.3, 1-4
		if reScore >= 121 {
			return core.ReParty
		}
		// 7.1.3 1, 2 und 3
		if contraScore >= 121 {
			return core.ContraParty
		}
		// both 120
		if maxContraBid == Contra && maxReBid == NoBid {
			return core.ReParty
		}
		return core.ContraParty
	}
	// now: at least one "keine 90"
	// 7.1.2, 5-6
	if maxReBid > Re && maxReBid.WinsOwnAbsage(reScore, reTricks == core.NumTricks) {
		return core.ReParty
	}
	// 7.1.3, 5-6
	if maxContraBid > Contra && maxContraBid.WinsOwnAbsage(contraScore, reTricks > 0) {
		return core.ContraParty
	}
	// 7.1.2, 7-8
	if maxReBid <= Re && maxContraBid > Contra && !maxContraBid.WinsOwnAbsage(contraScore, reTricks > 0) {
		return core.ReParty
	}
	// 7.1.3: 7-8
	if maxContraBid <= Contra && maxReBid > Re && !maxReBid.WinsOwnAbsage(reScore, reTricks == core.NumTricks) {
		return core.ContraParty
	}
	return core.NoParty

}
