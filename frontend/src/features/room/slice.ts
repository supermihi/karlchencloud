import { Table } from "model/table";
import { createSlice, createAsyncThunk } from "@reduxjs/toolkit";
import { selectClient } from "features/auth/slice";
import { RootState, AsyncThunkConfig } from "app/store";
import { Empty, TableData } from "api/karlchen_pb";
import { setLocation } from "core/routing";
import { setTable } from "features/table/slice";

export interface RoomState {
  tables: Table[];
  loading: boolean;
  loaded: boolean;
  error?: any;
}

const initialState: RoomState = { tables: [], loading: false, loaded: false };

export const fetchTables = createAsyncThunk<Table[], void, AsyncThunkConfig>(
  "room/fetchTables",
  async (_, thunkAPI) => {
    const { client, meta } = selectClient(thunkAPI.getState());
    if (!client) {
      throw new Error("client is null");
    }
    const result = await client.listTables(new Empty(), meta);
    const tables = result.getTablesList().map(toTable);
    const myTable = tables.find((t) => t.meAtTable);
    if (myTable) {
      thunkAPI.dispatch(setTable(myTable));
    }
    return tables;
  }
);

function toTable(t: TableData): Table {
  return {
    owner: t.getOwner(),
    id: t.getTableId(),
    meAtTable: t.getYouAtTable(),
  };
}
export const createTable = createAsyncThunk<Table, void, AsyncThunkConfig>(
  "room/createTable",
  async (_, thunkAPI) => {
    const { client, meta } = selectClient(thunkAPI.getState());
    const result = await client.createTable(new Empty(), meta);
    const table = toTable(result);
    thunkAPI.dispatch(setTable(table));
    thunkAPI.dispatch(setLocation("table"));
    return table;
  }
);
const slice = createSlice({
  name: "room",
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(fetchTables.fulfilled, (state, { payload }) => ({
        tables: payload,
        loaded: true,
        loading: false,
      }))
      .addCase(fetchTables.pending, (state) => {
        state.loading = true;
        state.error = false;
      })
      .addCase(fetchTables.rejected, (state, action) => {
        state.error = action.error;
        state.loading = false;
        state.loaded = false;
        state.tables = [];
      })
      .addCase(createTable.pending, (state) => {
        state.loading = true;
        state.error = false;
      })
      .addCase(createTable.fulfilled, (state, { payload: table }) => {
        state.tables.push(table);
        state.loading = false;
      })
      .addCase(createTable.rejected, (state, action) => {
        state.error = action.error;
        state.loading = false;
      });
  },
});
export const selectRoomState = (state: RootState) => state.room;
export default slice.reducer;
