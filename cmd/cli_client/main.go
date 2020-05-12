package main

import (
	"context"
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/server"
	"log"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	creds := server.NewClientCredentials("", "")
	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock(),
		grpc.WithPerRPCCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := api.NewKarlchencloudClient(conn)
	ans, err := c.Register(ctx, &api.RegisterRequest{Name: "michael"})
	if err != nil {
		log.Fatalf("could not register: %v", err)
	} else {
		log.Printf("registered with id %v", ans.Id)
	}
	creds.UpdateLogin(ans.Id, ans.Secret)
	_, err = c.CheckLogin(ctx, &api.Empty{})
	if err != nil {
		log.Fatalf("could not login: %v", err)
	}
	log.Print("login ok!")
	table, err := c.CreateTable(ctx, &api.Empty{})
	if err != nil {
		log.Fatalf("could not create table: %v", err)
	}
	log.Printf("created table with id %v", table.TableId)
	tables, _ := c.ListTables(ctx, &api.Empty{})
	log.Printf("there are %v tables:", len(tables.Tables))
	for _, t := range tables.Tables {
		log.Printf("- one with Id %v and owner %v", t.TableId, t.Owner)
	}
	id := &api.TableId{Value: table.TableId}
	_, err = c.StartTable(ctx, id)
	if err != nil {
		log.Fatalf("could not start table: %v", err)
	}
}
