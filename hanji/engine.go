package hanji

import (
	"github.com/fzipp/canvas"
)

type EnginePart interface {
	Draw(*canvas.Context)
	Handle(canvas.Event)
}
