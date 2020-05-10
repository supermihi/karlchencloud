package modes

import (
	"github.com/supermihi/karlchencloud/pkg/game/core"
	"github.com/supermihi/karlchencloud/pkg/game/match"
)

func OnlyHochzeit() match.Sonderspiele {
	return match.MakeSonderspiele(match.VorbehaltHochzeit{})
}

func StandardSonderspiele() match.Sonderspiele {
	return match.MakeSonderspiele(match.VorbehaltHochzeit{},
		VorbehaltFarbsolo{core.Karo},
		VorbehaltFarbsolo{core.Herz},
		VorbehaltFarbsolo{core.Pik},
		VorbehaltFarbsolo{core.Kreuz},
		VorbehaltFleischlos{})
}
