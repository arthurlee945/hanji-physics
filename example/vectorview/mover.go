package vectorview

import (
	"math"

	"github.com/arthurlee945/hanji-physics/suhag/vec"
	"github.com/fzipp/canvas"
)

type Mover struct {
	loc   *vec.Vec2
	speed *vec.Vec2
}

func NewMover(x, y float64) *Mover {
	return &Mover{
		loc:   &vec.Vec2{x, y},
		speed: &vec.Vec2{1, 3.3},
	}
}

func (m *Mover) Move(ctx *canvas.Context) {
	m.loc.Add(*m.speed)
	m.CheckEdges(ctx.CanvasWidth(), ctx.CanvasHeight())
	ctx.Arc(m.loc[0], m.loc[0], 16, 0, math.Pi*2, false)
	ctx.Fill()
}

func (m *Mover) CheckEdges(x, y int) {
	if m.loc[0] > float64(x) && m.loc[0] > 0 || m.loc[0] < float64(x) && m.loc[0] < 0 {
		m.loc[0] *= -1
	}
	if m.loc[1] > float64(y) && m.loc[1] > 0 || m.loc[1] < float64(y) && m.loc[1] < 0 {
		m.loc[1] *= -1
	}
}
