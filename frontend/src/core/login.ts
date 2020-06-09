import {
  createSlice,
  createAsyncThunk,
  createSelector,
} from "@reduxjs/toolkit";
import { AsyncThunkConfig, RootState } from "app/store";
import * as client from "api/client";
import { User } from "model/core";
import { getAuthenticatedClient } from "api/client";
import { Empty } from "../api/karlchen_pb";

interface LoginState {
  loggedIn: boolean;
  me: User;
  secret: string;
  loading: boolean;
  error?: any;
}

function getLoginDataFromLocalStorage(): LoginData | null {
  const [id, secret, name] = ["id", "secret", "name"].map((key) =>
    window.localStorage.getItem(key)
  );
  if (id !== null && secret !== null && name !== null) {
    return { id, secret, name };
  }
  return null;
}

function writeLoginDataToLocalStorage({ id, secret, name }: LoginData) {
  window.localStorage.setItem("id", id);
  window.localStorage.setItem("name", name);
  window.localStorage.setItem("secret", secret);
}

const loginFromLocalStorage = getLoginDataFromLocalStorage();
const initialState: LoginState = {
  loggedIn: false,
  loading: false,
  me: {
    id: loginFromLocalStorage?.id ?? "",
    name: loginFromLocalStorage?.name ?? "not logged in",
  },
  secret: loginFromLocalStorage?.secret ?? "",
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
    const ans = { name, id, secret };
    writeLoginDataToLocalStorage(ans);
    return ans;
  }
);

export const login = createAsyncThunk<string, void, AsyncThunkConfig>(
  "login/login",
  async (_, thunkAPI) => {
    const { client, meta } = selectClient(thunkAPI.getState());
    const user = await client.checkLogin(new Empty(), meta);
    return user.getName();
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
      })
      .addCase(login.fulfilled, (state, action) => {
        state.me.name = action.payload;
        state.loggedIn = true;
        state.loading = false;
      })
      .addCase(login.pending, (state) => {
        state.loading = true;
      })
      .addCase(login.rejected, (state, action) => {
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
