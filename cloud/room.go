package cloud

type Room struct {
	Users  Users
	Tables *Tables
}

func NewRoom(users Users) Room {
	return Room{users, NewTables()}
}

func (c *Room) CreateTable(owner UserId) *Table {
	return c.Tables.CreateTable(owner)
}

func (c *Room) TablePlayers(id string) []UserId {
	return c.Tables.ById[id].players
}
