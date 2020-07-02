package main

//go:generate protoc -I ../../api ../../api/karlchen.proto --go_out=plugins=grpc:../../api

import (
	crand "crypto/rand"
	"encoding/binary"
	"flag"
	"fmt"
	"github.com/supermihi/karlchencloud/server"
	"log"
	"math/rand"
	"net"
)

const (
	port     = ":50501"
	httpPort = ":8080"
)

func main() {
	var noHttpServer bool
	flag.BoolVar(&noHttpServer, "no-http-server", false, "")
	flag.Parse()
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
	table, cloudErr = room.JoinTable("dummy 2", table.InviteCode)
	if cloudErr != nil {
		log.Fatal(cloudErr)
	}
	srv := server.CreateServer(users, room)

	startServer := func() {
		log.Printf("starting grpc server")
		lis, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		if err := srv.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}
	startHttp := func() {
		httpServer := server.WrapServer(srv)
		lisHttp, err := net.Listen("tcp", httpPort)
		if err != nil {
			log.Fatalf("failed to listen HTTP: %v", err)
		}
		log.Printf("starting HTTP proxy server")
		if err := httpServer.Serve(lisHttp); err != nil {
			log.Fatalf("failed to serve HTTP: %v", err)
		}
	}
	if !noHttpServer {
		go startHttp()
	}
	startServer()

}
