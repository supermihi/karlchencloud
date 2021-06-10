import { afterPlayedCard, removeCard } from './mutations';
import { Card } from './core';
import * as c from './cards';
import { MatchPhase, Party } from '../api/karlchen_pb';
import { Pos } from './players';
import { fullHand } from 'mocks/cards';
import { createMatch } from '../mocks/match';
import { MatchInGame, Trick } from './match';
import { SpadeQ } from './cards';

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

test('afterPlayedCard when placing first card of game', () => {
  const before = createMatch({
    phase: MatchPhase.GAME,
    forehand: Pos.bottom,
    progress: 0,
    numPlayedTricks: 0,
  }) as MatchInGame;
  const playedCard = {
    card: fullHand[0],
    player: Pos.bottom,
    trickWinner: null,
    matchWinner: null,
  };
  const after = afterPlayedCard(before, playedCard);
  expect(after.cards).toEqual(fullHand.slice(1));
  expect(after.game.currentTrick.cards).toEqual([playedCard.card]);
  expect(after.turn).toEqual(Pos.left);
});

test('afterPlayedCard when placing first card of second trick', () => {
  const before = createMatch({
    phase: MatchPhase.GAME,
    forehand: Pos.top,
    progress: 4,
    numPlayedTricks: 0,
    currentTrickWinner: Pos.left,
  }) as MatchInGame;
  const handBefore = [...before.cards];
  const playedCard = { card: SpadeQ, player: Pos.left, trickWinner: null, matchWinner: null };
  const after = afterPlayedCard(before, playedCard);
  expect(after.cards).toEqual(handBefore);
  const expectedTrick: Trick = {
    forehand: Pos.left,
    cards: [SpadeQ],
    winner: null,
  };
  expect(after.game.currentTrick).toEqual(expectedTrick);
  expect(after.turn).toEqual(Pos.top);
});

test('afterPlayedCard when placing fourth card', () => {
  const before = createMatch({
    phase: MatchPhase.GAME,
    forehand: Pos.left,
    progress: 3,
    numPlayedTricks: 0,
  }) as MatchInGame;
  const cardsBefore = [...before.game.currentTrick.cards];
  const after = afterPlayedCard(before, {
    card: fullHand[0],
    player: Pos.bottom,
    trickWinner: Pos.top,
    matchWinner: null,
  });
  expect(after.cards).toEqual(fullHand.slice(1));
  expect(after.game.currentTrick.cards).toEqual([...cardsBefore, fullHand[0]]);
  expect(after.game.currentTrick.winner).toEqual(Pos.top);
  expect(after.turn).toEqual(Pos.top);
});

test('afterPlayedCard on last card of match', () => {
  const before = createMatch({
    phase: MatchPhase.GAME,
    forehand: Pos.left,
    progress: 3,
    numPlayedTricks: 12 - 1,
  }) as MatchInGame;
  const cardsBefore = [...before.game.currentTrick.cards];
  const handBefore = [...before.cards];
  const after = afterPlayedCard(before, {
    card: handBefore[0],
    player: Pos.bottom,
    trickWinner: Pos.right,
    matchWinner: Party.RE,
  });
  expect(after.cards).toEqual([]);
  expect(after.game.currentTrick.cards).toEqual([...cardsBefore, handBefore[0]]);
  expect(after.game.currentTrick.winner).toEqual(Pos.right);
  expect(after.turn).toBeNull();
  expect(after.winner).toEqual(Party.RE);
});
