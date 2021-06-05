package users

import (
	"github.com/supermihi/karlchencloud/server/errors"
	"github.com/supermihi/karlchencloud/utils/security"
	"sync"
)

type MemoryUserDb struct {
	users     map[Id]*User
	mtx       sync.RWMutex
	largestId int64
}

type User struct {
	Id           Id
	Email        string
	Name         string
	PasswordHash string
}

func ToUserData(user User) AccountData {
	return NewUserData(user.Id, user.Name, user.Email, user.Id.String())
}

func NewUser(id Id, email string, name string, passwordHash string) *User {
	return &User{id, email, name, passwordHash}
}

func NewMemoryUserDb() *MemoryUserDb {
	return &MemoryUserDb{users: make(map[Id]*User), largestId: 1}
}

func (m *MemoryUserDb) Add(email string, password string, name string) (AccountData, error) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	for _, user := range m.users {
		if user.Email == email {
			return AccountData{}, errors.NewCloudError(errors.UserAlreadyExists)
		}
	}
	m.largestId += 1
	id := Id(m.largestId)
	hash, err := security.HashAndSalt(password)
	if err != nil {
		return AccountData{}, err
	}
	user := NewUser(id, email, name, hash)
	m.users[id] = user
	return ToUserData(*user), nil
}

func (m *MemoryUserDb) ListIds() ([]Id, error) {
	m.mtx.RLock()
	ans := make([]Id, 0, len(m.users))
	for id := range m.users {
		ans = append(ans, id)
	}
	m.mtx.RUnlock()
	return ans, nil
}

func (m *MemoryUserDb) GetData(id Id) (AccountData, error) {
	user, ok := m.users[id]
	if !ok {
		return AccountData{}, errors.NewCloudError(errors.UserDoesNotExist)
	}
	return ToUserData(*user), nil
}

func (m *MemoryUserDb) VerifyToken(token string) (AccountData, error) {
	id, err := ParseId(token)
	if err != nil {
		return AccountData{}, errors.NewCloudError(errors.InvalidSessionToken)
	}
	user, ok := m.users[id]
	if !ok {
		return AccountData{}, errors.NewCloudError(errors.InvalidSessionToken)
	}
	return ToUserData(*user), nil
}

func (m *MemoryUserDb) ChangeName(id Id, newName string) error {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	existing, ok := m.users[id]
	if !ok {
		return errors.NewCloudError(errors.UserDoesNotExist)
	}
	existing.Name = newName
	return nil
}

func (m *MemoryUserDb) Authenticate(email string, password string) (AccountData, error) {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	for _, user := range m.users {
		if user.Email == email {
			ok := security.VerifyPassword(password, user.PasswordHash)
			if ok {
				return ToUserData(*user), nil
			} else {
				return AccountData{}, errors.NewCloudError(errors.AuthenticationFailed)
			}
		}
	}
	return AccountData{}, errors.NewCloudError(errors.UserDoesNotExist)
}
