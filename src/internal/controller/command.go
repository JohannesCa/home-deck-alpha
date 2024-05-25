package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
	"home-deck-tower/internal/model"
)

func SendCommand(w http.ResponseWriter, r *http.Request, sendMessage func(deviceIdDest, message string) error) {
	logger := log.With().Str("request_id", r.Header.Get("request-id")).Logger()
	logger.Info().Msgf("Staring http request: %s", r.Host)

	defer r.Body.Close()

	if r.Method != http.MethodPost {
		http.Error(w, "invalid method", http.StatusBadRequest)
		return
	}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Error().Err(err).Msg("failed to read request body")
		http.Error(w, "please contact support", http.StatusInternalServerError)
		return
	}

	var in model.Input
	err = json.Unmarshal(reqBody, &in)
	if err != nil {
		logger.Error().Err(err).Msg("failed to parse request body")
		http.Error(w, "please contact support", http.StatusInternalServerError)
		return
	}

	var response model.Output
	logger.Debug().Msgf("Sending command \"%s\" to device [%s]\n", in.Command, in.DeviceId)
	if err = sendMessage(in.DeviceId, in.Command); err != nil {
		logger.Error().Err(err).Msg("failed to send command")

		response = model.Output{
			DeviceId: in.DeviceId,
			Command:  in.Command,
			Error:    err.Error(),
		}

		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write(response.Serialize())
		if err != nil {
			logger.Error().Err(err).Msg("failed to write message body")
		}

		return
	}

	logger.Info().Msgf("Command sent successfully")

	response = model.Output{
		DeviceId: in.DeviceId,
		Command:  in.Command,
		Message:  "Lock was opened",
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response.Serialize())
}
