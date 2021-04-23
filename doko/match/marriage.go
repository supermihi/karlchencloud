package match

import "github.com/supermihi/karlchencloud/doko/game"

type ReservationMarriage struct {
}

func (v ReservationMarriage) Type() game.AnnouncedGameType {
	return game.MarriageType
}

func (v ReservationMarriage) CanAnnounceWith(handCards game.Hand) bool {
	return handCards.NumQueensOfClubs() == 2
}

const ReservationPriorityMarriage = 1

func (v ReservationMarriage) Priority() int {
	return ReservationPriorityMarriage
}

func (v ReservationMarriage) CreateMode(announcer game.Player) game.Mode {
	return game.NewMarriage(announcer)
}

func (v ReservationMarriage) AnnouncerTakesForehand() bool {
	return false
}
