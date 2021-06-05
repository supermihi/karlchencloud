package tables

import (
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
	"github.com/supermihi/karlchencloud/server/errors"
	u "github.com/supermihi/karlchencloud/server/users"
	"sync"
)

// Tables manages a collection of tables and exports an interface for creating, joining and playing at tables.
type Tables struct {
	byTableId map[TableId]tableWithMutex
	byUser    map[u.Id]tableWithMutex
	// for access to the above two maps
	mutex sync.RWMutex
}

type tableWithMutex struct {
	*Table
	m *sync.RWMutex
}

func NewTables() *Tables {
	return &Tables{
		byTableId: make(map[TableId]tableWithMutex),
		byUser:    make(map[u.Id]tableWithMutex),
	}
}

func (tb *Tables) CreateTable(owner u.Id, public bool, seed int64) (td *TableData, err error) {
	tb.mutex.Lock()
	defer tb.mutex.Unlock()
	if _, exists := tb.byUser[owner]; exists {
		return nil, errors.NewCloudError(errors.UserAlreadyAtOtherTable)
	}
	table := Initialize(owner, public, seed)
	tableAndMutex := tableWithMutex{table, &sync.RWMutex{}}
	tb.byTableId[table.Id] = tableAndMutex
	tb.byUser[owner] = tableAndMutex
	return NewTableData(table), nil
}

func (tb *Tables) GetTable(tableId TableId) (*TableData, error) {
	tb.mutex.RLock()
	defer tb.mutex.RUnlock()
	tableAndMutex, ok := tb.byTableId[tableId]
	if ok {
		return NewTableData(tableAndMutex.Table), nil
	}
	return nil, errors.NewCloudError(errors.TableDoesNotExist)
}

func (tb *Tables) JoinTableByInviteCode(user u.Id, inviteCode string) (*TableData, error) {
	tb.mutex.Lock()
	defer tb.mutex.Unlock()
	var tableAndMutex tableWithMutex
	ok := false
	for _, table := range tb.byTableId {
		if table.InviteCode == inviteCode {
			tableAndMutex = table
			ok = true
			break
		}
	}
	if !ok {
		return nil, errors.NewCloudError(errors.InvalidInviteCode)
	}
	if _, exists := tb.byUser[user]; exists {
		return nil, errors.NewCloudError(errors.UserAlreadyAtOtherTable)
	}
	tableAndMutex.m.Lock()
	defer tableAndMutex.m.Unlock()
	if err := tableAndMutex.Join(user); err != nil {
		return nil, err
	}
	tb.byUser[user] = tableAndMutex
	return NewTableData(tableAndMutex.Table), nil
}

func (tb *Tables) JoinTableByTableId(userId u.Id, id TableId) (*TableData, error) {
	tb.mutex.Lock()
	defer tb.mutex.Unlock()
	tableAndMutex, ok := tb.byTableId[id]
	if !ok {
		return nil, errors.NewCloudError(errors.TableDoesNotExist)
	}
	if _, exists := tb.byUser[userId]; exists {
		return nil, errors.NewCloudError(errors.UserAlreadyAtOtherTable)
	}
	tableAndMutex.m.Lock()
	defer tableAndMutex.m.Unlock()
	if err := tableAndMutex.Join(userId); err != nil {
		return nil, err
	}
	tb.byUser[userId] = tableAndMutex
	return NewTableData(tableAndMutex.Table), nil
}

func (tb *Tables) StartTable(tableId TableId, user u.Id) (*TableData, error) {
	tableAndMutex, err := tb.getTable(tableId)
	if err != nil {
		return nil, err
	}
	tableAndMutex.m.Lock()
	defer tableAndMutex.m.Unlock()
	err = ensureUserIsOwnerUnsafe(tableAndMutex.Table, user)
	if err != nil {
		return nil, err
	}
	err = tableAndMutex.Start()
	if err != nil {
		return nil, err
	}
	return NewTableData(tableAndMutex.Table), nil
}

func (tb *Tables) GetMatchData(tableId TableId) (m *MatchData, err error) {
	tableAndMutex, err := tb.getTable(tableId)
	if err != nil {
		return nil, err
	}
	tableAndMutex.m.RLock()
	defer tableAndMutex.m.RUnlock()
	tableMatch := tableAndMutex.CurrentMatch
	if tableMatch == nil {
		return nil, errors.NewCloudError(errors.NoCurrentMatch)
	}
	return NewMatchData(tableMatch), nil
}

