package engine

import "github.com/arthurlee945/hanji-physics/engine/canvas"

type EngineOpts func(*Engine)

type Setting struct {
	width  int
	height int
}

type Engine struct {
	Canvas *canvas.Canvas
	Setting
}

func NewEngine(opts ...EngineOpts) *Engine {
	engine := &Engine{}
	for _, opt := range opts {
		opt(engine)
	}
	return engine
}

func WithCanvas(x, y int) func(*Engine) {
	return func(engine *Engine) {
		engine.Canvas = canvas.NewCanvas(x, y)
	}
}

func SetSettings(setting Setting) func(*Engine) {
	return func(engine *Engine) {
		engine.Setting = setting
	}
}
