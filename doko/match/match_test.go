package match

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	. "github.com/supermihi/karlchencloud/doko/game"
	"testing"
)

func sampleCards() Cards {
	return [NumPlayers]Hand{
		[]Card{
			HeartsK, HeartsA, SpadesK, Spades10, SpadesA, ClubsK, DiamondsK, DiamondsA, HeartsJ, ClubsJ, ClubsQ, Hearts10,
		},
		[]Card{
			SpadesK, Spades10, Clubs10, Clubs10, Diamonds10, DiamondsJ, DiamondsJ, HeartsQ, HeartsQ, SpadesQ, SpadesQ, ClubsQ,
		},
		[]Card{
			Hearts9, Hearts9, HeartsK, Clubs9, Clubs9, Diamonds9, DiamondsK, DiamondsA, SpadesJ, DiamondsQ, DiamondsQ, Hearts10,
		},
		[]Card{
			HeartsA, Spades9, Spades9, SpadesA, ClubsK, ClubsA, ClubsA, Diamonds9, Diamonds10, HeartsJ, SpadesJ, ClubsJ,
		},
	}
}
func TestSampleMatch(t *testing.T) {
	match := NewMatch(Player3, sampleCards())
	play := func(player Player, card Card) {
		for _, otherPlayer := range Players() {
			if otherPlayer != player {
				assert.NotEqual(t, otherPlayer, match.Game.WhoseTurn())
			}
		}
		ans := match.PlayCard(player, card)
		require.Truef(t, ans, "error playing %v as %v", card, player)
	}
	sayHealthy := func(player Player) {
		ans := match.AnnounceGameType(player, NormalGameType)
		assert.Truef(t, ans, "error announcing healthy as %v", player)
	}
	expectTrickWinner := func(player Player) {
		tricks := match.Game.CompleteTricks
		winner := tricks[len(tricks)-1].Winner
		require.Equalf(t, player, winner, "expecting %v to win 1st trick instead of %v", player, winner)
		assert.True(t, match.Game.IsFinished() || match.Game.WhoseTurn() == player)
	}
	sayHealthy(Player3)
	sayHealthy(Player4)
	sayHealthy(Player1)
	sayHealthy(Player2)
	assert.Equal(t, InGame, match.Phase())
	mode := match.Game.Mode
	assert.Equal(t, ReParty, mode.PartyOf(Player1))
	assert.Equal(t, ReParty, mode.PartyOf(Player2))
	assert.Equal(t, ContraParty, mode.PartyOf(Player3))
	assert.Equal(t, ContraParty, mode.PartyOf(Player4))
	// trick 0
	play(Player3, SpadesJ)
	play(Player4, Diamonds9)
	play(Player1, ClubsQ)
	play(Player2, Diamonds10)
	expectTrickWinner(Player1)
	match.PlaceBid(Player1, Re)
	assert.False(t, match.PlaceBid(Player4, Re))
	// trick 1
	play(Player1, SpadesA)
	play(Player2, Spades10)
	play(Player3, DiamondsA)
	play(Player4, SpadesA) // doppelkopf
	expectTrickWinner(Player3)
	// trick 2
	play(Player3, Clubs9)
	play(Player4, ClubsA)
	play(Player1, ClubsK)
	play(Player2, Clubs10)
	expectTrickWinner(Player4)
	// trick 3
	play(Player4, HeartsA)
	play(Player1, HeartsK)
	play(Player2, DiamondsJ)
	play(Player3, Hearts9)
	expectTrickWinner(Player2)
	// trick 4
	play(Player2, SpadesQ)
	play(Player3, Hearts10)
	play(Player4, Diamonds10)
	play(Player1, DiamondsK)
	expectTrickWinner(Player3)
	// trick 5
	play(Player3, Diamonds9)
	play(Player4, SpadesJ)
	play(Player1, DiamondsA)
	play(Player2, HeartsQ)
	expectTrickWinner(Player2)
	// trick 6
	play(Player2, SpadesQ)
	play(Player3, DiamondsK)
	play(Player4, HeartsJ)
	play(Player1, HeartsJ)
	expectTrickWinner(Player2)
	// trick 7
	play(Player2, SpadesK)
	play(Player3, DiamondsQ)
	play(Player4, Spades9)
	play(Player1, SpadesK)
	expectTrickWinner(Player3)
	// trick 8
	play(Player3, Hearts9)
	play(Player4, Spades9)
	play(Player1, HeartsA)
	play(Player2, Clubs10)
	expectTrickWinner(Player1)
	// trick 9
	play(Player1, Spades10)
	play(Player2, ClubsQ)
	play(Player3, Clubs9)
	play(Player4, ClubsK)
	expectTrickWinner(Player2)
	// trick 10
	play(Player2, HeartsQ)
	play(Player3, DiamondsQ)
	play(Player4, ClubsJ)
	play(Player1, Hearts10)
	expectTrickWinner(Player1)
	// trick 11
	play(Player1, ClubsJ)
	play(Player2, DiamondsJ)
	play(Player3, HeartsK)
	play(Player4, ClubsA) // karlchen
	expectTrickWinner(Player1)
	assert.Equal(t, MatchFinished, match.Phase())
	result := match.Evaluate()
	assert.Equal(t, ReParty, result.Winner)
	assert.Equal(t, 134, result.TrickScoreRe)
	assert.ElementsMatch(t, [2]GamePoint{{Gewonnen, 1}, {ReAngesagt, 2}}, result.GamePoints)
	assert.Equal(t, 5, result.TotalValue)
	assert.ElementsMatch(t, [2]ExtraPoint{{Doppelkopf, Player3, 1}, {Charlie, Player1, 11}}, result.ExtraPoints)
	points := PointsByPlayer(&result, match.Mode())
	assert.Equal(t, 5, points[Player1])
	assert.Equal(t, 5, points[Player2])
	assert.Equal(t, -5, points[Player3])
	assert.Equal(t, -5, points[Player4])

}

func TestMatch_PlayCard_InAuction(t *testing.T) {
	cards := sampleCards()
	match := NewMatch(Player1, cards)
	ok := match.PlayCard(Player1, cards[Player1][0])
	assert.False(t, ok)
}
