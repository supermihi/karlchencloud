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
	auth := c.userId + ":" + c.secret
	enc := base64.StdEncoding.EncodeToString([]byte(auth))
	return map[string]string{
		"authorization": "basic " + enc,
	}, nil
}
