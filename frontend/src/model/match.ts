import { BidType, GameType, MatchPhase } from '../api/karlchen_pb';
import { Card } from './core';
import { Players } from './table';

export interface Match {
  phase: MatchPhase;
  turn?: string;
  players: Players;
  cards: Card[];
  details: Auction | Game;
}

export function isAuction(aog: Auction | Game): aog is Auction {
  return 'declarations' in aog;
}

export interface Game {
  bids: { [player: string]: BidType[] };
  completedTricks: number;
  currentTrick: Trick;
  previousTrick?: Trick;
  mode: Mode;
}

export interface Trick {
  forehand: string;
  winner?: string;
  cards: { [player: string]: Card };
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
