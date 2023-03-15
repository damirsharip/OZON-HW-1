package storage

import "fmt"

var lastId = uint(0)

type User struct {
	id       uint
	name     string
	password string
}

func NewUser(name, password string) (*User, error) {
	u := User{}
	err := u.SetName(name)
	if err != nil {
		return nil, err
	}

	err = u.SetPassword(password)
	if err != nil {
		return nil, err
	}
	lastId++
	u.id = lastId
	return &u, nil
}

func (u *User) SetName(name string) error {
	if len(name) == 0 || len(name) > 10 {
		return fmt.Errorf("bad name <%v>", name)
	}
	u.name = name
	return nil
}

func (u *User) SetPassword(pswd string) error {
	if len(pswd) < 6 || len(pswd) > 10 {
		return fmt.Errorf("bad password <%v>", pswd)
	}
	u.password = pswd
	return nil
}

func (u User) String() string {
	return fmt.Sprintf("%d: %s / %s", u.id, u.name, u.password)
}

func (u User) GetName() string {
	return u.name
}

func (u User) GetPassword() string {
	return u.password
}

func (u User) GetId() uint {
	return u.id
}
