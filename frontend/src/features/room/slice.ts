import { Table } from "model/table";
import { createSlice, createAsyncThunk } from "@reduxjs/toolkit";
import { selectClient } from "core/login";
import { RootState, AsyncThunkConfig } from "app/store";
import { Empty } from "api/karlchen_pb";

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
    const tables = await client.listTables(new Empty(), meta);
    return tables
      .getTablesList()
      .map((t) => ({ owner: t.getOwner(), id: t.getTableId() }));
  }
);
export const createTable = createAsyncThunk<Table, void, AsyncThunkConfig>(
  "room/createTable",
  async (_, thunkAPI) => {
    throw new Error("not implemented");
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
      })
      .addCase(fetchTables.rejected, (state, action) => {
        state.error = action.error;
        state.loading = false;
        state.loaded = false;
        state.tables = [];
      });
  },
});
export const selectRoomState = (state: RootState) => state.room;
export default slice.reducer;
