import { MyUserData, Credentials } from 'app/auth';
import { createSlice, createSelector, PayloadAction } from '@reduxjs/toolkit';
import { getClient, getAuthMeta, getAuthenticatedClient } from 'api/client';
import * as api from 'api/karlchen_pb';
import { AppThunk, RootState } from 'app/store';

import { ClientReadableStream, Error as GrpcError } from 'grpc-web';
import { onEvent } from './eventhandlers';
import { register } from 'app/auth/thunks';

export interface SessionState {
  session: MyUserData | null;
  starting: Credentials | null;
  error?: any;
}

const initialState: SessionState = { session: null, starting: null };

let _stream: ClientReadableStream<api.Event>;

export const startSession = (creds: Credentials): AppThunk => async (dispatch) => {
  dispatch(actions.sessionStarting(creds));
  const client = getClient();
  const { id, secret } = creds;
  const authMeta = getAuthMeta(id, secret);
  try {
    _stream = client.startSession(new api.Empty(), authMeta) as ClientReadableStream<api.Event>;
    _stream
      .on('data', (e) => dispatch(onEvent(e)))
      .on('error', (e) => {
        dispatch(actions.sessionError(e));
        console.log(`Stream error: ${e.message}`);
      });
  } catch (error) {
    console.log(`Session error: ${error}`);
    dispatch(actions.sessionError(error));
  }
};

export const endSession = (): AppThunk => (dispatch) => {
  if (_stream) {
    _stream.cancel();
    dispatch(actions.sessionEnded());
  }
};

const slice = createSlice({
  name: 'session',
  initialState,
  reducers: {
    resetError: (state) => {
      state.error = undefined;
    },
    sessionStarting: (_, { payload: creds }: PayloadAction<Credentials>) => ({
      ...initialState,
      starting: creds,
    }),
    sessionStarted: (state, { payload: name }: PayloadAction<string>) => {
      if (!state.starting) {
        return;
      }
      state.session = { ...state.starting, name };
      state.starting = null;
    },
    sessionEnded: () => initialState,
    sessionError: (state, { payload: error }: PayloadAction<GrpcError>) => {
      state.starting = null;
      state.session = null;
      state.error = error;
    },
  },
  extraReducers: (builder) => {
    builder.addCase(register.fulfilled, (state) => {
      state.error = undefined;
    });
  },
});
export const selectSession = (state: RootState) => state.session;
export const selectClient = createSelector(selectSession, ({ session }) =>
  session !== null ? getAuthenticatedClient(session.id, session.secret) : null
);
export const selectAuthenticatedClientOrThrow = createSelector(selectClient, (client) => {
  if (!client) {
    throw new Error('not authenticated');
  }
  return client;
});
export const actions = slice.actions;
export default slice.reducer;
