package match

import "github.com/supermihi/karlchencloud/doko/game"

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

func (bid Bid) Party() game.Party {
	if bid <= ReSchwarz {
		return game.ReParty
	}
	return game.ContraParty
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
	bids [game.NumPlayers][]Bid
}

func NewBids() Bids {
	bids := [game.NumPlayers][]Bid{
		{}, {}, {}, {},
	}
	return Bids{bids}

}

func (bids *Bids) placeBid(player game.Player, party game.Party, bid Bid) {
	bids.bids[player] = append(bids.bids[player], bid)
}

func (bids *Bids) BidsOf(p game.Player) []Bid {
	return bids.bids[p]
}

func (bids *Bids) AllBids() []Bid {
	var ans []Bid
	for _, p := range game.Players() {
		ans = append(ans, bids.bids[p]...)
	}
	return ans
}

func (bids *Bids) MaxPartyBid(p game.Party) Bid {
	max := NoBid
	for _, bid := range bids.AllBids() {
		if bid.Party() == p && bid > max {
			max = bid
		}
	}
	return max
}

func playedCardsSinceKlaerung(player game.Player, game *game.Game) int {
	playedCards := game.HandCards[player].NumPlayedCards()
	klearungsstich := game.Mode.Klaerungsstich()
	return playedCards - klearungsstich - 1
}
func CanPlaceBid(player game.Player, bid Bid, bids *Bids, g *game.Game) bool {
	if g.Mode.PartyOf(player) != bid.Party() {
		return false
	}
	if player == game.NoPlayer || bid == NoBid {
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

func TryPlaceBid(player game.Player, bid Bid, bids *Bids, g *game.Game) bool {
	if CanPlaceBid(player, bid, bids, g) {
		bids.placeBid(player, g.Mode.PartyOf(player), bid)
		return true
	}
	return false
}
