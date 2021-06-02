package server

import (
	"github.com/supermihi/karlchencloud/room"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func toGrpcError(err error) error {
	if _, ok := status.FromError(err); ok {
		return err // already a GRPC error
	}
	if cloudErr, ok := err.(room.CloudError); ok {
		return status.Error(codes.Internal, cloudErr.Error())
	}
	return status.Error(codes.Unknown, err.Error())
}
