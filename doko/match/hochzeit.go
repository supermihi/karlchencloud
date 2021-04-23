package match

import "github.com/supermihi/karlchencloud/doko/game"

type VorbehaltHochzeit struct {
}

func (v VorbehaltHochzeit) Type() game.AnnouncedGameType {
	return game.MarriageType
}

func (v VorbehaltHochzeit) CanAnnounceWith(handCards game.Hand) bool {
	return handCards.NumAlte() == 2
}

const VorbehaltPriorityHochzeit = 1

func (v VorbehaltHochzeit) Priority() int {
	return VorbehaltPriorityHochzeit
}

func (v VorbehaltHochzeit) CreateMode(announcer game.Player) game.Mode {
	return game.NewHochzeit(announcer)
}

func (v VorbehaltHochzeit) AnnouncerTakesForehand() bool {
	return false
}
