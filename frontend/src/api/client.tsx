import * as karlchen from './KarlchenServiceClientPb';
import * as proto from './karlchen_pb';
import * as grpc from 'grpc-web';
import { MyUserData } from 'session/model';

const url = 'http://localhost:8080';

let _client: karlchen.DokoClient | null = null;

export interface AuthenticatedClient {
  client: karlchen.DokoClient;
  meta: grpc.Metadata;
}
export function getPbClient(): karlchen.DokoClient {
  if (!_client) {
    _client = new karlchen.DokoClient(url, null, null);
  }
  return _client;
}

export function getAuthMeta(token: string): grpc.Metadata {
  return { authorization: `basic ${token}` };
}

export function getAuthenticatedClient(token: string): AuthenticatedClient {
  return { client: getPbClient(), meta: getAuthMeta(token) };
}

export async function register(email: string, name: string, password: string): Promise<MyUserData> {
  const request = new proto.RegisterRequest();
  request.setName(name);
  request.setEmail(email);
  request.setPassword(password);
  const ans = await getPbClient().register(request, null);
  const [id, token] = [ans.getUserId(), ans.getToken()];
  return { name, id, token, email };
}

export async function login(email: string, password: string): Promise<MyUserData> {
  const request = new proto.LoginRequest();
  request.setEmail(email);
  request.setPassword(password);
  const ans = await getPbClient().login(request, null);
  const [token, id, name] = [ans.getToken(), ans.getUserId(), ans.getName()];
  return { name, id, token, email };
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
