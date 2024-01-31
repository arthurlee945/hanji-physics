package client

import (
	"context"
	"log"
	"net/http"

	"github.com/arthurlee945/hanji-physics/websocket"
)

func setupAPI() {
	ctx := context.Background()
	manager := websocket.NewManager(ctx)

	http.Handle("/", http.FileServer(http.Dir("./client")))
	http.HandleFunc("/ws", manager.ServeWS)
	http.HandleFunc("/authorize", manager.AuthenticationHandler)
}

func Start() {
	setupAPI()
	log.Fatal(http.ListenAndServeTLS("localhost:8080", "server.crt", "server.key", nil))
}
