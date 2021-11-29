import { StatusCode } from 'grpc-web';
import { GrpcError } from 'shared/errors';

export function mockGrpcError(msg: string): GrpcError {
  return { status: StatusCode.UNIMPLEMENTED, message: msg };
}
