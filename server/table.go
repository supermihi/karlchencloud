package server

import (
	"fmt"
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
	"math/rand"
)

type TablePhase int

const (
	BeforeFirstGame TablePhase = iota
	Playing
	WaitingForNextGame
	AllGamesEnded
)

type Table struct {
	Id             string
	InviteCode     string
	Phase          TablePhase
	players        []string
	playersInOrder []string
	round          *match.Round
	CurrentMatch   *TableMatch
}

func (t *Table) Owner() string {
	return t.players[0]
}

func (t *Table) String() string {
	return fmt.Sprintf("Table %v", t.Id)
}

func NewTable(owner string) *Table {
	id := RandomLetters(6)
	inviteCode := RandomLetters(12)
	table := Table{id, inviteCode, BeforeFirstGame, []string{owner}, nil, nil, nil}
	return &table
}

func (t *Table) Start() error {
	if t.Phase != BeforeFirstGame {
		return Error("table already started")
	}
	if len(t.players) < game.NumPlayers || len(t.players) >= 7 {
		return Error(fmt.Sprintf("Cannot start table with %v players", len(t.players)))
	}
	t.playersInOrder = make([]string, len(t.players))
	copy(t.playersInOrder, t.players)
	rand.Shuffle(len(t.players), func(i int, j int) {
		t.playersInOrder[i], t.playersInOrder[j] = t.playersInOrder[j], t.playersInOrder[i]
	})
	t.round = match.NewRound(len(t.players), rand.Int63())
	t.Phase = WaitingForNextGame
	return t.StartMatch()
}
func (t *Table) StartMatch() error {
	if t.Phase != WaitingForNextGame {
		return Error("can only start match in phase WaitingForNextGame")
	}
	nextMatch := t.round.NextMatch()
	pa := t.round.CurrentPlayerAssignment()

	t.CurrentMatch = &TableMatch{nextMatch, getActivePlayerIds(t.playersInOrder, pa)}
	return nil
}

func (t *Table) ContainsPlayer(player string) bool {
	for _, p := range t.players {
		if p == player {
			return true
		}
	}
	return false
}

func (t *Table) Users() []string {
	return t.players
}

func (t *Table) Join(user string) error {
	if t.Phase != BeforeFirstGame {
		return Error("cannot join a started table")
	}
	if len(t.players) > match.MaxPlayersPerRound {
		return Error(fmt.Sprintf("only %v players supported per table", match.MaxPlayersPerRound))
	}
	if t.ContainsPlayer(user) {
		return Error("user already at table")
	}
	t.players = append(t.players, user)
	return nil
}