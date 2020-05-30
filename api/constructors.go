package api

func NewMemberEvent(user string, name string, eventType MemberEventType) *MatchEvent {
	return &MatchEvent{
		Event: &MatchEvent_Member{
			Member: &MemberEvent{
				UserId: user,
				Name:   name,
				Type:   eventType},
		},
	}
}
