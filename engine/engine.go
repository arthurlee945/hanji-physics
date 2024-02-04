package engine

import "github.com/arthurlee945/hanji-physics/engine/canvas"

type EngineOpts func(*Engine)

type Engine struct {
	Canvas canvas.Canvas
}

func NewEngine(opts ...EngineOpts) *Engine {
	engine := &Engine{}
	for _, opt := range opts {
		opt(engine)
	}
	return engine
}

func With2DCanvas(x, y uint16) func(*Engine) {
	return func(engine *Engine) {
		engine.Canvas = canvas.NewCanvas2D(x, y)
	}
}

func With3DScene(x, y, z uint16) func(*Engine) {
	return func(engine *Engine) {
		engine.Canvas = canvas.NewCanvas3D(x, y, z)
	}
}
