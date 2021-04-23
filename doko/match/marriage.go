package match

import "github.com/supermihi/karlchencloud/doko/game"

type VorbehaltMarriage struct {
}

func (v VorbehaltMarriage) Type() game.AnnouncedGameType {
	return game.MarriageType
}

func (v VorbehaltMarriage) CanAnnounceWith(handCards game.Hand) bool {
	return handCards.NumAlte() == 2
}

const VorbehaltPriorityMarriage = 1

func (v VorbehaltMarriage) Priority() int {
	return VorbehaltPriorityMarriage
}

func (v VorbehaltMarriage) CreateMode(announcer game.Player) game.Mode {
	return game.NewMarriage(announcer)
}

func (v VorbehaltMarriage) AnnouncerTakesForehand() bool {
	return false
}
