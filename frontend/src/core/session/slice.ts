import {
  createAsyncThunk,
  createSelector,
  createSlice,
  PayloadAction,
} from "@reduxjs/toolkit";
import { AppThunk, RootState } from "app/store";
import * as api from "api/client";
import { Empty } from "api/karlchen_pb";
import {
  LoginData,
  getLoginDataFromLocalStorage,
  writeLoginDataToLocalStorage,
  deleteLoginDataInLocalStorage,
} from "./localstorage";

export interface SessionState {
  storedLogin: LoginData | null;
  validLogin: LoginData | null;
  loading: boolean;
  error?: any;
}

const initialState = (): SessionState => {
  const existingLogin = getLoginDataFromLocalStorage();
  return {
    loading: false,
    storedLogin: existingLogin,
    validLogin: null,
  };
};

export const register = createAsyncThunk<LoginData, string>(
  "model/register",
  async (name, { dispatch }) => {
    const { id, secret } = await api.register(name);
    const ans = { name, id, secret };
    writeLoginDataToLocalStorage(ans);
    dispatch(localStorageUpdated(ans));
    return ans;
  }
);

const startSession = ({id, secret}: LoginData): AppThunk => (dispatch) =>  {
  try {
    const server = api.getClient().startSession(new Empty(), api.getAuthMeta(id, secret));
  } catch (error) {
    if (api.isGrpcError)
  }

}
export const tryLogin = createAsyncThunk<LoginData, LoginData>(
  "model/model",
  async (login, { dispatch }) => {
    const client = api.getClient();
    const meta = api.getAuthMeta(login.id, login.secret);
    const user = await client.checkLogin(new Empty(), meta);
    const newLogin = { ...login, name: user.getName() }; // name might have changed!
    writeLoginDataToLocalStorage(newLogin);
    dispatch(localStorageUpdated(newLogin));
    return newLogin;
  }
);

export const forgetLogin = (): AppThunk => (dispatch) => {
  deleteLoginDataInLocalStorage();
  dispatch(localStorageUpdated(null));
};
export const loginSlice = createSlice({
  name: "session",
  initialState: initialState(),
  reducers: {
    localStorageUpdated: (state, { payload }: PayloadAction<LoginData | null>) => {
      state.storedLogin = payload;
    },
  },
  extraReducers: (builder) => {
    builder
      .addCase(register.fulfilled, (state, { payload: login }) => {
        state.validLogin = login;
        state.loading = false;
      })
      .addCase(register.pending, (state) => {
        state.loading = true;
        state.error = null;
        state.validLogin = null;
      })
      .addCase(register.rejected, (state, action) => {
        state.loading = false;
        state.error = action.error;
      })
      .addCase(tryLogin.fulfilled, (state, { payload: login }) => {
        state.validLogin = login;
        state.loading = false;
      })
      .addCase(tryLogin.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(tryLogin.rejected, (state, action) => {
        state.loading = false;
        state.error = action.error;
      });
  },
});
const { localStorageUpdated } = loginSlice.actions;

export const selectAuth = (state: RootState) => state.core.session;
export const selectClient = createSelector(selectAuth, ({ validLogin }) =>
  validLogin !== null
    ? api.getAuthenticatedClient(validLogin.id, validLogin.secret)
    : null
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
export default loginSlice.reducer;
