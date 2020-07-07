import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { TableState } from 'model/table';
import { initialState, ActionKind } from './state';
import * as api from 'api/karlchen_pb';
import type { RootState } from 'app/store';
import { createTable, joinTable, gameActionPending, gameActionError } from './thunks';
import * as table from './table';

const gameSlice = createSlice({
  name: 'game',
  initialState,
  reducers: {
    currentTableChanged: (state, { payload: currentTable }: PayloadAction<TableState | null>) => {
      state.currentTable = currentTable;
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
      })
      .addCase(gameActionPending, (state, { payload }) => {
        state.pendingAction = payload;
      })
      .addCase(gameActionError, (state, { payload: { error, kind } }) => {
        state.pendingAction = ActionKind.noAction;
        state.error = { action: kind, error };
      });
    builder.addMatcher(table.isTableAction, (state, action) => {
      if (state.currentTable === null) {
        return;
      }
      table.reducer(state.currentTable, action);
    });
  },
});

function clearPendingAndError(state: typeof initialState) {
  state.pendingAction = ActionKind.noAction;
  state.error = undefined;
}

export const actions = gameSlice.actions;
export const selectGame = (state: RootState) => state.game;
export default gameSlice.reducer;
