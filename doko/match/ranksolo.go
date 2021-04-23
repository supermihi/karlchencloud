package match

import (
	"github.com/supermihi/karlchencloud/doko/game"
)

type RankSolo struct {
	Trumpf  game.Rank
	Soloist game.Player
}

func NewRankSolo(trump game.Rank, soloist game.Player) RankSolo {
	return RankSolo{trump, soloist}
}

func RankSoloSuit(trump game.Rank, card game.Card) game.GameSuit {
	if card.Rank == trump {
		return game.Trumpf
	}
	return card.Suit.AsFehl()
}

func (r RankSolo) GameSuit(card game.Card) game.GameSuit {
	return RankSoloSuit(r.Trumpf, card)
}

func (r RankSolo) PartyOf(p game.Player) game.Party {
	if p == r.Soloist {
		return game.ReParty
	}
	return game.ContraParty
}

func (RankSolo) Klaerungsstich() int {
	return -1
}

func (RankSolo) OnCompletedTrick(game.Trick, int) {
}

func (r RankSolo) Type() game.AnnouncedGameType {
	return gameTypeForRank(r.Trumpf)
}

func gameTypeForRank(rank game.Rank) game.AnnouncedGameType {
	switch rank {
	case game.Bube:
		return game.JacksSoloType
	case game.Dame:
		return game.QueensSoloType
	}
	panic("unexpected solo rank")
}

type VorbehaltRankSolo struct {
	rank game.Rank
}

func (v VorbehaltRankSolo) CanAnnounceWith(_ game.Hand) bool {
	return true
}
func (v VorbehaltRankSolo) Type() game.AnnouncedGameType {
	return gameTypeForRank(v.rank)
}

func (v VorbehaltRankSolo) Priority() int {
	return VorbehaltPriorityHochzeit + 1
}

func (v VorbehaltRankSolo) CreateMode(announcer game.Player) game.Mode {
	return NewRankSolo(v.rank, announcer)
}

func (v VorbehaltRankSolo) AnnouncerTakesForehand() bool {
	return false
}
