package client

import (
	"context"
	"github.com/supermihi/karlchencloud/api"
	"google.golang.org/grpc"
	"log"
)

type Client struct {
	Kc         api.KarlchencloudClient
	Connection *grpc.ClientConn
}

func (c *Client) Close() {
	_ = c.Connection.Close()
}

type ConnectData struct {
	DisplayName    string
	ExistingUserId *string
	ExistingSecret *string
	Address        string
}

func GetConnectedService(c ConnectData, ctx context.Context) (*Client, error) {
	creds := EmptyClientCredentials()
	if c.ExistingUserId != nil {
		creds.UpdateLogin(*c.ExistingUserId, *c.ExistingUserId)
	}
	conn, err := grpc.DialContext(ctx, c.Address, grpc.WithInsecure(), grpc.WithBlock(),
		grpc.WithPerRPCCredentials(creds))
	if err != nil {
		return nil, err
	}
	kc := api.NewKarlchencloudClient(conn)
	_, loginErr := kc.CheckLogin(ctx, &api.Empty{})
	if c.ExistingUserId == nil || loginErr != nil {
		ans, err := kc.Register(ctx, &api.RegisterRequest{Name: c.DisplayName})
		if err != nil {
			return nil, err
		}
		creds.UpdateLogin(ans.Id, ans.Secret)
		log.Printf("registered %v with id %v", c.DisplayName, ans.Id)
	}
	return &Client{kc, conn}, nil
}
