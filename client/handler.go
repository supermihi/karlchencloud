package client

import "github.com/supermihi/karlchencloud/api"

// ClientHandler is an interface for implementing karlchencloud clients
//
type ClientHandler interface {
	OnConnect()
	OnWelcome(s *api.UserState)
	OnMyTurnAuction()
	OnMyTurnGame()
	OnMemberJoin(userId string, name string)
	OnMatchStart()
	OnPlayedCard(card *api.PlayedCard)
}
