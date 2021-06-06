import { RootState } from './state';
import { TablePhase } from 'api/karlchen_pb';
import { selectSession } from 'session/selectors';
import { selectPlay } from 'play/selectors';
import { SessionPhase } from './session/model';

export enum Location {
  login,
  loginWithToken,
  lobby,
  table,
}

export function selectLocation(state: RootState): Location {
  const session = selectSession(state);
  switch (session.phase) {
    case SessionPhase.NoToken:
    case SessionPhase.ObtainingToken:
      return Location.login;
    case SessionPhase.TokenObtained:
    case SessionPhase.Starting:
      return Location.loginWithToken;
    case SessionPhase.Established:
      return selectPlay(state).table?.phase === TablePhase.PLAYING
        ? Location.table
        : Location.lobby;
  }
}
