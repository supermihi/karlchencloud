import { createSlice, createAsyncThunk } from "@reduxjs/toolkit";
import { RootState } from "app/store";
import * as karlchen from "api/KarlchenServiceClientPb";
import * as client from "api/client";

interface LoginState {
  loggedIn: boolean;
  name?: string;
  userId?: string;
  secret?: string;
  loading: boolean;
  client: null | karlchen.DokoClient;
  error?: any;
}

const initialState: LoginState = {
  loggedIn: false,
  loading: false,
  client: null,
};

interface LoginData {
  name: string;
  id: string;
  secret: string;
}
export const register = createAsyncThunk(
  "login/register",
  async (name: string, thunkAPI) => {
    const { id, secret } = await client.register(name);
    return { name, id, secret } as LoginData;
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
          name,
          id,
          secret,
          loggedIn: true,
          loading: false,
          client: client.getAuthenticatedClient(id, secret),
        })
      )
      .addCase(register.pending, (state, action) => {
        state.loading = true;
        state.name = action.meta.arg;
      })
      .addCase(register.rejected, (state, action) => {
        state.loading = false;
        state.error = action.error;
      });
  },
});

export const selectLogin = (state: RootState) => state.login;
export default loginSlice.reducer;
