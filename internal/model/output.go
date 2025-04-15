package model

import "encoding/json"

type Output struct {
	DeviceId string `json:"id"`
	Command  string `json:"command"`
	Message  string `json:"message,omitempty"`
	Error    string `json:"error,omitempty"`
}

func (o Output) Serialize() (ret []byte) {
	ret, _ = json.Marshal(o)
	return
}
