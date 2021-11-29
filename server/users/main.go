package users

import "strconv"

type Id int64

const InvalidId Id = -1

func (t Id) String() string {
	return strconv.FormatInt(int64(t), 10)
}

func ParseId(idStr string) (Id, error) {
	idInt, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return InvalidId, err
	}
	return Id(idInt), nil
}

type AccountData struct {
	Id    Id
	Name  string
	Email string
	Token string
}

func NewUserData(id Id, name string, email string, token string) AccountData {
	return AccountData{id, name, email, token}
}

func (d AccountData) String() string {
	return d.Name
}

type Users interface {
	Add(email string, password string, name string) (AccountData, error)
	ListIds() ([]Id, error)
	GetData(id Id) (AccountData, error)
	ChangeName(id Id, newName string) error
  ChangePassword(id Id, newPassword string) error
	VerifyToken(token string) (AccountData, error)
	Authenticate(email string, password string) (AccountData, error)
}
