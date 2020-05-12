package match

import "github.com/supermihi/karlchencloud/doko/game"

type VorbehaltHochzeit struct {
}

func (v VorbehaltHochzeit) Identifier() ModeId {
	return "HOCHZEIT"
}

func (v VorbehaltHochzeit) CanAnnounceWith(handCards game.Hand) bool {
	return handCards.NumAlte() == 2
}

const VORBEHALT_PRIORITY_HOCHZEIT = 1

func (v VorbehaltHochzeit) Priority() int {
	return VORBEHALT_PRIORITY_HOCHZEIT
}

func (v VorbehaltHochzeit) CreateMode(announcer game.Player) game.Mode {
	return game.NewHochzeit(announcer)
}

func (v VorbehaltHochzeit) AnnouncerTakesForehand() bool {
	return false
}
