package users

import (
	"errors"
	"github.com/supermihi/karlchencloud/room"
	"github.com/supermihi/karlchencloud/utils/security"
	"sync"
)

type MemoryUserDb struct {
	users     map[room.UserId]*User
	mtx       sync.RWMutex
	largestId int64
}

type User struct {
	Id           room.UserId
	Email        string
	Name         string
	PasswordHash string
}

func NewUser(id room.UserId, email string, name string, passwordHash string) *User {
	return &User{id, name, name, passwordHash}
}

func NewMemoryUserDb() *MemoryUserDb {
	return &MemoryUserDb{users: make(map[room.UserId]*User), largestId: 1}
}

func (m *MemoryUserDb) Add(email string, password string, name string, isAdmin bool) (id room.UserId, err error) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	for _, user := range m.users {
		if user.Email == email {
			return room.InvalidUserId, errors.New("user exists")
		}
	}
	m.largestId += 1
	id = room.UserId(m.largestId)
	hash, err := security.HashAndSalt(password)
	if err != nil {
		return room.InvalidUserId, err
	}
	m.users[id] = NewUser(id, email, name, hash)
	return
}

func (m *MemoryUserDb) ListIds() ([]room.UserId, error) {
	m.mtx.RLock()
	ans := make([]room.UserId, 0, len(m.users))
	for id := range m.users {
		ans = append(ans, id)
	}
	m.mtx.RUnlock()
	return ans, nil
}

func (m *MemoryUserDb) GetName(id room.UserId) (name string, err error) {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	user, ok := m.users[id]
	if !ok {
		return "", errors.New("user does not exist")
	}
	return user.Name, nil
}

func (m *MemoryUserDb) ChangeName(id room.UserId, newName string) error {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	existing, ok := m.users[id]
	if !ok {
		return errors.New("user does not exist")
	}
	existing.Name = newName
	return nil
}

func (m *MemoryUserDb) Authenticate(id room.UserId, password string) bool {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	user, ok := m.users[id]
	if !ok {
		return false
	}

	return security.VerifyPassword(password, user.PasswordHash)
}