func (tb *Tables) PlayCard(tableId TableId, userId u.Id, card game.Card) (matchData *MatchData, err error) {
	tableAndMatch, err := tb.getTable(tableId)
	if err != nil {
		return nil, err
	}
	tableAndMatch.m.Lock()
	defer tableAndMatch.m.Unlock()
	m, player, err := getMatchAndPlayerUnsafe(tableAndMatch.Table, userId)
	if err != nil {
		return nil, err
	}
	if !m.Match.PlayCard(player, card) {
		return nil, errors.NewCloudError(errors.CannotPlayCard)
	}
	if m.Match.Phase() == match.MatchFinished {
		err = tableAndMatch.EndMatch()
		if err != nil {
			return nil, err
		}
	}
	return NewMatchData(m), nil
}

func (tb *Tables) Declare(tableId TableId, userId u.Id, gameType game.AnnouncedGameType) (*MatchData, error) {
	tableAndMatch, err := tb.getTable(tableId)
	if err != nil {
		return nil, err
	}
	tableAndMatch.m.Lock()
	defer tableAndMatch.m.Unlock()
	m, player, err := getMatchAndPlayerUnsafe(tableAndMatch.Table, userId)
	if err != nil {
		return nil, err
	}
	if !m.Match.AnnounceGameType(player, gameType) {
		return nil, errors.NewCloudError(errors.CannotAnnounce)
	}
	return NewMatchData(m), nil
}

func (tb *Tables) PlaceBid(tableId TableId, userId u.Id, bid match.Bid) (*MatchData, error) {
	tableAndMatch, err := tb.getTable(tableId)
	if err != nil {
		return nil, err
	}
	tableAndMatch.m.Lock()
	defer tableAndMatch.m.Unlock()
	m, player, err := getMatchAndPlayerUnsafe(tableAndMatch.Table, userId)
	if err != nil {
		return nil, err
	}
	if !m.Match.PlaceBid(player, bid) {
		return nil, errors.NewCloudError(errors.CannotPlaceBid)
	}
	return NewMatchData(m), nil
}

func (tb *Tables) StartNextMatch(tableId TableId, userId u.Id) (*MatchData, error) {
	tableAndMatch, err := tb.getTable(tableId)
	if err != nil {
		return nil, err
	}
	tableAndMatch.m.Lock()
	defer tableAndMatch.m.Unlock()
	err = ensureUserIsOwnerUnsafe(tableAndMatch.Table, userId)
	if err != nil {
		return nil, err
	}
	err = tableAndMatch.StartMatch()
	if err != nil {
		return nil, err
	}
	return NewMatchData(tableAndMatch.CurrentMatch), nil
}

func (tb *Tables) ActiveTableOf(user u.Id) *TableData {
	tb.mutex.RLock()
	tableAndMutex, ok := tb.byUser[user]
	tb.mutex.RUnlock()
	if !ok {
		return nil
	}
	tableAndMutex.m.RLock()
	defer tableAndMutex.m.RUnlock()
	return NewTableData(tableAndMutex.Table)
}

func (tb *Tables) IsAtAnyTable(user u.Id) bool {
	tb.mutex.RLock()
	_, ok := tb.byUser[user]
	tb.mutex.RUnlock()
	return ok
}

func (tb *Tables) UsersAtSameTable(user u.Id) []u.Id {
	tb.mutex.RLock()
	tableAndMutex, ok := tb.byUser[user]
	tb.mutex.RUnlock()
	if !ok {
		return nil
	}
	tableAndMutex.m.RLock()
	defer tableAndMutex.m.RUnlock()
	return u.IdsExcept(tableAndMutex.Players, user)
}

func (tb *Tables) GetOpenTables(_ u.Id) []*TableData {
	ans := make([]*TableData, 0)
	tb.mutex.RLock()
	defer tb.mutex.RUnlock()
	for _, tableAndMutex := range tb.byTableId {
		tableAndMutex.m.RLock()
		if tableAndMutex.Phase != api.TablePhase_NOT_STARTED || !tableAndMutex.Public {
			tableAndMutex.m.RUnlock()
			continue
		}
		ans = append(ans, NewTableData(tableAndMutex.Table))
		tableAndMutex.m.RUnlock()
	}
	return ans
}

func (tb *Tables) getTable(tableId TableId) (result tableWithMutex, err error) {
	tb.mutex.RLock()
	result, ok := tb.byTableId[tableId]
	tb.mutex.RUnlock()
	if !ok {
		err = errors.NewCloudError(errors.TableDoesNotExist)
	}
	return
}
