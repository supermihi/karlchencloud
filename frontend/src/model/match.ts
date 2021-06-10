import { BidType, GameType, MatchPhase, Party } from 'api/karlchen_pb';
import { Auction } from './auction';
import { Card } from './core';
import { Pos, PlayerIds, newPlayerMap } from './players';

export interface Match {
  phase: MatchPhase;
  turn: Pos | null;
  winner: Party | null;
  players: PlayerIds;
  cards: Card[];
  game: Game | null;
  auction: Auction | null;
}
export interface MatchInGame extends Match {
  game: Game;
}
export interface MatchInAuction extends Match {
  auction: Auction;
}

export function inGame(match: Match): match is MatchInGame {
  return match.game !== null;
}
export function inAuction(match: Match): match is MatchInAuction {
  return match.auction !== null;
}

export interface Game {
  bids: Bids;
  completedTricks: number;
  currentTrick: Trick;
  previousTrick?: Trick;
  mode: Mode;
}
export type Bids = Record<Pos, BidType[]>;
export function emptyBids(): Bids {
  return newPlayerMap(() => [] as BidType[]);
}

export function newGame(mode: Mode): Game {
  return {
    mode,
    bids: newPlayerMap(() => [] as BidType[]),
    completedTricks: 0,
    currentTrick: newTrick(mode.forehand),
  };
}

export interface PlayedCard {
  card: Card;
  player: Pos;
  trickWinner: Pos | null;
  matchWinner: Party | null;
}
export interface Trick {
  forehand: Pos;
  winner: Pos | null;
  cards: Card[];
}

export function newTrick(forehand: Pos): Trick {
  return { forehand, cards: [], winner: null };
}

export interface Mode {
  type: GameType;
  soloist?: Pos;
  spouse?: Pos;
  forehand: Pos;
}
