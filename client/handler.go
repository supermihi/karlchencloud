package client

import pb "github.com/supermihi/karlchencloud/api"

// ClientHandler is an interface for implementing karlchencloud clients
//
type ClientHandler interface {
	OnConnect()
	OnWelcome(*pb.UserState)
	OnNewTable(TableInfo)
	OnMyTurnAuction()
	OnMyTurnGame()
	OnMemberJoin(userId string, name string)
	OnMatchStart()
	OnPlayedCard(*pb.PlayedCard)
}
