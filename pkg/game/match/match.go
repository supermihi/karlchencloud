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
	InAuction Phase = iota
	InGame
	MatchFinished
)

func (m *Match) Phase() Phase {
	switch m.auction.Phase() {
	case VorbehaltAbfrage:
	case VorbehaltSpezifikation:
		return InAuction
	}
	if !m.game.IsFinished() {
		return InGame
	}
	return MatchFinished
}

func (m *Match) DealtCards() core.Cards {
	return m.auction.cards
}

func (m *Match) Auction() *Auction {
	return m.auction
}

func (m *Match) proceedToGame() {
	result := m.auction.GetResult()
	if result.IsSonderspiel {
		m.game = core.NewGame(m.DealtCards(), result.Forehand, result.Sonderspiel)
	} else {
		spiel := core.NewNormalspiel(m.DealtCards())
		m.game = core.NewGame(m.DealtCards(), m.auction.forehand, spiel)
	}
}

func (m *Match) AnnounceGesundOrVorbehalt(player core.Player, vorbehalt bool) bool {
	if m.Phase() != InAuction || m.auction.Phase() != VorbehaltAbfrage {
		return false
	}
	if m.auction.WhoseTurn() != player {
		return false
	}
	m.auction.Announce(player, vorbehalt)
	if m.auction.Phase() == AuctionFinished {
		m.proceedToGame()
	}
	return true
}

func (m *Match) SpecifyVorbehalt(player core.Player, id ModeId) bool {
	if m.Phase() != InAuction || m.auction.Phase() != VorbehaltSpezifikation {
		return false
	}
	if m.auction.WhoseTurn() != player {
		return false
	}
	result := m.auction.SpecifyVorbehalt(player, id)
	if result != Ok {
		return false
	}
	if m.auction.Phase() == AuctionFinished {
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

func (m *Match) WhoseTurn() core.Player {
	switch m.Phase() {
	case InAuction:
		return m.auction.WhoseTurn()
	case InGame:
		return m.game.WhoseTurn()
	case MatchFinished:
		return core.NoPlayer
	}
	panic("unexpected match phase")
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
