package room

import (
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
	"github.com/supermihi/karlchencloud/server/errors"
	t "github.com/supermihi/karlchencloud/server/table"
	u "github.com/supermihi/karlchencloud/server/users"
	"sync"
)

type TableAndMutex struct {
	table *t.Table
	mutex *sync.RWMutex
}

// A Room manages a collection of Doppelkopf tables and users.
type Room struct {
	tables      map[t.Id]TableAndMutex
	tableByUser map[u.Id]TableAndMutex
	// for access to the above two maps
	mutex sync.RWMutex
}

func NewRoom() *Room {
	return &Room{
		tables:      make(map[t.Id]TableAndMutex),
		tableByUser: make(map[u.Id]TableAndMutex),
	}
}

func (r *Room) CreateTable(owner u.Id, public bool, seed int64) (td *TableData, err error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if _, exists := r.tableByUser[owner]; exists {
		return nil, errors.NewCloudError(errors.UserAlreadyAtOtherTable)
	}
	table := t.Initialize(owner, public, seed)
	tableAndMutex := TableAndMutex{table, &sync.RWMutex{}}
	r.tables[table.Id] = tableAndMutex
	r.tableByUser[owner] = tableAndMutex
	return NewTableData(table), nil
}

func (r *Room) GetTable(tableId t.Id) (*TableData, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	tableAndMutex, ok := r.tables[tableId]
	if ok {
		return NewTableData(tableAndMutex.table), nil
	}
	return nil, errors.NewCloudError(errors.TableDoesNotExist)
}

func (r *Room) findTableAndMutexWithInviteCodeUnsafe(inviteCode string) (TableAndMutex, bool) {
	for _, tableAndMutex := range r.tables {
		if tableAndMutex.table.InviteCode == inviteCode {
			return tableAndMutex, true
		}
	}
	return TableAndMutex{}, false
}

func (r *Room) JoinTableByInviteCode(user u.Id, inviteCode string) (*TableData, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	tableAndMutex, ok := r.findTableAndMutexWithInviteCodeUnsafe(inviteCode)
	if !ok {
		return nil, errors.NewCloudError(errors.InvalidInviteCode)
	}
	if _, exists := r.tableByUser[user]; exists {
		return nil, errors.NewCloudError(errors.UserAlreadyAtOtherTable)
	}
	tableAndMutex.mutex.Lock()
	defer tableAndMutex.mutex.Unlock()
	if err := tableAndMutex.table.Join(user); err != nil {
		return nil, err
	}
	r.tableByUser[user] = tableAndMutex
	return NewTableData(tableAndMutex.table), nil
}

func (r *Room) JoinTableByTableId(userId u.Id, id t.Id) (*TableData, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	tableAndMutex, ok := r.tables[id]
	if !ok {
		return nil, errors.NewCloudError(errors.TableDoesNotExist)
	}
	if _, exists := r.tableByUser[userId]; exists {
		return nil, errors.NewCloudError(errors.UserAlreadyAtOtherTable)
	}
	tableAndMutex.mutex.Lock()
	defer tableAndMutex.mutex.Unlock()
	if err := tableAndMutex.table.Join(userId); err != nil {
		return nil, err
	}
	r.tableByUser[userId] = tableAndMutex
	return NewTableData(tableAndMutex.table), nil
}

func (r *Room) getTable(tableId t.Id) (TableAndMutex, error) {
	r.mutex.RLock()
	tableAndMutex, ok := r.tables[tableId]
	r.mutex.RUnlock()
	if !ok {
		return TableAndMutex{}, errors.NewCloudError(errors.TableDoesNotExist)
	}
	return tableAndMutex, nil
}

func ensureUserIsOwnerUnsafe(table *t.Table, userId u.Id) error {
	if !table.ContainsPlayer(userId) {
		return errors.NewCloudError(errors.UserNotAtTable)
	}
	if table.Owner() != userId {
		return errors.NewCloudError(errors.NotOwnerOfTable)
	}
	return nil
}

func (r *Room) StartTable(tableId t.Id, user u.Id) (*TableData, error) {
	tableAndMutex, err := r.getTable(tableId)
	if err != nil {
		return nil, err
	}
	tableAndMutex.mutex.Lock()
	defer tableAndMutex.mutex.Unlock()
	err = ensureUserIsOwnerUnsafe(tableAndMutex.table, user)
	if err != nil {
		return nil, err
	}
	err = tableAndMutex.table.Start()
	if err != nil {
		return nil, err
	}
	return NewTableData(tableAndMutex.table), nil
}

func (r *Room) GetMatchData(tableId t.Id) (m *MatchData, err error) {
	tableAndMutex, err := r.getTable(tableId)
	if err != nil {
		return nil, err
	}
	tableAndMutex.mutex.RLock()
	defer tableAndMutex.mutex.RUnlock()
	tableMatch := tableAndMutex.table.CurrentMatch
	if tableMatch == nil {
		return nil, errors.NewCloudError(errors.NoCurrentMatch)
	}
	return NewMatchData(tableMatch), nil
}

