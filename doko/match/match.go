package match

import (
	"github.com/supermihi/karlchencloud/doko/game"
)

type Match struct {
	Auction *Auction
	Game    *game.Game
	Bids    *Bids
}

func NewMatch(forehand game.Player, cards game.Cards) Match {
	auct := NewAuction(forehand, cards)
	bids := NewBids()
	return Match{auct, nil, &bids}
}

type Phase int

const (
	InAuction Phase = iota
	InGame
	MatchFinished
)

func (m *Match) InitialForehand() game.Player {
	return m.Auction.forehand
}

func (m *Match) InitialHandCards(p game.Player) game.Hand {
	return m.Auction.cards[p]
}

func (m *Match) Phase() Phase {
	switch {
	case !m.Auction.IsFinished():
		return InAuction
	case !m.Game.IsFinished():
		return InGame
	}
	return MatchFinished
}

func (m *Match) DealtCards() game.Cards {
	return m.Auction.cards
}

func (m *Match) proceedToGame() {
	result := m.Auction.GetResult()
	m.Game = game.NewGame(m.DealtCards(), result.Forehand, result.Mode)
}

func (m *Match) AnnounceGameType(player game.Player, t game.AnnouncedGameType) bool {
	if m.Phase() != InAuction || m.Auction.WhoseTurn() != player {
		return false
	}
	result := m.Auction.Declare(player, t)
	if !result {
		return false
	}
	if m.Auction.IsFinished() {
		m.proceedToGame()
	}
	return true
}

func (m *Match) PlayCard(player game.Player, card game.Card) bool {
	if m.Phase() != InGame {
		return false
	}
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
