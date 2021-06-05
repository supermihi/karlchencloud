import {
  Action,
  createAction,
  createReducer,
  PayloadActionCreator,
  PrepareAction,
  ThunkAction,
} from '@reduxjs/toolkit';
import { AuthenticatedClient } from 'api/client';
import { selectAuthenticatedClientOrThrow } from 'session/selectors';
import { RootState, AppThunk } from 'state';

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
  error: unknown;
}

interface PlayThunkConfig extends AuthenticatedClient {
  getState: () => RootState;
}
export const actionStarted = createAction<ActionKind>('ACTION_STARTED');
export const actionSucceeded = createAction<ActionKind>('ACTION_SUCCEEDED');
export const actionFailed = createAction<ActionError>('ACTION_FAILED');
export interface ActionState {
  pending: boolean;
  lastAction?: ActionKind;
  error?: unknown;
}

export const reducer = createReducer<ActionState>(
  { pending: false },
  {
    actionStarted: (_, { payload }) => ({ pending: true, lastAction: payload }),
    actionSucceded: (_, { payload }) => ({ pending: false, lastAction: payload }),
    actionFailed: (_, { payload: { action, error } }) => ({
      pending: false,
      lastAction: action,
      error,
    }),
  }
);
export type PlayThunk<Returned, ThunkArg> = ((
  arg: ThunkArg
) => ThunkAction<unknown, RootState, unknown, Action<string>>) & {
  fulfilled: PayloadActionCreator<Returned, string, PrepareAction<Returned>>;
};
export function createPlayThunk<TArg, Returned = void>(
  actionKind: ActionKind,
  creator: (thunkArg: TArg, api: PlayThunkConfig) => Promise<Returned>
): PlayThunk<Returned, TArg> {
  const action = createAction('play/' + actionKind, (result: Returned) => ({
    payload: result,
  }));
  const thunk: (arg: TArg) => AppThunk = (arg) => async (dispatch, getState) => {
    dispatch(actionStarted(actionKind));
    const state = getState();
    const { client, meta } = selectAuthenticatedClientOrThrow(state);
    const config = { client, meta, getState };
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
