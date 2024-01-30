package client

import (
	"log"
	"net/http"

	"github.com/arthurlee945/hanji-physics/websocket"
)

func setupAPI() {

	manager := websocket.NewManager()

	http.Handle("/", http.FileServer(http.Dir("./client")))
	http.HandleFunc("/ws", manager.ServeWS)
}

func Start() {
	setupAPI()
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
