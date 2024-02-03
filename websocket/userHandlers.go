package websocket

func handlePointerPositionEvent(evt RequestEvent, u *User) error {
	pointerPositionEvent := &UserResponseEvent{
		UserId:   u.id,
		Type:     PointerPositionEvent,
		Position: evt.Position,
	}
	for user := range u.manager.users {
		if u == user {
			continue
		}
		user.egress <- pointerPositionEvent
	}
	return nil
}

func handleStartEvent(evt RequestEvent, u *User) error {
	startEvent := &UserResponseEvent{
		UserId: u.id,
		Type:   StartEvent,
	}
	u.egress <- startEvent
	return nil
}
