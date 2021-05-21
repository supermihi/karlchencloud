package server

import (
	"context"
	"encoding/base64"
	grpcAuth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/status"
	"log"
	"strconv"
	"strings"
	"time"
)

type Auth struct {
	Users Users
}

func NewAuth(users Users) Auth {
	return Auth{users}
}

type userMDKey struct{}

func parseUserIdSecret(auth string) (UserId, string, error) {
	contentB, err := base64.StdEncoding.DecodeString(auth)
	if err != nil {
		return -1, "", status.Error(codes.Unauthenticated, "invalid base64 in header")
	}
	content := string(contentB)
	colonPos := strings.IndexByte(content, ':')
	if colonPos < 0 {
		return -1, "", status.Error(codes.Unauthenticated, "invalid basic auth format")
	}
	idStr, secret := content[:colonPos], content[colonPos+1:]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return -1, "", status.Error(codes.Unauthenticated, "cannot parse user id")
	}
	return UserId(id), secret, nil
}

func (a *Auth) Authenticate(ctx context.Context) (newCtx context.Context, err error) {
	meth, ok := grpc.Method(ctx)
	if ok && meth == "/api.Doko/Register" {
		return ctx, nil // ok to call register without auth
	}
	basic, err := grpcAuth.AuthFromMD(ctx, "basic")
	if err != nil {
		log.Printf("no basic auth: %v", err)
		return nil, err
	}
	userId, secret, err := parseUserIdSecret(basic)
	if err != nil {
		log.Printf("could not parse user/secret in %s: %v", basic, err)
		return nil, err
	}
	if !a.Users.Authenticate(userId, secret) {
		log.Printf("invalid user/secret combination %s/%s", userId, secret)
		return ctx, status.Error(codes.Unauthenticated, "invalid user/secret combination")
	}
	name, ok := a.Users.GetName(userId)
	if !ok {
		log.Printf("unknown error: user does not exist in auth")
		return ctx, status.Error(codes.Internal, "unknown error: did not find authenticated user")
	}
	userMd := UserData{Id: userId, Name: name}
	return context.WithValue(ctx, userMDKey{}, userMd), nil

}

func GetAuthenticatedUser(ctx context.Context) (UserData, bool) {
	userMD := ctx.Value(userMDKey{})

	switch md := userMD.(type) {
	case UserData:
		return md, true
	default:
		return UserData{}, false
	}
}

func CreateGrpcServerForAuth(auth Auth) *grpc.Server {
	return grpc.NewServer(grpc.UnaryInterceptor(grpcAuth.UnaryServerInterceptor(auth.Authenticate)),
		grpc.StreamInterceptor(grpcAuth.StreamServerInterceptor(auth.Authenticate)),
		grpc.KeepaliveParams(keepalive.ServerParameters{Timeout: time.Hour * 24}))
}
