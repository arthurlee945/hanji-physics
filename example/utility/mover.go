package utility

import (
	"math"
	"sync"

	"github.com/arthurlee945/suhag/vec"
	"github.com/fzipp/canvas"
)

type Mover struct {
	Loc          *vec.Vec2
	Velocity     *vec.Vec2
	Acceleration *vec.Vec2
	Mass         float64
	Angle        *Angle
}

func NewMover(x, y, mass float64) *Mover {
	return &Mover{
		Loc:          &vec.Vec2{x, y},
		Velocity:     &vec.Vec2{0, 0},
		Acceleration: &vec.Vec2{0, 0},
		Mass:         mass,
		Angle:        &Angle{0, 0, 0.001},
	}
}

func (m *Mover) ApplyForce(force *vec.Vec2) {
	m.Acceleration.Add(*vec.Div(*force, m.Mass))
}

func (m *Mover) Update() {
	m.Velocity.Add(*m.Acceleration)
	m.Loc.Add(*m.Velocity)
	m.Acceleration.Mult(0)
}

func (m *Mover) Display(ctx *canvas.Context, mu *sync.RWMutex) {
	mu.Lock()
	defer mu.Unlock()
	ctx.BeginPath()
	ctx.Arc(m.Loc[0], m.Loc[1], m.Mass, 0, math.Pi*2, false)
	ctx.Fill()
}

func (m *Mover) CheckEdges(x, y float64) {
	if m.Loc[0] > x {
		m.Loc[0] = x
		m.Velocity[0] *= -1
	} else if m.Loc[0] < 0 {
		m.Velocity[0] *= -1
		m.Loc[0] = 0
	}

	if m.Loc[1] > y {
		m.Velocity[1] *= -1
		m.Loc[1] = y
	} else if m.Loc[1] < 0 {
		m.Velocity[1] *= -1
		m.Loc[1] = 0
	}
}
