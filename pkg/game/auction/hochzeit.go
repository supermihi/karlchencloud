package auction

import "github.com/supermihi/karlchencloud/pkg/game/core"

type VorbehaltHochzeit struct {
}

func (v VorbehaltHochzeit) Identifier() ModeId {
	return "HOCHZEIT"
}

func (v VorbehaltHochzeit) CanAnnounceWith(handCards core.Hand) bool {
	return handCards.NumAlte() == 2
}

func (v VorbehaltHochzeit) Priority() int {
	return 1
}

func (v VorbehaltHochzeit) CreateMode(announcer core.Player) core.Mode {
	return core.NewHochzeit(announcer)
}

func (v VorbehaltHochzeit) AnnouncerTakesForehand() bool {
	return false
}
