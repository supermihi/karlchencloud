package match

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	. "github.com/supermihi/karlchencloud/pkg/game/core"
	"testing"
)

func TestSampleMatch(t *testing.T) {
	var cards Cards = [NumPlayers]Hand{
		[]Card{
			HerzK, HerzA, PikK, Pik10, PikA, KreuzK, KaroK, KaroA, HerzB, KreuzB, KreuzD, Herz10,
		},
		[]Card{
			PikK, Pik10, Kreuz10, Kreuz10, Karo10, KaroB, KaroB, HerzD, HerzD, PikD, PikD, KreuzD,
		},
		[]Card{
			Herz9, Herz9, HerzK, Kreuz9, Kreuz9, Karo9, KaroK, KaroA, PikB, KaroD, KaroD, Herz10,
		},
		[]Card{
			HerzA, Pik9, Pik9, PikA, KreuzK, KreuzA, KreuzA, Karo9, Karo10, HerzB, PikB, KreuzB,
		},
	}
	sonderspiele := MakeSonderspiele(VorbehaltHochzeit{})
	match := NewMatch(Player3, sonderspiele, cards)
	play := func(player Player, card Card) {
		for _, otherPlayer := range Players() {
			if otherPlayer != player {
				assert.NotEqual(t, otherPlayer, match.game.WhoseTurn())
			}
		}
		ans := match.PlayCard(player, card)
		require.Truef(t, ans, "error playing %v as %v", card, player)
	}
	sayGesund := func(player Player) {
		ans := match.AnnounceGesundOrVorbehalt(player, false)
		assert.Truef(t, ans, "error announcing gesund as %v", player)
	}
	expectTrickWinner := func(player Player) {
		tricks := match.game.CompleteTricks
		winner := tricks[len(tricks)-1].Winner
		require.Equalf(t, player, winner, "expecting %v to win 1st trick instead of %v", player, winner)
		assert.True(t, match.game.IsFinished() || match.game.WhoseTurn() == player)
	}
	sayGesund(Player3)
	sayGesund(Player4)
	sayGesund(Player1)
	sayGesund(Player2)
	assert.Equal(t, GamePhase, match.Phase())
	mode := match.game.Mode
	assert.Equal(t, ReParty, mode.PartyOf(Player1))
	assert.Equal(t, ReParty, mode.PartyOf(Player2))
	assert.Equal(t, ContraParty, mode.PartyOf(Player3))
	assert.Equal(t, ContraParty, mode.PartyOf(Player4))
	// trick 0
	play(Player3, PikB)
	play(Player4, Karo9)
	play(Player1, KreuzD)
	play(Player2, Karo10)
	expectTrickWinner(Player1)
	match.PlaceBid(Player1, Re)
	assert.False(t, match.PlaceBid(Player4, Re))
	// trick 1
	play(Player1, PikA)
	play(Player2, Pik10)
	play(Player3, KaroA)
	play(Player4, PikA) // doppelkopf
	expectTrickWinner(Player3)
	// trick 2
	play(Player3, Kreuz9)
	play(Player4, KreuzA)
	play(Player1, KreuzK)
	play(Player2, Kreuz10)
	expectTrickWinner(Player4)
	// trick 3
	play(Player4, HerzA)
	play(Player1, HerzK)
	play(Player2, KaroB)
	play(Player3, Herz9)
	expectTrickWinner(Player2)
	// trick 4
	play(Player2, PikD)
	play(Player3, Herz10)
	play(Player4, Karo10)
	play(Player1, KaroK)
	expectTrickWinner(Player3)
	// trick 5
	play(Player3, Karo9)
	play(Player4, PikB)
	play(Player1, KaroA)
	play(Player2, HerzD)
	expectTrickWinner(Player2)
	// trick 6
	play(Player2, PikD)
	play(Player3, KaroK)
	play(Player4, HerzB)
	play(Player1, HerzB)
	expectTrickWinner(Player2)
	// trick 7
	play(Player2, PikK)
	play(Player3, KaroD)
	play(Player4, Pik9)
	play(Player1, PikK)
	expectTrickWinner(Player3)
	// trick 8
	play(Player3, Herz9)
	play(Player4, Pik9)
	play(Player1, HerzA)
	play(Player2, Kreuz10)
	expectTrickWinner(Player1)
	// trick 9
	play(Player1, Pik10)
	play(Player2, KreuzD)
	play(Player3, Kreuz9)
	play(Player4, KreuzK)
	expectTrickWinner(Player2)
	// trick 10
	play(Player2, HerzD)
	play(Player3, KaroD)
	play(Player4, KreuzB)
	play(Player1, Herz10)
	expectTrickWinner(Player1)
	// trick 11
	play(Player1, KreuzB)
	play(Player2, KaroB)
	play(Player3, HerzK)
	play(Player4, KreuzA) // karlchen
	expectTrickWinner(Player1)
	assert.Equal(t, MatchFinished, match.Phase())
	result := match.Evaluate()
	assert.Equal(t, ReParty, result.Winner)
	assert.Equal(t, 134, result.TrickScoreRe)
	assert.ElementsMatch(t, [2]GamePoint{{Gewonnen, 1}, {ReAngesagt, 2}}, result.GamePoints)
	assert.Equal(t, 5, result.TotalValue)
	assert.ElementsMatch(t, [2]ExtraPoint{{Doppelkopf, Player3, 1}, {Karlchen, Player1, 11}}, result.ExtraPoints)
	points := PointsByPlayer(&result, match.Mode())
	assert.Equal(t, 5, points[Player1])
	assert.Equal(t, 5, points[Player2])
	assert.Equal(t, -5, points[Player3])
	assert.Equal(t, -5, points[Player4])

}
