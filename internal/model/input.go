package model

const (
	Healthcheck = "HEALTHCHECK"
	Open        = "OPEN"
)

type Input struct {
	DeviceId string            `json:"id"`
	Command  string            `json:"command"`
	Params   map[string]string `json:"params"`
}
