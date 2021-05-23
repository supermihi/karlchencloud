package users

import (
	"errors"
	"github.com/supermihi/karlchencloud/server"
	"github.com/supermihi/karlchencloud/utils/security"
	"sync"
)

type MemoryUserDb struct {
	users     map[server.UserId]*User
	mtx       sync.RWMutex
	largestId int64
}

type User struct {
	Id           server.UserId
	Email        string
	Name         string
	PasswordHash string
}

func NewUser(id server.UserId, email string, name string, passwordHash string) *User {
	return &User{id, name, name, passwordHash}
}

func NewMemoryUserDb() *MemoryUserDb {
	return &MemoryUserDb{users: make(map[server.UserId]*User), largestId: 1}
}

func (m *MemoryUserDb) Add(email string, password string, name string, isAdmin bool) (id server.UserId, err error) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	for _, user := range m.users {
		if user.Email == email {
			return server.InvalidUserId, errors.New("user exists")
		}
	}
	m.largestId += 1
	id = server.UserId(m.largestId)
	hash, err := security.HashAndSalt(password)
	if err != nil {
		return server.InvalidUserId, err
	}
	m.users[id] = NewUser(id, email, name, hash)
	return
}

func (m *MemoryUserDb) ListIds() ([]server.UserId, error) {
	m.mtx.RLock()
	ans := make([]server.UserId, 0, len(m.users))
	for id := range m.users {
		ans = append(ans, id)
	}
	m.mtx.RUnlock()
	return ans, nil
}

func (m *MemoryUserDb) GetName(id server.UserId) (name string, err error) {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	user, ok := m.users[id]
	if !ok {
		return "", errors.New("user does not exist")
	}
	return user.Name, nil
}

func (m *MemoryUserDb) ChangeName(id server.UserId, newName string) error {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	existing, ok := m.users[id]
	if !ok {
		return errors.New("user does not exist")
	}
	existing.Name = newName
	return nil
}

func (m *MemoryUserDb) Authenticate(id server.UserId, password string) bool {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	user, ok := m.users[id]
	if !ok {
		return false
	}

	return security.VerifyPassword(password, user.PasswordHash)
}
