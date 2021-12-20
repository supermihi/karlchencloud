package api

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
