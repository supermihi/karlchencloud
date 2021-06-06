import { afterPlayedCard, removeCard } from './mutations';
import { Card } from './core';
import * as c from './cards';
import { emptyBids, MatchInGame, newTrick } from './match';
import * as mp from 'mocks/players';
import { GameType, MatchPhase } from '../api/karlchen_pb';
import { Pos } from './players';
import { fullHand } from 'mocks/cards';

test('removeCard removes a card', () => {
  const cards: Card[] = [c.Diamond9, c.Diamond10, c.DiamondJ];
  const afterRemove = removeCard(cards, c.Diamond10);
  expect(afterRemove).toEqual([c.Diamond9, c.DiamondJ]);
  // assert immutability
  expect(cards).toEqual([c.Diamond9, c.Diamond10, c.DiamondJ]);
});

test('removeCard does nothing when card does not exist', () => {
  const cards: Card[] = [c.Diamond9, c.Diamond10, c.DiamondJ];
  const afterRemove = removeCard(cards, c.DiamondA);
  expect(afterRemove).toBe(cards);
});

const getMatch = (): MatchInGame => ({
  players: mp.players,
  phase: MatchPhase.GAME,
  auction: null,
  cards: fullHand,
  turn: Pos.bottom,
  game: {
    bids: emptyBids(),
    completedTricks: 0,
    mode: {
      type: GameType.NORMAL_GAME,
      forehand: Pos.bottom,
    },
    currentTrick: newTrick(Pos.bottom),
  },
});
test('afterPlayedCard when self draws first card', () => {
  const before = getMatch();
  const after = afterPlayedCard(before, { card: fullHand[0], player: Pos.bottom });
  expect(after.cards).toEqual(fullHand.slice(1));
  expect(after.game.currentTrick.cards).toEqual([fullHand[0]]);
  expect(after.turn).toEqual(Pos.left);
});
