package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"home-deck-tower/internal"
	"home-deck-tower/internal/model"

	"github.com/gorilla/websocket"
)

var up = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebSocketConnect(w http.ResponseWriter, r *http.Request, espManager *internal.Esp) {
	c, err := up.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	var (
		registrationChan = make(chan string)
		writeChan        = make(chan string)
		errChan          = make(chan error)

		registeredDeviceId string
	)

	defer func() {
		if registeredDeviceId != "" {
			espManager.UnregisterEsp(registeredDeviceId)
		}

		if err = c.Close(); err != nil {
			fmt.Printf("failed to close channel: %s", err.Error())
		}
	}()

	go func(c *websocket.Conn) {
		for {
			mt, message, err := c.ReadMessage()
			if err != nil {
				errChan <- err
				return
			}

			fmt.Printf("> Received a message from client: %s\n", string(message))

			var ev model.EspEvent
			err = json.Unmarshal(message, &ev)
			if err != nil {
				log.Printf("err readin: %v\n", err)
				continue
			}

			if ev.Type == model.Registration {
				regEv := ev.AsRegistrationEvent()
				espManager.RegisterEsp(regEv.DeviceId, writeChan)

				registrationChan <- regEv.DeviceId
				err = c.WriteMessage(mt, []byte("reg ok"))
				if err != nil {
					fmt.Printf("[WARN] error writing registration response")
				}
			}
		}
	}(c)

	for {
		select {
		case registeredDeviceId = <-registrationChan:
			fmt.Printf("Registered device on ESP manager [%s]\n", registeredDeviceId)

		case message := <-writeChan:
			fmt.Printf("< Sending message to client: %s\n", message)
			err = c.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				errChan <- err
			}

		case err := <-errChan:
			fmt.Printf("received error: %s\n", err.Error())
			return
		}
	}
}
