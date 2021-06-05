package errors

import "fmt"

type CloudError struct {
	msg  string
	code ErrorCode
}

func (c CloudError) Error() string {
	return fmt.Sprintf("Error %d: %s", c.code, c.msg)
}

type ErrorCode int

const (
	TableDoesNotExist ErrorCode = iota
	TableAlreadyExists
	UserAlreadyAtTable
	UserAlreadyAtOtherTable
	UserNotAtTable
	NotOwnerOfTable
	UnableToJoinStartedTable
	TablePlayerLimitReached
	NoCurrentMatch
	UserNotPlayingInMatch
	NoActiveTable
	UserDoesNotExist
	InvalidInviteCode
	TableAlreadyStarted
	InvalidNumberOfPlayers
	CannotStartTableNow
	CannotEndMatch
	CannotPlayCard
	CannotAnnounce
	CannotPlaceBid
	StartedAnotherSession
	InvalidSessionToken
	AuthenticationFailed
)

func NewCloudError(c ErrorCode) error {
	return &CloudError{c.Message(), c}
}
func (c ErrorCode) Message() string {
	switch c {
	case TableDoesNotExist:
		return "Table does not exist"
	case NotOwnerOfTable:
		return "The user does not own this table"
	case NoCurrentMatch:
		return "No current match at table"
	case TableAlreadyExists:
		return "The table you are trying to create already exists"
	case UserAlreadyAtTable:
		return "The user is already at this table"
	case UserAlreadyAtOtherTable:
		return "The user is already at another table"
	case UserNotAtTable:
		return "The user is not a member of this table"
	case UnableToJoinStartedTable:
		return "Cannot join a table that has already started"
	case TablePlayerLimitReached:
		return "The maximum number of players at the table has been reached"
	case NoActiveTable:
		return "The user is not currently at any table"
	case UserDoesNotExist:
		return "The user does not exist"
	case InvalidInviteCode:
		return "The invite code is invalid"
	case TableAlreadyStarted:
		return "The table was already started"
	case InvalidNumberOfPlayers:
		return "Cannot start table: invalid number of players"
	case CannotStartTableNow:
		return "can only start match in phase 'WaitingForNextGame'"
	case UserNotPlayingInMatch:
		return "The user is not playing in the current match but spectating"
	case CannotPlayCard:
		return "Cannot play card"
	case CannotAnnounce:
		return "Cannot annonuce game type"
	case CannotPlaceBid:
		return "Cannot place bid"
	case StartedAnotherSession:
		return "Closing session because another one was started"
	case InvalidSessionToken:
		return "Session token is invalid"
	case AuthenticationFailed:
		return "Authentication failed"
	}
	return "unknown error code"
}
