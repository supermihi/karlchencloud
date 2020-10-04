import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { TableState } from 'model/table';
import { User } from 'model/core';
import { Match } from 'model/match';
import { ActionKind } from './state';
import { toMatch } from 'model/apiconv';
import { tableId } from 'api/helpers';
import * as pb from 'api/karlchen_pb';
import { createGameThunk, isMatchAction } from './constants';
import matchReducer from './match';

export const startTable = createGameThunk(
  ActionKind.startTable,
  async (id: string, { client: { client, meta } }) => {
    const match = await client.startTable(tableId(id), meta);
    return toMatch(match);
  }
);

const slice = createSlice({
  name: 'game/table',
  initialState: {} as TableState,
  reducers: {
    memberJoined: (state, { payload }: PayloadAction<User>) => {
      state.table.players.push(payload);
    },
    memberLeft: ({ table }, { payload: id }: PayloadAction<string>) => {
      const index = table.players.findIndex((p) => p.id === id);
      if (index !== -1) {
        table.players.splice(index, 1);
      }
    },
    memberStatusChanged: ({ table }, { payload: user }: PayloadAction<User>) => {
      const index = table.players.findIndex((p) => p.id === user.id);
      if (index !== -1) {
        table.players.splice(index, 1, user);
      }
    },
    matchStarted: (table, { payload: match }: PayloadAction<Match>) => {
      table.match = match;
      table.phase = pb.TablePhase.PLAYING;
    },
  },
  extraReducers: (builder) => {
    builder.addCase(startTable.fulfilled, (table, { payload }) => {
      table.match = payload;
      table.phase = pb.TablePhase.PLAYING;
    });
    builder.addMatcher(isMatchAction, (state, action) => {
      if (state.match === null) {
        return;
      }
      matchReducer(state.match, action);
    });
  },
});
export const { actions, reducer } = slice;
