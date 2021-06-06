import { CaseReducer, createSlice, PayloadAction, SerializedError } from '@reduxjs/toolkit';
import { SessionPhase } from './model';
import { initialState, SessionState } from './state';
import * as events from './events';
import { forgetLogin, login, register } from './thunks/authenticate';
import { sessionError, sessionStarting } from './thunks/session';

const reduceLoginOrRegisterError: CaseReducer<
  SessionState,
  PayloadAction<unknown, string, never, SerializedError>
> = (state, action) => {
  state.error = action.error;
  state.phase = SessionPhase.NoToken;
  state.userData = null;
};

const sessionSlice = createSlice({
  name: 'session',
  initialState: initialState(),
  reducers: {
    resetError: (state) => {
      state.error = undefined;
    },
  },
  extraReducers: (builder) => {
    builder
      .addCase(sessionStarting, (state, { payload: userData }) => {
        state.userData = userData;
        state.phase = SessionPhase.Starting;
      })
      .addCase(sessionError, (state, { payload: error }) => {
        state.phase = SessionPhase.TokenObtained;
        state.error = error;
      })
      .addCase(login.fulfilled, (state, { payload }) => {
        state.phase = SessionPhase.TokenObtained;
        state.userData = payload;
      })
      .addCase(register.fulfilled, (state, { payload }) => {
        state.phase = SessionPhase.TokenObtained;
        state.userData = payload;
      })
      .addCase(login.pending, (state) => {
        state.phase = SessionPhase.ObtainingToken;
        state.error = null;
      })
      .addCase(register.pending, (state) => {
        state.phase = SessionPhase.ObtainingToken;
        state.error = null;
      })
      .addCase(register.rejected, reduceLoginOrRegisterError)
      .addCase(login.rejected, reduceLoginOrRegisterError)
      .addCase(events.sessionStarted, (state) => {
        state.phase = SessionPhase.Established;
      })
      .addCase(forgetLogin, (state) => {
        state.userData = null;
        state.phase = SessionPhase.NoToken;
      });
  },
});

export const { actions, reducer } = sessionSlice;
