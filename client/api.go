package client

import "github.com/supermihi/karlchencloud/doko/game"

// ClientApi is the interface used by client implementations.
// It encapsulates all GRPC details.
type ClientApi interface {
	Logf(format string, v ...interface{})
	Table() *TableView
	Match() *MatchView
	User() UserData
	CreateTable() error
	JoinTable(invite string) (err error)
	StartTable() error
	PlayCard(i int) error
	Declare(t game.AnnouncedGameType) error
	StartNextMatch() error
}
