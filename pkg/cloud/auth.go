package cloud

import (
	"context"
	"encoding/base64"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

type Auth struct {
	Users Users
}

func NewAuth(users Users) Auth {
	return Auth{users}
}

const userIdHeader = "auth_user"
const secretHeader = "auth_secret"

type userMDKey struct{}
type userMetadata struct {
	user UserId
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
	basic, err := grpc_auth.AuthFromMD(ctx, "basic")
	if err != nil {
		return nil, err
	}
	userId, secret, err := parseUserIdSecret(basic)
	if err != nil {
		return nil, err
	}
	if !a.Users.Authenticate(UserId(userId), secret) {
		return ctx, status.Error(codes.Unauthenticated, "invalid user/secret combination")
	}
	userMd := userMetadata{user: UserId(userId)}
	return context.WithValue(ctx, userMDKey{}, userMd), nil

}

func GetAuthenticatedUser(ctx context.Context) (UserId, bool) {
	userMD := ctx.Value(userMDKey{})

	switch md := userMD.(type) {
	case userMetadata:
		return md.user, true
	default:
		return UserId(""), false
	}
}

type ClientCredentials struct {
	userId string
	secret string
}

func NewClientCredentials(userId string, secret string) ClientCredentials {
	return ClientCredentials{userId, secret}
}
func (r ClientCredentials) RequireTransportSecurity() bool {
	return false
}

func (r ClientCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	auth := r.userId + ":" + r.secret
	enc := base64.StdEncoding.EncodeToString([]byte(auth))
	return map[string]string{
		"authorization": "basic " + enc,
	}, nil
}