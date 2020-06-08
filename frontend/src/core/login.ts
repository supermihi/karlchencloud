import {
  createSlice,
  createAsyncThunk,
  createSelector,
} from "@reduxjs/toolkit";
import { RootState } from "app/store";
import * as client from "api/client";
import { User } from "model/core";
import { getAuthenticatedClient } from "api/client";

interface LoginState {
  loggedIn: boolean;
  me: User;
  secret: string;
  loading: boolean;
  error?: any;
}

const initialState: LoginState = {
  loggedIn: false,
  loading: false,
  me: { id: "", name: "not logged in" },
  secret: "",
};

interface LoginData {
  name: string;
  id: string;
  secret: string;
}

export const register = createAsyncThunk<LoginData, string>(
  "login/register",
  async (name) => {
    const { id, secret } = await client.register(name);
    return { name, id, secret };
  }
);

export const loginSlice = createSlice({
  name: "login",
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(
        register.fulfilled,
        (state, { payload: { name, id, secret } }) => ({
          me: { name, id },
          secret,
          loggedIn: true,
          loading: false,
        })
      )
      .addCase(register.pending, (state) => {
        state.loading = true;
      })
      .addCase(register.rejected, (state, action) => {
        state.loading = false;
        state.error = action.error;
      });
  },
});

export const selectLogin = (state: RootState) => state.login;
export const selectClient = createSelector(selectLogin, (l) =>
  getAuthenticatedClient(l.me.id, l.secret)
);
export const selectMe = createSelector(selectLogin, (l) => l.me);
export default loginSlice.reducer;
