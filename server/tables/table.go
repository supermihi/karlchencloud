package tables

import (
	"fmt"
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
	"github.com/supermihi/karlchencloud/server/errors"
	"github.com/supermihi/karlchencloud/server/users"
	"log"
	"math/rand"
	"time"
)

type Table struct {
	Id             TableId
	Created        time.Time
	InviteCode     string
	Public         bool
	Phase          api.TablePhase
	Players        []users.Id
	PlayersInOrder []users.Id
	round          *match.Round
	CurrentMatch   *TableMatch
	Rng            *rand.Rand
}

func Initialize(owner users.Id, public bool, seed int64) *Table {
	log.Printf("Creating table (seed: %d)", seed)
	tableId := randomTableId()
	return &Table{
		tableId,
		time.Now(),
		randomInviteCode(),
		public,
		api.TablePhase_NOT_STARTED,
		[]users.Id{owner},
		nil,
		nil,
		nil,
		rand.New(rand.NewSource(seed)),
	}
}

func (t *Table) Owner() users.Id {
	return t.Players[0]
}

func (t *Table) String() string {
	return fmt.Sprintf("Table %v", t.Id)
}

func (t *Table) Start() error {
	if t.Phase != api.TablePhase_NOT_STARTED {
		return errors.NewCloudError(errors.TableAlreadyStarted)
	}
	if len(t.Players) < game.NumPlayers || len(t.Players) >= 7 {
		return errors.NewCloudError(errors.InvalidNumberOfPlayers)
	}
	t.PlayersInOrder = make([]users.Id, len(t.Players))
	copy(t.PlayersInOrder, t.Players)
	t.Rng.Shuffle(len(t.Players), func(i int, j int) {
		t.PlayersInOrder[i], t.PlayersInOrder[j] = t.PlayersInOrder[j], t.PlayersInOrder[i]
	})
	t.round = match.NewRound(len(t.Players), t.Rng)
	t.Phase = api.TablePhase_BETWEEN_GAMES
	return t.StartMatch()
}

func (t *Table) EndMatch() error {
	if t.Phase != api.TablePhase_PLAYING || t.CurrentMatch == nil || t.CurrentMatch.Match.Phase() != match.MatchFinished {
		return errors.NewCloudError(errors.CannotEndMatch)
	}
	t.Phase = api.TablePhase_BETWEEN_GAMES
	evaluation := t.CurrentMatch.Match.Evaluate()
	t.round.AddScores(&evaluation)
	return nil
}

func (t *Table) StartMatch() error {
	if t.Phase != api.TablePhase_BETWEEN_GAMES {
		return errors.NewCloudError(errors.CannotStartTableNow)
	}
	nextMatch := t.round.NextMatch()
	pa := t.round.CurrentPlayerAssignment()

	t.CurrentMatch = &TableMatch{nextMatch, getActivePlayerIds(t.PlayersInOrder, pa)}
	t.Phase = api.TablePhase_PLAYING
	return nil
}

func (t *Table) ContainsPlayer(player users.Id) bool {
	for _, p := range t.Players {
		if p == player {
			return true
		}
	}
	return false
}

func (t *Table) Users() []users.Id {
	return t.Players
}

func (t *Table) Join(user users.Id) error {
	if t.Phase != api.TablePhase_NOT_STARTED {
		return errors.NewCloudError(errors.UnableToJoinStartedTable)
	}
	if len(t.Players) > match.MaxPlayersPerRound {
		return errors.NewCloudError(errors.TablePlayerLimitReached)
	}
	if t.ContainsPlayer(user) {
		return errors.NewCloudError(errors.UserAlreadyAtTable)
	}
	t.Players = append(t.Players, user)
	return nil
}
