package match

import "github.com/supermihi/karlchencloud/pkg/game/core"

type VorbehaltHochzeit struct {
}

func (v VorbehaltHochzeit) Identifier() ModeId {
	return "HOCHZEIT"
}

func (v VorbehaltHochzeit) CanAnnounceWith(handCards core.Hand) bool {
	return handCards.NumAlte() == 2
}

const VORBEHALT_PRIORITY_HOCHZEIT = 1

func (v VorbehaltHochzeit) Priority() int {
	return VORBEHALT_PRIORITY_HOCHZEIT
}

func (v VorbehaltHochzeit) CreateMode(announcer core.Player) core.Mode {
	return core.NewHochzeit(announcer)
}

func (v VorbehaltHochzeit) AnnouncerTakesForehand() bool {
	return false
}
