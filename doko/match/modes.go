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
	case game.JacksSoloType:
		return VorbehaltRankSolo{game.Bube}
	case game.QueensSoloType:
		return VorbehaltRankSolo{game.Dame}
	}
	panic(fmt.Sprintf("unexpected game type %v in GetVorbehalt", t))
}

func GameSuitOf(card game.Card, t game.AnnouncedGameType) game.GameSuit {
	switch t {
	case game.NormalspielType:
		fallthrough
	case game.HochzeitType:
		return game.NormalGameSuit(card)
	case game.FleischlosType:
		return FleischlosSuite(card)
	case game.KaroSoloType:
		return FarbsoloSuit(game.Karo, card)
	case game.HerzSoloType:
		return FarbsoloSuit(game.Herz, card)
	case game.PikSoloType:
		return FarbsoloSuit(game.Pik, card)
	case game.KreuzSoloType:
		return FarbsoloSuit(game.Kreuz, card)
	case game.JacksSoloType:
		return RankSoloSuit(game.Bube, card)
	case game.QueensSoloType:
		return RankSoloSuit(game.Dame, card)
	}
	panic(fmt.Sprintf("unsupported game type: %v", t))
}
