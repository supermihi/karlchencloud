import * as karlchen from './KarlchenServiceClientPb';
import * as proto from './karlchen_pb';
import { Error, Metadata } from 'grpc-web';
import { Base64 } from 'js-base64';
import { MyUserData } from 'app/auth';

const url = 'http://localhost:8080';

let _client: karlchen.DokoClient | null = null;

export interface AuthenticatedClient {
  client: karlchen.DokoClient;
  meta: Metadata;
}
export function getClient() {
  if (!_client) {
    _client = new karlchen.DokoClient(url, null, null);
  }
  return _client;
}

export function getAuthMeta(user: string, secret: string): Metadata {
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

export function isGrpcError(error: any): error is Error {
  return error.code && error.message;
}

export function formatError(error: any) {
  if (!error) {
    return '';
  }
  if (isGrpcError(error)) {
    return `error ${error.code}: ${error.message}`;
  }
  if (error.message) {
    return error.message;
  }
  return `${error}`;
}
