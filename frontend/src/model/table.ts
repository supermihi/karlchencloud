import { User } from './core';
import { TablePhase } from 'api/karlchen_pb';

export interface Table {
  id: string;
  owner: string;
  invite?: string;
  members: User[];
  created: string;
  phase: TablePhase;
}

export function canStartTable(t: Table, me: User): boolean {
  return t.owner === me.id && t.phase === TablePhase.NOT_STARTED && t.members.length >= 4;
}
export function canContinueTable(t: Table): boolean {
  return t.phase === TablePhase.BETWEEN_GAMES || t.phase === TablePhase.PLAYING;
}
export function waitingForPlayers(t: Table): boolean {
  return t.phase === TablePhase.NOT_STARTED && t.members.length < 4;
}
