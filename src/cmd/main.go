package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"home-deck-tower/internal"
	"home-deck-tower/web"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "0.0.0.0:8080", "http service address")

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var EspManager internal.Esp

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
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
			EspManager.UnregisterEsp(registeredDeviceId)
		}

		c.Close()
	}()

	go func(c *websocket.Conn) {
		for {
			mt, message, err := c.ReadMessage()
			if err != nil {
				errChan <- err
				return
			}

			fmt.Printf("> Received a message from client: %s\n", string(message))

			var ev internal.EspEvent
			err = json.Unmarshal(message, &ev)
			if err != nil {
				log.Printf("err readin: %v\n", err)
				continue
			}

			if ev.Type == internal.Registration {
				regEv := ev.AsRegistrationEvent()
				EspManager.RegisterEsp(regEv.DeviceId, writeChan)

				registrationChan <- regEv.DeviceId
				c.WriteMessage(mt, []byte("registered successfully"))
			}
		}
	}(c)

	for {
		select {
		case registeredDeviceId = <-registrationChan:
			fmt.Printf("Registered device on ESP manager [%s]\n", registeredDeviceId)

		case message := <-writeChan:
			fmt.Printf("< Sending message to client: %s\n", message)
			c.WriteMessage(websocket.TextMessage, []byte(message))

		case err := <-errChan:
			fmt.Printf("received error: %s\n", err.Error())
			return
		}
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://"+r.Host+"/echo")
}

type input struct {
	DeviceId string `json:"device"`
	Command  string `json:"command"`
}

func sendCommand(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Printf("method %s not accepted\n", r.Method)
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Failed to read request body: %s\n", err.Error())
		return
	}

	var in input
	err = json.Unmarshal(reqBody, &in)
	if err != nil {
		fmt.Printf("Failed to parse request body: %s\n", err.Error())
		return
	}

	fmt.Printf("Sending command \"%s\" to device [%s]\n", in.Command, in.DeviceId)
	err = EspManager.SendMessage(in.DeviceId, in.Command)
	if err != nil {
		fmt.Printf("Failed to send command: %s\n", err.Error())
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/command", sendCommand)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

var homeTemplate = template.Must(template.New("").Parse(web.IndexHtml))
