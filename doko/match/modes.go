package match

import (
	"fmt"
	"github.com/supermihi/karlchencloud/doko/game"
)

type Reservation interface {
	CanAnnounceWith(handCards game.Hand) bool
	Type() game.AnnouncedGameType
	Priority() int
	CreateMode(announcer game.Player) game.Mode
	AnnouncerTakesForehand() bool
}

func GetReservation(t game.AnnouncedGameType) Reservation {
	switch t {
	case game.NormalGameType:
		return nil
	case game.MarriageType:
		return ReservationMarriage{}
	case game.AceSoloType:
		return ReservationAceSolo{}
	case game.DiamondSoloType:
		return ReservationSuitSolo{game.Diamonds}
	case game.HeartSoloType:
		return ReservationSuitSolo{game.Hearts}
	case game.SpadeSoloType:
		return ReservationSuitSolo{game.Spades}
	case game.ClubSoloType:
		return ReservationSuitSolo{game.Clubs}
	case game.JackSoloType:
		return ReservationRankSolo{game.Jack}
	case game.QueenSoloType:
		return ReservationRankSolo{game.Queen}
	}
	panic(fmt.Sprintf("unexpected game type %v in GetReservation", t))
}

func GameSuitOf(card game.Card, t game.AnnouncedGameType) game.GameSuit {
	switch t {
	case game.NormalGameType:
		fallthrough
	case game.MarriageType:
		return game.NormalGameSuit(card)
	case game.AceSoloType:
		return AceSoloSuite(card)
	case game.DiamondSoloType:
		return SuitSoloSuit(game.Diamonds, card)
	case game.HeartSoloType:
		return SuitSoloSuit(game.Hearts, card)
	case game.SpadeSoloType:
		return SuitSoloSuit(game.Spades, card)
	case game.ClubSoloType:
		return SuitSoloSuit(game.Clubs, card)
	case game.JackSoloType:
		return RankSoloSuit(game.Jack, card)
	case game.QueenSoloType:
		return RankSoloSuit(game.Queen, card)
	}
	panic(fmt.Sprintf("unsupported game type: %v", t))
}
