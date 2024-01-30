package websocket

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type Manager struct {
	users UserList
	sync.RWMutex
}

func NewManager() *Manager {
	return &Manager{
		users: make(UserList),
	}
}

func (m *Manager) ServeWS(w http.ResponseWriter, r *http.Request) {
	log.Println("new connection")

	//upgrade regular http conn to websocket
	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	user := NewUser(conn, m)

	m.addUser(user)

	// start User process
	go user.readPosition()
	go user.writePosition()
}

func (m *Manager) addUser(User *User) {
	m.Lock()
	defer m.Unlock()
	m.users[User] = true
}

func (m *Manager) removeUser(user *User) {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.users[user]; ok {
		user.connection.Close()
		delete(m.users, user)
	}
}
