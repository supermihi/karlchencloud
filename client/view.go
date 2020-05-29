package client

import (
	"fmt"
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/common"
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
	"log"
)

type TrickView struct {
	Forehand string
	Cards    map[string]game.Card
}

func NewTrickView(forehand string) TrickView {
	return TrickView{Forehand: forehand, Cards: make(map[string]game.Card)}
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
	TableId string
	MyTurn  bool
	Names   map[string]string
	Trick   *TrickView
	Mode    *ModeView
}

func (m *MatchView) PlayerNames() string {
	p := m.Players
	return fmt.Sprintf("Left: %s, Face: %s, Right: %s", m.Names[p.Left], m.Names[p.Face], m.Names[p.Right])
}

func NewMatchView(tableId string) MatchView {
	return MatchView{Phase: match.MatchFinished, TableId: tableId,
		Names: make(map[string]string)}
}

func (v *MatchView) InitFromMatchState(state *api.MatchState) {
	switch r := state.Role.(type) {
	case *api.MatchState_Spectator:
		log.Fatalf("unexpected role spectator")
	case *api.MatchState_OwnCards:
		v.Cards = ToHand(r.OwnCards.Cards)
	}
	v.Players.Left = state.Players.UserIdLeft
	v.Players.Face = state.Players.UserIdFace
	v.Players.Right = state.Players.UserIdRight
	v.Players.Me = state.Players.UserIdSelf
	v.Phase = common.ToMatchPhase(state.Phase)
	v.MyTurn = state.Turn.UserId == v.Players.Me

}

func (v *MatchView) AddName(user string, name string) {
	v.Names[user] = name
}

func (v *MatchView) UpdateTrick(pc *api.PlayedCard) {
	if v.Trick == nil {
		tv := NewTrickView(pc.UserId)
		v.Trick = &tv
	}
	v.Trick.Cards[pc.UserId] = common.ToCard(pc.Card)
	v.MyTurn = pc.UserId == v.Players.Right
	if pc.TrickWinner != nil {
		v.Trick = nil
		v.MyTurn = pc.TrickWinner.UserId == v.Players.Me
	}
}

func (v *MatchView) DrawCard(index int) game.Card {
	card := v.Cards[index]
	v.Cards = append(v.Cards[:index], v.Cards[index+1:]...)
	return card
}

func (v *MatchView) setMode(m *api.Mode) {
	v.Mode = &ModeView{Type: common.ToGameType(m.Type)}
	if m.Soloist != nil {
		v.Mode.Soloist = &m.Soloist.UserId
	}
	if m.Spouse != nil {
		v.Mode.Spouse = &m.Spouse.UserId
	}
	v.MyTurn = m.Forehand == v.Players.Me
	tv := NewTrickView(m.Forehand)
	v.Trick = &tv
}

func (v *MatchView) GameSuit(card game.Card) game.GameSuit {
	return match.GameSuitOf(card, v.Mode.Type)
}

func (v *MatchView) UpdateOnDeclare(d *api.Declaration) {
	if d.DefinedGameMode != nil {
		v.setMode(d.DefinedGameMode)
		v.Phase = match.InGame
	} else {
		v.MyTurn = d.UserId == v.Players.Right
	}
}
