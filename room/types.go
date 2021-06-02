package room

import (
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
)

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
