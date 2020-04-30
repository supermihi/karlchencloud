package match

import (
	"github.com/supermihi/karlchencloud/pkg/game/auction"
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
	sonderspiele := auction.MakeSonderspiele(auction.VorbehaltHochzeit{})
	match := NewMatch(Player3, sonderspiele, cards)
	play := func(player Player, card Card) {
		ans := match.PerformAction(PlayCardAction(player, card))
		if ans.Type != Ok {
			t.Errorf("error playing %v as %v: %v", card, player, ans.ErrorMsg)
			t.FailNow()
		}
	}
	sayGesund := func(player Player) {
		ans := match.PerformAction(AnnounceGesundAction(player))
		if ans.Type != Ok {
			t.Errorf("error announcing gesund as %v: %v", player, ans.ErrorMsg)
		}
	}
	expectTrickWinner := func(player Player) {
		tricks := match.game.CompleteTricks
		winner := tricks[len(tricks)-1].Winner
		if winner != player {
			t.Errorf("expecting %v to win 1st trick instead of %v", player, winner)
			t.FailNow()
		}
		if !match.game.IsFinished() && match.game.WhoseTurn() != player {
			t.Errorf("expecting winner of trick to be next on turn")
		}
	}
	sayGesund(Player3)
	sayGesund(Player4)
	sayGesund(Player1)
	sayGesund(Player2)
	if match.Phase() != GamePhase {
		t.Error("expecting game phase after four 'gesund'")
	}
	mode := match.game.Mode
	if mode.PartyOf(Player1) != ReParty || mode.PartyOf(Player2) != ReParty || mode.PartyOf(Player3) != ContraParty || mode.PartyOf(Player4) != ContraParty {
		t.Error("unexpected parties")
	}
	if match.PerformAction(PlayCardAction(Player2, KreuzD)).Type != WrongPlayerTurn {
		t.Errorf("not Player2's turn")
	}
	play(Player3, PikB)
	play(Player4, Karo9)
	play(Player1, KreuzD)
	play(Player2, Karo10)
	expectTrickWinner(Player1)
	play(Player1, PikA)
	play(Player2, Pik10)
	play(Player3, KaroA)
	play(Player4, PikA) // doppelkopf
	expectTrickWinner(Player3)
	play(Player3, Kreuz9)
	play(Player4, KreuzA)
	play(Player1, KreuzK)
	play(Player2, Kreuz10)
	expectTrickWinner(Player4)
	play(Player4, HerzA)
	play(Player1, HerzK)
	play(Player2, KaroB)
	play(Player3, Herz9)
	expectTrickWinner(Player2)
	play(Player2, PikD)
	play(Player3, Herz10)
	play(Player4, Karo10)
	play(Player1, KaroK)
	expectTrickWinner(Player3)
	play(Player3, Karo9)
	play(Player4, PikB)
	play(Player1, KaroA)
	play(Player2, HerzD)
	expectTrickWinner(Player2)
	play(Player2, PikD)
	play(Player3, KaroK)
	play(Player4, HerzB)
	play(Player1, HerzB)
	expectTrickWinner(Player2)
	play(Player2, PikK)
	play(Player3, KaroD)
	play(Player4, Pik9)
	play(Player1, PikK)
	expectTrickWinner(Player3)
	play(Player3, Herz9)
	play(Player4, Pik9)
	play(Player1, HerzA)
	play(Player2, Kreuz10)
	expectTrickWinner(Player1)
	play(Player1, Pik10)
	play(Player2, KreuzD)
	play(Player3, Kreuz9)
	play(Player4, KreuzK)
	expectTrickWinner(Player2)
	play(Player2, HerzD)
	play(Player3, KaroD)
	play(Player4, KreuzB)
	play(Player1, Herz10)
	expectTrickWinner(Player1)
	play(Player1, KreuzB)
	play(Player2, KaroB)
	play(Player3, HerzK)
	play(Player4, KreuzA) // karlchen
	expectTrickWinner(Player1)
	if match.Phase() != MatchFinished {
		t.Errorf("expecting match finish after last card played")
	}
	result := EvaluateGame(match.game, match.bids)
	if result.Winner != ReParty {
		t.Error("expecting Re has won")
	}
	if result.TrickScoreRe != 134 {
		t.Error("Expecting score 134 for Re")
	}
	if len(result.GamePoints) != 1 || result.GamePoints[0].Type != Gewonnen {
		t.Errorf("unexpected game points")
	}
	if result.TotalValue != 3 {
		t.Errorf("Expecting game value of 3, not %v", result.TotalValue)
	}
	if len(result.ExtraPoints) != 2 {
		t.Errorf("Expecting 2 extra points instead of %v", len(result.ExtraPoints))
	}

}
