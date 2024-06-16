package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

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
		ctx, cancel      = context.WithCancel(r.Context())
		registrationChan = make(chan string)
		writeChan        = make(chan string)

		registeredDeviceId string
	)

	defer func() {
		if registeredDeviceId != "" {
			espManager.UnregisterEsp(registeredDeviceId)
		}

		if err = c.Close(); err != nil {
			fmt.Printf("failed to close channel: %s", err.Error())
		}

		close(registrationChan)
		// close(writeChan)
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("ending go routine")
				return

			default:
				mt, message, err := c.ReadMessage()
				if err != nil {
					log.Printf("Failed to read message: %s\n", err.Error())
					cancel()
					return
				}

				fmt.Printf("> Received a message from client: %s\n", string(message))

				var ev model.EspEvent
				err = json.Unmarshal(message, &ev)
				if err != nil {
					log.Printf("err readin: %v\n", err)
					continue
				}

				switch ev.Type {
				case model.EventTypeRegistration:
					regEv := ev.AsRegistrationEvent()
					espManager.RegisterEsp(regEv.DeviceId, writeChan)

					registrationChan <- regEv.DeviceId

					regEvResponse := ev
					regEvResponse.UnixTimestamp = int(time.Now().Unix())
					regEvResponse.Data = "ok"

					response, _ := json.Marshal(regEvResponse)
					err = c.WriteMessage(mt, response)
					if err != nil {
						fmt.Printf("[WARN] error writing registration response")
					}

				case model.EventTypeCommand:
					commEv := ev.AsCommandEvent()
					log.Printf("Response from command \"%s\": %s\n", commEv.Verb, commEv.Message)
				}
			}
		}
	}()

	for {
		select {
		case registeredDeviceId = <-registrationChan:
			fmt.Printf("Registered device on ESP manager [%s]\n", registeredDeviceId)

		case message := <-writeChan:
			fmt.Printf("< Sending message to client: %s\n", message)
			err = c.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				writeChan <- err.Error()
				log.Printf("Failed to send message: %s\n", err.Error())
				cancel()
			}
			writeChan <- "ok"

		case <-ctx.Done():
			fmt.Println("Context deadline")
			return
		}
	}
}
