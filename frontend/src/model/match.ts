import { BidType, GameType, MatchPhase } from '../api/karlchen_pb';
import { Auction } from './auction';
import { Card } from './core';
import { Pos, PlayerIds, newPlayerMap } from './players';

export interface Match {
  phase: MatchPhase;
  turn?: Pos;
  players: PlayerIds;
  cards: Card[];
  game: Game;
  auction: Auction;
}

export interface Game {
  bids: Record<Pos, BidType[]>;
  completedTricks: number;
  currentTrick: Trick;
  previousTrick?: Trick;
  mode: Mode;
}

export function newGame(mode: Mode): Game {
  return {
    mode,
    bids: newPlayerMap(() => [] as BidType[]),
    completedTricks: 0,
    currentTrick: newTrick(mode.forehand),
  };
}
export function dummyGame(): Game {
  return newGame({ type: GameType.NORMAL_GAME, forehand: Pos.bottom });
}

export interface PlayedCard {
  card: Card;
  player: Pos;
  trickWinner?: Pos;
}
export interface Trick {
  forehand: Pos;
  winner?: Pos;
  cards: Card[];
}

export function newTrick(forehand: Pos): Trick {
  return { forehand, cards: [] };
}

export interface Mode {
  type: GameType;
  soloist?: Pos;
  spouse?: Pos;
  forehand: Pos;
}
