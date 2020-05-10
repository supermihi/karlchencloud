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
	users.Add(cloud.UserId("6ba7b810-9dad-11d1-80b4-00c04fd430c8"), "Woldemar", "geheim")
	cloud.StartServer(users, port)
}
