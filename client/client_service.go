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

type ClientService struct {
	Api        api.KarlchencloudClient
	connection *grpc.ClientConn
	Creds      *ClientCredentials
	Name       string
	Context    context.Context
}

func (c *ClientService) CloseConnection() {
	_ = c.connection.Close()
}

func (c *ClientService) Logf(format string, v ...interface{}) {
	log.Printf(c.Name+": "+format, v...)
}

type ConnectData struct {
	DisplayName    string
	ExistingUserId *string
	ExistingSecret *string
	Address        string
}

func GetClientService(c ConnectData, ctx context.Context) (ClientService, error) {
	creds := EmptyClientCredentials()
	if c.ExistingUserId != nil {
		creds.UpdateLogin(*c.ExistingUserId, *c.ExistingUserId)
	}
	conn, err := grpc.DialContext(ctx, c.Address, grpc.WithInsecure(), grpc.WithBlock(),
		grpc.WithPerRPCCredentials(creds))
	if err != nil {
		return ClientService{}, err
	}
	kc := api.NewKarlchencloudClient(conn)
	ans, loginErr := kc.CheckLogin(ctx, &api.Empty{})
	var username string
	if ans != nil {
		username = ans.Name
	}
	if c.ExistingUserId == nil || loginErr != nil {
		ans, err := kc.Register(ctx, &api.UserName{Name: c.DisplayName})
		if err != nil {
			return ClientService{}, err
		}
		creds.UpdateLogin(ans.Id, ans.Secret)
		username = c.DisplayName
		log.Printf("registered %v with id %v", c.DisplayName, ans.Id)
	}

	return ClientService{kc, conn, creds, username, ctx}, nil
}

func (c *ClientService) UserId() string {
	return c.Creds.UserId()
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

func (c *ClientService) Play(card game.Card, tableId string) (err error) {
	_, err = c.Api.Play(c.Context, &api.PlayRequest{Table: tableId,
		Request: &api.PlayRequest_Card{Card: common.ToApiCard(card)}})
	return
}
