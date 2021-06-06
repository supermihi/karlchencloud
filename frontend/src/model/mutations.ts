import { Card } from './core';
import { Match, MatchInAuction, MatchInGame, newGame, PlayedCard } from './match';
import { nextPos, Pos } from './players';
import update from 'immutability-helper';
import { DeclareResult } from './auction';
import { MatchPhase } from 'api/karlchen_pb';

export function afterPlayedCard(
  match: MatchInGame,
  { card, player, trickWinner }: PlayedCard
): MatchInGame {
  const cards = player === Pos.bottom ? removeCard(match.cards, card) : match.cards;
  return update(match, {
    game: {
      currentTrick: {
        winner: { $set: trickWinner },
        cards: trickWinner === undefined ? { $push: [card] } : { $set: [] },
      },
    },
    cards: { $set: cards },
    turn: { $set: nextPos(player) },
  });
}

export function removeCard(cards: Card[], card: Card): Card[] {
  const cardIndex = cards.findIndex((c) => c.rank === card.rank && c.suit === card.suit);
  if (cardIndex === -1) {
    return cards;
  }
  return update(cards, { $splice: [[cardIndex, 1]] });
}

export function afterDeclaration(
  match: MatchInAuction,
  { player, declaration, mode }: DeclareResult
): Match {
  const result = update(match, {
    auction: {
      declarations: { [player]: { $set: declaration } },
    },
  });
  if (mode !== null) {
    return update(result, {
      phase: { $set: MatchPhase.GAME },
      game: { $set: newGame(mode) },
      turn: { $set: mode.forehand },
    });
  } else {
    return update(result, { turn: { $set: nextPos(player) } });
  }
}
