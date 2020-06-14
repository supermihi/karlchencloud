import { createAsyncThunk, createSlice, PayloadAction } from "@reduxjs/toolkit";
import { Table } from "model/table";
import { AsyncThunkConfig } from "app/store";
import { selectClient } from "features/auth/slice";
import { TableId } from "api/karlchen_pb";
import { tableId } from "../../api/helpers";

interface State {
  table: Table | null;
  loading: boolean;
  fetched: boolean;
  details: TableDetails | null;
}
interface TableDetails {
  users: string[];
}
const fetchTableDetails = createAsyncThunk<
  TableDetails,
  string,
  AsyncThunkConfig
>("table/fetchDetails", async (id, thunkAPI) => {
  const state = thunkAPI.getState();
  const { client, meta } = selectClient(state);
  const tableState = await client.getTableState(tableId(id), meta);
  const members = tableState.getMembersList();
  return { users: members.map((m) => m.getName()) };
});
const initialState: State = {
  table: null,
  fetched: false,
  loading: false,
  details: null,
};
export const tableSlice = createSlice({
  initialState,
  name: "table",
  reducers: {
    setTable: (state, { payload: table }: PayloadAction<Table>) => {
      state.table = table;
    },
  },
});

export const { setTable } = tableSlice.actions;
export default tableSlice.reducer;
