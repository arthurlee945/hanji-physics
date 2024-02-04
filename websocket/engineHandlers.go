package websocket

func handleWalkerEvent(evt RequestEvent, u *User) error {
	return nil
}

func handleEngineEvent(evt RequestEvent, u *User) error {
	matrix := u.manager.engine.Canvas.Matrix
	for user := range u.manager.users {
		engineEvent := &EngineResponseEvent{
			Type:   EngineEvent,
			Matrix: matrix,
		}
		user.egress <- engineEvent
	}
	return nil
}
