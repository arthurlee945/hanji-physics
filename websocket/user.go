package websocket

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type UserPosition map[string]float64

type UserList map[*User]bool

type User struct {
	connection *websocket.Conn
	manager    *Manager

	//egress is used to avoid concurrent writes on the ws conn
	egress chan []byte
}

func NewUser(conn *websocket.Conn, manager *Manager) *User {
	return &User{
		connection: conn,
		manager:    manager,
		egress:     make(chan []byte),
	}
}

func (u *User) readPosition() {
	defer func() {
		//clean up connection

		u.manager.removeUser(u)
	}()
	for {
		var uPosition map[string]float64
		err := u.connection.ReadJSON(&uPosition)

		if err != nil {
			//when connection closed
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error reading message(format : {x:float64, y:float64}): %v", err)
			}
			break
		}

		for wsUser := range u.manager.users {
			posBytes, err := json.Marshal(uPosition)
			if err != nil {
				log.Printf("error serializing json user position data: %v", err)
			}
			wsUser.egress <- posBytes
		}

		x, xOk := uPosition["x"]
		y, yOk := uPosition["y"]
		if !xOk || !yOk {
			log.Printf("Not all required data passed required x and y position")
		}
		log.Printf("x:%f, y: %f", x, y)

	}
}

func (u *User) writePosition() {
	defer func() {
		u.manager.removeUser(u)
	}()
	for {
		select {
		case payload, ok := <-u.egress:
			if !ok {
				if err := u.connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Println("Connection Closed: ", err)
				}
				return
			}
			if err := u.connection.WriteMessage(websocket.TextMessage, payload); err != nil {
				log.Printf("failed to send message: %v", err)
			}
			log.Println("message sent")

		}
	}
}
