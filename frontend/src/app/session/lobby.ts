import { createAsyncThunk, ActionReducerMapBuilder } from "@reduxjs/toolkit";
import { Table } from "model/table";
import { AsyncThunkConfig } from "app/store";
import { selectAuthenticatedClientOrThrow, SessionState } from ".";
import { toTable } from "model/apiconv";
import { Empty } from "api/karlchen_pb";

export const createTable = createAsyncThunk<Table, void, AsyncThunkConfig>(
  "lobby/createTable",
  async (_, thunkAPI) => {
    const { client, meta } = selectAuthenticatedClientOrThrow(
      thunkAPI.getState()
    );
    const result = await client.createTable(new Empty(), meta);
    const table = toTable(result);
    return table;
  }
);

export const addReducerCases = (
  builder: ActionReducerMapBuilder<SessionState>
) => {
  builder.addCase(createTable.fulfilled, (state, { payload: table }) => {
    state.currentTable = {
      table,
    };
  });
};
