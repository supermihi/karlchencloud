import { createAsyncThunk, createSlice, PayloadAction } from "@reduxjs/toolkit";
import { Table } from "model/table";
import { AsyncThunkConfig } from "app/store";
import {
  selectClient,
  selectAuthenticatedClientOrThrow,
} from "features/auth/slice";
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
