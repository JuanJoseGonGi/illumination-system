package main

import (
	"io"

	"github.com/charmbracelet/log"
)

func closeOrLog(clsr io.Closer) {
	if clsr == nil {
		return
	}

	err := clsr.Close()
	if err != nil {
		log.Error("failed to close", "err", err)
	}
}
