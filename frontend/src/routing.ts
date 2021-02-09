import { RootState } from './state';
import { TablePhase } from 'api/karlchen_pb';
import { selectSession } from 'session/selectors';
import { selectGame } from 'game/selectors';

export enum Location {
  register,
  login,
  lobby,
  table,
}
export function selectLocation(state: RootState): Location {
  const session = selectSession(state);
  if (session.session) {
    if (selectGame(state).table?.phase === TablePhase.PLAYING) {
      return Location.table;
    }
    return Location.lobby;
  }
  return session.storedLogin ? Location.login : Location.register;
}
