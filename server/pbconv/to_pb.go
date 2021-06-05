package pbconv

import (
	"fmt"
	pb "github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
	r "github.com/supermihi/karlchencloud/server/room"
	t "github.com/supermihi/karlchencloud/server/table"
	u "github.com/supermihi/karlchencloud/server/users"
)

func ToApiUserId(p game.Player, users t.PlayerUserMap) string {
	if p == game.NoPlayer {
		panic("cannot convert NoPlayer to user id")
	}
	return users[p].String()
}
func ToPbPlayerValue(p game.Player, users t.PlayerUserMap) *pb.PlayerValue {
	if p != game.NoPlayer {
		return &pb.PlayerValue{UserId: users[p].String()}
	}
	return nil
}

func ToPbBid(b match.Bid) pb.BidType {
	switch b {
	case match.Re:
		return pb.BidType_RE_BID
	case match.ReKeine90:
		return pb.BidType_RE_NO_NINETY
	case match.ReKeine60:
		return pb.BidType_CONTRA_NO_SIXTY
	case match.ReKeine30:
		return pb.BidType_RE_NO_THIRTY
	case match.ReSchwarz:
		return pb.BidType_RE_SCHWARZ
	case match.Contra:
		return pb.BidType_CONTRA_BID
	case match.ContraKeine90:
		return pb.BidType_CONTRA_NO_NINETY
	case match.ContraKeine60:
		return pb.BidType_CONTRA_NO_SIXTY
	case match.ContraKeine30:
		return pb.BidType_CONTRA_NO_THIRTY
	case match.ContraSchwarz:
		return pb.BidType_CONTRA_SCHWARZ
	default:
		panic(fmt.Sprintf("unexpected bid %v in ToPbBid", b))
	}
}

func ToPbGameType(t game.AnnouncedGameType) pb.GameType {
	switch t {
	case game.NormalGameType:
		return pb.GameType_NORMAL_GAME
	case game.MarriageType:
		return pb.GameType_MARRIAGE
	case game.DiamondSoloType:
		return pb.GameType_DIAMONDS_SOLO
	case game.HeartSoloType:
		return pb.GameType_HEARTS_SOLO
	case game.SpadeSoloType:
		return pb.GameType_SPADES_SOLO
	case game.ClubSoloType:
		return pb.GameType_CLUBS_SOLO
	case game.AceSoloType:
		return pb.GameType_MEATLESS_SOLO
	}
	panic(fmt.Sprintf("not a solo type: %v", t))
}

func ToPbMode(mode game.Mode, forehand game.Player, users t.PlayerUserMap) *pb.Mode {
	soloist := ToPbPlayerValue(game.Soloist(mode), users)
	var spouse *pb.PlayerValue
	switch h := mode.(type) {
	case game.Marriage:
		spouse = ToPbPlayerValue(h.Partner(), users)
	}
	return &pb.Mode{Type: ToPbGameType(mode.Type()), Soloist: soloist, Spouse: spouse,
		Forehand: users[int(forehand)].String()}
}

func ToPbSuit(s game.Suit) pb.Suit {
	switch s {
	case game.Diamonds:
		return pb.Suit_DIAMONDS
	case game.Hearts:
		return pb.Suit_HEARTS
	case game.Spades:
		return pb.Suit_SPADES
	case game.Clubs:
		return pb.Suit_CLUBS
	}
	panic(fmt.Sprintf("unexpected suit %v", s))
}

func ToPbRank(r game.Rank) pb.Rank {
	switch r {
	case game.Nine:
		return pb.Rank_NINE
	case game.Jack:
		return pb.Rank_JACK
	case game.Queen:
		return pb.Rank_QUEEN
	case game.King:
		return pb.Rank_KING
	case game.Ten:
		return pb.Rank_TEN
	case game.Ace:
		return pb.Rank_ACE
	}
	panic(fmt.Sprintf("unexpected rank %v", r))
}

func ToPbCard(c game.Card) *pb.Card {
	return &pb.Card{Suit: ToPbSuit(c.Suit), Rank: ToPbRank(c.Rank)}
}

func ToPbTrick(t *game.IncompleteTrick, m game.Mode, users t.PlayerUserMap) *pb.Trick {
	result := &pb.Trick{UserIdForehand: ToApiUserId(t.Forehand, users)}
	cards := make([]*pb.Card, t.NumCardsPlayed())
	for i := 0; i < len(cards); i++ {
		cards[i] = ToPbCard(t.NthCard(i))
	}
	result.Cards = cards
	if t.IsComplete() {
		result.UserIdWinner = ToPbPlayerValue(game.WinnerOfTrick(t.CardsByPlayer(), t.Forehand, m), users)
	}
	return result
}

