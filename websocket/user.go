package websocket

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

var (
	pongWait     = 10 * time.Second
	pingInterval = (pongWait * 9) / 10
)

type UserList map[*User]bool

type User struct {
	id         string
	connection *websocket.Conn
	manager    *Manager
	//egress is used to avoid concurrent writes on the ws conn
	egress chan ResponseEvent
}

func NewUser(conn *websocket.Conn, manager *Manager, id string) *User {
	return &User{
		id:         id,
		connection: conn,
		manager:    manager,
		egress:     make(chan ResponseEvent, 1),
	}
}

func (u *User) readPosition() {
	defer func() {
		//clean up connection
		u.manager.removeUser(u)
	}()

	if err := u.connection.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		log.Printf("error occured waiting on pong msg: %v", err)
		return
	}

	u.connection.SetReadLimit(1024)
	u.connection.SetPongHandler(u.pongHandler)
	for {
		var request RequestEvent
		err := u.connection.ReadJSON(&request)

		if err != nil {
			//when connection closed
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error reading message(format : %T (%+v)): %v", request, request, err)
			}
			break
		}

		if err := u.manager.routeEvent(request, u); err != nil {
			log.Println("error occured: ", err)
		}
	}
}

func (u *User) writePosition() {
	defer func() {
		u.manager.removeUser(u)
	}()

	ticker := time.NewTicker(pingInterval)
	for {
		select {
		case resEvent, ok := <-u.egress:
			if !ok {
				if err := u.connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Println("Connection Closed: ", err)
				}
				return
			}
			data, err := json.Marshal(resEvent)
			if err != nil {
				log.Println("Error occured: ", err)
				break
			}
			if err := u.connection.WriteMessage(websocket.TextMessage, data); err != nil {
				log.Printf("failed to send message: %v", err)
			}
		case <-ticker.C:
			log.Println("ping")
			if err := u.connection.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				log.Printf("error occured sending ping msg: %v", err)
				return
			}
		}
	}
}

func (u *User) pongHandler(msg string) error {
	log.Println("pong")
	return u.connection.SetReadDeadline(time.Now().Add(pongWait))
}
