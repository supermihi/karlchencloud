package cloud

import "sync"

type MemoryUserDb struct {
	users map[UserId]*UserData
	mtx   sync.RWMutex
}

type UserData struct {
	id     UserId
	name   string
	secret string
}

func RandomSecret() string {
	return RandomLetters(16)
}

func NewUserData(id UserId, name string, secret string) *UserData {
	return &UserData{id, name, secret}
}

func NewMemoryUserDb() *MemoryUserDb {
	users := make(map[UserId]*UserData)
	ans := MemoryUserDb{users: users}
	return &ans
}

func (m *MemoryUserDb) Add(user UserId, name string, secret string) bool {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	_, exists := m.users[user]
	if exists {
		return false
	}
	m.users[user] = NewUserData(user, name, secret)
	return true
}

func (m *MemoryUserDb) List() []UserId {
	m.mtx.RLock()
	ans := make([]UserId, 0, len(m.users))
	for id := range m.users {
		ans = append(ans, id)
	}
	m.mtx.RUnlock()
	return ans
}

func (m *MemoryUserDb) GetName(id UserId) string {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	return m.users[id].name
}

func (m *MemoryUserDb) ChangeName(id UserId, newName string) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	_, exists := m.users[id]
	if !exists {
		panic("cannot change newName of not existing user")
	}
	m.users[id].name = newName
}

func (m *MemoryUserDb) Authenticate(id UserId, secret string) bool {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	user, ok := m.users[id]
	if !ok {
		return false
	}
	return user.secret == secret
}
