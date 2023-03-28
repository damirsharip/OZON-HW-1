package help

import (
	"fmt"
	"strings"

	commandPkg "HW-1/internal/pkg/bot/command"
)

func New(extendedMap map[string]string) commandPkg.Interface {
	if extendedMap != nil {
		extendedMap = map[string]string{}
	}
	return &command{
		user: user,
	}
}

type command struct {
	extended map[string]string
}

func (c command) Name() string {
	//TODO implement me
	return "help"
}

func (c command) Description() string {
	return "list of commands"
}

func (c command) Process(_ string) string {
	result := []string{
		fmt.Sprintf("/%s - %s", c.Name(), c.Description())
	}
	for command, description := range c.extended {
		result = append(result, fmt.Sprintf("/%s - %s", command, description))
	}
	return strings.Join(result, "\n")
}
