package websocket

type EventType string

const (
	StartEvent        = "start"
	PointerMoveEvent  = "pointermove"
	PointerDownEvent  = "pointerdown"
	PointerUpEvent    = "pointerup"
	PointerLeaveEvent = "mousedown"
	MatrixEvent       = "test"
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

// TODO: update matrix reflect correct data type
type ResponseEvent struct {
	UserId   string       `json:"userId"`
	Type     EventType    `json:"type"`
	Position UserPosition `json:"position"`
	Matrix   []byte       `json:"matrix"`
}

type EventHandler func(evt RequestEvent, u *User) error
