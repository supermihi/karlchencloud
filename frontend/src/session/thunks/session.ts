import { getPbClient, getAuthMeta } from 'api/client';
import { ClientReadableStream } from 'grpc-web';
import { MyUserData } from 'session/model';
import { AppThunk } from 'state';
import * as api from 'api/karlchen_pb';
import { createEventAction } from './streamEvents';
import { actions } from 'session/slice';
import { selectSession } from 'session/selectors';
import { Action } from '@reduxjs/toolkit';

let _stream: ClientReadableStream<api.Event>;

export const startSession = (): AppThunk => async (dispatch, getState) => {
  const pbClient = getPbClient();
  const userData = selectSession(getState()).storedLogin as MyUserData;
  const authMeta = getAuthMeta(userData.token);
  dispatch(actions.sessionStarting(userData));
  try {
    _stream = pbClient.startSession(new api.Empty(), authMeta) as ClientReadableStream<api.Event>;
    _stream
      .on('data', (e) => {
        const dispatchable = createEventAction(e);
        dispatch(dispatchable as Action);
      })
      .on('error', (e) => {
        dispatch(actions.sessionError(e));
        _stream.cancel();
        console.log(`Stream error: ${e.message}`);
      });
  } catch (error) {
    console.log(`Session error: ${error}`);
    dispatch(actions.sessionError(error));
  }
};
