package client

import (
	"context"
	"fmt"
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/common"
	"github.com/supermihi/karlchencloud/doko/game"
	"google.golang.org/grpc"
	"log"
)

type Client struct {
	Kc         api.KarlchencloudClient
	Connection *grpc.ClientConn
	Creds      *ClientCredentials
}

func (c *Client) Close() {
	_ = c.Connection.Close()
}

type ConnectData struct {
	DisplayName    string
	ExistingUserId *string
	ExistingSecret *string
	Address        string
}

func GetConnectedService(c ConnectData, ctx context.Context) (*Client, error) {
	creds := EmptyClientCredentials()
	if c.ExistingUserId != nil {
		creds.UpdateLogin(*c.ExistingUserId, *c.ExistingUserId)
	}
	conn, err := grpc.DialContext(ctx, c.Address, grpc.WithInsecure(), grpc.WithBlock(),
		grpc.WithPerRPCCredentials(creds))
	if err != nil {
		return nil, err
	}
	kc := api.NewKarlchencloudClient(conn)
	_, loginErr := kc.CheckLogin(ctx, &api.Empty{})
	if c.ExistingUserId == nil || loginErr != nil {
		ans, err := kc.Register(ctx, &api.RegisterRequest{Name: c.DisplayName})
		if err != nil {
			return nil, err
		}
		creds.UpdateLogin(ans.Id, ans.Secret)
		log.Printf("registered %v with id %v", c.DisplayName, ans.Id)
	}
	return &Client{kc, conn, creds}, nil
}

func MatchEventString(ev *api.MatchEventStream) string {
	switch e := ev.Event.(type) {
	case *api.MatchEventStream_Member:
		return fmt.Sprintf("user %s status: %s", e.Member.UserId, e.Member.Type)
	case *api.MatchEventStream_Start:
		return fmt.Sprintf("game started with players %v", e.Start.Players)
	}

	return ev.String()
}

func ToHand(cards []*api.Card) game.Hand {
	ans := make([]game.Card, len(cards))
	for i := 0; i < len(ans); i++ {
		ans[i] = common.ToCard(cards[i])
	}
	return ans
}
