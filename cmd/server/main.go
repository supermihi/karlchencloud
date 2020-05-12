package main

//go:generate protoc -I ../../api ../../api/karlchen.proto --go_out=plugins=grpc:../../api

import (
	"github.com/supermihi/karlchencloud/cloud"
	"github.com/supermihi/karlchencloud/server"
)

const (
	port = ":50051"
)

func main() {
	users := cloud.NewMemoryUserDb()
	server.StartServer(users, port)
}
