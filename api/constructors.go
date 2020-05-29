package api

func NewMemberEvent(user string, name string, eventType MemberEventType) *MatchEventStream {
	return &MatchEventStream{
		Event: &MatchEventStream_Member{
			Member: &MemberEvent{
				UserId: string(user),
				Name:   name,
				Type:   eventType},
		},
	}
}
