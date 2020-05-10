package main

import (
	"context"
	"github.com/supermihi/karlchencloud/grpc_api"
	"github.com/supermihi/karlchencloud/pkg/cloud"
	"log"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

type Cred struct {
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	creds := cloud.NewClientCredentials("6ba7b810-9dad-11d1-80b4-00c04fd430c8", "geheim")
	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock(),
		grpc.WithPerRPCCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := grpc_api.NewKarlchencloudClient(conn)

	_, err = c.CheckLogin(ctx, &grpc_api.EmptyRequest{})
	if err != nil {
		log.Fatalf("could not login: %v", err)
	}
	log.Print("login ok!")
}
