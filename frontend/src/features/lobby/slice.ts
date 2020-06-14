import { Table, toTable, toUser } from "model/table";
import { createSlice, createAsyncThunk } from "@reduxjs/toolkit";
import { selectAuthenticatedClientOrThrow } from "features/auth/slice";
import { RootState, AsyncThunkConfig } from "app/store";
import { Empty } from "api/karlchen_pb";

export interface LobbyState {
  activeTable: Table | null;
  loading: boolean;
  loaded: boolean;
  error?: any;
}

const initialState: LobbyState = {
  activeTable: null,
  loading: false,
  loaded: false,
};

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

export const fetchUserState = createAsyncThunk<
  Table | null,
  void,
  AsyncThunkConfig
>("lobby/fetchUserState", async (_, thunkAPI) => {
  const { client, meta } = selectAuthenticatedClientOrThrow(
    thunkAPI.getState()
  );
  const result = await client.getUserState(new Empty(), meta);
  const table = result.getCurrenttable()?.getData();
  if (!table) {
    return null;
  }

  return {
    id: table.getTableId(),
    owner: table.getOwner(),
    players: table.getMembersList().map(toUser),
  };
});
const slice = createSlice({
  name: "room",
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(fetchUserState.fulfilled, (state, { payload }) => {
        state.loaded = true;
        state.loading = false;
        state.activeTable = payload;
      })
      .addCase(fetchUserState.pending, () => ({
        ...initialState,
        loading: true,
      }))
      .addCase(fetchUserState.rejected, (state, { error }) => ({
        ...initialState,
        error,
      }))
      .addCase(createTable.pending, (state) => {
        state.loading = true;
        state.error = false;
      })
      .addCase(createTable.fulfilled, (state, { payload: table }) => {
        state.activeTable = table;
        state.loading = false;
      })
      .addCase(createTable.rejected, (state, action) => {
        state.error = action.error;
        state.loading = false;
      });
  },
});
export const selectRoomState = (state: RootState) => state.lobby;
export default slice.reducer;
