package main

import (
	"log"

	botPkg "HW-1/internal/pkg/bot"
	cmdAddPkg "HW-1/internal/pkg/bot/command/add"
	cmdHelpPkg "HW-1/internal/pkg/bot/command/help"
	userPkg "HW-1/internal/pkg/core/user"
)

func main() {
	var user userPkg.Interface
	{
		userPkg.New()
	}

	var bot botPkg.Interface
	{
		bot = botPkg.MustNew()

		commandAdd := cmdAddPkg.New(user)
		bot.RegisterHandler(commandAdd)

		commandHelp := cmdHelpPkg.New(map[string]string{
			commandAdd.Name(): commandAdd.Description(),
		})
		bot.RegisterHandler(commandHelp)

	}

	if err := bot.Run(); err != nil {
		log.Panic(err)
	}

}
