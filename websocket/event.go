package websocket

type EventType string

const (
	EventHover = "hover"
	EventReply = "Reply from server"
)

// MOST LIKELY WONT USE
type UserPosition struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Event struct {
	Type     EventType    `json:"type"`
	Position UserPosition `json:"position"`
}

type EventHandler func(evt Event, u *User) error

type SendHoverEvent struct {
	Message string `json:"message"`
	FROM    string `json:"from"`
}

type NewHoverEvent struct {
	Position UserPosition
	Message  string `json:"message"`
}
