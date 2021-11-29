import { RpcError, StatusCode } from 'grpc-web';

export interface GrpcError {
  status: StatusCode;
  message: string;
}

export function toGrpcError(error: RpcError): GrpcError {
  return { status: error.code, message: error.message };
}

export function isGrpcError(error: unknown): error is GrpcError {
  return typeof error === 'object' && error !== null && 'message' in error && 'status' in error;
}

export function isRpcError(error: unknown): error is RpcError {
  return typeof error === 'object' && error !== null && 'message' in error && 'code' in error;
}
