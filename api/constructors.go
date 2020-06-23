package api

import "time"

func NewMemberEvent(user string, name string, eventType MemberEventType) *Event {
	return &Event{
		Event: &Event_Member{
			Member: &MemberEvent{
				UserId: user,
				Name:   name,
				Type:   eventType},
		},
	}
}

func NewTimestamp(time time.Time) *Timestamp {
	return &Timestamp{Nanos: time.UnixNano()}
}