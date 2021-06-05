package server

import (
	"context"
	grpcAuth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/supermihi/karlchencloud/server/users"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

type Auth struct {
	Users users.Users
}

func NewAuth(users users.Users) Auth {
	return Auth{users}
}

type userMDKey struct{}

func (a *Auth) Authenticate(ctx context.Context) (newCtx context.Context, err error) {
	meth, ok := grpc.Method(ctx)
	if ok && (meth == "/api.Doko/Register" || meth == "/api.Doko/Login") {
		return ctx, nil // ok to call register without auth
	}
	token, err := grpcAuth.AuthFromMD(ctx, "basic")
	if err != nil {
		log.Printf("no basic auth: %v", err)
		return nil, err
	}
	user, err := a.Users.VerifyToken(token)
	if err != nil {
		log.Printf("invalid token %s", token)
		return ctx, status.Error(codes.Unauthenticated, "invalid token")
	}
	return context.WithValue(ctx, userMDKey{}, user), nil

}

func GetAuthenticatedUser(ctx context.Context) (users.AccountData, bool) {
	user := ctx.Value(userMDKey{})
	switch md := user.(type) {
	case users.AccountData:
		return md, true
	default:
		return users.AccountData{}, false
	}
}

func CreateGrpcServerForAuth(auth Auth) *grpc.Server {
	return grpc.NewServer(grpc.UnaryInterceptor(grpcAuth.UnaryServerInterceptor(auth.Authenticate)),
		grpc.StreamInterceptor(grpcAuth.StreamServerInterceptor(auth.Authenticate)),
		grpc.KeepaliveParams(keepalive.ServerParameters{Timeout: time.Hour * 24}))
}
