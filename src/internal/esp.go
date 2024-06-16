package internal

import (
	"errors"
	"fmt"
	"home-deck-tower/internal/model"
	"sync"
)

type Esp struct {
	mu      sync.Mutex
	entries map[string]chan string
}

func (e *Esp) RegisterEsp(deviceId string, newChan chan string) {
	e.mu.Lock()
	if e.entries == nil {
		e.entries = make(map[string]chan string)
	}

	espChan, ok := e.entries[deviceId]
	if ok {
		close(espChan)
	}

	e.entries[deviceId] = newChan
	e.mu.Unlock()
}

func (e *Esp) UnregisterEsp(deviceId string) {
	e.mu.Lock()
	espChan, ok := e.entries[deviceId]
	if ok {
		close(espChan)
	}
	delete(e.entries, deviceId)
	e.mu.Unlock()
}

func (e *Esp) SendMessage(deviceIdDest, message string) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	channel, ok := e.entries[deviceIdDest]
	if !ok {
		return fmt.Errorf("no device with id [%s] found", deviceIdDest)
	}

	if message != model.Open {
		return errors.ErrUnsupported
	}

	openRequest := model.NewCommandEvent("open")
	channel <- string(openRequest.Serialize())

	stringErr, ok := <-channel
	if !ok {
		return errors.New("channel is closed")
	}

	if stringErr != "ok" {
		return fmt.Errorf("error sending message: %s", stringErr)
	}
	return nil
}
