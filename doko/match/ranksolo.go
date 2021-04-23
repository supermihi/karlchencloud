package match

import (
	"github.com/supermihi/karlchencloud/doko/game"
)

type RankSolo struct {
	Trump   game.Rank
	Soloist game.Player
}

func NewRankSolo(trump game.Rank, soloist game.Player) RankSolo {
	return RankSolo{trump, soloist}
}

func RankSoloSuit(trump game.Rank, card game.Card) game.GameSuit {
	if card.Rank == trump {
		return game.Trump
	}
	return card.Suit.AsNonTrump()
}

func (r RankSolo) GameSuit(card game.Card) game.GameSuit {
	return RankSoloSuit(r.Trump, card)
}

func (r RankSolo) PartyOf(p game.Player) game.Party {
	if p == r.Soloist {
		return game.ReParty
	}
	return game.ContraParty
}

func (RankSolo) ClarificationTrick() int {
	return -1
}

func (RankSolo) OnCompletedTrick(game.Trick, int) {
}

func (r RankSolo) Type() game.AnnouncedGameType {
	return gameTypeForRank(r.Trump)
}

func gameTypeForRank(rank game.Rank) game.AnnouncedGameType {
	switch rank {
	case game.Jack:
		return game.JackSoloType
	case game.Queen:
		return game.QueenSoloType
	}
	panic("unexpected solo rank")
}

type ReservationRankSolo struct {
	rank game.Rank
}

func (v ReservationRankSolo) CanAnnounceWith(_ game.Hand) bool {
	return true
}
func (v ReservationRankSolo) Type() game.AnnouncedGameType {
	return gameTypeForRank(v.rank)
}

func (v ReservationRankSolo) Priority() int {
	return ReservationPriorityMarriage + 1
}

func (v ReservationRankSolo) CreateMode(announcer game.Player) game.Mode {
	return NewRankSolo(v.rank, announcer)
}

func (v ReservationRankSolo) AnnouncerTakesForehand() bool {
	return false
}