func ToPbTrickComplete(t game.Trick, users t.PlayerUserMap) *pb.Trick {
	result := &pb.Trick{UserIdForehand: ToApiUserId(t.Forehand, users)}
	cards := make([]*pb.Card, game.NumPlayers)
	for i := 0; i < len(cards); i++ {
		cards[i] = ToPbCard(t.NthCard(i))
	}
	result.UserIdWinner = ToPbPlayerValue(t.Winner, users)
	return result
}

func toPbBids(bids *match.Bids, users t.PlayerUserMap) []*pb.Bid {
	var ans []*pb.Bid
	for _, player := range game.Players() {
		bidsOf := bids.BidsOf(player)
		apiBidsOf := make([]*pb.Bid, len(bidsOf))
		for i, b := range bidsOf {
			apiBidsOf[i] = &pb.Bid{UserId: ToApiUserId(player, users), Bid: ToPbBid(b)}
		}
		ans = append(ans, apiBidsOf...)
	}
	return ans
}

func ToPbGameState(m *r.MatchData) *pb.GameState {
	var prevTrick *pb.Trick
	prevTrickGame := m.PreviousTrick
	if prevTrickGame != nil {
		prevTrick = ToPbTrickComplete(*prevTrickGame, m.Players)
	}
	return &pb.GameState{Mode: ToPbMode(m.Mode, m.InitialForehand, m.Players),
		Bids:            toPbBids(&m.Bids, m.Players),
		CompletedTricks: int32(m.CompletedTricks),
		CurrentTrick:    ToPbTrick(m.CurrentTrick, m.Mode, m.Players),
		PreviousTrick:   prevTrick,
	}
}
func ToPbMatchState(matchData *r.MatchData, user u.Id) *pb.MatchState {
	turn := &pb.PlayerValue{}
	if matchData.Turn != game.NoPlayer {
		turn.UserId = matchData.Players[matchData.Turn].String()
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
	players := &pb.Players{
		UserIdSelf:  matchData.Players[effectiveSelf].String(),
		UserIdLeft:  matchData.Players[effectiveSelf.NextPlayer()].String(),
		UserIdFace:  matchData.Players[effectiveSelf.NthNext(2)].String(),
		UserIdRight: matchData.Players[effectiveSelf.NthNext(3)].String(),
	}
	ans := &pb.MatchState{Turn: turn, Players: players}
	addDetails(ans, matchData)
	addRole(ans, matchData, self)
	return ans
}

func ToPbCards(cards game.Hand) []*pb.Card {
	ans := make([]*pb.Card, len(cards))
	for i, card := range cards {
		ans[i] = ToPbCard(card)
	}
	return ans
}

type MemberResolver func(u.Id) *pb.TableMember

func ToPbTableData(table *r.TableData, user u.Id, getPbTableMember MemberResolver) *pb.TableData {
	pbMembers := make([]*pb.TableMember, len(table.Players))
	for i, player := range table.Players {
		pbMembers[i] = getPbTableMember(player)
	}
	ans := &pb.TableData{TableId: table.Id.String(), Owner: table.Owner.String(), Members: pbMembers,
		Created: pb.NewTimestamp(table.Created), Public: table.Public}
	if table.Owner == user {
		ans.InviteCode = table.InviteCode
	}
	return ans
}

func ToPbTables(tables []*r.TableData, user u.Id, getPbTableMember MemberResolver) []*pb.TableData {
	result := make([]*pb.TableData, len(tables))
	for i, table := range tables {
		result[i] = ToPbTableData(table, user, getPbTableMember)
	}
	return result
}

func toPbAuctionState(data *r.MatchData) *pb.AuctionState {
	declarations := make([]*pb.Declaration, len(data.Declarations))
	i := 0
	for player, decl := range data.Declarations {
		declarations[i] = &pb.Declaration{UserId: data.Players[player].String(), Vorbehalt: !decl.Healthy}
		i++
	}
	return &pb.AuctionState{Declarations: declarations}
}

func addDetails(state *pb.MatchState, md *r.MatchData) {
	switch md.Phase {
	case match.InAuction:
		state.Phase = pb.MatchPhase_AUCTION
		state.Details = &pb.MatchState_AuctionState{AuctionState: toPbAuctionState(md)}
	case match.InGame:
		state.Phase = pb.MatchPhase_GAME
		gameState := ToPbGameState(md)
		state.Details = &pb.MatchState_GameState{GameState: gameState}
	case match.MatchFinished:
		state.Phase = pb.MatchPhase_FINISHED
	default:
		panic(fmt.Sprintf("ToPbMatchState called with invalid match phase %v", md.Phase))
	}
}
