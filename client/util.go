package client

import (
	"context"
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/server"
	"google.golang.org/grpc"
	"log"
)

func ToHand(cards []*api.Card) game.Hand {
	ans := make([]game.Card, len(cards))
	for i := 0; i < len(ans); i++ {
		ans[i] = server.ToCard(cards[i])
	}
	return ans
}

func GetGrpcClientConn(ctx context.Context, server string, creds *ClientCredentials) (conn *grpc.ClientConn, err error) {
	return grpc.DialContext(ctx, server, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithPerRPCCredentials(creds))
}

func CheckExistingLogin(client api.DokoClient, ctx context.Context) bool {
	ans, err := client.CheckLogin(ctx, &api.Empty{})
	if err != nil {
		log.Printf("error in CheckLogin: %v", err)
	}
	return err == nil && ans != nil
}
