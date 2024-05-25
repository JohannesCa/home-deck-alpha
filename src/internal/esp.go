package internal

import (
	"errors"
	"fmt"
)

type espChan chan string

type Esp struct {
	entries map[string]espChan
}

func (e *Esp) RegisterEsp(deviceId string, messageChan espChan) {
	if e.entries == nil {
		e.entries = make(map[string]espChan)
	}

	_, ok := e.entries[deviceId]
	if !ok {
		e.entries[deviceId] = messageChan
	}
}

func (e *Esp) UnregisterEsp(deviceId string) {
	delete(e.entries, deviceId)
}

func (e *Esp) SendMessage(deviceIdDest, message string) error {
	channel, ok := e.entries[deviceIdDest]
	if !ok {
		return errors.New(fmt.Sprintf("no device with id [%s] found", deviceIdDest))
	}

	channel <- message
	return nil
}
