package client

import (
	"context"
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/server"
	"google.golang.org/grpc"
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
