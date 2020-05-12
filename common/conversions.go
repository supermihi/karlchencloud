package common

import (
	"fmt"
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
)

func ToApiPlayer(p game.Player, mapNoPlayerToPlayer1 bool) api.Player {
	switch p {
	case game.Player1:
		return api.Player_PLAYER_1
	case game.Player2:
		return api.Player_PLAYER_2
	case game.Player3:
		return api.Player_PLAYER_3
	case game.Player4:
		return api.Player_PLAYER_4
	default:
		if mapNoPlayerToPlayer1 {
			return api.Player_PLAYER_1
		}
		panic("unexpected NoPlayer in ToApiPlayer")
	}
}

func ToAuctionPhase(p match.AuctionPhase) api.AuctionPhase {
	switch p {
	case match.VorbehaltAbfrage:
		return api.AuctionPhase_DECLARATION
	case match.VorbehaltSpezifikation:
		return api.AuctionPhase_SPECIFICATION
	}
	panic("unsupported auction phase")
}

func toSoloType(t game.AnnouncedGameType) api.SoloType {
	switch t {
	case game.KaroSoloType:
		return api.SoloType_DIAMONDS_SOLO
	case game.HerzSoloType:
		return api.SoloType_HEARTS_SOLO
	case game.PikSoloType:
		return api.SoloType_SPADES_SOLO
	case game.KreuzSoloType:
		return api.SoloType_CLUBS_SOLO
	}
	panic(fmt.Sprintf("not a solo type: %v", t))
}
func ToApiMode(mode game.Mode) *api.Mode {
	var gameType api.GameType
	var soloInfo *api.SoloInfo
	switch mode.(type) {
	case game.NormalspielMode:
		gameType = api.GameType_NORMAL_GAME
	case game.Hochzeit:
		gameType = api.GameType_MARRIAGE
	default:
		gameType = api.GameType_VOLUNTARY_SOLO
		soloInfo = &api.SoloInfo{Soloist: ToApiPlayer(game.Soloist(mode), false), SoloType: toSoloType(mode.Type())}
	}
	return &api.Mode{Type: gameType, SoloInfo: soloInfo}
}

func ToMatchPhase(phase match.Phase) api.MatchPhase {
	switch phase {
	case match.InAuction:
		return api.MatchPhase_AUCTION
	case match.InGame:
		return api.MatchPhase_GAME
	case match.MatchFinished:
		return api.MatchPhase_FINISHED
	default:
		panic(fmt.Sprintf("unexpected phase %v in toMatchPhase", phase))
	}
}

func ToApiSuit(s game.Suit) api.Suit {
	switch s {
	case game.Karo:
		return api.Suit_DIAMONDS
	case game.Herz:
		return api.Suit_HEARTS
	case game.Pik:
		return api.Suit_SPADES
	case game.Kreuz:
		return api.Suit_CLUBS
	}
	panic(fmt.Sprintf("unexpected suit %v", s))
}

func ToApiRank(r game.Rank) api.Rank {
	switch r {
	case game.Neun:
		return api.Rank_NINE
	case game.Bube:
		return api.Rank_JACK
	case game.Dame:
		return api.Rank_QUEEN
	case game.Koenig:
		return api.Rank_KING
	case game.Zehn:
		return api.Rank_TEN
	case game.Ass:
		return api.Rank_ACE
	}
	panic(fmt.Sprintf("unexpected rank %v", r))
}

func ToApiCard(c game.Card) *api.Card {
	return &api.Card{Suit: ToApiSuit(c.Suit), Rank: ToApiRank(c.Rank)}
}

func ToApiTrick(t *game.IncompleteTrick, m game.Mode) *api.Trick {
	result := &api.Trick{Forehand: ToApiPlayer(t.Forehand, false)}
	if c, ok := t.TryCardOf(game.Player1); ok {
		result.CardPlayer_1 = ToApiCard(c)
	}
	if c, ok := t.TryCardOf(game.Player2); ok {
		result.CardPlayer_2 = ToApiCard(c)
	}
	if c, ok := t.TryCardOf(game.Player3); ok {
		result.CardPlayer_3 = ToApiCard(c)
	}
	if c, ok := t.TryCardOf(game.Player4); ok {
		result.CardPlayer_4 = ToApiCard(c)
	}
	if t.IsComplete() {
		result.Winner = ToApiPlayer(game.WinnerOfTrick(t.CardsByPlayer(), t.Forehand, m), false)
	}
	return result
}
