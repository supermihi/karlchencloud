import { CaseReducer, createSlice, PayloadAction } from '@reduxjs/toolkit';
import { MyUserData } from './model';
import { initialState, SessionState } from './state';
import * as grpc from 'grpc-web';
import * as events from './events';
import { localStorageUpdated, login, register } from './thunks/authenticate';

const reduceSessionStarted: CaseReducer<SessionState, PayloadAction> = (
  state,
) => {
  if (!state.startingSession) {
    return;
  }
  state.activeSession = state.startingSession;
  state.startingSession = null;
};

const sessionSlice = createSlice({
  name: 'session',
  initialState: initialState(),
  reducers: {
    sessionStarting: (_, { payload: userData }: PayloadAction<MyUserData>) => ({
      ...initialState(),
      startingSession: userData,
    }),
    sessionError: (state, { payload: error }: PayloadAction<grpc.Error>) => {
      state.loading = false;
      state.activeSession = null;
      state.startingSession = null;
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
      .addCase(login.fulfilled, (state) => {
        state.loading = false;
      })
      .addCase(register.fulfilled, (state) => {
        state.loading = false;
      })
      .addCase(login.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(register.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(register.rejected, (state, action) => {
        state.loading = false;
        state.error = action.error;
      })
      .addCase(login.rejected, (state, action) => {
        state.loading = false;
        state.error = action.error;
      })
      .addCase(events.sessionStarted, reduceSessionStarted);
  },
});

export const { actions, reducer } = sessionSlice;
