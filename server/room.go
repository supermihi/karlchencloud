package server

import (
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
)

type Room struct {
	users  Users
	tables map[TableId]*Table
}

func NewRoom(users Users) *Room {
	return &Room{users, make(map[TableId]*Table)}
}

func (r *Room) CreateTable(owner UserId, fixedTableId TableId, fixedInviteCode *string, seed int64) (table *TableData, err error) {
	if r.ActiveTableOf(owner) != nil {
		return nil, NewCloudError(TableAlreadyExists)
	}
	t := NewTable(owner, fixedTableId, fixedInviteCode, seed)
	r.tables[t.Id] = t
	return GetData(t), nil
}

func (r *Room) GetTable(tableId TableId) (table *TableData, err error) {
	t, ok := r.tables[tableId]
	if ok {
		return GetData(t), nil
	}
	return nil, NewCloudError(TableDoesNotExist)
}

func (r *Room) findTableWithInviteCode(inviteCode string) *Table {
	for _, t := range r.tables {
		if t.InviteCode == inviteCode {
			return t
		}
	}
	return nil
}

func (r *Room) JoinTable(userId UserId, inviteCode string) (*TableData, error) {
	t := r.findTableWithInviteCode(inviteCode)
	if t == nil {
		return nil, NewCloudError(InvalidInviteCode)
	}
	if r.ActiveTableOf(userId) != nil {
		return nil, NewCloudError(UserAlreadyAtOtherTable)
	}
	if err := t.Join(userId); err != nil {
		return nil, err
	}
	return GetData(t), nil

}

func (r *Room) ensureIsOwner(tableId TableId, userId UserId) (*Table, error) {
	t, tableExists := r.tables[tableId]
	if !tableExists {
		return nil, NewCloudError(TableDoesNotExist)
	}
	if !t.ContainsPlayer(userId) {
		return nil, NewCloudError(UserNotAtTable)
	}
	if t.Owner() != userId {
		return nil, NewCloudError(NotOwnerOfTable)
	}
	return t, nil
}

func (r *Room) StartTable(tableId TableId, userId UserId) (*TableData, error) {
	t, err := r.ensureIsOwner(tableId, userId)
	if err != nil {
		return nil, err
	}
	err = t.Start()
	if err != nil {
		return nil, err
	}
	return GetData(t), nil
}

type Declaration struct {
	Healthy     bool
	Reservation game.AnnouncedGameType
}
type MatchData struct {
	Phase           match.Phase
	Turn            game.Player
	Players         [game.NumPlayers]UserId
	InitialForehand game.Player
	Cards           game.Cards
	Declarations    map[game.Player]Declaration
	Bids            match.Bids
	CompletedTricks int
	CurrentTrick    *game.IncompleteTrick
	PreviousTrick   *game.Trick
	Mode            game.Mode
	Evaluation      *match.GameEvaluation
}

func GetMatchData(tm *TableMatch) *MatchData {
	decls := tm.Match.Auction.Declarations
	declarations := make(map[game.Player]Declaration, len(decls))
	for k, v := range decls {

		if v.Healthy {
			declarations[k] = Declaration{true, game.NormalGameType}
		} else {
			declarations[k] = Declaration{false, v.Reservation.Type()}
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
func (r *Room) GetMatchData(tableId TableId) (m *MatchData, err error) {
	t, ok := r.tables[tableId]
	if !ok {
		return nil, NewCloudError(TableDoesNotExist)
	}
	tm := t.CurrentMatch
	if tm == nil {
		return nil, NewCloudError(NoCurrentMatch)
	}
	return GetMatchData(tm), nil
}

func (r *Room) PlayCard(tableId TableId, userId UserId, card game.Card) (matchData *MatchData, err error) {
	m, player, err := r.getMatchAndPlayer(tableId, userId)
	if err != nil {
		return nil, err
	}
	if !m.Match.PlayCard(player, card) {
		return nil, NewCloudError(CannotPlayCard)
	}
	if m.Match.Phase() == match.MatchFinished {
		table := r.activeTableOf(userId)
		return nil, table.EndMatch()
	}
	return GetMatchData(m), nil
}

func (r *Room) Declare(tableId TableId, userId UserId, gameType game.AnnouncedGameType) (*MatchData, error) {
	m, player, err := r.getMatchAndPlayer(tableId, userId)
	if err != nil {
		return nil, err
	}
	if !m.Match.AnnounceGameType(player, gameType) {
		return nil, NewCloudError(CannotAnnounce)
	}
	return GetMatchData(m), nil
}

func (r *Room) PlaceBid(tableId TableId, userId UserId, bid match.Bid) (*MatchData, error) {
	m, player, err := r.getMatchAndPlayer(tableId, userId)
	if err != nil {
		return nil, err
	}
	if !m.Match.PlaceBid(player, bid) {
		return nil, NewCloudError(CannotPlaceBid)
	}
	return GetMatchData(m), nil
}

func (r *Room) StartNextMatch(tableId TableId, userId UserId) (*MatchData, error) {
	table, err := r.ensureIsOwner(tableId, userId)
	if err != nil {
		return nil, err
	}
	err = table.StartMatch()
	if err != nil {
		return nil, err
	}
	return r.GetMatchData(tableId)
}

func (r *Room) getMatchAndPlayer(tableId TableId, user UserId) (match *TableMatch, p game.Player, err error) {
	table, exists := r.tables[tableId]
	if !exists {
		err = NewCloudError(TableDoesNotExist)
		return
	}
	if table.CurrentMatch == nil {
		err = NewCloudError(NoCurrentMatch)
		return
	}
	players := table.CurrentMatch.Players
	p = players.PlayerFor(user)
	if p == game.NoPlayer {
		err = NewCloudError(UserNotPlayingInMatch)
		return
	}
	return table.CurrentMatch, p, nil
}

func (r *Room) activeTableOf(user UserId) *Table {
	for _, table := range r.tables {
		if table.ContainsPlayer(user) {
			return table
		}
	}
	return nil
}
func (r *Room) ActiveTableOf(user UserId) *TableData {
	t := r.activeTableOf(user)
	if t != nil {
		return GetData(t)
	}
	return nil
}

func (r *Room) RelatedUsers(userId UserId) []UserId {
	table := r.activeTableOf(userId)
	if table == nil {
		return []UserId{}
	}
	return usersExcept(table.players, userId)
}
