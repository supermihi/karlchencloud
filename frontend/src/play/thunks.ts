import * as apiconv from 'model/apiconv';
import * as api from 'api/karlchen_pb';
import { ActionKind, createPlayThunk } from './playActions';
import { Table } from 'model/table';
import { Match, PlayedCard } from 'model/match';
import { Card } from 'model/core';
import { selectCurrentTableOrThrow, selectPlayers } from './selectors';
import { newDeclareRequest, newPlayCardRequest } from 'api/modelToPb';
import { selectAuthenticatedClientOrThrow } from 'session/selectors';
import { DeclareResult } from 'model/auction';

export const createTable = createPlayThunk<void, Table>(
  ActionKind.createTable,
  async (_, { client, meta }) => {
    const request = new api.CreateTableRequest();
    request.setPublic(true);
    const result = await client.createTable(request, meta);
    return apiconv.toTable(result, api.TablePhase.NOT_STARTED);
  }
);

export const joinTable = createPlayThunk(
  ActionKind.joinTable,
  async (inviteCode: string, { client, meta }) => {
    const req = new api.JoinTableRequest();
    req.setInviteCode(inviteCode);
    const table = await client.joinTable(req, meta);
    return apiconv.toTableState(table);
  }
);

export const startTable = createPlayThunk<void, Match>(
  ActionKind.startTable,
  async (_, { client, meta, getState }) => {
    const id = selectCurrentTableOrThrow(getState()).id;
    const request = new api.StartTableRequest();
    request.setTableId(id);
    const match = await client.startTable(request, meta);
    return apiconv.toMatch(match);
  }
);

export const playCard = createPlayThunk<Card, PlayedCard>(
  ActionKind.playCard,
  async (card, { client, meta, getState }) => {
    const tableId = selectCurrentTableOrThrow(getState()).id;
    const req = newPlayCardRequest(card, tableId);
    const result = await client.playCard(req, meta);
    const players = selectPlayers(getState());
    return apiconv.toPlayedCard(result, players);
  }
);

export const declare = createPlayThunk<api.GameType, DeclareResult & { gametype: api.GameType }>(
  ActionKind.declare,
  async (gametype: api.GameType, thunkAPI) => {
    const state = thunkAPI.getState();
    const table = selectCurrentTableOrThrow(state);
    const req = newDeclareRequest(gametype, table.id);
    const { client, meta } = selectAuthenticatedClientOrThrow(state);
    const ans = await client.declare(req, meta);
    const players = selectPlayers(state);
    return { ...apiconv.toDeclareResult(ans, players), gametype };
  }
);
