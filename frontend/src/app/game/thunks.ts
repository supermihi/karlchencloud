import { AsyncThunkConfig } from 'app/store';
import { selectAuthenticatedClientOrThrow } from 'app/session';
import { toTable, toTableState } from 'model/apiconv';
import { AuthenticatedClient } from 'api/client';
import * as api from 'api/karlchen_pb';
import { createAsyncThunk } from '@reduxjs/toolkit';

export const createTable = createGameThunk('lobby/createTable', async (_, { client, meta }) => {
  const result = await client.createTable(new api.Empty(), meta);
  return toTable(result);
});

type GameThunkPayloadCreator<Returned, ThunkArg = void> = (
  arg: ThunkArg,
  client: AuthenticatedClient
) => Promise<Returned>;

export const joinTable = createGameThunk(
  'game/joinTable',
  async (inviteCode: string, { client, meta }) => {
    const req = new api.JoinTableRequest();
    req.setInviteCode(inviteCode);
    const table = await client.joinTable(req, meta);
    return toTableState(table);
  }
);

export function createGameThunk<Returned, ThunkArg = void>(
  name: string,
  payloadCreator: GameThunkPayloadCreator<Returned, ThunkArg>
) {
  return createAsyncThunk<Returned, ThunkArg, AsyncThunkConfig>(name, (arg, api) => {
    const AuthenticatedClient = selectAuthenticatedClientOrThrow(api.getState());
    return payloadCreator(arg, AuthenticatedClient);
  });
}
