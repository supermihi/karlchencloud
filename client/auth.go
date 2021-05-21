package client

import (
	"context"
	"encoding/base64"
)

type ClientCredentials struct {
	userId   string
	password string
}

func (c *ClientCredentials) UserId() string {
	return c.userId
}

func (c *ClientCredentials) RequireTransportSecurity() bool {
	return false
}

func (c *ClientCredentials) GetRequestMetadata(_ context.Context, _ ...string) (map[string]string, error) {
	return getAuthMeta(c.userId, c.password), nil
}

func getAuthMeta(userId string, password string) map[string]string {
	auth := userId + ":" + password
	enc := base64.StdEncoding.EncodeToString([]byte(auth))
	return map[string]string{
		"authorization": "basic " + enc,
	}
}
