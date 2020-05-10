package cloud

import (
	"context"
	grpcauth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	uuid "github.com/satori/go.uuid"
	"github.com/supermihi/karlchencloud/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

type grpcserver struct {
	api.UnimplementedKarlchencloudServer
	room Room
	auth Auth
}

func (s *grpcserver) Register(ctx context.Context, in *api.RegisterRequest) (*api.RegisterReply, error) {
	log.Printf("Received: %v", in.GetName())
	id := uuid.NewV4()
	s.room.Users.Add(UserId(id.String()), in.GetName(), RandomSecret())
	return &api.RegisterReply{Id: id.String(), Secret: "oh my"}, nil
}

func (s *grpcserver) CheckLogin(ctx context.Context, request *api.EmptyRequest) (*api.OkOrNot, error) {
	user, ok := GetAuthenticatedUser(ctx)
	if !ok {
		log.Print("check login failed")
		return &api.OkOrNot{Value: false}, nil
	}
	log.Printf("user %v ok", user)
	return &api.OkOrNot{Value: true}, nil
}

func (s *grpcserver) CreateTable(ctx context.Context, params *api.EmptyRequest) (*api.TableId, error) {
	user, ok := GetAuthenticatedUser(ctx)
	if !ok {
		panic("should not be here")
	}
	id := s.room.CreateTable(user)
	log.Printf("user %v created new table with id %v", s.room.Users.GetName(user), id)
	return &api.TableId{Value: id}, nil
}

func (s *grpcserver) ListTables(context.Context, *api.EmptyRequest) (*api.TableList, error) {
	tables := s.room.tables.List()
	return &api.TableList{Ids: tables}, nil
}

func (s *grpcserver) StartTable(ctx context.Context, id *api.TableId) (*api.EmptyReply, error) {
	user, _ := GetAuthenticatedUser(ctx)
	table, ok := s.room.tables.GetTable(id.Value)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "table does not exist")
	}
	if table.Owner() != user {
		return nil, status.Error(codes.PermissionDenied, "you are not owner of the table")
	}
	err := table.Start()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &api.EmptyReply{}, nil

}
func StartServer(users Users, port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	room := NewRoom(users)
	auth := NewAuth(users)
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpcauth.UnaryServerInterceptor(auth.Authenticate)),
		grpc.StreamInterceptor(grpcauth.StreamServerInterceptor(auth.Authenticate)))

	serv := grpcserver{room: room, auth: auth}
	api.RegisterKarlchencloudServer(grpcServer, &serv)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
