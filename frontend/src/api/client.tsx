import * as karlchen from './KarlchenServiceClientPb';
import * as proto from './karlchen_pb';
import * as grpc from 'grpc-web';
import { Base64 } from 'js-base64';
import { MyUserData } from 'session/model';

const url = 'http://localhost:8080';

let _client: karlchen.DokoClient | null = null;

export interface AuthenticatedClient {
  client: karlchen.DokoClient;
  meta: grpc.Metadata;
}
export function getClient(): karlchen.DokoClient {
  if (!_client) {
    _client = new karlchen.DokoClient(url, null, null);
  }
  return _client;
}

export function getAuthMeta(user: string, secret: string): grpc.Metadata {
  const encoded = Base64.encode(`${user}:${secret}`);
  return { authorization: `basic ${encoded}` };
}

export function getAuthenticatedClient(user: string, secret: string): AuthenticatedClient {
  return { client: getClient(), meta: getAuthMeta(user, secret) };
}

export async function register(name: string): Promise<MyUserData> {
  const userName = new proto.UserName();
  userName.setName(name);
  const ans = await getClient().register(userName, null);
  const [id, secret] = [ans.getId(), ans.getSecret()];
  return { name, id, secret };
}

export function isGrpcError(error: unknown): error is grpc.Error {
  return typeof error === 'object' && error !== null && 'code' in error && 'message' in error;
}

export function isError(error: unknown): error is { message: string } {
  return typeof error === 'object' && error !== null && 'message' in error;
}

export function formatError(error: unknown): string {
  if (!error) {
    return '';
  }
  if (isGrpcError(error)) {
    return `error ${error.code}: ${error.message}`;
  }
  if (isError(error)) {
    return error.message;
  }
  return `${error}`;
}
