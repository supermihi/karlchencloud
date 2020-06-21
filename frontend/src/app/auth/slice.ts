import { createAsyncThunk, createSlice, PayloadAction } from "@reduxjs/toolkit";
import { AppThunk, RootState } from "app/store";
import * as api from "api/client";
import {
  getLoginDataFromLocalStorage,
  writeLoginDataToLocalStorage,
  deleteLoginDataInLocalStorage,
} from "./localstorage";
import { LoginData } from ".";

export interface SessionState {
  storedLogin: LoginData | null;
  registering: boolean;
  registerError?: any;
}

const initialState = (): SessionState => {
  const existingLogin = getLoginDataFromLocalStorage();
  return {
    registering: false,
    storedLogin: existingLogin,
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

export const forgetLogin = (): AppThunk => (dispatch) => {
  deleteLoginDataInLocalStorage();
  dispatch(localStorageUpdated(null));
};
export const authSlice = createSlice({
  name: "auth",
  initialState: initialState(),
  reducers: {
    localStorageUpdated: (
      state,
      { payload }: PayloadAction<LoginData | null>
    ) => {
      state.storedLogin = payload;
    },
  },
  extraReducers: (builder) => {
    builder
      .addCase(register.fulfilled, (state, { payload: login }) => {
        state.registering = false;
      })
      .addCase(register.pending, (state) => {
        state.registering = true;
        state.registerError = null;
      })
      .addCase(register.rejected, (state, action) => {
        state.registering = false;
        state.registerError = action.error;
      });
  },
});
const { localStorageUpdated } = authSlice.actions;
export const selectAuth = (state: RootState) => state.auth;
export default authSlice.reducer;
