import { ActionReducerMapBuilder, createAsyncThunk } from '@reduxjs/toolkit';
import { Table } from 'model/table';
import { AsyncThunkConfig } from 'app/store';
import { selectAuthenticatedClientOrThrow, SessionState } from '.';
import { toTable } from 'model/apiconv';
import { Empty, TablePhase } from 'api/karlchen_pb';

export const createTable = createAsyncThunk<Table, void, AsyncThunkConfig>(
  'lobby/createTable',
  async (_, thunkAPI) => {
    const { client, meta } = selectAuthenticatedClientOrThrow(
      thunkAPI.getState()
    );
    const result = await client.createTable(new Empty(), meta);
    return toTable(result);
  }
);

export const addReducerCases = (
  builder: ActionReducerMapBuilder<SessionState>
) => {
  builder.addCase(createTable.fulfilled, (state, { payload: table }) => {
    state.currentTable = {
      table,
      phase: TablePhase.NOT_STARTED,
    };
  });
};
