package server

import (
	"fmt"
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
	"log"
	"math/rand"
	"time"
)

type Table struct {
	Id             TableId
	Created        time.Time
	InviteCode     string
	Phase          api.TablePhase
	players        []UserId
	playersInOrder []UserId
	round          *match.Round
	CurrentMatch   *TableMatch
	Rng            *rand.Rand
}

func (t *Table) Owner() UserId {
	return t.players[0]
}

func (t *Table) String() string {
	return fmt.Sprintf("Table %v", t.Id)
}

func NewTable(owner UserId, fixedTableId TableId, fixedInviteCode *string, seed int64) *Table {
	log.Printf("Creating table (seed: %d)", seed)
	tableId := fixedTableId
	if tableId == InvalidTableId {
		tableId = randomTableId()
	}
	return &Table{
		tableId,
		time.Now(),
		getStringWithDefault(fixedInviteCode, randomInviteCode),
		api.TablePhase_NOT_STARTED,
		[]UserId{owner},
		nil,
		nil,
		nil,
		rand.New(rand.NewSource(seed)),
	}
}

func randomTableId() TableId {
	return TableId(rand.Int63())
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
	t.playersInOrder = make([]UserId, len(t.players))
	copy(t.playersInOrder, t.players)
	t.Rng.Shuffle(len(t.players), func(i int, j int) {
		t.playersInOrder[i], t.playersInOrder[j] = t.playersInOrder[j], t.playersInOrder[i]
	})
	t.round = match.NewRound(len(t.players), t.Rng)
	t.Phase = api.TablePhase_BETWEEN_GAMES
	return t.StartMatch()
}

func (t *Table) EndMatch() error {
	if t.Phase != api.TablePhase_PLAYING || t.CurrentMatch == nil || t.CurrentMatch.Match.Phase() != match.MatchFinished {
		return NewCloudError(CannotEndMatch)
	}
	t.Phase = api.TablePhase_BETWEEN_GAMES
	evaluation := t.CurrentMatch.Match.Evaluate()
	t.round.AddScores(&evaluation)
	return nil
}

func (t *Table) StartMatch() error {
	if t.Phase != api.TablePhase_BETWEEN_GAMES {
		return NewCloudError(CannotStartTableNow)
	}
	nextMatch := t.round.NextMatch()
	pa := t.round.CurrentPlayerAssignment()

	t.CurrentMatch = &TableMatch{nextMatch, getActivePlayerIds(t.playersInOrder, pa)}
	t.Phase = api.TablePhase_PLAYING
	return nil
}

func (t *Table) ContainsPlayer(player UserId) bool {
	for _, p := range t.players {
		if p == player {
			return true
		}
	}
	return false
}

func (t *Table) Users() []UserId {
	return t.players
}

func (t *Table) Join(user UserId) error {
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
