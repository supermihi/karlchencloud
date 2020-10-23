import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { Table } from 'model/table';
import { User } from 'model/core';
import { ActionKind, createGameThunk } from './asyncs';
import { toMatch } from 'model/apiconv';
import { tableId } from 'api/helpers';
import * as api from 'api/karlchen_pb';
import * as events from 'app/session/events';
import { createTable, joinTable } from './thunks';
import { AsyncState } from './asyncs';

export const startTable = createGameThunk(
  ActionKind.startTable,
  async (id: string, { client: { client, meta } }) => {
    const match = await client.startTable(tableId(id), meta);
    return toMatch(match);
  }
);

export type TableState = (Table & AsyncState) | null;
const slice = createSlice({
  name: 'game/table',
  initialState: null as TableState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(events.tableChanged, (_, { payload }) => payload?.table ?? null)
      .addCase(events.memberJoined, (table, { payload }) => {
        table?.players.push(payload);
      })
      .addCase(events.memberLeft, (table, { payload: id }: PayloadAction<string>) => {
        if (table === null) return;
        const index = table.players.findIndex((p) => p.id === id);
        if (index !== -1) {
          table.players.splice(index, 1);
        }
      })
      .addCase(events.memberStatusChanged, (table, { payload: user }: PayloadAction<User>) => {
        if (table === null) return;
        const index = table.players.findIndex((p) => p.id === user.id);
        if (index !== -1) {
          table.players.splice(index, 1, user);
        }
      })
      .addCase(events.matchStarted, (table) => {
        if (table === null) return;
        table.phase = api.TablePhase.PLAYING;
      })
      .addCase(startTable.fulfilled, (table) => {
        if (table === null) return;
        table.phase = api.TablePhase.PLAYING;
      })
      .addCase(createTable.fulfilled, (_, { payload: table }) => table)
      .addCase(joinTable.fulfilled, (_, { payload: table }) => table.table);
  },
});
export const { actions, reducer } = slice;