func (r *Room) PlayCard(tableId t.Id, userId u.Id, card game.Card) (matchData *MatchData, err error) {
	tableAndMatch, err := r.getTable(tableId)
	if err != nil {
		return nil, err
	}
	tableAndMatch.mutex.Lock()
	defer tableAndMatch.mutex.Unlock()
	m, player, err := getMatchAndPlayerUnsafe(tableAndMatch.table, userId)
	if err != nil {
		return nil, err
	}
	if !m.Match.PlayCard(player, card) {
		return nil, errors.NewCloudError(errors.CannotPlayCard)
	}
	if m.Match.Phase() == match.MatchFinished {
		err = tableAndMatch.table.EndMatch()
		if err != nil {
			return nil, err
		}
	}
	return NewMatchData(m), nil
}

func (r *Room) Declare(tableId t.Id, userId u.Id, gameType game.AnnouncedGameType) (*MatchData, error) {
	tableAndMatch, err := r.getTable(tableId)
	if err != nil {
		return nil, err
	}
	tableAndMatch.mutex.Lock()
	defer tableAndMatch.mutex.Unlock()
	m, player, err := getMatchAndPlayerUnsafe(tableAndMatch.table, userId)
	if err != nil {
		return nil, err
	}
	if !m.Match.AnnounceGameType(player, gameType) {
		return nil, errors.NewCloudError(errors.CannotAnnounce)
	}
	return NewMatchData(m), nil
}

func (r *Room) PlaceBid(tableId t.Id, userId u.Id, bid match.Bid) (*MatchData, error) {
	tableAndMatch, err := r.getTable(tableId)
	if err != nil {
		return nil, err
	}
	tableAndMatch.mutex.Lock()
	defer tableAndMatch.mutex.Unlock()
	m, player, err := getMatchAndPlayerUnsafe(tableAndMatch.table, userId)
	if err != nil {
		return nil, err
	}
	if !m.Match.PlaceBid(player, bid) {
		return nil, errors.NewCloudError(errors.CannotPlaceBid)
	}
	return NewMatchData(m), nil
}

func (r *Room) StartNextMatch(tableId t.Id, userId u.Id) (*MatchData, error) {
	tableAndMatch, err := r.getTable(tableId)
	if err != nil {
		return nil, err
	}
	tableAndMatch.mutex.Lock()
	defer tableAndMatch.mutex.Unlock()
	err = ensureUserIsOwnerUnsafe(tableAndMatch.table, userId)
	if err != nil {
		return nil, err
	}
	err = tableAndMatch.table.StartMatch()
	if err != nil {
		return nil, err
	}
	return NewMatchData(tableAndMatch.table.CurrentMatch), nil
}

func getMatchAndPlayerUnsafe(table *t.Table, user u.Id) (match *t.TableMatch, p game.Player, err error) {
	if table.CurrentMatch == nil {
		err = errors.NewCloudError(errors.NoCurrentMatch)
		return
	}
	players := table.CurrentMatch.Players
	p = players.PlayerFor(user)
	if p == game.NoPlayer {
		err = errors.NewCloudError(errors.UserNotPlayingInMatch)
		return
	}
	return table.CurrentMatch, p, nil
}

func (r *Room) ActiveTableOf(user u.Id) *TableData {
	r.mutex.RLock()
	tableAndMutex, ok := r.tableByUser[user]
	r.mutex.RUnlock()
	if !ok {
		return nil
	}
	tableAndMutex.mutex.RLock()
	defer tableAndMutex.mutex.RUnlock()
	return NewTableData(tableAndMutex.table)
}

func (r *Room) IsAtAnyTable(user u.Id) bool {
	r.mutex.RLock()
	_, ok := r.tableByUser[user]
	r.mutex.RUnlock()
	return ok
}

func (r *Room) UsersAtSameTable(user u.Id) []u.Id {
	r.mutex.RLock()
	tableAndMutex, ok := r.tableByUser[user]
	r.mutex.RUnlock()
	if !ok {
		return nil
	}
	tableAndMutex.mutex.RLock()
	defer tableAndMutex.mutex.RUnlock()
	return u.IdsExcept(tableAndMutex.table.Players, user)
}

func (r *Room) GetOpenTables(_ u.Id) []*TableData {
	ans := make([]*TableData, 0)
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	for _, tableAndMutex := range r.tables {
		tableAndMutex.mutex.RLock()
		if tableAndMutex.table.Phase != api.TablePhase_NOT_STARTED || !tableAndMutex.table.Public {
			tableAndMutex.mutex.RUnlock()
			continue
		}
		ans = append(ans, NewTableData(tableAndMutex.table))
		tableAndMutex.mutex.RUnlock()
	}
	return ans
}
