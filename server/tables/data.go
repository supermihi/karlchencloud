package tables

import (
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
	u "github.com/supermihi/karlchencloud/server/users"
	"time"
)

// This file contains data types returned from the public interface of Tables.

// TableData gives information on a table (not including match data)
type TableData struct {
	Id         TableId
	Owner      u.Id
	InviteCode string
	Players    []u.Id
	Phase      api.TablePhase
	Created    time.Time
	Public     bool
}

// NewTableData extracts TableData from a given Table.
func NewTableData(t *Table) *TableData {
	players := t.PlayersInOrder
	if t.PlayersInOrder == nil {
		players = t.Players
	}
	return &TableData{
		Id:         t.Id,
		Owner:      t.Owner(),
		InviteCode: t.InviteCode,
		Players:    players,
		Phase:      t.Phase,
		Created:    t.Created,
		Public:     t.Public,
	}
}

type DeclarationData struct {
	Healthy     bool
	Reservation game.AnnouncedGameType
}

// MatchData contains data about the current state of a match.
type MatchData struct {
	Phase           match.Phase
	Turn            game.Player
	Players         [game.NumPlayers]u.Id
	InitialForehand game.Player
	Cards           game.Cards
	Declarations    map[game.Player]DeclarationData
	Bids            match.Bids
	CompletedTricks int
	CurrentTrick    *game.IncompleteTrick
	PreviousTrick   *game.Trick
	Mode            game.Mode
	Evaluation      *match.GameEvaluation
}

func NewMatchData(tm *TableMatch) *MatchData {
	decls := tm.Match.Auction.Declarations
	declarations := make(map[game.Player]DeclarationData, len(decls))
	for k, v := range decls {
		if v.Healthy {
			declarations[k] = DeclarationData{true, game.NormalGameType}
		} else {
			declarations[k] = DeclarationData{false, v.Reservation.Type()}
		}
	}
	ans := &MatchData{Phase: tm.Match.Phase(), Turn: tm.Match.WhoseTurn(),
		Players: tm.Players, InitialForehand: tm.Match.InitialForehand(), Declarations: declarations,
		Bids: *tm.Match.Bids}
	switch tm.Match.Phase() {
	case match.InGame:
		g := tm.Match.Game
		ans.CompletedTricks = g.NumCompletedTricks()
		if g.PreviousTrick() != nil {
			tmp := *g.PreviousTrick()
			ans.PreviousTrick = &tmp
		}
		tmp := *g.CurrentTrick
		ans.CurrentTrick = &tmp
		ans.Mode = g.Mode
		ans.Cards = g.HandCards
	case match.InAuction:
		ans.Cards = tm.Match.DealtCards()
	case match.MatchFinished:
		evaluation := match.EvaluateGame(tm.Match.Game, tm.Match.Bids)
		ans.Evaluation = &evaluation
	}
	return ans
}
