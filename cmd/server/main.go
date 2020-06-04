package main

//go:generate protoc -I ../../api ../../api/karlchen.proto --go_out=plugins=grpc:../../api

import (
	"github.com/supermihi/karlchencloud/server"
	"log"
	"net"
)

const (
	port = "0.0.0.0:9090"
)

func main() {
	users := server.NewMemoryUserDb()

	srv := server.CreateServer(users)

	log.Printf("starting raw")

	lis, err := net.Listen("tcp", ":50501")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
