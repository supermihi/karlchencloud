package room

import (
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
)

type IRoom interface {
	CreateTable(owner UserId, public bool, seed int64) (*TableData, error)
	GetTable(TableId) (*TableData, error)
	// GetOpenTables lists all tables that are not yet started and visible to the given user
	GetOpenTables(user UserId) []*TableData
	ActiveTableOf(UserId) *TableData
	JoinTableByInviteCode(userId UserId, inviteCode string) (*TableData, error)
	JoinTableByTableId(userId UserId, tableId TableId) (*TableData, error)
	StartTable(TableId, UserId) (*TableData, error)
	GetMatchData(TableId) (*MatchData, error)
	PlayCard(TableId, UserId, game.Card) (*MatchData, error)
	StartNextMatch(TableId, UserId) (*MatchData, error)
	PlaceBid(TableId, UserId, match.Bid) (*MatchData, error)
	Declare(TableId, UserId, game.AnnouncedGameType) (*MatchData, error)
	RelatedUsers(userId UserId) []UserId
}
