import * as karlchen from "./KarlchenServiceClientPb";
import * as proto from "./karlchen_pb";
import { Error } from "grpc-web";
const url = "http://localhost:8080";
export function getAuthenticatedClient(user: string, secret: string) {
  let auth = "test";
  return new karlchen.DokoClient(url, { authorization: auth }, null);
}

export async function register(name: string) {
  const client = new karlchen.DokoClient(url, null, null);
  const userName = new proto.UserName();
  userName.setName(name);
  const ans = await client.register(userName, null);
  const [id, secret] = [ans.getId(), ans.getSecret()];
  return { id, secret };
}

export function isGrpcError(error: any): error is Error {
  return error.code && error.message;
}

export function formatError(error: any) {
  if (!error) {
    return "";
  }
  if (isGrpcError(error)) {
    return `error ${error.code}: ${error.message}`;
  }
  if (error.message) {
    return error.message;
  }
  return `${error}`;
}
