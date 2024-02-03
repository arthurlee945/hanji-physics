package websocket

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/arthurlee945/hanji-physics/engine"
	"github.com/gorilla/websocket"
)

var (
	websocketUpgrader = websocket.Upgrader{
		CheckOrigin:     checkOrigin,
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type Manager struct {
	users    UserList
	engine   *engine.Engine
	mu       sync.RWMutex
	otps     RetentionMap
	handlers map[EventType]EventHandler
}

func NewManager(ctx context.Context) *Manager {
	engine := engine.NewEngine(engine.With2DCanvas(250, 250))
	m := &Manager{
		users:    make(UserList),
		engine:   engine,
		handlers: make(map[EventType]EventHandler),
		otps:     NewRetentionMap(ctx, 5*time.Second),
	}
	m.setupEventHandlers()
	return m
}

func (m *Manager) AuthenticationHandler(w http.ResponseWriter, r *http.Request) {
	type response struct {
		OTP string `json:"otp"`
	}
	token := r.Header.Get("token")
	if token == "" {
		http.Error(w, "no token passed", http.StatusBadRequest)
	}
	//TODO: NEED TO USE DYNAMIC INSTEAD OF HARD CODED
	if token != "MonkeySaysHi" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
	}
	otp := m.otps.NewOTP()

	resp := response{
		OTP: otp.Key,
	}
	data, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (m *Manager) ServeWS(w http.ResponseWriter, r *http.Request) {
	otp := r.URL.Query().Get("otp")
	if otp == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if !m.otps.VerifyOTP(otp) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	//upgrade regular http conn to websocket
	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	user := NewUser(conn, m, otp)

	m.addUser(user)

	// start User process
	go user.readPosition()
	go user.writePosition()

	log.Println("new connection - ", otp)
}

func (m *Manager) broadcastEngineMatrix() {
	// ticker := time.NewTicker(25 * time.Millisecond);
	// for {
	// 	<-ticker.C

	// }
}

func (m *Manager) setupEventHandlers() {
	m.handlers[PointerPositionEvent] = handlePointerPositionEvent
	m.handlers[StartEvent] = handleStartEvent
	//physics event
	m.handlers[WalkerEvent] = handleWalkerEvent
	m.handlers[EngineEvent] = handle2DEngineEvent
}

func (m *Manager) routeEvent(evt RequestEvent, u *User) error {
	if handler, ok := m.handlers[evt.Type]; ok {
		if err := handler(evt, u); err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("there is no such event type. got=%s", evt.Type)
}

func (m *Manager) addUser(User *User) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.users[User] = true
}

func (m *Manager) removeUser(user *User) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.users[user]; ok {
		user.connection.Close()
		delete(m.users, user)
	}
}

// TODO: UPDATE THIS TO USE ENV
func checkOrigin(r *http.Request) bool {
	origin := r.Header.Get("Origin")
	switch origin {
	case "https://localhost:8080":
		return true
	default:
		return false
	}
}
