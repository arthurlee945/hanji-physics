package websocket

import "github.com/arthurlee945/hanji-physics/engine/canvas"

func handleWalkerEvent(evt RequestEvent, u *User) error {
	return nil
}

func handle2DEngineEvent(evt RequestEvent, u *User) error {
	matrix := u.manager.engine.Canvas.(*canvas.Canvas2D).Matrix
	for user := range u.manager.users {
		engineEvent := &EngineResponseEvent{
			Type:   EngineEvent,
			Matrix: matrix,
		}
		user.egress <- engineEvent
	}
	return nil
}
