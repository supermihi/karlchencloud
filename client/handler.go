package client

import "github.com/supermihi/karlchencloud/api"

// ClientHandler is an interface for implementing karlchencloud clients
//
type ClientHandler interface {
	OnConnect(client ClientApi)
	OnWelcome(client ClientApi, s *api.UserState)
	OnMyTurnAuction(client ClientApi)
	OnMyTurnGame(client ClientApi)
	OnMemberJoin(client ClientApi, userId string, name string)
	OnMatchStart(client ClientApi)
	OnPlayedCard(client ClientApi, card *api.PlayedCard)
}
