import * as grpc from "grpc-web";
export function mockGrpcError(msg: string): grpc.Error {
  return {
    code: 42,
    message: msg,
  };
}
