package api

import "github.com/supermihi/karlchencloud/cloud"

func NewMemberEvent(user cloud.UserId, eventType MemberEventType) *MatchEventStream {
	return &MatchEventStream{
		Event: &MatchEventStream_Member{
			Member: &MemberEvent{
				UserId: string(user),
				Type:   eventType},
		},
	}
}
