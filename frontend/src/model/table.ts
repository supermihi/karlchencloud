import { User } from './core';
import { Match } from './match';
import { TablePhase } from '../api/karlchen_pb';

export interface Table {
  id: string;
  owner: string;
  invite?: string;
  players: User[];
  created: string;
  phase: TablePhase;
}

export function canStartTable(t: Table) {
  return t.phase === TablePhase.NOT_STARTED && t.players.length >= 4;
}
export function canContinueTable(t: Table) {
  return t.phase === TablePhase.BETWEEN_GAMES || t.phase === TablePhase.PLAYING;
}
export function waitingForPlayers(t: Table) {
  return t.phase === TablePhase.NOT_STARTED && t.players.length < 4;
}

export interface TableState {
  table: Table;
  match?: Match;
}
