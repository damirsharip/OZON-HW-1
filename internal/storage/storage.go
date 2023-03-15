package storage

import (
	"fmt"
	"log"
	"strconv"

	"github.com/pkg/errors"
)

var data map[uint]*User

var UserExists = errors.New("user exists")
var UserNotExist = errors.New("user does not exist")

func init() {
	log.Println("init storage")
	data = make(map[uint]*User)
	u, err := NewUser("name", "qwerty123")
	if err != nil {
		fmt.Println(err.Error())
	}
	data[u.GetId()] = u
}

func List() []*User {
	res := make([]*User, 0, len(data))
	for _, v := range data {
		res = append(res, v)
	}

	return res
}

func Add(u *User) error {
	if _, ok := data[u.GetId()]; ok {
		return errors.Wrap(UserExists, strconv.FormatUint(uint64(u.GetId()), 10))
	}
	data[u.GetId()] = u
	return nil
}

func Update(u *User) error {
	if _, ok := data[u.GetId()]; !ok {
		return errors.Wrap(UserNotExist, strconv.FormatUint(uint64(u.GetId()), 10))
	}
	data[u.GetId()] = u
	return nil
}

func Delete(id uint) error {
	if _, ok := data[id]; !ok {
		return errors.Wrap(UserNotExist, strconv.FormatUint(uint64(id), 10))
	}
	delete(data, id)
	return nil
}
