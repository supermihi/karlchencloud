import { toDeclareResult, toMatch, toTable, toTableState } from 'model/apiconv';
import * as api from 'api/karlchen_pb';
import { ActionKind, createGameThunk } from './gameActions';
import { Table } from 'model/table';
import { Match, PlayedCard } from 'model/match';
import { tableId } from 'api/helpers';
import { Card } from 'model/core';
import { selectCurrentTableOrThrow, selectPlayers } from './selectors';
import { newDeclareRequest, newPlayCardRequest } from 'api/modelToPb';
import { getPosition, Pos } from 'model/players';
import { selectAuthenticatedClientOrThrow } from 'session/selectors';
import { DeclareResult } from 'model/auction';

export const createTable = createGameThunk<void, Table>(
  ActionKind.createTable,
  async (_, { client, meta }) => {
    const result = await client.createTable(new api.Empty(), meta);
    return toTable(result, api.TablePhase.NOT_STARTED);
  }
);

export const joinTable = createGameThunk(
  ActionKind.joinTable,
  async (inviteCode: string, { client, meta }) => {
    const req = new api.JoinTableRequest();
    req.setInviteCode(inviteCode);
    const table = await client.joinTable(req, meta);
    return toTableState(table);
  }
);

export const startTable = createGameThunk<void, Match>(
  ActionKind.startTable,
  async (_, { client, meta, getState }) => {
    const id = selectCurrentTableOrThrow(getState()).id;
    const match = await client.startTable(tableId(id), meta);
    return toMatch(match);
  }
);

export const playCard = createGameThunk<Card, PlayedCard>(
  ActionKind.playCard,
  async (card, { client, meta, getState }) => {
    const tableId = selectCurrentTableOrThrow(getState()).id;
    const req = newPlayCardRequest(card, tableId);
    const result = await client.playCard(req, meta);
    let winner: Pos | undefined = undefined;
    if (result.hasTrickWinner()) {
      const winnerId = (result.getTrickWinner() as api.PlayerValue).getUserId();
      const players = selectPlayers(getState());
      winner = getPosition(players, winnerId);
    }
    return { card, player: Pos.bottom, trickWinner: winner };
  }
);

export const declare = createGameThunk<api.GameType, DeclareResult & { gametype: api.GameType }>(
  ActionKind.declare,
  async (gametype: api.GameType, thunkAPI) => {
    const state = thunkAPI.getState();
    const table = selectCurrentTableOrThrow(state);
    const req = newDeclareRequest(gametype, table.id);
    const { client, meta } = selectAuthenticatedClientOrThrow(state);
    const ans = await client.declare(req, meta);
    const players = selectPlayers(state);
    return { ...toDeclareResult(ans, players), gametype };
  }
);
