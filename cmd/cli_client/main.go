package main

import (
	"github.com/supermihi/karlchencloud/client"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	c := client.GetConnectedService(address, "me", 10*time.Second)
	defer c.Cancel()
	defer func() { _ = c.Connection.Close() }()

}
