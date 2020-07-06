import { createSlice, PayloadAction, ActionReducerMapBuilder } from '@reduxjs/toolkit';
import { TableState } from 'model/table';
import { initialState, ActionKind } from './state';
import * as api from 'api/karlchen_pb';
import type { RootState } from 'app/store';
import { createTable, joinTable, GameThunk } from './thunks';
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
      .addCase(createTable.thunk.fulfilled, (state, { payload: table }) => {
        clearPendingAndError(state);
        state.currentTable = { table, phase: api.TablePhase.NOT_STARTED };
      })
      .addCase(joinTable.thunk.fulfilled, (state, { payload: table }) => {
        clearPendingAndError(state);
        state.currentTable = table;
      });
    reducePendingAndRejected(builder, createTable, joinTable, table.startTable);
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

function reducePendingAndRejected(
  builder: ActionReducerMapBuilder<typeof initialState>,
  ...thunks: GameThunk<any, any>[]
) {
  for (const gt of thunks) {
    reduceGameThunkPendingAndRejected(builder, gt);
  }
}
function reduceGameThunkPendingAndRejected(
  builder: ActionReducerMapBuilder<typeof initialState>,
  { thunk, kind }: GameThunk<any, any>
) {
  return builder
    .addCase(thunk.pending, (state) => {
      state.pendingAction = kind;
    })
    .addCase(thunk.rejected, (state, { error }) => {
      state.pendingAction = ActionKind.noAction;
      state.error = { action: kind, error };
    });
}

export const actions = gameSlice.actions;
export const selectGame = (state: RootState) => state.game;
export default gameSlice.reducer;
