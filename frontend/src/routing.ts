import { RootState } from './state';
import { TablePhase } from 'api/karlchen_pb';
import { selectSession } from 'session/selectors';
import { selectPlay } from 'play/selectors';

export enum Location {
  register,
  login,
  loginWithToken,
  lobby,
  table,
}
export function selectLocation(state: RootState): Location {
  const session = selectSession(state);
  if (session.activeSession) {
    if (selectPlay(state).table?.phase === TablePhase.PLAYING) {
      return Location.table;
    }
    return Location.lobby;
  }
  if (session.storedLogin) {
    return Location.loginWithToken;
  }
  return Location.login;
}
