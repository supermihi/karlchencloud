package server

import (
	"errors"
	"github.com/supermihi/karlchencloud/utils/security"
	"sync"
)

type MemoryUserDb struct {
	users     map[UserId]*User
	mtx       sync.RWMutex
	largestId int64
}

type User struct {
	Id           UserId
	Email        string
	Name         string
	PasswordHash string
}

func NewUser(id UserId, email string, name string, passwordHash string) *User {
	return &User{id, name, name, passwordHash}
}

func NewMemoryUserDb() *MemoryUserDb {
	return &MemoryUserDb{users: make(map[UserId]*User), largestId: 1}
}

func (m *MemoryUserDb) Add(email string, password string, name string, isAdmin bool) (id UserId, err error) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	for _, user := range m.users {
		if user.Email == email {
			return InvalidUserId, errors.New("user exists")
		}
	}
	m.largestId += 1
	id = UserId(m.largestId)
	hash, err := security.HashAndSalt(password)
	if err != nil {
		return InvalidUserId, err
	}
	m.users[id] = NewUser(id, email, name, hash)
	return
}

func (m *MemoryUserDb) ListIds() ([]UserId, error) {
	m.mtx.RLock()
	ans := make([]UserId, 0, len(m.users))
	for id := range m.users {
		ans = append(ans, id)
	}
	m.mtx.RUnlock()
	return ans, nil
}

func (m *MemoryUserDb) GetName(id UserId) (name string, ok bool) {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	user, ok := m.users[id]
	if !ok {
		return "", false
	}
	return user.Name, true
}

func (m *MemoryUserDb) ChangeName(id UserId, newName string) (ok bool) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	existing, ok := m.users[id]
	if !ok {
		return false
	}
	existing.Name = newName
	return true
}

func (m *MemoryUserDb) Authenticate(id UserId, password string) bool {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	user, ok := m.users[id]
	if !ok {
		return false
	}

	return security.VerifyPassword(password, user.PasswordHash)
}
