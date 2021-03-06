package server

import (
	"fmt"
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
	"math/rand"
	"time"
)

type Table struct {
	Id             string
	Created        time.Time
	InviteCode     string
	Phase          api.TablePhase
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

func NewTable(owner string, fixedTableId *string, fixedInviteCode *string) *Table {
	table := Table{
		getStringWithDefault(fixedTableId, randomTableId),
		time.Now(),
		getStringWithDefault(fixedInviteCode, randomInviteCode),
		api.TablePhase_NOT_STARTED,
		[]string{owner},
		nil, nil, nil}
	return &table
}

func randomTableId() string {
	return RandomLetters(6)
}

func randomInviteCode() string {
	return RandomLetters(12)
}

func (t *Table) Start() error {
	if t.Phase != api.TablePhase_NOT_STARTED {
		return NewCloudError(TableAlreadyStarted)
	}
	if len(t.players) < game.NumPlayers || len(t.players) >= 7 {
		return NewCloudError(InvalidNumberOfPlayers)
	}
	t.playersInOrder = make([]string, len(t.players))
	copy(t.playersInOrder, t.players)
	rand.Shuffle(len(t.players), func(i int, j int) {
		t.playersInOrder[i], t.playersInOrder[j] = t.playersInOrder[j], t.playersInOrder[i]
	})
	t.round = match.NewRound(len(t.players), rand.Int63())
	t.Phase = api.TablePhase_BETWEEN_GAMES
	return t.StartMatch()
}
func (t *Table) StartMatch() error {
	if t.Phase != api.TablePhase_BETWEEN_GAMES {
		return NewCloudError(CannotStartTableNow)
	}
	nextMatch := t.round.NextMatch()
	pa := t.round.CurrentPlayerAssignment()

	t.CurrentMatch = &TableMatch{nextMatch, getActivePlayerIds(t.playersInOrder, pa)}
	t.Phase = api.TablePhase_PLAYING;
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
	if t.Phase != api.TablePhase_NOT_STARTED {
		return NewCloudError(UnableToJoinStartedTable)
	}
	if len(t.players) > match.MaxPlayersPerRound {
		return NewCloudError(TablePlayerLimitReached)
	}
	if t.ContainsPlayer(user) {
		return NewCloudError(UserAlreadyAtTable)
	}
	t.players = append(t.players, user)
	return nil
}
