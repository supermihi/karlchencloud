import { User } from './core';
import { Match } from './match';
import { TablePhase } from '../api/karlchen_pb';

export interface Table {
  id: string;
  owner: string;
  invite?: string;
  players: User[];
  created: string;
}

export function canStartTable(t: TableState) {
  return t.phase === TablePhase.NOT_STARTED && t.table.players.length >= 4;
}

export interface TableState {
  table: Table;
  match?: Match;
  phase: TablePhase;
}

export type Players = [string, string, string, string]; // from self
export function nthNext(players: Players, player: string, i: number): string {
  const playerIndex = players.indexOf(player);
  if (playerIndex === -1) {
    throw new Error('user id not in players list');
  }
  return players[(playerIndex + i) % players.length];
}
