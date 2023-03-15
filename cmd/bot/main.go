package main

import (
	"fmt"
	"log"

	"HW-1/internal/commander"
	"HW-1/internal/handlers"
)

func main() {
	fmt.Println("start main")
	cmd, err := commander.Init()
	if err != nil {
		log.Panic(err)
	}

	handlers.AddHandlers(cmd)

	if err = cmd.Run(); err != nil {
		log.Panic(err)
	}

}
