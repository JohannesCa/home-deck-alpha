package internal

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type EventType string

const (
	Registration EventType = "registration"
)

type EspEvent struct {
	Type      EventType  `json:"event"`
	Timestamp CustomTime `json:"timestamp"`
	Data      string     `json:"data"`
}

/*
{
	"event":"registration",
	"timestamp":1716513384,
	"data":"{\"device_id\":\"1\",\"device_name\":\"test\"}"
}
*/

func (e EspEvent) AsRegistrationEvent() (ev RegistrationEvent) {
	json.Unmarshal([]byte(e.Data), &ev)
	return
}

type RegistrationEvent struct {
	DeviceId   string `json:"device_id"`
	DeviceName string `json:"device_name"`
	Location   struct {
		Name string `json:"name"`
		Lat  string `json:"lat"`
		Long string `json:"long"`
	} `json:"location"`
}

type CustomTime struct {
	time.Time
}

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	n, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return err
	}
	ct.Time = time.Unix(n, 0)
	return
}

func (ct *CustomTime) MarshalJSON() ([]byte, error) {
	if ct.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(time.RFC3339Nano))), nil
}

func (ct *CustomTime) String() string {
	return ct.Time.Format(time.RFC3339)
}
