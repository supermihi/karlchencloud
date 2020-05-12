package match

import "github.com/supermihi/karlchencloud/doko/game"

type ModeId string

type SonderspielMode interface {
	CanAnnounceWith(handCards game.Hand) bool
	Identifier() ModeId
	Priority() int
	CreateMode(announcer game.Player) game.Mode
	AnnouncerTakesForehand() bool
}

type Sonderspiele map[ModeId]SonderspielMode

func (r *Sonderspiele) FindSonderspiel(id ModeId) SonderspielMode {
	var result, ok = (*r)[id]
	if !ok {
		return nil
	}
	return result
}

func MakeSonderspiele(modes []SonderspielMode) Sonderspiele {
	result := make(Sonderspiele)
	for _, m := range modes {
		result[m.Identifier()] = m
	}
	return result
}
