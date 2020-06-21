import { User, Card } from "./core";
import { MatchPhase, BidType, GameType } from "api/karlchen_pb";

export interface Table {
  id: string;
  owner: string;
  invite?: string;
  players: User[];
}

export interface TableState {
  table: Table;
  match?: Match;
}
export interface Match {
  phase: MatchPhase;
  turn?: string;
  players: Players;
  cards?: Card[];
  details: Auction | Game;
}


export function isAuction(aog: Auction | Game): aog is Auction {
  return "declarations" in aog;
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


export type Players = [string, string, string, string]; // from self
export function nthNext(players: Players, player: string, i: number): string {
  const playerIndex = players.indexOf(player);
  if (playerIndex === -1) {
    throw new Error("user id not in players list");
  }
  return players[(playerIndex + i) % players.length];
}


export enum Declaration {
  gesund,
  vorbehalt,
}
export interface Auction {
  declarations: { [player: string]: Declaration };
}

