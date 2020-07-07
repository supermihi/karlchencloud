import { RootState } from 'app/store';
import { selectAuthenticatedClientOrThrow } from 'app/session';
import { toTable, toTableState } from 'model/apiconv';
import { AuthenticatedClient } from 'api/client';
import * as api from 'api/karlchen_pb';
import {
  createAction,
  ThunkAction,
  Action,
  PayloadActionCreator,
  PrepareAction,
} from '@reduxjs/toolkit';
import { ActionKind } from './state';

export const createTable = createGameThunk(ActionKind.createTable, async (_, { client, meta }) => {
  const result = await client.createTable(new api.Empty(), meta);
  return toTable(result);
});

export const joinTable = createGameThunk(
  ActionKind.joinTable,
  async (inviteCode: string, { client, meta }) => {
    const req = new api.JoinTableRequest();
    req.setInviteCode(inviteCode);
    const table = await client.joinTable(req, meta);
    return toTableState(table);
  }
);

type GameThunkPayloadCreator<Returned, ThunkArg = void> = (
  arg: ThunkArg,
  client: AuthenticatedClient
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
        const authenticatedClient = selectAuthenticatedClientOrThrow(getState());
        const result = await payloadCreator(arg, authenticatedClient);
        dispatch(fulfilled(result));
      } catch (error) {
        dispatch(gameActionError(kind, error));
      }
    };
  };
  return Object.assign(actionCreator, { fulfilled });
}
