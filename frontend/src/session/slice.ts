import { CaseReducer, createSlice, PayloadAction } from '@reduxjs/toolkit';
import { Credentials } from './model';
import { initialState, SessionState } from './state';
import * as grpc from 'grpc-web';
import * as events from './events';
import { localStorageUpdated, register } from './thunks/register';

const reduceSessionStarted: CaseReducer<SessionState, PayloadAction<string>> = (
  state,
  { payload: name }
) => {
  if (!state.currentLoginCredentials) {
    return;
  }
  state.session = { ...state.currentLoginCredentials, name };
  state.currentLoginCredentials = null;
};

const sessionSlice = createSlice({
  name: 'session',
  initialState: initialState(),
  reducers: {
    sessionStarting: (_, { payload: creds }: PayloadAction<Credentials>) => ({
      ...initialState(),
      currentLoginCredentials: creds,
    }),
    sessionError: (state, { payload: error }: PayloadAction<grpc.Error>) => {
      state.loading = false;
      state.session = null;
      state.error = error;
    },
    resetError: (state) => {
      state.error = undefined;
    },
  },
  extraReducers: (builder) => {
    builder
      .addCase(localStorageUpdated, (state, { payload }) => {
        state.storedLogin = payload;
      })
      .addCase(register.fulfilled, (state) => {
        state.loading = false;
      })
      .addCase(register.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(register.rejected, (state, action) => {
        state.loading = false;
        state.error = action.error;
      })
      .addCase(events.sessionStarted, reduceSessionStarted);
  },
});

export const { actions, reducer } = sessionSlice;
