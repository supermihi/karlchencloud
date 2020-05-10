package cloud

import uuid "github.com/satori/go.uuid"

type MemoryUserDb struct {
	users map[UserId]*UserData
}

type UserData struct {
	id     UserId
	name   string
	secret string
}

func RandomSecret() string {
	return uuid.NewV4().String()
}

func NewUserData(id UserId, name string, secret string) *UserData {
	return &UserData{id, name, secret}
}

func NewMemoryUserDb() *MemoryUserDb {
	users := make(map[UserId]*UserData)
	ans := MemoryUserDb{users}
	return &ans
}

func (m *MemoryUserDb) Add(user UserId, name string, secret string) bool {
	_, exists := m.users[user]
	if exists {
		return false
	}
	m.users[user] = NewUserData(user, name, secret)
	return true
}

func (m *MemoryUserDb) List() []UserId {
	ans := make([]UserId, 0, len(m.users))
	for id := range m.users {
		ans = append(ans, id)
	}
	return ans
}

func (m *MemoryUserDb) GetName(id UserId) string {
	return m.users[id].name
}

func (m *MemoryUserDb) ChangeName(id UserId, newName string) {
	_, exists := m.users[id]
	if !exists {
		panic("cannot change newName of not existing user")
	}
	m.users[id].name = newName
}

func (m *MemoryUserDb) Authenticate(id UserId, secret string) bool {
	user, ok := m.users[id]
	if !ok {
		return false
	}
	return user.secret == secret
}
