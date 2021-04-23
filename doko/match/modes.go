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
	case game.NormalGameType:
		return nil
	case game.MarriageType:
		return VorbehaltMarriage{}
	case game.AceSoloType:
		return VorbehaltFleischlos{}
	case game.DiamondSoloType:
		return VorbehaltSuitSolo{game.Karo}
	case game.HeartSoloType:
		return VorbehaltSuitSolo{game.Herz}
	case game.SpadeSoloType:
		return VorbehaltSuitSolo{game.Pik}
	case game.ClubSoloType:
		return VorbehaltSuitSolo{game.Kreuz}
	case game.JackSoloType:
		return VorbehaltRankSolo{game.Bube}
	case game.QueenSoloType:
		return VorbehaltRankSolo{game.Dame}
	}
	panic(fmt.Sprintf("unexpected game type %v in GetVorbehalt", t))
}

func GameSuitOf(card game.Card, t game.AnnouncedGameType) game.GameSuit {
	switch t {
	case game.NormalGameType:
		fallthrough
	case game.MarriageType:
		return game.NormalGameSuit(card)
	case game.AceSoloType:
		return FleischlosSuite(card)
	case game.DiamondSoloType:
		return SuitSoloSuit(game.Karo, card)
	case game.HeartSoloType:
		return SuitSoloSuit(game.Herz, card)
	case game.SpadeSoloType:
		return SuitSoloSuit(game.Pik, card)
	case game.ClubSoloType:
		return SuitSoloSuit(game.Kreuz, card)
	case game.JackSoloType:
		return RankSoloSuit(game.Bube, card)
	case game.QueenSoloType:
		return RankSoloSuit(game.Dame, card)
	}
	panic(fmt.Sprintf("unsupported game type: %v", t))
}
