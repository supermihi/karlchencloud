package match

import (
	"fmt"
	"github.com/supermihi/karlchencloud/doko/game"
)

type Vorbehalt interface {
	CanAnnounceWith(handCards game.Hand) bool
	Type() game.AnnouncedGameType
	Priority() int
	CreateMode(announcer game.Player) game.Mode
	AnnouncerTakesForehand() bool
}

func GetVorbehalt(t game.AnnouncedGameType) Vorbehalt {
	switch t {
	case game.NormalspielType:
		return nil
	case game.HochzeitType:
		return VorbehaltHochzeit{}
	case game.FleischlosType:
		return VorbehaltFleischlos{}
	case game.KaroSoloType:
		return VorbehaltFarbsolo{game.Karo}
	case game.HerzSoloType:
		return VorbehaltFarbsolo{game.Herz}
	case game.PikSoloType:
		return VorbehaltFarbsolo{game.Pik}
	case game.KreuzSoloType:
		return VorbehaltFarbsolo{game.Kreuz}
	}
	panic(fmt.Sprintf("unexpected game type %v in GetVorbehalt", t))
}
