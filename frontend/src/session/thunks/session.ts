import { getPbClient, getAuthMeta } from 'api/client';
import { ClientReadableStream } from 'grpc-web';
import { AppThunk } from 'state';
import * as api from 'api/karlchen_pb';
import { createEventAction } from './streamEvents';
import { selectSession } from 'session/selectors';
import { Action, createAction } from '@reduxjs/toolkit';
import { MyUserData } from '../model';
import { GrpcError, toGrpcError } from 'shared/errors';

let _stream: ClientReadableStream<api.Event>;
export const sessionStarting = createAction<MyUserData>('session/starting');
export const sessionError = createAction<GrpcError>('session/error');

export const startSession =
  (givenUserData?: MyUserData): AppThunk =>
  async (dispatch, getState) => {
    const userData = givenUserData ?? selectSession(getState()).userData;
    if (userData === null) {
      return;
    }
    const pbClient = getPbClient();

    const authMeta = getAuthMeta(userData.token);
    dispatch(sessionStarting(userData));
    try {
      _stream = pbClient.startSession(new api.Empty(), authMeta) as ClientReadableStream<api.Event>;
      _stream
        .on('data', (e) => {
          const dispatchable = createEventAction(e);
          dispatch(dispatchable as Action);
        })
        .on('error', (e) => {
          dispatch(sessionError(toGrpcError(e)));
          _stream.cancel();
          console.log(`Stream error: ${e.message}`);
        });
    } catch (error) {
      console.log(`Session error: ${error}`);
      dispatch(sessionError(error));
    }
  };
