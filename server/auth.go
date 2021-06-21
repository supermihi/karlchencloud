package server

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/supermihi/karlchencloud/server/users"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

// CreateAuthenticatingGrpcServer returns a grpc.Server set up with authenticating interceptors from the grpc_auth
// package.
func CreateAuthenticatingGrpcServer(users users.Users) *grpc.Server {
	authFunc := createAuthFunc(users)
	return grpc.NewServer(grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(authFunc)),
		grpc.StreamInterceptor(grpc_auth.StreamServerInterceptor(authFunc)),
		grpc.KeepaliveParams(keepalive.ServerParameters{Timeout: time.Hour * 24}))
}

// GetAuthenticatedUser returns the users.AccountData identified by the valid session token in the grpc metadata.
// If the token is missing or invalid, the success flag returned is false.
func GetAuthenticatedUser(ctx context.Context) (user users.AccountData, success bool) {
	userValue := ctx.Value(authenticatedUserContextKey{})
	switch md := userValue.(type) {
	case users.AccountData:
		return md, true
	default:
		return users.AccountData{}, false
	}
}

type authenticatedUserContextKey struct{}

func createAuthFunc(users users.Users) grpc_auth.AuthFunc {
	return func(ctx context.Context) (newCtx context.Context, err error) {
		return authenticate(users, ctx)
	}
}

func authenticate(users users.Users, ctx context.Context) (newCtx context.Context, err error) {
	meth, ok := grpc.Method(ctx)
	if ok && (meth == "/api.Doko/Register" || meth == "/api.Doko/Login") {
		return ctx, nil // ok to call these methods without token
	}
	token, err := grpc_auth.AuthFromMD(ctx, "basic")
	if err != nil {
		log.Printf("no basic auth: %v", err)
		return nil, err
	}
	user, err := users.VerifyToken(token)
	if err != nil {
		log.Printf("invalid token %s", token)
		return ctx, status.Error(codes.Unauthenticated, "invalid token")
	}
	return context.WithValue(ctx, authenticatedUserContextKey{}, user), nil

}
