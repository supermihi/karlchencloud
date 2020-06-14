package main

//go:generate protoc -I ../../api ../../api/karlchen.proto --go_out=plugins=grpc:../../api

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"github.com/supermihi/karlchencloud/server"
	"log"
	"math/rand"
	"net"
)

const (
	port = ":50501"
)

func main() {
	var v int64
	randErr := binary.Read(crand.Reader, binary.BigEndian, &v)
	if randErr != nil {
		log.Fatal(randErr)
	}
	rand.Seed(v)
	users, err := server.NewExportedMemoryUserDb("users.json")
	if err != nil {
		log.Fatalf("error creating users: %v", err)
	}
	if len(users.List()) == 0 {
		for i := 1; i < 5; i++ {
			users.Add(fmt.Sprintf("%d", i), fmt.Sprintf("dummy %d", i), "123")
		}
	}
	room := server.NewRoom(users)
	table, cloudErr := room.CreateTable("dummy 1")
	if cloudErr != nil {
		log.Fatal(cloudErr)
	}
	table, cloudErr = room.JoinTable(table.Id, "dummy 2", table.InviteCode)
	if cloudErr != nil {
		log.Fatal(cloudErr)
	}
	srv := server.CreateServer(users, room)
	log.Printf("starting raw")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
