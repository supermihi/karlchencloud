package server

import "sync"

type MemoryUserDb struct {
	users map[string]*User
	mtx   sync.RWMutex
}

type User struct {
	id     string
	name   string
	secret string
}

func RandomSecret() string {
	return RandomLetters(16)
}

func NewUser(id string, name string, secret string) *User {
	return &User{id, name, secret}
}

func NewMemoryUserDb() *MemoryUserDb {
	users := make(map[string]*User)
	ans := MemoryUserDb{users: users}
	return &ans
}

func (m *MemoryUserDb) Add(user string, name string, secret string) bool {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	_, exists := m.users[user]
	if exists {
		return false
	}
	m.users[user] = NewUser(user, name, secret)
	return true
}

func (m *MemoryUserDb) List() []string {
	m.mtx.RLock()
	ans := make([]string, 0, len(m.users))
	for id := range m.users {
		ans = append(ans, id)
	}
	m.mtx.RUnlock()
	return ans
}

func (m *MemoryUserDb) GetName(id string) string {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	return m.users[id].name
}

func (m *MemoryUserDb) ChangeName(id string, newName string) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	_, exists := m.users[id]
	if !exists {
		panic("cannot change newName of not existing user")
	}
	m.users[id].name = newName
}

func (m *MemoryUserDb) Authenticate(id string, secret string) bool {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	user, ok := m.users[id]
	if !ok {
		return false
	}
	return user.secret == secret
}
