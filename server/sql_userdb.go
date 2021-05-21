package server

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/supermihi/karlchencloud/utils/security"
)

var schema = `
CREATE TABLE IF NOT EXISTS user (
	id integer PRIMARY KEY,
    email text NOT NULL UNIQUE,
    name text NOT NULL,
    secret text,
    is_admin boolean
);
`

type SqlUser struct {
	Id              int    `db:"user_id"`
	Email           string `db:"email"`
	Name            string `db:"name"`
	PasswordHash    string `db:"secret"`
	IsAdministrator bool   `db:"is_admin"`
}

type SqlUserDatabase struct {
	db *sqlx.DB
}

func (s *SqlUserDatabase) Add(email string, password string, name string, isAdmin bool) (id UserId, err error) {
	hash, err := security.HashAndSalt(password)
	result, err := s.db.Exec("INSERT INTO user (email, name, secret, is_admin) VALUES (?, ?, ?, ?)",
		email, name, hash, isAdmin)
	if err != nil {
		return 0, err
	}
	userId, err := result.LastInsertId()
	return UserId(userId), err
}

func (s *SqlUserDatabase) ListIds() ([]UserId, error) {
	rows, err := s.db.Query("SELECT id FROM user")
	if err != nil {
		return nil, err
	}
	var ids []UserId
	var id int64
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, UserId(id))
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return ids, nil
}

func (s *SqlUserDatabase) GetName(id UserId) (name string, err error) {
	row := s.db.QueryRow("SELECT name FROM user WHERE id = ?", id)
	err = row.Scan(&name)
	return
}

func (s *SqlUserDatabase) ChangeName(id UserId, newName string) (ok bool) {
	panic("implement me")
}

func (s *SqlUserDatabase) Authenticate(id UserId, secret string) bool {
	panic("implement me")
}

func (s *SqlUserDatabase) Close() error {
	return s.db.Close()
}

func NewSqlUserDatabase(filename string) (*SqlUserDatabase, error) {
	db, err := sqlx.Connect("sqlite3", filename)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(schema)
	if err != nil {
		return nil, err
	}
	users := SqlUserDatabase{db}
	return &users, nil
}
