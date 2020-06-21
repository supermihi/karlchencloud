import { LoginData } from "app/auth";
import {
  createAsyncThunk,
  createSlice,
  createSelector,
} from "@reduxjs/toolkit";
import { TableState } from "model/table";
import { getClient, getAuthMeta, getAuthenticatedClient } from "api/client";
import * as api from "api/karlchen_pb";
import { toTableState } from "model/apiconv";
import { RootState } from "app/store";
import * as lobby from "./lobby";

export interface SessionState {
  session: LoginData | null;
  currentTable: TableState | null;
  starting: boolean;
  error?: any;
}

const initialState: SessionState = {
  session: null,
  currentTable: null,
  starting: false,
};
interface SessionStart {
  login: LoginData;
  currentTable: TableState | null;
}
export const startSession = createAsyncThunk<
  SessionStart,
  { id: string; secret: string }
>("session/start", async ({ id, secret }, { dispatch }) => {
  const client = getClient();
  const authMeta = getAuthMeta(id, secret);
  const userStateReceived = new Promise<api.UserState>((resolve, reject) => {
    try {
      const stream = client.startSession(new api.Empty(), authMeta);
      stream.on("data", (e) =>
        e.hasWelcome() ? resolve(e.getWelcome()) : null
      );
    } catch (error) {
      reject(error);
    }
  });
  const userState = await userStateReceived;
  const name = userState.getName();
  const currentTable = userState.hasCurrenttable()
    ? toTableState(userState.getCurrenttable() as api.TableState)
    : null;
  return { login: { name, id, secret }, currentTable };
});

const slice = createSlice({
  name: "session",
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(startSession.pending, (state) => {
        state.starting = true;
        state.error = null;
      })
      .addCase(startSession.rejected, (state, { error }) => {
        state.starting = false;
        state.error = error;
      })
      .addCase(
        startSession.fulfilled,
        (state, { payload: { login, currentTable } }) => {
          state.starting = false;
          state.session = login;
          state.currentTable = currentTable;
        }
      );
    lobby.addReducerCases(builder);
  },
});
export const selectSession = (state: RootState) => state.session;
export const selectClient = createSelector(selectSession, ({ session }) =>
  session !== null ? getAuthenticatedClient(session.id, session.secret) : null
);
export const selectAuthenticatedClientOrThrow = createSelector(
  selectClient,
  (client) => {
    if (!client) {
      throw new Error("not authenticated");
    }
    return client;
  }
);
export default slice.reducer;
