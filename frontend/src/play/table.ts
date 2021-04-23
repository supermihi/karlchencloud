import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { User } from 'model/core';
import * as api from 'api/karlchen_pb';
import * as events from 'session/events';
import { createTable, joinTable, startTable } from './thunks';
import { Table } from 'model/table';

const tableSlice = createSlice({
  name: 'game/table',
  initialState: null as Table | null,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(events.tableChanged, (_, { payload }) => payload?.table ?? null)
      .addCase(events.memberJoined, (table, { payload }) => {
        table?.members.push(payload);
      })
      .addCase(events.memberLeft, (table, { payload: id }: PayloadAction<string>) => {
        if (table === null) return;
        const index = table.members.findIndex((p) => p.id === id);
        if (index !== -1) {
          table.members.splice(index, 1);
        }
      })
      .addCase(events.memberStatusChanged, (table, { payload: user }: PayloadAction<User>) => {
        if (table === null) return;
        const index = table.members.findIndex((p) => p.id === user.id);
        if (index !== -1) {
          table.members.splice(index, 1, user);
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
      .addCase(joinTable.fulfilled, (_, { payload: { table } }) => table);
  },
});

export const { actions, reducer } = tableSlice;
