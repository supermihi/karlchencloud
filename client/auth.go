package client

import (
	"context"
)

type ClientCredentials struct {
	token string
}

func (c *ClientCredentials) RequireTransportSecurity() bool {
	return false
}

func (c *ClientCredentials) GetRequestMetadata(_ context.Context, _ ...string) (map[string]string, error) {
	return getAuthMeta(c.token), nil
}

func getAuthMeta(token string) map[string]string {
	return map[string]string{
		"authorization": "basic " + token,
	}
}
