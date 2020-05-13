package client

import (
	"context"
	"github.com/supermihi/karlchencloud/api"
	"google.golang.org/grpc"
	"log"
	"time"
)

type Client struct {
	Kc         api.KarlchencloudClient
	Ctx        context.Context
	Cancel     context.CancelFunc
	Connection *grpc.ClientConn
}

func GetConnectedService(addr string, name string, timeout time.Duration) *Client {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	creds := EmptyClientCredentials()
	conn, err := grpc.DialContext(ctx, addr, grpc.WithInsecure(), grpc.WithBlock(),
		grpc.WithPerRPCCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	kc := api.NewKarlchencloudClient(conn)
	ans, err := kc.Register(ctx, &api.RegisterRequest{Name: name})
	if err != nil {
		log.Fatalf("could not register: %v", err)
	} else {
		log.Printf("registered %v with id %v", name, ans.Id)
	}
	creds.UpdateLogin(ans.Id, ans.Secret)
	return &Client{kc, ctx, cancel, conn}
}
