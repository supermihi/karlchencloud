import { RootState } from './store';
import { selectAuth } from 'app/auth/slice';
import { selectSession } from './session';
import { selectGame } from './game';
import { TablePhase } from 'api/karlchen_pb';

export enum Location {
  register,
  login,
  lobby,
  table,
}
export function selectLocation(state: RootState): Location {
  if (selectSession(state).session) {
    if (selectGame(state).table?.phase === TablePhase.PLAYING) {
      return Location.table;
    }
    return Location.lobby;
  }
  return selectAuth(state).storedLogin ? Location.login : Location.register;
}
