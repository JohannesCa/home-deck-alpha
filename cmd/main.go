package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"home-deck-tower/internal"
	"home-deck-tower/internal/controller"
)

var addr = flag.String("addr", "0.0.0.0:8080", "http service address")

var EspManager internal.Esp

func main() {
	flag.Parse()
	log.SetFlags(0)

	http.HandleFunc("/device/connect", func(w http.ResponseWriter, r *http.Request) {
		controller.WebSocketConnect(w, r, &EspManager)
	})
	http.HandleFunc("/command", func(w http.ResponseWriter, r *http.Request) {
		controller.SendCommand(w, r, EspManager.SendMessage)
	})

	// http.HandleFunc("/", home)
	fmt.Printf(":: Started serving on %s\n\n", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal(err)
	}
}

func HttpEspManagerWrapper(w http.ResponseWriter, r *http.Request) {

}

// var homeTemplate = template.Must(template.New("").Parse(web.IndexHtml))

// func home(w http.ResponseWriter, r *http.Request) {
// 	homeTemplate.Execute(w, "ws://"+r.Host+"/webSocketConnect")
// }
