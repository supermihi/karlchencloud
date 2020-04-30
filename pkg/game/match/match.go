package match

import (
	"fmt"
	"github.com/supermihi/karlchencloud/pkg/game/auction"
	"github.com/supermihi/karlchencloud/pkg/game/core"
)

type Match struct {
	auction *auction.Auction
	game    *core.Game
	bids    *Bids
}

func NewMatch(forehand core.Player, sonderspiele auction.Sonderspiele, cards core.Cards) Match {
	auct := auction.NewAuction(forehand, cards, sonderspiele)
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
	case auction.VorbehaltAbfrage:
		return AuctionAbfragePhase
	case auction.VorbehaltSpezifikation:
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

func (m *Match) ProceedToGame() {
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

type ActionType int

const (
	AnnounceGesund ActionType = iota
	AnnounceVorbehalt
	SpecifyVorbehalt
	PlayCard
)

type PlayerAction struct {
	Player      core.Player
	Type        ActionType
	VorbehaltId auction.ModeId
	Card        core.Card
}

func PlayCardAction(player core.Player, card core.Card) PlayerAction {
	return PlayerAction{Player: player, Type: PlayCard, Card: card}
}
func AnnounceGesundAction(player core.Player) PlayerAction {
	return PlayerAction{Player: player, Type: AnnounceGesund}
}
func AnnounceVorbehaltAction(player core.Player) PlayerAction {
	return PlayerAction{Player: player, Type: AnnounceVorbehalt}
}

type ActionResultType int

const (
	Ok ActionResultType = iota
	WrongPhase
	WrongPlayerTurn
	CannotPlayCard
	OtherError
)

type ActionResult struct {
	Type     ActionResultType
	ErrorMsg string
}

var OkResult = ActionResult{Ok, ""}

func WrongPlayerResult(actual core.Player) ActionResult {
	return ActionResult{WrongPlayerTurn, fmt.Sprintf("it's %v's turn", actual)}
}
func (m *Match) PerformAction(act PlayerAction) ActionResult {
	switch m.Phase() {
	case AuctionAbfragePhase:
		return m.performActionAbfrage(act)
	case AuctionSpezifikationPhase:
		return m.performActionSpezifikation(act)
	case GamePhase:
		return m.performActionGame(act)
	case MatchFinished:
		return m.performActionFinished()
	default:
		panic("unknown match phase")
	}
}

func (m *Match) performActionAbfrage(act PlayerAction) ActionResult {
	if act.Type != AnnounceGesund && act.Type != AnnounceVorbehalt {
		return ActionResult{WrongPhase, "expected gesund/vorbehalt"}
	}
	if m.auction.WhoseTurn() != act.Player {
		return WrongPlayerResult(m.auction.WhoseTurn())
	}
	m.auction.Announce(act.Player, act.Type == AnnounceVorbehalt)
	if m.auction.Phase() == auction.Finished {
		m.ProceedToGame()
	}
	return OkResult
}

func (m *Match) performActionSpezifikation(act PlayerAction) ActionResult {
	if act.Type != SpecifyVorbehalt {
		return ActionResult{WrongPhase, "expected vorbehalt specification"}
	}
	if m.auction.WhoseTurn() != act.Player {
		return WrongPlayerResult(m.auction.WhoseTurn())
	}
	result := m.auction.SpecifyVorbehalt(act.Player, act.VorbehaltId)
	switch result {
	case auction.UnknownSonderspielId:
		return ActionResult{OtherError, fmt.Sprintf("unknown sonderspiel id: %v", act.VorbehaltId)}
	case auction.SonderspielRequirementsUnmet:
		return ActionResult{OtherError, fmt.Sprintf("cannot play %v with your cards", act.VorbehaltId)}
	case auction.Ok:
		if m.auction.Phase() == auction.Finished {
			m.ProceedToGame()
		}
		return OkResult
	}
	panic("should not be here in performActionSpezifikation")
}

func (m *Match) performActionGame(act PlayerAction) ActionResult {
	if act.Type != PlayCard {
		return ActionResult{WrongPhase, "expected card"}
	}
	if m.game.WhoseTurn() != act.Player {
		return WrongPlayerResult(m.game.WhoseTurn())
	}
	if m.game.TryPlayCard(act.Player, act.Card) == core.CardPlayed {
		return OkResult
	} else {
		return ActionResult{CannotPlayCard, fmt.Sprintf("could not play card")}
	}

}

func (m *Match) performActionFinished() ActionResult {
	return ActionResult{WrongPhase, "match has finished"}
}
