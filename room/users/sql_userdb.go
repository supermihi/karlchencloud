package users

import (
	"errors"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/supermihi/karlchencloud/room"
	"github.com/supermihi/karlchencloud/utils/security"
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

func (s *SqlUserDatabase) Add(email string, password string, name string) (user room.UserData, err error) {
	hash, err := security.HashAndSalt(password)
	result, err := s.db.Exec("INSERT INTO user (email, name, secret) VALUES (?, ?, ?)",
		email, name, hash)
	if err != nil {
		return
	}
	userId, err := result.LastInsertId()
	return room.NewUserData(room.UserId(userId), name, email, room.UserId(userId).String()), nil
}

func (s *SqlUserDatabase) ListIds() ([]room.UserId, error) {
	rows, err := s.db.Query("SELECT id FROM user")
	if err != nil {
		return nil, err
	}
	var ids []room.UserId
	var id int64
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, room.UserId(id))
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return ids, nil
}

func (s *SqlUserDatabase) GetData(id room.UserId) (data room.UserData, err error) {
	row := s.db.QueryRow("SELECT name, email FROM user WHERE id = ?", id)
	data.Id = id
	data.Token = id.String()
	err = row.Scan(&data.Name, &data.Email)
	return
}

func (s *SqlUserDatabase) ChangeName(id room.UserId, newName string) error {
	panic("implement me")
}

func (s *SqlUserDatabase) Authenticate(email string, password string) (user room.UserData, err error) {
	row := s.db.QueryRow("SELECT id, email, name, secret FROM user WHERE email = ?", email)
	var hash string
	err = row.Scan(&user.Id, &user.Email, &user.Name, &hash)
	if err != nil {
		return
	}
	if !security.VerifyPassword(password, hash) {
		err = errors.New("authentication failed")
		return
	}
	user.Token = user.Id.String()
	return
}

func (s *SqlUserDatabase) VerifyToken(token string) (user room.UserData, err error) {
	row := s.db.QueryRow("SELECT id, email, name FROM user WHERE id = ?", token)
	err = row.Scan(&user.Id, &user.Email, &user.Name)
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
