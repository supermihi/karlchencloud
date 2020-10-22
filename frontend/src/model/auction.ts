import { GameType } from 'api/karlchen_pb';
import { Mode } from './match';
import { PartialPlayerMap, Pos } from './players';

export interface DeclareResult {
  mode: Mode | null;
  player: Pos;
  declaration: Declaration;
}

export enum Declaration {
  gesund,
  vorbehalt,
}

export interface Auction {
  declarations: PartialPlayerMap<Declaration>;
  ownDeclaration?: GameType;
}

export function isVorbehalt(gt: GameType): boolean {
  return !(gt === GameType.NORMAL_GAME || gt === GameType.MARRIAGE);
}

export function emptyAuction(): Auction {
  return { declarations: {} };
}
