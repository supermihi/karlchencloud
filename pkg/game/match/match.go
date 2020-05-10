package match

import (
	"github.com/supermihi/karlchencloud/pkg/game/core"
)

type Match struct {
	auction *Auction
	game    *core.Game
	bids    *Bids
}

func NewMatch(forehand core.Player, sonderspiele Sonderspiele, cards core.Cards) Match {
	auct := NewAuction(forehand, cards, sonderspiele)
	bids := NewBids()
	return Match{auct, nil, &bids}
}

type Phase int

const (
	AuctionAbfragePhase Phase = iota
	AuctionSpezifikationPhase
	GamePhase
	MatchFinished
)

func (m *Match) Phase() Phase {
	switch m.auction.Phase() {
	case VorbehaltAbfrage:
		return AuctionAbfragePhase
	case VorbehaltSpezifikation:
		return AuctionSpezifikationPhase
	}
	if !m.game.IsFinished() {
		return GamePhase
	}
	return MatchFinished
}

func (m *Match) DealtCards() core.Cards {
	return m.auction.Cards
}

func (m *Match) proceedToGame() {
	result := m.auction.GetResult()
	if result.IsSonderspiel {
		m.game = core.NewGame(m.DealtCards(), result.Forehand, result.Sonderspiel)
	} else {
		var spiel core.Mode = core.NewNormalspiel(m.DealtCards())
		for _, player := range core.Players() {
			if m.auction.Cards[player].NumAlte() == 2 {
				spiel = core.StilleHochzeit{Soloist: player}
			}
		}
		m.game = core.NewGame(m.DealtCards(), m.auction.Forehand, spiel)
	}
}

func (m *Match) AnnounceGesundOrVorbehalt(player core.Player, vorbehalt bool) bool {
	if m.Phase() != AuctionAbfragePhase {
		return false
	}
	if m.auction.WhoseTurn() != player {
		return false
	}
	m.auction.Announce(player, vorbehalt)
	if m.auction.Phase() == Finished {
		m.proceedToGame()
	}
	return true
}

func (m *Match) SpecifyVorbehalt(player core.Player, id ModeId) bool {
	if m.Phase() != AuctionSpezifikationPhase {
		return false
	}
	if m.auction.WhoseTurn() != player {
		return false
	}
	result := m.auction.SpecifyVorbehalt(player, id)
	if result != Ok {
		return false
	}
	if m.auction.Phase() == Finished {
		m.proceedToGame()
	}
	return true
}

func (m *Match) PlayCard(player core.Player, card core.Card) bool {
	result := m.game.TryPlayCard(player, card)
	return result == core.CardPlayed
}

func (m *Match) PlaceBid(player core.Player, bid Bid) bool {
	return TryPlaceBid(player, bid, m.bids, m.game)
}

func (m *Match) Evaluate() GameEvaluation {
	if m.Phase() != MatchFinished {
		panic("cannot get points of unfinished match")
	}
	return EvaluateGame(m.game, m.bids)
}

func (m *Match) Mode() core.Mode {
	return m.game.Mode
}
