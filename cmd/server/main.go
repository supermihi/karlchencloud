package main

//go:generate protoc -I ../../api ../../api/karlchen.proto --go_out=plugins=grpc:../../api

import (
	"github.com/supermihi/karlchencloud/pkg/cloud"
)

const (
	port = ":50051"
)

func main() {
	users := cloud.NewMemoryUserDb()
	cloud.StartServer(users, port)
}
