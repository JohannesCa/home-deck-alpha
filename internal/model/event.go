package model

import (
	"encoding/json"
	"time"
)

const (
	EventTypeRegistration string = "registration"
	EventTypeCommand      string = "command"
	EventTypeHeartbeat    string = "heartbeat"
)

type EspEvent struct {
	Type          string `json:"event"`
	UnixTimestamp int    `json:"timestamp"`
	Data          string `json:"data"`
}

func (e EspEvent) AsRegistrationEvent() (ev RegistrationEvent) {
	json.Unmarshal([]byte(e.Data), &ev)
	return
}

func (e EspEvent) AsCommandEvent() (ev CommandEvent) {
	json.Unmarshal([]byte(e.Data), &ev)
	return
}

func (e EspEvent) Serialize() (res []byte) {
	res, _ = json.Marshal(e)
	return
}

type RegistrationEvent struct {
	DeviceId   string `json:"device_id"`
	DeviceName string `json:"device_name"`
}

type CommandEvent struct {
	Verb    string `json:"command"`
	Message string `json:"message,omitempty"`
}

func NewCommandEvent(verb string) EspEvent {
	return EspEvent{
		Type:          EventTypeCommand,
		UnixTimestamp: int(time.Now().Unix()),
		Data:          verb,
	}
}
