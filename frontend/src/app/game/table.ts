import { createSlice, PayloadAction, AnyAction, Action } from '@reduxjs/toolkit';
import { TableState } from 'model/table';
import { User } from 'model/core';
import { Match } from 'model/match';
import { createGameThunk } from './thunks';
import { ActionKind } from './state';
import { toMatch } from 'model/apiconv';
import { tableId } from 'api/helpers';
import { TablePhase } from 'api/karlchen_pb';

export const name = 'game/table';

export const startTable = createGameThunk(
  ActionKind.startTable,
  async (id: string, { client, meta }) => {
    const match = await client.startTable(tableId(id), meta);
    return toMatch(match);
  }
);
const slice = createSlice({
  name,
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
      table.phase = TablePhase.PLAYING;
    },
  },
  extraReducers: (builder) => {
    builder.addCase(startTable.fulfilled, (table, { payload }) => {
      table.match = payload;
      table.phase = TablePhase.PLAYING;
    });
  },
});

export interface TableAction extends Action<string> {}
export function isTableAction(action: AnyAction): action is TableAction {
  return (
    typeof action.type === 'string' &&
    (action.type.startsWith(name) || action.type === startTable.fulfilled.type)
  );
}

export const { actions, reducer } = slice;
