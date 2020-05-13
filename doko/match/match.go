package match

import (
	"github.com/supermihi/karlchencloud/doko/game"
)

type Match struct {
	Auction *Auction
	Game    *game.Game
	Bids    *Bids
}

func NewMatch(forehand game.Player, sonderspiele Sonderspiele, cards game.Cards) Match {
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
	switch m.Auction.Phase() {
	case VorbehaltAbfrage:
		fallthrough
	case VorbehaltSpezifikation:
		return InAuction
	}
	if !m.Game.IsFinished() {
		return InGame
	}
	return MatchFinished
}

func (m *Match) DealtCards() game.Cards {
	return m.Auction.cards
}

func (m *Match) proceedToGame() {
	result := m.Auction.GetResult()
	if result.IsSonderspiel {
		m.Game = game.NewGame(m.DealtCards(), result.Forehand, result.Sonderspiel)
	} else {
		spiel := game.NewNormalspiel(m.DealtCards())
		m.Game = game.NewGame(m.DealtCards(), m.Auction.forehand, spiel)
	}
}

func (m *Match) AnnounceGesundOrVorbehalt(player game.Player, vorbehalt bool) bool {
	if m.Phase() != InAuction || m.Auction.Phase() != VorbehaltAbfrage {
		return false
	}
	if m.Auction.WhoseTurn() != player {
		return false
	}
	m.Auction.Announce(player, vorbehalt)
	if m.Auction.Phase() == AuctionFinished {
		m.proceedToGame()
	}
	return true
}

func (m *Match) SpecifyVorbehalt(player game.Player, id ModeId) bool {
	if m.Phase() != InAuction || m.Auction.Phase() != VorbehaltSpezifikation {
		return false
	}
	if m.Auction.WhoseTurn() != player {
		return false
	}
	result := m.Auction.SpecifyVorbehalt(player, id)
	if result != Ok {
		return false
	}
	if m.Auction.Phase() == AuctionFinished {
		m.proceedToGame()
	}
	return true
}

func (m *Match) PlayCard(player game.Player, card game.Card) bool {
	result := m.Game.TryPlayCard(player, card)
	return result == game.CardPlayed
}

func (m *Match) PlaceBid(player game.Player, bid Bid) bool {
	return TryPlaceBid(player, bid, m.Bids, m.Game)
}

func (m *Match) WhoseTurn() game.Player {
	switch m.Phase() {
	case InAuction:
		return m.Auction.WhoseTurn()
	case InGame:
		return m.Game.WhoseTurn()
	case MatchFinished:
		return game.NoPlayer
	}
	panic("unexpected match phase")
}

func (m *Match) Evaluate() GameEvaluation {
	if m.Phase() != MatchFinished {
		panic("cannot get points of unfinished match")
	}
	return EvaluateGame(m.Game, m.Bids)
}

func (m *Match) Mode() game.Mode {
	return m.Game.Mode
}
