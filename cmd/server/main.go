package main

//go:generate protoc -I ../../api ../../api/karlchen.proto --go_out=plugins=grpc:../../api

import (
	"fmt"
	"github.com/supermihi/karlchencloud/server"
	"log"
	"net"
)

const (
	port = "0.0.0.0:9090"
)

func main() {
	users, err := server.NewMemoryUserDb("users.json")
	if err != nil {
		log.Fatalf("error creating users: %v", err)
	}
	if len(users.List()) == 0 {
		for i := 1; i < 5; i++ {
			users.Add(fmt.Sprintf("%d", i), fmt.Sprintf("dummy %d", i), "123")
		}
	}
	room := server.NewRoom(users)
	table := room.CreateTable("dummy 1")
	err = table.Join("dummy 2")
	if err != nil {
		log.Fatalf("mäh")
	}
	srv := server.CreateServer(users, room)
	log.Printf("starting raw")

	lis, err := net.Listen("tcp", ":50501")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
