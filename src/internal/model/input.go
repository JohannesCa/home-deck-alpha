package model

type Command string

const (
	Healthcheck Command = "HEALTHCHECK"
	Open        Command = "OPEN"
)

type Input struct {
	DeviceId string            `json:"id"`
	Command  string            `json:"command"`
	Params   map[string]string `json:"params"`
}
