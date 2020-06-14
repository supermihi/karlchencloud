package server

import (
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
)

type Room struct {
	users  Users
	Tables *Tables
}

func NewRoom(users Users) *Room {
	return &Room{users, NewTables()}
}

func (r *Room) CreateTable(owner string) (table *TableData, err error) {
	if r.Tables.ActiveTableOf(owner) != nil {
		return nil, NewCloudError(TableAlreadyExists)
	}
	t := r.Tables.CreateTable(owner)
	return GetData(t), nil
}

func (r *Room) TablePlayers(id string) []string {
	return r.Tables.ById[id].players
}

func (r *Room) AddUser(user string, name string, secret string) bool {
	return r.users.Add(user, name, secret)
}

func (r *Room) GetUserData(userId string) (data UserData, err error) {
	name, ok := r.users.GetName(userId)
	if ok {
		return UserData{Id: userId, Name: name}, nil
	}
	return UserData{}, NewCloudError(UserDoesNotExist)
}

func (r *Room) GetTable(tableId string) (table *TableData, err error) {
	t, ok := r.Tables.ById[tableId]
	if ok {
		return GetData(t), nil
	}
	return nil, NewCloudError(TableDoesNotExist)
}

func (r *Room) JoinTable(tableId string, userId string, inviteCode string) (*TableData, error) {
	t, tableExists := r.Tables.ById[tableId]
	if !tableExists {
		return nil, NewCloudError(TableDoesNotExist)
	}
	if inviteCode != t.InviteCode {
		return nil, NewCloudError(InvalidInviteCode)
	}
	if r.Tables.ActiveTableOf(userId) != nil {
		return nil, NewCloudError(UserAlreadyAtOtherTable)
	}
	if err := t.Join(userId); err != nil {
		return nil, err
	}
	return GetData(t), nil

}
func (r *Room) StartTable(tableId string) error {
	t, tableExists := r.Tables.ById[tableId]
	if !tableExists {
		return NewCloudError(TableDoesNotExist)
	}
	return t.Start()
}

type Declaration struct {
	Gesund    bool
	Vorbehalt game.AnnouncedGameType
}
type MatchData struct {
	Phase           match.Phase
	Turn            game.Player
	Players         [game.NumPlayers]string
	InitialForehand game.Player
	Cards           game.Cards
	Declarations    map[game.Player]Declaration
	Bids            match.Bids
	CompletedTricks int
	CurrentTrick    *game.IncompleteTrick
	PreviousTrick   *game.Trick
	Mode            game.Mode
}

func GetMatchData(tm *TableMatch) *MatchData {
	decls := tm.Match.Auction.Declarations
	declarations := make(map[game.Player]Declaration, len(decls))
	for k, v := range decls {

		if v.Gesund {
			declarations[k] = Declaration{true, game.NormalspielType}
		} else {
			declarations[k] = Declaration{false, v.Vorbehalt.Type()}
		}
	}
	ans := &MatchData{Phase: tm.Match.Phase(), Turn: tm.Match.WhoseTurn(),
		Players: tm.Players, InitialForehand: tm.Match.InitialForehand(), Declarations: declarations,
		Bids: *tm.Match.Bids}
	if tm.Match.Phase() == match.InGame {
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
	} else {
		ans.Cards = tm.Match.DealtCards()
	}
	return ans
}
func (r *Room) GetMatchData(tableId string) (m *MatchData, err error) {
	t, ok := r.Tables.ById[tableId]
	if !ok {
		return nil, NewCloudError(TableDoesNotExist)
	}
	tm := t.CurrentMatch
	if tm == nil {
		return nil, NewCloudError(NoCurrentMatch)
	}
	return GetMatchData(tm), nil
}

func (r *Room) PlayCard(tableId string, userId string, card game.Card) (matchData *MatchData, err error) {
	m, player, err := r.getMatchAndPlayer(tableId, userId)
	if err != nil {
		return nil, err
	}
	if !m.Match.PlayCard(player, card) {
		return nil, NewCloudError(CannotPlayCard)
	}
	return GetMatchData(m), nil

}

func (r *Room) Declare(tableId string, userId string, gameType game.AnnouncedGameType) (*MatchData, error) {
	m, player, err := r.getMatchAndPlayer(tableId, userId)
	if err != nil {
		return nil, err
	}
	if !m.Match.AnnounceGameType(player, gameType) {
		return nil, NewCloudError(CannotAnnounce)
	}
	return GetMatchData(m), nil
}

func (r *Room) PlaceBid(tableId string, userId string, bid match.Bid) (*MatchData, error) {
	m, player, err := r.getMatchAndPlayer(tableId, userId)
	if err != nil {
		return nil, err
	}
	if !m.Match.PlaceBid(player, bid) {
		return nil, NewCloudError(CannotPlaceBid)
	}
	return GetMatchData(m), nil
}

func (r *Room) getMatchAndPlayer(tableId string, user string) (match *TableMatch, p game.Player, err error) {
	table, exists := r.Tables.ById[tableId]
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
