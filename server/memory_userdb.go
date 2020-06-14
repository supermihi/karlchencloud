package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

type MemoryUserDb struct {
	users    map[string]*User
	mtx      sync.RWMutex
	filename string
}

type User struct {
	Id     string
	Name   string
	Secret string
}

func RandomSecret() string {
	return RandomLetters(16)
}

func NewUser(id string, name string, secret string) *User {
	return &User{id, name, secret}
}

func NewMemoryUserDb(filename string) (*MemoryUserDb, error) {
	users := make(map[string]*User)
	if filename != "" {
		userList, err := importUsers(filename)
		if err != nil {
			return nil, err
		}
		for _, u := range userList {
			users[u.Id] = &u
		}
	}
	ans := MemoryUserDb{users: users, filename: filename}
	return &ans, nil
}

func importUsers(filename string) (users []User, err error) {
	if _, statErr := os.Stat(filename); statErr != nil {
		return
	}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &users)
	return
}

func (m *MemoryUserDb) Export() {
	if m.filename == "" {
		return
	}
	log.Printf("exporting user db to %s", m.filename)
	users := make([]User, len(m.users))
	i := 0
	for _, u := range m.users {
		users[i] = *u
		i++
	}
	ans, err := json.Marshal(users)
	if err != nil {
		log.Fatalf("error exporting user db: %v", err)
	}
	if err := ioutil.WriteFile(m.filename, ans, 0644); err != nil {
		log.Fatalf("error writing user db: %v", err)
	}

}

func (m *MemoryUserDb) Add(user string, name string, secret string) bool {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	_, exists := m.users[user]
	if exists {
		return false
	}
	m.users[user] = NewUser(user, name, secret)
	m.Export()
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

func (m *MemoryUserDb) GetName(id string) (name string, ok bool) {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	user, ok := m.users[id]
	if !ok {
		return "", false
	}
	return user.Name, true
}

func (m *MemoryUserDb) ChangeName(id string, newName string) (ok bool) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	existing, ok := m.users[id]
	if !ok {
		return false
	}
	m.Export()
	existing.Name = newName
	return true
}

func (m *MemoryUserDb) Authenticate(id string, secret string) bool {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	user, ok := m.users[id]
	if !ok {
		return false
	}
	return user.Secret == secret
}
