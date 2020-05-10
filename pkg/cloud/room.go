package cloud

type Room struct {
	Users  Users
	tables *Tables
}

func NewRoom(users Users) Room {
	return Room{users, NewTables()}
}

func (c *Room) CreateTable(owner UserId) string {
	return c.tables.CreateTable(owner)
}

func (c *Room) ListTables() []string {
	return c.tables.List()
}

func (c *Room) TablePlayers(id string) []UserId {
	return c.tables.tables[id].players
}
