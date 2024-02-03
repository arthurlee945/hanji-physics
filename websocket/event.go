package websocket

type EventType string

const (
	//USER EVENT
	StartEvent           = "start"
	PointerPositionEvent = "pointerposition"
	PointerDownEvent     = "pointerdown"
	PointerUpEvent       = "pointerup"
	PointerLeaveEvent    = "pointerleave"
	//PHYSICS EVENT
	WalkerEvent = "walker"
	EngineEvent = "engine"
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

type ResponseEvent interface {
	isResponseEvent()
}

type UserResponseEvent struct {
	UserId   string       `json:"userId"`
	Type     EventType    `json:"type"`
	Position UserPosition `json:"position"`
}

func (ure *UserResponseEvent) isResponseEvent() {}

type EngineResponseEvent struct {
	Type   EventType   `json:"type"`
	Matrix interface{} `json:"matrix"`
}

func (ere *EngineResponseEvent) isResponseEvent() {}

type EventHandler func(evt RequestEvent, u *User) error
