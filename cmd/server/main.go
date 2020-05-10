package main

//go:generate protoc -I ../../grpc_api ../../grpc_api/karlchen.proto --go_out=plugins=grpc:../../grpc_api

import (
	"context"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	uuid "github.com/satori/go.uuid"
	grpc_api "github.com/supermihi/karlchencloud/grpc_api"
	"github.com/supermihi/karlchencloud/pkg/cloud"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

type grpcserver struct {
	grpc_api.UnimplementedKarlchencloudServer
	s    cloud.Cloud
	auth cloud.Auth
}

func (s *grpcserver) Register(ctx context.Context, in *grpc_api.RegisterRequest) (*grpc_api.RegisterReply, error) {
	log.Printf("Received: %v", in.GetName())
	id := uuid.NewV4()
	s.s.Users.Add(cloud.UserId(id.String()), in.GetName(), cloud.RandomSecret())
	return &grpc_api.RegisterReply{Id: id.String(), Secret: "oh my"}, nil
}

func (s *grpcserver) CheckLogin(ctx context.Context, request *grpc_api.EmptyRequest) (*grpc_api.Bool, error) {
	user, ok := cloud.GetAuthenticatedUser(ctx)
	if !ok {
		log.Print("check login failed")
		return &grpc_api.Bool{Value: false}, nil
	}
	log.Printf("user %v ok", user)
	return &grpc_api.Bool{Value: true}, nil

}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	users := cloud.NewMemoryUserDb()
	users.Add(cloud.UserId("6ba7b810-9dad-11d1-80b4-00c04fd430c8"), "Woldemar", "geheim")
	theCloud := cloud.NewCloud(users)
	auth := cloud.NewAuth(users)
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(auth.Authenticate)),
		grpc.StreamInterceptor(grpc_auth.StreamServerInterceptor(auth.Authenticate)))

	serv := grpcserver{s: theCloud, auth: auth}
	grpc_api.RegisterKarlchencloudServer(grpcServer, &serv)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
