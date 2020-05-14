package main

import (
	"context"
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/client"
	"io/ioutil"
	"log"
	"time"
)

const address = "localhost:50051"

func main() {
	conn := client.ConnectData{
		DisplayName:    "client",
		ExistingUserId: nil,
		ExistingSecret: nil,
		Address:        address,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	c, err := client.GetConnectedService(conn, ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	table, err := c.Kc.CreateTable(ctx, &api.Empty{})
	if err != nil {
		log.Fatalf("%s could not create table: %v", conn.DisplayName, err)
	}
	log.Printf("table %s created with invite code %s", table.TableId, table.InviteCode)
	err = ioutil.WriteFile("table.config", []byte(table.TableId+":"+table.InviteCode), 0644)
	if err != nil {
		log.Fatalf("error wrtiing table.config: %v", err)
	}
	log.Printf("wrote table.config")
	stream, err := c.Kc.SubscribeMatchEvents(ctx, &api.TableId{Value: table.TableId})
	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Print(err)
			break
		}
		log.Printf("new match event %v", msg)
	}
	_, err = c.Kc.StartTable(ctx, &api.TableId{Value: table.TableId})
	if err != nil {
		log.Fatalf("%s could not start table: %v", conn.DisplayName, err)
	}

}
