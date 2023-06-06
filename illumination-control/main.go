package main

import (
	"os"

	"github.com/charmbracelet/log"
)

func main() {
	err := initSerial()
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	initServer()
}
