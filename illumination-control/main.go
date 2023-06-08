package main

import (
	"os"

	"github.com/charmbracelet/log"
)

func main() {
	openedPort, err := initSerial()
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	defer openedPort.Close()

	initServer()
}
