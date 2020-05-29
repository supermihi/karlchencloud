package match

import (
	"fmt"
	"github.com/supermihi/karlchencloud/doko/game"
)

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

	}
	panic(fmt.Sprintf("unsupported game type: %v", t))
}
