package users

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/supermihi/karlchencloud/server/errors"
	"github.com/supermihi/karlchencloud/utils/security"
	"log"
)

var schema = `
CREATE TABLE IF NOT EXISTS user (
	id integer PRIMARY KEY,
    email text NOT NULL UNIQUE,
    name text NOT NULL,
    secret text
);
`

type SqlUserDatabase struct {
	db *sqlx.DB
}

func (s *SqlUserDatabase) Add(email string, password string, name string) (user AccountData, err error) {
	hash, err := security.HashAndSalt(password)
	result, err := s.db.Exec("INSERT INTO user (email, name, secret) VALUES (?, ?, ?)",
		email, name, hash)
	if err != nil {
		return
	}
	userId, err := result.LastInsertId()
	return NewUserData(Id(userId), name, email, Id(userId).String()), nil
}

func (s *SqlUserDatabase) ListIds() ([]Id, error) {
	rows, err := s.db.Query("SELECT id FROM user")
	if err != nil {
		return nil, err
	}
	var ids []Id
	var id int64
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, Id(id))
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return ids, nil
}

func (s *SqlUserDatabase) GetData(id Id) (data AccountData, err error) {
	row := s.db.QueryRow("SELECT name, email FROM user WHERE id = ?", id)
	data.Id = id
	data.Token = id.String()
	err = row.Scan(&data.Name, &data.Email)
	return
}

func (s *SqlUserDatabase) ChangeName(id Id, newName string) error {
	panic("implement me")
}

func (s *SqlUserDatabase) Authenticate(email string, password string) (user AccountData, err error) {
	row := s.db.QueryRow("SELECT id, email, name, secret FROM user WHERE email = ?", email)
	var hash string
	err = row.Scan(&user.Id, &user.Email, &user.Name, &hash)
	if err != nil {
		log.Printf("user not found: %v", err)
		err = errors.NewCloudError(errors.AuthenticationFailed)
		return
	}
	if !security.VerifyPassword(password, hash) {
		err = errors.NewCloudError(errors.AuthenticationFailed)
		return
	}
	user.Token = user.Id.String()
	return
}

func (s *SqlUserDatabase) VerifyToken(token string) (user AccountData, err error) {
	row := s.db.QueryRow("SELECT id, email, name FROM user WHERE id = ?", token)
	err = row.Scan(&user.Id, &user.Email, &user.Name)
	if err != nil {
		return AccountData{}, errors.NewCloudError(errors.InvalidSessionToken)
	}
	user.Token = user.Id.String()
	return
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
