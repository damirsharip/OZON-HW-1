package handlers

import (
	"fmt"
	"strconv"
	"strings"

	"HW-1/internal/commander"

	"github.com/pkg/errors"

	"HW-1/internal/storage"
)

var BadArgument = errors.New("bad argument")

const (
	listCmd   = "list"
	addCmd    = "add"
	deleteCmd = "delete"
	updateCmd = "update"
	helpCmd   = "help"
)

func listFunc(s string) string {
	data := storage.List()
	res := make([]string, 0, len(data))
	for _, v := range data {
		res = append(res, v.String())
	}
	return strings.Join(res, "\n")
}

func helpFunc(s string) string {
	return "/help - list commands\n" +
		"/list - list users\n" +
		"/add <name> <password> - add a new user with name and password\n"
}

func addFunc(data string) string {
	params := strings.Split(data, " ")
	if len(params) != 2 {
		return errors.Wrapf(BadArgument, "%d items: <%v>", len(params), params).Error()
	}
	u, err := storage.NewUser(params[0], params[1])
	if err != nil {
		return err.Error()
	}

	err = storage.Add(u)
	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("user %v added", u)
}

//func updateFunc(data string) string {
//	params := strings.Split(data, " ")
//	if len(params) != 2 {
//		return errors.Wrapf(BadArgument, "%d items: <%v>", len(params), params).Error()
//	}
//	u, err := storage.Update(params[0], params[1])
//	if err != nil {
//		return err.Error()
//	}
//
//	err = storage.Add(u)
//	if err != nil {
//		return err.Error()
//	}
//
//	return fmt.Sprintf("user %v added", u)
//}

func deleteFunc(s string) string {
	iid, _ := strconv.Atoi(s)
	id := uint(iid)

	err := storage.Delete(id)
	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("user %v deleted", id)
}

func AddHandlers(c *commander.Commander) {
	c.RegisterHandler(helpCmd, helpFunc)
	c.RegisterHandler(listCmd, listFunc)
	c.RegisterHandler(addCmd, addFunc)
	//c.RegisterHandler(updateCmd, updateFunc)
	c.RegisterHandler(deleteCmd, deleteFunc)
}
