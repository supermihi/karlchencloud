package client

import (
	"context"
	"encoding/base64"
)

type ClientCredentials struct {
	userId string
	secret string
}

func NewClientCredentials(userId string, secret string) *ClientCredentials {
	return &ClientCredentials{userId, secret}
}

func (c *ClientCredentials) UserId() string {
	return c.userId
}

func EmptyClientCredentials() *ClientCredentials {
	return &ClientCredentials{"", ""}
}

func (c *ClientCredentials) UpdateLogin(userId string, secret string) {
	c.userId = userId
	c.secret = secret
}

func (c *ClientCredentials) RequireTransportSecurity() bool {
	return false
}

func (c *ClientCredentials) GetRequestMetadata(_ context.Context, _ ...string) (map[string]string, error) {
	return getAuthMeta(c.userId, c.secret), nil
}

func getAuthMeta(userId string, secret string) map[string]string {
	auth := userId + ":" + secret
	enc := base64.StdEncoding.EncodeToString([]byte(auth))
	return map[string]string{
		"authorization": "basic " + enc,
	}
}