import {
  Action,
  createAction,
  Draft,
  PayloadActionCreator,
  PrepareAction,
  ThunkAction,
} from '@reduxjs/toolkit';
import { AuthenticatedClient } from 'api/client';
import { selectAuthenticatedClientOrThrow } from 'app/session';
import { RootState, AppThunk } from 'app/store';

export enum ActionKind {
  noAction = 'noAction',
  joinTable = 'joinTable',
  createTable = 'createTable',
  startTable = 'startTable',
  playCard = 'playCard',
  placeBid = 'placeBid',
  declare = 'declare',
}

export interface ActionError {
  action: ActionKind;
  error: any;
}

export interface AsyncState {
  pendingAction?: ActionKind;
  error?: ActionError;
}
export function clearPendingAndError(state: Draft<AsyncState>) {
  state.pendingAction = ActionKind.noAction;
  state.error = undefined;
}

interface GameThunkConfig {
  client: AuthenticatedClient;
  getState: () => RootState;
}
export const actionStarted = createAction<ActionKind>('ACTION_STARTED');
export const actionSucceeded = createAction<ActionKind>('ACTION_SUCCEEDED');
export const actionFailed = createAction<ActionError>('ACTION_FAILED');

export type GameThunk<Returned, ThunkArg> = ((
  arg: ThunkArg
) => ThunkAction<any, RootState, any, Action<string>>) & {
  fulfilled: PayloadActionCreator<Returned, string, PrepareAction<Returned>>;
};
export function createGameThunk<TArg, Returned = void>(
  actionKind: ActionKind,
  creator: (thunkArg: TArg, api: GameThunkConfig) => Promise<Returned>
): GameThunk<Returned, TArg> {
  const action = createAction('game' + actionKind, (result: Returned) => ({
    payload: result,
  }));
  const thunk: (arg: TArg) => AppThunk = (arg) => async (dispatch, getState) => {
    dispatch(actionStarted(actionKind));
    const client = selectAuthenticatedClientOrThrow(getState());
    const config = { client, getState };
    try {
      const result = await creator(arg, config);
      dispatch(action(result));
      dispatch(actionSucceeded(actionKind));
      return result;
    } catch (err) {
      dispatch(actionFailed({ error: err, action: actionKind }));
    }
  };
  return Object.assign(thunk, { fulfilled: action });
}
