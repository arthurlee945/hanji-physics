package websocket

func handlePointerMoveEvent(evt RequestEvent, u *User) error {
	pointerMoveEvent := ResponseEvent{
		UserId:   u.id,
		Type:     PointerMoveEvent,
		Position: evt.Position,
		Matrix:   []byte("matrix"),
	}
	for user := range u.manager.users {
		if u == user {
			continue
		}
		user.egress <- pointerMoveEvent
	}
	return nil
}

func handleStartEvent(evt RequestEvent, u *User) error {
	startEvent := ResponseEvent{
		UserId: u.id,
		Type:   StartEvent,
	}
	u.egress <- startEvent
	return nil
}
