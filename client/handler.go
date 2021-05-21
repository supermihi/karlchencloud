package client

import "github.com/supermihi/karlchencloud/api"

// ClientHandler is an interface for implementing karlchencloud clients
//
type ClientHandler interface {
	OnConnect(service *ClientService)
	OnWelcome(client *KarlchenClient, s *api.UserState)
	OnMyTurn()
	OnTableStateReceived(state *api.TableState)
	OnMemberEvent(ev *api.MemberEvent)
	OnMatchStart(s *api.MatchState)
	OnDeclaration(d *api.Declaration)
	OnPlayedCard(card *api.PlayedCard)
}
