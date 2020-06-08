package server

type Room struct {
	Users  Users
	Tables *Tables
}

func NewRoom(users Users) *Room {
	return &Room{users, NewTables()}
}

func (c *Room) CreateTable(owner string) *Table {
	return c.Tables.CreateTable(owner)
}

func (c *Room) TablePlayers(id string) []string {
	return c.Tables.ById[id].players
}
