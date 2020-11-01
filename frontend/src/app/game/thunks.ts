import { toTable, toTableState } from 'model/apiconv';
import * as api from 'api/karlchen_pb';
import { ActionKind, createGameThunk } from './asyncs';
import { Table } from 'model/table';

export const createTable = createGameThunk<void, Table>(
  ActionKind.createTable,
  async (_, { client: { client, meta } }) => {
    const result = await client.createTable(new api.Empty(), meta);
    return toTable(result, api.TablePhase.NOT_STARTED);
  }
);

export const joinTable = createGameThunk(
  ActionKind.joinTable,
  async (inviteCode: string, { client: { client, meta } }) => {
    const req = new api.JoinTableRequest();
    req.setInviteCode(inviteCode);
    const table = await client.joinTable(req, meta);
    return toTableState(table);
  }
);
