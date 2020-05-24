package common

import (
	"fmt"
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/cloud"
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

func ToApiBid(b match.Bid) api.BidType {
	switch b {
	case match.Re:
		return api.BidType_RE_BID
	case match.ReKeine90:
		return api.BidType_RE_NO_NINETY
	case match.ReKeine60:
		return api.BidType_CONTRA_NO_SIXTY
	case match.ReKeine30:
		return api.BidType_RE_NO_THIRTY
	case match.ReSchwarz:
		return api.BidType_RE_SCHWARZ
	case match.Contra:
		return api.BidType_CONTRA_BID
	case match.ContraKeine90:
		return api.BidType_CONTRA_NO_NINETY
	case match.ContraKeine60:
		return api.BidType_CONTRA_NO_SIXTY
	case match.ContraKeine30:
		return api.BidType_CONTRA_NO_THIRTY
	case match.ContraSchwarz:
		return api.BidType_CONTRA_SCHWARZ
	default:
		panic(fmt.Sprintf("unexpected bid %v in ToApiBid", b))
	}
}

func ToBid(b api.BidType) match.Bid {
	switch b {
	case api.BidType_RE_BID:
		return match.Re
	case api.BidType_CONTRA_BID:
		return match.Contra
	case api.BidType_RE_NO_NINETY:
		return match.ReKeine90
	case api.BidType_CONTRA_NO_NINETY:
		return match.ContraKeine90
	case api.BidType_RE_NO_SIXTY:
		return match.ReKeine60
	case api.BidType_CONTRA_NO_SIXTY:
		return match.ContraKeine60
	case api.BidType_RE_NO_THIRTY:
		return match.ReKeine30
	case api.BidType_CONTRA_NO_THIRTY:
		return match.ContraKeine30
	case api.BidType_RE_SCHWARZ:
		return match.ReSchwarz
	case api.BidType_CONTRA_SCHWARZ:
		return match.ContraSchwarz
	}
	panic(fmt.Sprintf("unexpected api bid %v in ToBid()", b))
}

func ToApiGameType(t game.AnnouncedGameType) api.GameType {
	switch t {
	case game.NormalspielType:
		return api.GameType_NORMAL_GAME
	case game.HochzeitType:
		return api.GameType_MARRIAGE
	case game.KaroSoloType:
		return api.GameType_DIAMONDS_SOLO
	case game.HerzSoloType:
		return api.GameType_HEARTS_SOLO
	case game.PikSoloType:
		return api.GameType_SPADES_SOLO
	case game.KreuzSoloType:
		return api.GameType_CLUBS_SOLO
	case game.FleischlosType:
		return api.GameType_MEATLESS_SOLO
	}
	panic(fmt.Sprintf("not a solo type: %v", t))
}

func ToGameType(t api.GameType) game.AnnouncedGameType {
	switch t {
	case api.GameType_NORMAL_GAME:
		return game.NormalspielType
	case api.GameType_MARRIAGE:
		return game.HochzeitType
	case api.GameType_DIAMONDS_SOLO:
		return game.KaroSoloType
	case api.GameType_HEARTS_SOLO:
		return game.HerzSoloType
	case api.GameType_SPADES_SOLO:
		return game.PikSoloType
	case api.GameType_CLUBS_SOLO:
		return game.KreuzSoloType
	case api.GameType_MEATLESS_SOLO:
		return game.FleischlosType
	}
	panic(fmt.Sprintf("not an api game type: %s", t))
}

func ToApiMode(mode game.Mode) *api.Mode {
	soloist := api.Player_PLAYER_1
	if !game.IsNormalspiel(mode.Type()) {
		soloist = ToApiPlayer(game.Soloist(mode), false)
	}
	spouse := api.Player_PLAYER_1
	switch h := mode.(type) {
	case game.Hochzeit:
		if h.PartnerFound() {
			spouse = ToApiPlayer(h.Partner(), false)
		}

	}
	return &api.Mode{Type: ToApiGameType(mode.Type()), Soloist: soloist, Spouse: spouse}
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

func ToSuit(s api.Suit) game.Suit {
	switch s {
	case api.Suit_DIAMONDS:
		return game.Karo
	case api.Suit_HEARTS:
		return game.Herz
	case api.Suit_SPADES:
		return game.Pik
	case api.Suit_CLUBS:
		return game.Kreuz
	}
	panic(fmt.Sprintf("unexpected apisuit %v in ToSuit()", s))
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

func ToRank(r api.Rank) game.Rank {
	switch r {
	case api.Rank_NINE:
		return game.Neun
	case api.Rank_JACK:
		return game.Bube
	case api.Rank_QUEEN:
		return game.Dame
	case api.Rank_KING:
		return game.Koenig
	case api.Rank_TEN:
		return game.Zehn
	case api.Rank_ACE:
		return game.Ass
	}
	panic(fmt.Sprintf("unexpected apirank %v in ToRank()", r))
}

func ToApiCard(c game.Card) *api.Card {
	return &api.Card{Suit: ToApiSuit(c.Suit), Rank: ToApiRank(c.Rank)}
}

func ToCard(c *api.Card) game.Card {
	return game.Card{Suit: ToSuit(c.Suit), Rank: ToRank(c.Rank)}
}

func ToApiTrick(t *game.IncompleteTrick, m game.Mode) *api.Trick {
	result := &api.Trick{Forehand: ToApiPlayer(t.Forehand, false)}
	if c, ok := t.CardOf(game.Player1); ok {
		result.CardPlayer_1 = ToApiCard(c)
	}
	if c, ok := t.CardOf(game.Player2); ok {
		result.CardPlayer_2 = ToApiCard(c)
	}
	if c, ok := t.CardOf(game.Player3); ok {
		result.CardPlayer_3 = ToApiCard(c)
	}
	if c, ok := t.CardOf(game.Player4); ok {
		result.CardPlayer_4 = ToApiCard(c)
	}
	if t.IsComplete() {
		result.Winner = ToApiPlayer(game.WinnerOfTrick(t.CardsByPlayer(), t.Forehand, m), false)
	}
	return result
}

func toAuctionState(a *match.Auction) *api.AuctionState {
	declarations := make([]*api.Declaration, 0)
	for _, p := range game.Players() {
		decl, hasDeclared := a.DeclarationOf(p)
		if hasDeclared {
			apiDecl := &api.Declaration{Player: ToApiPlayer(p, false), Vorbehalt: !decl.Gesund}
			declarations = append(declarations, apiDecl)
		}
	}
	return &api.AuctionState{Declarations: declarations}
}

func toApiBids(bids *match.Bids) []*api.Bid {
	var ans []*api.Bid
	for _, player := range game.Players() {
		bidsOf := bids.BidsOf(player)
		apiBidsOf := make([]*api.Bid, len(bidsOf))
		for i, b := range bidsOf {
			apiBidsOf[i] = &api.Bid{Player: ToApiPlayer(player, false), Bid: ToApiBid(b)}
		}
		ans = append(ans, apiBidsOf...)
	}
	return ans
}

func ToGameState(m *match.Match) *api.GameState {
	return &api.GameState{Mode: ToApiMode(m.Mode()),
		Bids:            toApiBids(m.Bids),
		CompletedTricks: int32(m.Game.NumCompletedTricks()),
		CurrentTrick:    ToApiTrick(m.Game.CurrentTrick, m.Mode())}
}
func ToMatchState(tm *cloud.TableMatch, user cloud.UserId) *api.MatchState {
	m := tm.Match
	turn := ToApiPlayer(m.WhoseTurn(), true)
	players := &api.Players{
		Player_1: string(tm.Players[game.Player1]),
		Player_2: string(tm.Players[game.Player2]),
		Player_3: string(tm.Players[game.Player3]),
		Player_4: string(tm.Players[game.Player4]),
	}
	ans := &api.MatchState{Turn: turn, Players: players}
	addDetails(ans, m)
	addRole(ans, tm, user)
	return ans

}

func addRole(state *api.MatchState, tm *cloud.TableMatch, user cloud.UserId) {
	isPlayer := false
	for pInd, playerUser := range tm.Players {
		if user == playerUser {
			player := game.Player(pInd)
			playerData := &api.MatchPlayerData{
				Cards:  GetHandCards(tm.Match, player),
				Player: ToApiPlayer(player, false)}
			state.Role = &api.MatchState_PlayerData{PlayerData: playerData}
			isPlayer = true
			break
		}
	}
	if !isPlayer {
		state.Role = &api.MatchState_Spectator{Spectator: &api.Empty{}}
	}
}

func addDetails(state *api.MatchState, m *match.Match) {
	switch m.Phase() {
	case match.InAuction:
		state.Phase = api.MatchPhase_AUCTION
		auctionState := toAuctionState(m.Auction)
		state.Details = &api.MatchState_AuctionState{AuctionState: auctionState}
	case match.InGame:
		state.Phase = api.MatchPhase_GAME
		gameState := ToGameState(m)
		state.Details = &api.MatchState_GameState{GameState: gameState}
	case match.MatchFinished:
		state.Phase = api.MatchPhase_FINISHED
	default:
		panic(fmt.Sprintf("ToMatchState called with invalid match phase %v", m.Phase()))
	}
}

func GetHandCards(m *match.Match, p game.Player) []*api.Card {
	if m.Phase() != match.InGame {
		return nil
	}
	cards := m.Game.HandCards[p]
	ans := make([]*api.Card, len(cards))
	for i, card := range cards {
		ans[i] = ToApiCard(card)
	}
	return ans
}

func ToTableData(table *cloud.Table, user cloud.UserId) *api.TableData {
	exposedInviteCode := ""
	if table.Owner() == user {
		exposedInviteCode = table.InviteCode
	}
	return &api.TableData{TableId: table.Id, Owner: string(table.Owner()), InviteCode: exposedInviteCode}
}
