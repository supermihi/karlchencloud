import { BidType, GameType, MatchPhase } from '../api/karlchen_pb';
import { Card } from './core';
import { Pos, Players } from './players';

export interface Match {
  phase: MatchPhase;
  turn?: Pos;
  players: Players;
  cards: Card[];
  details: Auction | Game;
}

export function isAuction(aog: Auction | Game): aog is Auction {
  return 'declarations' in aog;
}

export interface Game {
  bids: Record<Pos, BidType[]>;
  completedTricks: number;
  currentTrick: Trick;
  previousTrick?: Trick;
  mode: Mode;
}

export interface Trick {
  forehand: Pos;
  winner?: Pos;
  cards: Card[];
}

export interface Mode {
  type: GameType;
  soloist?: string;
  spouse?: string;
  forehand: string;
}

export enum Declaration {
  gesund,
  vorbehalt,
}

export interface Auction {
  declarations: { [player: string]: Declaration };
}
