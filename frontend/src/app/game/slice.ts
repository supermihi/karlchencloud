import { createSlice, PayloadAction, AsyncThunk, ActionReducerMapBuilder } from '@reduxjs/toolkit';
import { TableState } from 'model/table';
import { initialState, ActionKind } from './state';
import * as api from 'api/karlchen_pb';
import type { RootState } from 'app/store';
import { createTable, joinTable } from './thunks';
import { User } from 'model/core';
import { Match } from 'model/match';

const gameSlice = createSlice({
  name: 'game',
  initialState,
  reducers: {
    currentTableChanged: (state, { payload: currentTable }: PayloadAction<TableState | null>) => {
      state.currentTable = currentTable;
    },
    memberJoined: (state, { payload }: PayloadAction<User>) => {
      state.currentTable?.table.players.push(payload);
    },
    memberLeft: ({ currentTable }, { payload: id }: PayloadAction<string>) => {
      if (!currentTable) {
        return;
      }
      const index = currentTable.table.players.findIndex((p) => p.id === id);
      if (index !== -1) {
        currentTable.table.players.splice(index, 1);
      }
    },
    memberStatusChanged: ({ currentTable }, { payload: user }: PayloadAction<User>) => {
      if (!currentTable) {
        return;
      }
      const index = currentTable.table.players.findIndex((p) => p.id === user.id);
      if (index !== -1) {
        currentTable.table.players.splice(index, 1, user);
      }
    },
    matchStarted: ({ currentTable }, { payload: match }: PayloadAction<Match>) => {
      currentTable && (currentTable.match = match);
    },
  },
  extraReducers: (builder) => {
    builder
      .addCase(createTable.fulfilled, (state, { payload: table }) => {
        clearPendingAndError(state);
        state.currentTable = { table, phase: api.TablePhase.NOT_STARTED };
      })
      .addCase(joinTable.fulfilled, (state, { payload: table }) => {
        clearPendingAndError(state);
        state.currentTable = table;
      });
    reduceGameThunkPendingAndRejected(builder, createTable, ActionKind.createTable);
    reduceGameThunkPendingAndRejected(builder, joinTable, ActionKind.joinTable);
  },
});
function clearPendingAndError(state: typeof initialState) {
  state.pendingAction = ActionKind.noAction;
  state.error = undefined;
}
function reduceGameThunkPendingAndRejected<Returned, ThunkArg = void>(
  builder: ActionReducerMapBuilder<typeof initialState>,
  thunk: AsyncThunk<Returned, ThunkArg, {}>,
  actionKind: ActionKind
) {
  return builder
    .addCase(thunk.pending, (state) => {
      state.pendingAction = actionKind;
    })
    .addCase(thunk.rejected, (state, { error }) => {
      state.pendingAction = ActionKind.noAction;
      state.error = { action: actionKind, error };
    });
}

export const actions = gameSlice.actions;
export const selectGame = (state: RootState) => state.game;
export default gameSlice.reducer;
