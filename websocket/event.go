package websocket

type EventType string

const (
	//USER EVENT
	StartEvent        = "start"
	PointerMoveEvent  = "pointermove"
	PointerDownEvent  = "pointerdown"
	PointerUpEvent    = "pointerup"
	PointerLeaveEvent = "mousedown"
	//PHYSICS EVENT
	WalkerEvent = "walker"
)

type UserPosition struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type RequestEvent struct {
	UserId   string       `json:"userId"`
	Type     EventType    `json:"type"`
	Position UserPosition `json:"position"`
}

type ResponseEvent struct {
	UserId   string       `json:"userId"`
	Type     EventType    `json:"type"`
	Position UserPosition `json:"position"`
	Matrix   interface{}  `json:"matrix"`
}

type EventHandler func(evt RequestEvent, u *User) error
