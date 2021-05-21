package client

import (
	"context"
	"errors"
	"github.com/supermihi/karlchencloud/api"
	"google.golang.org/grpc"
	"log"
)

type LoginData struct {
	Name     string
	Email    string
	Password string
	// set to empty string if no account yet
	UserId                string
	RegisterIfEmptyUserId bool
	ServerAddress         string
}

type ClientService struct {
	Grpc       api.DokoClient
	connection *grpc.ClientConn
	creds      *ClientCredentials
	Name       string
	Context    context.Context
}

func GetConnectedClientService(login LoginData, ctx context.Context) (*ClientService, error) {
	creds := &ClientCredentials{userId: login.UserId, password: login.Password}
	log.Printf("connecting to %s ...", login.ServerAddress)
	conn, err := GetGrpcClientConn(ctx, login.ServerAddress, creds)
	if err != nil {
		return nil, err
	}
	log.Print("connected")
	dokoClient := api.NewDokoClient(conn)
	loggedIn := login.UserId != "" && CheckExistingLogin(dokoClient, ctx)
	if !loggedIn {
		if !login.RegisterIfEmptyUserId {
			return nil, errors.New("given credentials wrong or missing and RegisterIfEmptyUserId not set")
		}
		request := &api.RegisterRequest{Name: login.Name, Email: login.Email, Password: login.Password}
		ans, err := dokoClient.Register(ctx, request)
		if err != nil {
			return nil, err
		}
		creds.userId = ans.Id
		log.Printf("registered %v with id %v", login.Name, ans.Id)
	}
	return &ClientService{dokoClient, conn, creds, login.Name, ctx}, nil
}

func (c *ClientService) CloseConnection() {
	_ = c.connection.Close()
}

func (c *ClientService) Logf(format string, v ...interface{}) {
	log.Printf(c.Name+": "+format, v...)
}

func (c *ClientService) UserId() string {
	return c.creds.UserId()
}
