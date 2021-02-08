import { getClient, getAuthMeta } from 'api/client';
import { ClientReadableStream } from 'grpc-web';
import { Credentials } from 'session/model';
import { AppThunk } from 'state';
import * as api from 'api/karlchen_pb';
import { createEventAction } from './streamEvents';
import { actions } from 'session/slice';

let _stream: ClientReadableStream<api.Event>;

export const startSession = (creds: Credentials): AppThunk => async (dispatch) => {
  const client = getClient();
  const { id, secret } = creds;
  const authMeta = getAuthMeta(id, secret);
  try {
    _stream = client.startSession(new api.Empty(), authMeta) as ClientReadableStream<api.Event>;
    _stream
      .on('data', (e) => {
        const dispatchable = createEventAction(e);
        dispatch(dispatchable as any);
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
