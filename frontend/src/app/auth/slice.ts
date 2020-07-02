import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { RootState } from 'app/store';
import { getLoginDataFromLocalStorage } from './localstorage';
import { MyUserData } from '.';
import { register } from './thunks';

export interface AuthState {
  storedLogin: MyUserData | null;
  registering: boolean;
  registerError?: any;
}

const initialState = (): AuthState => {
  const existingLogin = getLoginDataFromLocalStorage();
  return {
    registering: false,
    storedLogin: existingLogin,
  };
};

export const authSlice = createSlice({
  name: 'auth',
  initialState: initialState(),
  reducers: {
    localStorageUpdated: (state, { payload }: PayloadAction<MyUserData | null>) => {
      state.storedLogin = payload;
    },
  },
  extraReducers: (builder) => {
    builder
      .addCase(register.fulfilled, (state) => {
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
export const actions = authSlice.actions;
export const selectAuth = (state: RootState) => state.auth;
export default authSlice.reducer;
