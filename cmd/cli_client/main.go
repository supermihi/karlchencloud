package main

import (
	"context"
	"github.com/supermihi/karlchencloud/api"
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
	c := api.NewKarlchencloudClient(conn)

	_, err = c.CheckLogin(ctx, &api.EmptyRequest{})
	if err != nil {
		log.Fatalf("could not login: %v", err)
	}
	log.Print("login ok!")
	tableId, err := c.CreateTable(ctx, &api.EmptyRequest{})
	if err != nil {
		log.Fatalf("could not create table: %v", err)
	}
	log.Printf("created table with id %v", tableId.Value)
	tables, _ := c.ListTables(ctx, &api.EmptyRequest{})
	log.Printf("there are %v tables:", len(tables.Ids))
	for _, tableId := range tables.Ids {
		log.Printf("- one with id %v", tableId)
	}
	_, err = c.StartTable(ctx, tableId)
	if err != nil {
		log.Fatalf("could not start table: %v", err)
	}
}
