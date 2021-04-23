package server

import (
	"fmt"
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
)

func ToApiUserId(p game.Player, users PlayerUserMap) string {
	if p == game.NoPlayer {
		panic("cannot convert NoPlayer to user id")
	}
	return users[p]
}
func ToPlayerValue(p game.Player, users PlayerUserMap) *api.PlayerValue {
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
	case game.NormalGameType:
		return api.GameType_NORMAL_GAME
	case game.MarriageType:
		return api.GameType_MARRIAGE
	case game.DiamondSoloType:
		return api.GameType_DIAMONDS_SOLO
	case game.HeartSoloType:
		return api.GameType_HEARTS_SOLO
	case game.SpadeSoloType:
		return api.GameType_SPADES_SOLO
	case game.ClubSoloType:
		return api.GameType_CLUBS_SOLO
	case game.AceSoloType:
		return api.GameType_MEATLESS_SOLO
	}
	panic(fmt.Sprintf("not a solo type: %v", t))
}

func ToGameType(t api.GameType) game.AnnouncedGameType {
	switch t {
	case api.GameType_NORMAL_GAME:
		return game.NormalGameType
	case api.GameType_MARRIAGE:
		return game.MarriageType
	case api.GameType_DIAMONDS_SOLO:
		return game.DiamondSoloType
	case api.GameType_HEARTS_SOLO:
		return game.HeartSoloType
	case api.GameType_SPADES_SOLO:
		return game.SpadeSoloType
	case api.GameType_CLUBS_SOLO:
		return game.ClubSoloType
	case api.GameType_MEATLESS_SOLO:
		return game.AceSoloType
	}
	panic(fmt.Sprintf("not an api game type: %s", t))
}

func ToApiMode(mode game.Mode, forehand game.Player, users PlayerUserMap) *api.Mode {
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

func ToApiTrick(t *game.IncompleteTrick, m game.Mode, users PlayerUserMap) *api.Trick {
	result := &api.Trick{UserIdForehand: ToApiUserId(t.Forehand, users)}
	cards := make([]*api.Card, t.NumCardsPlayed())
	for i := 0; i < len(cards); i++ {
		cards[i] = ToApiCard(t.NthCard(i))
	}
	result.Cards = cards
	if t.IsComplete() {
		result.UserIdWinner = ToPlayerValue(game.WinnerOfTrick(t.CardsByPlayer(), t.Forehand, m), users)
	}
	return result
}

func ToApiTrickC(t game.Trick, users PlayerUserMap) *api.Trick {
	result := &api.Trick{UserIdForehand: ToApiUserId(t.Forehand, users)}
	cards := make([]*api.Card, game.NumPlayers)
	for i := 0; i < len(cards); i++ {
		cards[i] = ToApiCard(t.NthCard(i))
	}
	result.UserIdWinner = ToPlayerValue(t.Winner, users)
	return result
}

func toAuctionState(data *MatchData) *api.AuctionState {
	declarations := make([]*api.Declaration, len(data.Declarations))
	i := 0
	for player, decl := range data.Declarations {
		declarations[i] = &api.Declaration{UserId: data.Players[player], Vorbehalt: !decl.Gesund}
		i++
	}
	return &api.AuctionState{Declarations: declarations}
}

func toApiBids(bids *match.Bids, users PlayerUserMap) []*api.Bid {
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

func ToGameState(m *MatchData) *api.GameState {
	var prevTrick *api.Trick
	prevTrickGame := m.PreviousTrick
	if prevTrickGame != nil {
		prevTrick = ToApiTrickC(*prevTrickGame, m.Players)
	}
	return &api.GameState{Mode: ToApiMode(m.Mode, m.InitialForehand, m.Players),
		Bids:            toApiBids(&m.Bids, m.Players),
		CompletedTricks: int32(m.CompletedTricks),
		CurrentTrick:    ToApiTrick(m.CurrentTrick, m.Mode, m.Players),
		PreviousTrick:   prevTrick,
	}
}
func ToMatchState(matchData *MatchData, user string) *api.MatchState {
	turn := &api.PlayerValue{}
	if matchData.Turn != game.NoPlayer {
		turn.UserId = matchData.Players[matchData.Turn]
	}
	self := game.NoPlayer
	for i, p := range game.Players() {
		if matchData.Players[i] == user {
			self = p
			break
		}
	}
	effectiveSelf := self
	if self == game.NoPlayer { // spectator -> initial forehand at place "self"
		effectiveSelf = matchData.InitialForehand
	}
	players := &api.Players{
		UserIdSelf:  matchData.Players[effectiveSelf],
		UserIdLeft:  matchData.Players[effectiveSelf.NextPlayer()],
		UserIdFace:  matchData.Players[effectiveSelf.NthNext(2)],
		UserIdRight: matchData.Players[effectiveSelf.NthNext(3)],
	}
	ans := &api.MatchState{Turn: turn, Players: players}
	addDetails(ans, matchData)
	addRole(ans, matchData, self)
	return ans
}

func addRole(state *api.MatchState, m *MatchData, self game.Player) {
	if self == game.NoPlayer {
		state.Role = &api.MatchState_Spectator{Spectator: &api.Empty{}}
	} else {
		cards := ToApiCards(m.Cards[self])
		state.Role = &api.MatchState_OwnCards{OwnCards: &api.Cards{Cards: cards}}
	}
}

func addDetails(state *api.MatchState, md *MatchData) {
	switch md.Phase {
	case match.InAuction:
		state.Phase = api.MatchPhase_AUCTION
		state.Details = &api.MatchState_AuctionState{AuctionState: toAuctionState(md)}
	case match.InGame:
		state.Phase = api.MatchPhase_GAME
		gameState := ToGameState(md)
		state.Details = &api.MatchState_GameState{GameState: gameState}
	case match.MatchFinished:
		state.Phase = api.MatchPhase_FINISHED
	default:
		panic(fmt.Sprintf("ToMatchState called with invalid match phase %v", md.Phase))
	}
}

func ToApiCards(cards game.Hand) []*api.Card {
	ans := make([]*api.Card, len(cards))
	for i, card := range cards {
		ans[i] = ToApiCard(card)
	}
	return ans
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

func ToTableData(table *TableData, user string, members []*api.TableMember) *api.TableData {
	ans := &api.TableData{TableId: table.Id, Owner: table.Owner, Members: members,
		Created: api.NewTimestamp(table.Created)}
	if table.Owner == user {
		ans.InviteCode = table.InviteCode
	}
	return ans
}
