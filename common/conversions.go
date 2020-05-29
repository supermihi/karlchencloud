package common

import (
	"fmt"
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/cloud"
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
)

func ToApiUserId(p game.Player, users cloud.PlayerUserMap) string {
	if p == game.NoPlayer {
		panic("cannot convert NoPlayer to user id")
	}
	return string(users[p])
}
func ToPlayerValue(p game.Player, users cloud.PlayerUserMap) *api.PlayerValue {
	if p != game.NoPlayer {
		return &api.PlayerValue{UserId: string(users[p])}
	}
	return nil
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

func ToApiMode(mode game.Mode, forehand game.Player, users cloud.PlayerUserMap) *api.Mode {
	soloist := ToPlayerValue(game.Soloist(mode), users)
	var spouse *api.PlayerValue
	switch h := mode.(type) {
	case game.Hochzeit:
		spouse = ToPlayerValue(h.Partner(), users)
	}
	return &api.Mode{Type: ToApiGameType(mode.Type()), Soloist: soloist, Spouse: spouse,
		Forehand: users[int(forehand)]}
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

func ToApiTrick(t *game.IncompleteTrick, m game.Mode, users cloud.PlayerUserMap) *api.Trick {
	result := &api.Trick{UserIdForehand: ToApiUserId(t.Forehand, users)}
	cards := make([]*api.Card, t.NumCardsPlayed())
	for i := 0; i < len(cards); i++ {
		cards[i] = ToApiCard(t.NthCard(i))
	}
	if t.IsComplete() {
		result.UserIdWinner = ToPlayerValue(game.WinnerOfTrick(t.CardsByPlayer(), t.Forehand, m), users)
	}
	return result
}

func ToApiTrickC(t game.Trick, users cloud.PlayerUserMap) *api.Trick {
	result := &api.Trick{UserIdForehand: ToApiUserId(t.Forehand, users)}
	cards := make([]*api.Card, game.NumPlayers)
	for i := 0; i < len(cards); i++ {
		cards[i] = ToApiCard(t.NthCard(i))
	}
	result.UserIdWinner = ToPlayerValue(t.Winner, users)
	return result
}

func toAuctionState(a *match.Auction, users cloud.PlayerUserMap) *api.AuctionState {
	declarations := make([]*api.Declaration, 0)
	for _, p := range game.Players() {
		decl, hasDeclared := a.DeclarationOf(p)
		if hasDeclared {
			apiDecl := &api.Declaration{UserId: ToApiUserId(p, users), Vorbehalt: !decl.Gesund}
			declarations = append(declarations, apiDecl)
		}
	}
	return &api.AuctionState{Declarations: declarations}
}

func toApiBids(bids *match.Bids, users cloud.PlayerUserMap) []*api.Bid {
	var ans []*api.Bid
	for _, player := range game.Players() {
		bidsOf := bids.BidsOf(player)
		apiBidsOf := make([]*api.Bid, len(bidsOf))
		for i, b := range bidsOf {
			apiBidsOf[i] = &api.Bid{UserId: ToApiUserId(player, users), Bid: ToApiBid(b)}
		}
		ans = append(ans, apiBidsOf...)
	}
	return ans
}

func ToGameState(m *match.Match, users cloud.PlayerUserMap) *api.GameState {
	var prevTrick *api.Trick
	prevTrickGame := m.Game.PreviousTrick()
	if prevTrickGame != nil {
		prevTrick = ToApiTrickC(*prevTrickGame, users)
	}
	return &api.GameState{Mode: ToApiMode(m.Mode(), m.InitialForehand(), users),
		Bids:            toApiBids(m.Bids, users),
		CompletedTricks: int32(m.Game.NumCompletedTricks()),
		CurrentTrick:    ToApiTrick(m.Game.CurrentTrick, m.Mode(), users),
		PreviousTrick:   prevTrick,
	}
}
func ToMatchState(tm *cloud.TableMatch, user string) *api.MatchState {
	m := tm.Match
	turn := ToPlayerValue(m.WhoseTurn(), tm.Players)
	self := tm.Players.PlayerFor(user)
	if self == game.NoPlayer { // spectator
		self = m.InitialForehand()
	}
	players := &api.Players{
		UserIdSelf:  ToApiUserId(self, tm.Players),
		UserIdLeft:  ToApiUserId(self.NextPlayer(), tm.Players),
		UserIdFace:  ToApiUserId(self.NthNext(2), tm.Players),
		UserIdRight: ToApiUserId(self.NthNext(3), tm.Players),
	}
	ans := &api.MatchState{Turn: turn, Players: players}
	addDetails(ans, *tm)
	addRole(ans, *tm, user)
	return ans

}

func addRole(state *api.MatchState, tm cloud.TableMatch, user string) {
	self := tm.Players.PlayerFor(user)
	if self == game.NoPlayer {
		state.Role = &api.MatchState_Spectator{Spectator: &api.Empty{}}
	} else {
		cards := GetHandCards(tm.Match, self)
		state.Role = &api.MatchState_OwnCards{OwnCards: &api.Cards{Cards: cards}}
	}
}

func addDetails(state *api.MatchState, tm cloud.TableMatch) {
	switch tm.Match.Phase() {
	case match.InAuction:
		state.Phase = api.MatchPhase_AUCTION
		auctionState := toAuctionState(tm.Match.Auction, tm.Players)
		state.Details = &api.MatchState_AuctionState{AuctionState: auctionState}
	case match.InGame:
		state.Phase = api.MatchPhase_GAME
		gameState := ToGameState(tm.Match, tm.Players)
		state.Details = &api.MatchState_GameState{GameState: gameState}
	case match.MatchFinished:
		state.Phase = api.MatchPhase_FINISHED
	default:
		panic(fmt.Sprintf("ToMatchState called with invalid match phase %v", tm.Match.Phase()))
	}
}

func GetHandCards(m *match.Match, p game.Player) []*api.Card {
	var cards game.Hand
	if m.Phase() != match.InGame {
		cards = m.InitialHandCards(p)
	} else {
		cards = m.Game.HandCards[p]
	}
	ans := make([]*api.Card, len(cards))
	for i, card := range cards {
		ans[i] = ToApiCard(card)
	}
	return ans
}

func ToTableData(table *cloud.Table, user string) *api.TableData {
	exposedInviteCode := ""
	if table.Owner() == user {
		exposedInviteCode = table.InviteCode
	}
	return &api.TableData{TableId: table.Id, Owner: string(table.Owner()), InviteCode: exposedInviteCode}
}

func ToMatchPhase(p api.MatchPhase) match.Phase {
	switch p {
	case api.MatchPhase_AUCTION:
		return match.InAuction
	case api.MatchPhase_GAME:
		return match.InGame
	case api.MatchPhase_FINISHED:
		return match.MatchFinished
	}
	panic(fmt.Sprintf("unknown match phase: %v", p))
}
