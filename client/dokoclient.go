package client

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/supermihi/karlchencloud/api"
	"google.golang.org/grpc"
	"log"
)

type LoginData struct {
	Name               string
	Email              string
	Password           string
	RegisterOnAuthFail bool
	ServerAddress      string
}

type UserData struct {
	Name  string
	Id    string
	Email string
}

// DokoClient encapuslates technical details such as auth from the gRPC client pb.DokoClient
type DokoClient struct {
	Grpc       pb.DokoClient
	connection *grpc.ClientConn
	creds      *ClientCredentials
	user       UserData
	Context    context.Context
}

func GetConnectedDokoClient(login LoginData, ctx context.Context) (*DokoClient, error) {
	fmt.Printf("%s: connecting to %s ... ", login.Name, login.ServerAddress)
	creds := &ClientCredentials{}
	conn, err := GetGrpcClientConn(ctx, login.ServerAddress, creds)
	if err != nil {
		return nil, err
	}
	fmt.Print("connected\n")
	dokoClient := pb.NewDokoClient(conn)
	userData := UserData{Email: login.Email}
	loginResponse, err := dokoClient.Login(ctx, &pb.LoginRequest{Email: login.Email, Password: login.Password})
	if err != nil {
		if !login.RegisterOnAuthFail {
			return nil, errors.New("given credentials wrong or missing and RegisterIfEmptyUserId not set")
		}
		request := &pb.RegisterRequest{Name: login.Name, Email: login.Email, Password: login.Password}
		ans, err := dokoClient.Register(ctx, request)
		if err != nil {
			return nil, err
		}
		creds.token = ans.Token
		userData.Id = ans.UserId
		userData.Name = login.Name
		log.Printf("registered %v with token %v", login.Name, ans.Token)
	} else {
		creds.token = loginResponse.Token
		userData.Name = loginResponse.Name
		userData.Id = loginResponse.UserId
	}
	return &DokoClient{dokoClient, conn, creds, userData, ctx}, nil
}

func (c *DokoClient) CloseConnection() {
	_ = c.connection.Close()
}
