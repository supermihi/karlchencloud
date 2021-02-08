import { createAsyncThunk, createSlice, PayloadAction } from '@reduxjs/toolkit';
import { Table } from 'model/table';
import { User } from 'model/core';
import { toMatch } from 'model/apiconv';
import { tableId } from 'api/helpers';
import * as api from 'api/karlchen_pb';
import * as events from 'session/events';
import { createTable, joinTable } from '../game/thunks';

export interface TableState {
  table: Table | null;
  loading: boolean;
  error?: unknown;
}
const initialState: TableState = { table: null, loading: false };

const slice = createSlice({
  name: 'game/table',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(events.tableChanged, (state, { payload }) => ({
        ...state,
        table: payload?.table ?? null,
      }))
      .addCase(events.memberJoined, ({ table }, { payload }) => {
        table?.players.push(payload);
      })
      .addCase(events.memberLeft, ({ table }, { payload: id }: PayloadAction<string>) => {
        if (table === null) return;
        const index = table.players.findIndex((p) => p.id === id);
        if (index !== -1) {
          table.players.splice(index, 1);
        }
      })
      .addCase(events.memberStatusChanged, ({ table }, { payload: user }: PayloadAction<User>) => {
        if (table === null) return;
        const index = table.players.findIndex((p) => p.id === user.id);
        if (index !== -1) {
          table.players.splice(index, 1, user);
        }
      })
      .addCase(events.matchStarted, ({ table }) => {
        if (table === null) return;
        table.phase = api.TablePhase.PLAYING;
      })
      .addCase(startTable.fulfilled, ({ table }) => {
        if (table === null) return;
        table.phase = api.TablePhase.PLAYING;
      })
      .addCase(createTable.fulfilled, (state, { payload: table }) => {
        state.table = table;
      })
      .addCase(joinTable.fulfilled, (state, { payload: table }) => (state.table = table));
  },
});
export const { actions, reducer } = slice;
