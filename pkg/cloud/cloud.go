package cloud

type Cloud struct {
	Users  Users
	tables *Tables
}

func NewCloud(users Users) Cloud {
	return Cloud{users, NewTables()}
}

func (c *Cloud) CreateTable(owner UserId) TableId {
	return c.tables.CreateTable(owner)
}

func (c *Cloud) ListTables() []TableId {
	return c.tables.List()
}

func (c *Cloud) TablePlayers(id TableId) []UserId {
	return c.tables.tables[id].players
}
