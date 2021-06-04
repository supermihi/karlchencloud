package client

import (
	"fmt"
	pb "github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/api/pbconv"
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
	"log"
	"sort"
)

type TrickView struct {
	Forehand string
	Cards    map[string]game.Card
}

func NewTrickView(forehand string) *TrickView {
	return &TrickView{Forehand: forehand, Cards: make(map[string]game.Card)}
}

type Players struct {
	Left  string
	Face  string
	Right string
	Me    string
}

type ModeView struct {
	Type    game.AnnouncedGameType
	Soloist *string
	Spouse  *string
}

type MatchView struct {
	Cards   game.Hand
	Players Players
	Phase   match.Phase
	MyTurn  bool
	Trick   *TrickView
	Mode    *ModeView
}
type TableView struct {
	TableInfo
	Match *MatchView
}

type TableInfo struct {
	Id              string
	Invite          string
	MemberNamesById map[string]string
	Public          bool
}

func NewTableInfo(table *pb.TableData) TableInfo {
	ans := TableInfo{Id: table.TableId, Invite: table.InviteCode, MemberNamesById: make(map[string]string),
		Public: table.Public}
	for _, member := range table.Members {
		ans.MemberNamesById[member.UserId] = member.Name
	}
	return ans
}

func (m *TableView) PlayerNames() string {
	if m.Match == nil {
		return "- no match -"
	}
	p := m.Match.Players
	return fmt.Sprintf("Left: %s, Face: %s, Right: %s", m.MemberNamesById[p.Left], m.MemberNamesById[p.Face], m.MemberNamesById[p.Right])
}

func NewMatchView(state *pb.MatchState) *MatchView {
	v := &MatchView{}
	switch r := state.Role.(type) {
	case *pb.MatchState_Spectator:
		log.Fatalf("unexpected role spectator")
	case *pb.MatchState_OwnCards:
		v.Cards = ToHand(r.OwnCards.Cards)
		sort.Sort(game.BySuitAndRank(v.Cards))
	}
	v.Players.Left = state.Players.UserIdLeft
	v.Players.Face = state.Players.UserIdFace
	v.Players.Right = state.Players.UserIdRight
	v.Players.Me = state.Players.UserIdSelf
	switch details := state.Details.(type) {
	case *pb.MatchState_AuctionState:
		v.Phase = match.InAuction
	case *pb.MatchState_GameState:
		v.Phase = match.InGame
		gs := details.GameState
		if gs.Mode != nil {
			v.setMode(details.GameState.Mode)
		}
		if gs.CurrentTrick != nil {
			v.Trick = NewTrickView(gs.CurrentTrick.UserIdForehand)
			if len(gs.CurrentTrick.Cards) > 0 {
				v.Trick.Cards[v.Trick.Forehand] = pbconv.ToCard(gs.CurrentTrick.Cards[0])
			}
		}
	}
	v.Phase = pbconv.ToMatchPhase(state.Phase)
	v.MyTurn = state.Turn.UserId == v.Players.Me
	return v
}

func NewTableView(ts *pb.TableState) *TableView {
	ans := TableView{TableInfo: NewTableInfo(ts.Data)}
	if ts.CurrentMatch != nil {
		ans.Match = NewMatchView(ts.CurrentMatch)
	}
	return &ans
}

func (v *MatchView) UpdateTrick(pc *pb.PlayedCard) {
	if v.Trick == nil {
		v.Trick = NewTrickView(pc.UserId)
	}
	v.Trick.Cards[pc.UserId] = pbconv.ToCard(pc.Card)
	v.MyTurn = pc.UserId == v.Players.Right
	if pc.TrickWinner != nil {
		v.Trick = NewTrickView(pc.TrickWinner.UserId)
		v.MyTurn = pc.TrickWinner.UserId == v.Players.Me
	}
}

func (v *MatchView) DrawCard(index int) game.Card {
	card := v.Cards[index]
	v.Cards = append(v.Cards[:index], v.Cards[index+1:]...)
	return card
}

func (v *MatchView) setMode(m *pb.Mode) {
	v.Mode = &ModeView{Type: pbconv.ToGameType(m.Type)}
	if m.Soloist != nil {
		v.Mode.Soloist = &m.Soloist.UserId
	}
	if m.Spouse != nil {
		v.Mode.Spouse = &m.Spouse.UserId
	}
	v.MyTurn = m.Forehand == v.Players.Me
	v.Trick = NewTrickView(m.Forehand)
}

func (v *MatchView) GameSuit(card game.Card) game.GameSuit {
	return match.GameSuitOf(card, v.Mode.Type)
}

func (v *MatchView) UpdateOnDeclare(d *pb.Declaration) {
	if d.DefinedGameMode != nil {
		v.setMode(d.DefinedGameMode)
		v.Phase = match.InGame
	} else {
		v.MyTurn = d.UserId == v.Players.Right
	}
}
