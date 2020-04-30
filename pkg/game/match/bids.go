package match

import "github.com/supermihi/karlchencloud/pkg/game/core"

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
		return core.ReParty
	}
	return core.ContraParty
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

func NewBids() Bids {
	bids := [core.NumPlayers][]Bid{
		{}, {}, {}, {},
	}
	partyBids := map[core.Party][]Bid{
		core.ReParty:     {},
		core.ContraParty: {},
	}
	return Bids{bids, partyBids}

}

func (bids *Bids) placeBid(player core.Player, party core.Party, bid Bid) {
	bids.bids[player] = append(bids.bids[player], bid)
	bids.partyBids[party] = append(bids.partyBids[party], bid)
}

func (bids *Bids) AllBids() []Bid {
	ans := bids.partyBids[core.ReParty]
	return append(ans, bids.partyBids[core.ContraParty]...)
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
		bids.placeBid(player, g.Mode.PartyOf(player), bid)
		return true
	}
	return false
}
