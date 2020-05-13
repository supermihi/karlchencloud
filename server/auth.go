package server

import (
	"context"
	"encoding/base64"
	grpcAuth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/supermihi/karlchencloud/cloud"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

type Auth struct {
	Users cloud.Users
}

func NewAuth(users cloud.Users) Auth {
	return Auth{users}
}

type userMDKey struct{}
type userMetadata struct {
	user cloud.UserId
}

func parseUserIdSecret(auth string) (string, string, error) {
	contentB, err := base64.StdEncoding.DecodeString(auth)
	if err != nil {
		return "", "", status.Error(codes.Unauthenticated, "invalid base64 in header")
	}
	content := string(contentB)
	colonPos := strings.IndexByte(content, ':')
	if colonPos < 0 {
		return "", "", status.Error(codes.Unauthenticated, "invalid basic auth format")
	}
	user, secret := content[:colonPos], content[colonPos+1:]
	return user, secret, nil
}

func (a *Auth) Authenticate(ctx context.Context) (newCtx context.Context, err error) {
	meth, ok := grpc.Method(ctx)
	if ok && meth == "/api.Karlchencloud/Register" {
		return ctx, nil // ok to call register without auth
	}
	basic, err := grpcAuth.AuthFromMD(ctx, "basic")
	if err != nil {
		return nil, err
	}
	userId, secret, err := parseUserIdSecret(basic)
	if err != nil {
		return nil, err
	}
	if !a.Users.Authenticate(cloud.UserId(userId), secret) {
		return ctx, status.Error(codes.Unauthenticated, "invalid user/secret combination")
	}
	userMd := userMetadata{user: cloud.UserId(userId)}
	return context.WithValue(ctx, userMDKey{}, userMd), nil

}

func GetAuthenticatedUser(ctx context.Context) (cloud.UserId, bool) {
	userMD := ctx.Value(userMDKey{})

	switch md := userMD.(type) {
	case userMetadata:
		return md.user, true
	default:
		return cloud.UserId(""), false
	}
}
