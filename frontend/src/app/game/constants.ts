import { createAction, PayloadActionCreator, PrepareAction, ThunkAction } from '@reduxjs/toolkit';
import { AuthenticatedClient } from 'api/client';
import { selectAuthenticatedClientOrThrow } from 'app/session';
import { RootState } from 'app/store';
import { Action, AnyAction } from 'redux';
import { ActionKind } from './state';
import { startTable } from './table';

export interface TableAction extends Action<string> {}
export interface MatchAction extends TableAction {}

export function isTableAction(action: AnyAction): action is TableAction {
  return (
    typeof action.type === 'string' &&
    (action.type.startsWith('game/table') || action.type === startTable.fulfilled.type)
  );
}

export function isMatchAction(action: AnyAction): action is MatchAction {
  return typeof action.type === 'string' && action.type.startsWith('game/table/match');
}

interface GameThunkPayloadCreatorMeta {
  client: AuthenticatedClient;
  tableId?: string;
}

type GameThunkPayloadCreator<Returned, ThunkArg = void> = (
  arg: ThunkArg,
  meta: GameThunkPayloadCreatorMeta
) => Promise<Returned>;

export type GameThunk<Returned, ThunkArg> = ((
  arg: ThunkArg
) => ThunkAction<any, RootState, any, Action<string>>) & {
  fulfilled: PayloadActionCreator<Returned, string, PrepareAction<Returned>>;
};
export const gameActionPending = createAction('game/action/pending', (kind: ActionKind) => ({
  payload: kind,
}));
export const gameActionError = createAction(
  'game/action/error',
  (kind: ActionKind, error: any) => ({
    payload: { kind, error },
  })
);

export function createGameThunk<Returned, ThunkArg = void>(
  kind: ActionKind,
  payloadCreator: GameThunkPayloadCreator<Returned, ThunkArg>
): GameThunk<Returned, ThunkArg> {
  const fulfilled = createAction('game/action/' + kind, (result: Returned) => {
    return { payload: result };
  });
  const actionCreator = function (
    arg: ThunkArg
  ): ThunkAction<any, RootState, undefined, Action<string>> {
    return async (dispatch, getState) => {
      dispatch(gameActionPending(kind));
      try {
        const state = getState();
        const authenticatedClient = selectAuthenticatedClientOrThrow(state);
        const selectors = await require('./selectors'); //import { selectTable } from './selectors';
        const table = selectors.selectTable(state);
        const result = await payloadCreator(arg, {
          client: authenticatedClient,
          tableId: table?.table.id,
        });
        dispatch(fulfilled(result));
      } catch (error) {
        dispatch(gameActionError(kind, error));
      }
    };
  };
  return Object.assign(actionCreator, { fulfilled });
}
