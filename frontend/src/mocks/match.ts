import { emptyBids, Game, Match } from '../model/match';
import * as mp from './players';
import * as mc from './cards';
import { fullHand } from './cards';
import { GameType, MatchPhase, Party } from '../api/karlchen_pb';
import { nthNext, PartialPlayerMap, Pos } from 'model/players';
import { Auction, Declaration } from 'model/auction';

export interface MatchConfig {
  phase: MatchPhase;
  forehand: Pos;
  progress: number;
  numPlayedTricks?: number;
}

export interface AuctionConfig {
  forehand: Pos;
  progress: number;
}

function createDeclarations(forehand: Pos, progress: number): PartialPlayerMap<Declaration> {
  const result: PartialPlayerMap<Declaration> = {};
  for (let i = 0; i < progress; i++) {
    result[nthNext(forehand, i)] = Declaration.gesund;
  }
  return result;
}

export function createAuction({ forehand, progress }: AuctionConfig): Auction {
  const declarations = createDeclarations(forehand, progress);
  return {
    declarations,
    ownDeclaration: declarations[Pos.bottom] !== undefined ? GameType.MARRIAGE : undefined,
  };
}

export interface GameConfig {
  numPlayedTricks?: number;
  forehand: Pos;
  progress: number;
}

export function createGame({ numPlayedTricks, forehand, progress }: GameConfig): Game {
  const trick = mc.trick(forehand, progress, progress === 4 ? Pos.left : undefined);
  return {
    bids: emptyBids(),
    completedTricks: numPlayedTricks ?? 0,
    mode: {
      type: GameType.NORMAL_GAME,
      forehand,
    },
    currentTrick: trick,
  };
}

export function createMatch({ phase, forehand, progress, numPlayedTricks }: MatchConfig): Match {
  return {
    players: mp.players,
    phase: phase,
    auction: phase === MatchPhase.AUCTION ? createAuction({ forehand, progress }) : null,
    cards: fullHand.slice(numPlayedTricks ?? 0),
    turn: nthNext(forehand, progress),
    game: phase === MatchPhase.GAME ? createGame({ numPlayedTricks, forehand, progress }) : null,
    winner: numPlayedTricks === 12 && progress === 4 ? Party.RE : null,
  };
}
