package pbconv

import (
	"fmt"
	pb "github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
	"github.com/supermihi/karlchencloud/server/tables"
)

func ToBid(b pb.BidType) match.Bid {
	switch b {
	case pb.BidType_RE_BID:
		return match.Re
	case pb.BidType_CONTRA_BID:
		return match.Contra
	case pb.BidType_RE_NO_NINETY:
		return match.ReKeine90
	case pb.BidType_CONTRA_NO_NINETY:
		return match.ContraKeine90
	case pb.BidType_RE_NO_SIXTY:
		return match.ReKeine60
	case pb.BidType_CONTRA_NO_SIXTY:
		return match.ContraKeine60
	case pb.BidType_RE_NO_THIRTY:
		return match.ReKeine30
	case pb.BidType_CONTRA_NO_THIRTY:
		return match.ContraKeine30
	case pb.BidType_RE_SCHWARZ:
		return match.ReSchwarz
	case pb.BidType_CONTRA_SCHWARZ:
		return match.ContraSchwarz
	}
	panic(fmt.Sprintf("unexpected api bid %v in ToBid()", b))
}

func ToGameType(t pb.GameType) game.AnnouncedGameType {
	switch t {
	case pb.GameType_NORMAL_GAME:
		return game.NormalGameType
	case pb.GameType_MARRIAGE:
		return game.MarriageType
	case pb.GameType_DIAMONDS_SOLO:
		return game.DiamondSoloType
	case pb.GameType_HEARTS_SOLO:
		return game.HeartSoloType
	case pb.GameType_SPADES_SOLO:
		return game.SpadeSoloType
	case pb.GameType_CLUBS_SOLO:
		return game.ClubSoloType
	case pb.GameType_MEATLESS_SOLO:
		return game.AceSoloType
	}
	panic(fmt.Sprintf("not an api game type: %s", t))
}

func ToSuit(s pb.Suit) game.Suit {
	switch s {
	case pb.Suit_DIAMONDS:
		return game.Diamonds
	case pb.Suit_HEARTS:
		return game.Hearts
	case pb.Suit_SPADES:
		return game.Spades
	case pb.Suit_CLUBS:
		return game.Clubs
	}
	panic(fmt.Sprintf("unexpected apisuit %v in ToSuit()", s))
}

func ToRank(r pb.Rank) game.Rank {
	switch r {
	case pb.Rank_NINE:
		return game.Nine
	case pb.Rank_JACK:
		return game.Jack
	case pb.Rank_QUEEN:
		return game.Queen
	case pb.Rank_KING:
		return game.King
	case pb.Rank_TEN:
		return game.Ten
	case pb.Rank_ACE:
		return game.Ace
	}
	panic(fmt.Sprintf("unexpected apirank %v in ToRank()", r))
}

func ToCard(c *pb.Card) game.Card {
	return game.Card{Suit: ToSuit(c.Suit), Rank: ToRank(c.Rank)}
}

func addRole(state *pb.MatchState, m *tables.MatchData, self game.Player) {
	if self == game.NoPlayer {
		state.Role = &pb.MatchState_Spectator{Spectator: &pb.Empty{}}
	} else {
		cards := ToPbCards(m.Cards[self])
		state.Role = &pb.MatchState_OwnCards{OwnCards: &pb.Cards{Cards: cards}}
	}
}

func ToMatchPhase(p pb.MatchPhase) match.Phase {
	switch p {
	case pb.MatchPhase_AUCTION:
		return match.InAuction
	case pb.MatchPhase_GAME:
		return match.InGame
	case pb.MatchPhase_FINISHED:
		return match.MatchFinished
	}
	panic(fmt.Sprintf("unknown match phase: %v", p))
}
