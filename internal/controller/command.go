package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"home-deck-tower/internal/model"
)

func SendCommand(w http.ResponseWriter, r *http.Request, sendMessage func(deviceIdDest, message string) error) {
	log.Printf("Staring http request: %s\n", r.Host)

	defer r.Body.Close()

	if r.Method != http.MethodPost {
		http.Error(w, "invalid method", http.StatusBadRequest)
		return
	}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("failed to read request body: %s\n", err.Error())
		http.Error(w, "please contact support", http.StatusInternalServerError)
		return
	}

	var in model.Input
	err = json.Unmarshal(reqBody, &in)
	if err != nil {
		log.Printf("failed to parse request body: %s\n", err.Error())
		http.Error(w, "please contact support", http.StatusInternalServerError)
		return
	}

	var response model.Output
	log.Printf("Sending command \"%s\" to device [%s]\n", in.Command, in.DeviceId)
	if err = sendMessage(in.DeviceId, in.Command); err != nil {
		log.Printf("failed to send command: %s\n", err.Error())

		response = model.Output{
			DeviceId: in.DeviceId,
			Command:  in.Command,
			Error:    err.Error(),
		}

		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write(response.Serialize())
		if err != nil {
			log.Printf("failed to write message body: %s\n", err.Error())
		}

		return
	}

	log.Println("Command sent successfully")

	response = model.Output{
		DeviceId: in.DeviceId,
		Command:  in.Command,
		Message:  "Lock was opened",
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response.Serialize())
}
