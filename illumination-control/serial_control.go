package main

import (
	"fmt"

	"github.com/charmbracelet/log"
	"go.bug.st/serial"
)

var mode = &serial.Mode{
	BaudRate: 9600,
}

var openedPort serial.Port

func getPortsList() ([]string, error) {
	ports, err := serial.GetPortsList()
	if err != nil {
		return nil, fmt.Errorf("failed to get ports list: %w", err)
	}

	if len(ports) == 0 {
		return nil, fmt.Errorf("no ports found")
	}

	log.Info("found ports", "ports", ports)

	return ports, nil
}

func openPort(portName string) (serial.Port, error) {
	port, err := serial.Open(portName, mode)
	if err != nil {
		return nil, fmt.Errorf("failed to open port: %w", err)
	}

	log.Info("opened port", "port", portName)

	return port, nil
}

func listenSerial() {
	buf := make([]byte, 128)

	for {
		n, err := openedPort.Read(buf)
		if err != nil {
			log.Error("failed to read from port", "err", err)
			continue
		}

		log.Info("received message", "msg", string(buf[:n]))
	}
}

func initSerial() (serial.Port, error) {
	ports, err := getPortsList()
	if err != nil {
		return nil, err
	}

	openedPort, err = openPort(ports[0])
	if err != nil {
		return nil, err
	}

	go listenSerial()

	return openedPort, nil
}

func sendSerial(msg []byte) error {
	_, err := openedPort.Write(msg)
	if err != nil {
		return fmt.Errorf("failed to write to port: %w", err)
	}

	log.Info("sent message", "msg", string(msg))

	return nil
}

func setLightsModeState(state state) error {
	msg := []byte(fmt.Sprintf("lights_mode %s\n", state))
	err := sendSerial(msg)
	if err != nil {
		return fmt.Errorf("failed to set lights mode state: %w", err)
	}

	return nil
}

func setLightsState(state state) error {
	msg := []byte(fmt.Sprintf("lights %s\n", state))
	err := sendSerial(msg)
	if err != nil {
		return fmt.Errorf("failed to set lights state: %w", err)
	}

	return nil
}

func setCurtainsState(state state) error {
	msg := []byte(fmt.Sprintf("curtains %s\n", state))
	err := sendSerial(msg)
	if err != nil {
		return fmt.Errorf("failed to set curtains state: %w", err)
	}

	return nil
}

func setBlindsState(state state) error {
	msg := []byte(fmt.Sprintf("blinds %s\n", state))
	err := sendSerial(msg)
	if err != nil {
		return fmt.Errorf("failed to set blinds state: %w", err)
	}

	return nil
}
