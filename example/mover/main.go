package main

import (
	"flag"
	"fmt"
	"image/color"
	"log"
	"math"
	"math/rand/v2"
	"time"

	"github.com/arthurlee945/suhag/force"
	"github.com/arthurlee945/suhag/vec"
	"github.com/fzipp/canvas"
)

func main() {
	http := flag.String("http", ":8080", "HTTP service address (e.g.. '127.0.0.1:8080' or ':8080')")
	flag.Parse()

	fmt.Println("Listening on " + httpLink(*http))
	err := canvas.ListenAndServe(*http, runCanvas, &canvas.Options{
		Title:          "Vector View",
		Width:          700,
		Height:         500,
		PageBackground: color.RGBA{R: 0xFA, G: 0xF9, B: 0xF6, A: 0xFF},
		EnabledEvents: []canvas.Event{
			canvas.MouseMoveEvent{},
		},
	})
	if err != nil {
		log.Fatalf("Failed on starting canvas server: %v", err)
	}
}

func runCanvas(ctx *canvas.Context) {
	ctx.SetFillStyle(color.RGBA{0x08, 0x08, 0x08, 0xff})

	engine := NewMoverView(float64(ctx.CanvasWidth()), float64(ctx.CanvasHeight()))

	for {
		select {
		case event := <-ctx.Events():
			if _, ok := event.(canvas.CloseEvent); ok {
				return
			}
			engine.Handle(event)
		default:
			engine.Draw(ctx)
			ctx.Flush()
			time.Sleep(5 * time.Millisecond)
		}
	}
}

func httpLink(addr string) string {
	if addr[0] == ':' {
		addr = "localhost" + addr
	}
	return "http://" + addr
}

type MoverView struct {
	movers [100]*Mover
	liquid *Liquid
}

func NewMoverView(w, h float64) *MoverView {
	movers := &MoverView{
		liquid: NewLiquid(0, h/2, w, h/2, 0.01),
	}
	for i := range 100 {
		movers.movers[i] = NewMover(0, 0, rand.Float64()*8+2)
	}
	return movers
}

func (mv *MoverView) Draw(ctx *canvas.Context) {
	width, height := float64(ctx.CanvasWidth()), float64(ctx.CanvasHeight())
	ctx.ClearRect(0, 0, width, height)
	ctx.BeginPath()
	ctx.Rect(0, 0, width, height)
	ctx.Stroke()

	mv.liquid.Display(ctx)
	wind := vec.NewVec2(0.01, 0)

	for _, m := range mv.movers {
		// c := 0.01
		// friction := force.Friction(m.velocity, c, 1)
		// m.ApplyForce(friction)
		if m.isIniside(*mv.liquid) {
			drag := force.Drag(m.velocity, 1, 1, mv.liquid.coefficient)
			m.ApplyForce(drag)
		}

		gravity := vec.NewVec2(0, 0.01*m.mass)
		m.ApplyForce(wind)
		m.ApplyForce(gravity)

		m.Move()
		m.Display(ctx)
		m.CheckEdges(width, height)
	}
}
func (mv *MoverView) Handle(evt canvas.Event) {}

type Mover struct {
	loc          *vec.Vec2
	velocity     *vec.Vec2
	acceleration *vec.Vec2
	mass         float64
}

func NewMover(x, y, mass float64) *Mover {
	return &Mover{
		loc:          &vec.Vec2{x, y},
		velocity:     &vec.Vec2{0, 0},
		acceleration: &vec.Vec2{0, 0},
		mass:         mass,
	}
}

func (m *Mover) ApplyForce(force *vec.Vec2) {
	m.acceleration.Add(*vec.Div(*force, m.mass))
}

func (m *Mover) Move() {
	m.velocity.Add(*m.acceleration)
	m.loc.Add(*m.velocity)
	m.acceleration.Mult(0)
}

func (m *Mover) Display(ctx *canvas.Context) {
	ctx.BeginPath()
	ctx.Arc(m.loc[0], m.loc[1], m.mass, 0, math.Pi*2, false)
	ctx.Fill()
}

func (m *Mover) CheckEdges(x, y float64) {
	if m.loc[0] > x {
		m.loc[0] = x
		m.velocity[0] *= -1
	} else if m.loc[0] < 0 {
		m.velocity[0] *= -1
		m.loc[0] = 0
	}
	if m.loc[1] > y {
		m.velocity[1] *= -1
		m.loc[1] = y
	}
}

func (m *Mover) isIniside(liq Liquid) bool {
	if m.loc[0] >= liq.loc[0] && m.loc[0] <= liq.loc[0]+liq.size[0] && m.loc[1] >= liq.loc[1] && m.loc[1] <= liq.loc[1]+liq.size[1] {
		return true
	}
	return false
}

type Liquid struct {
	loc         *vec.Vec2
	size        *vec.Vec2
	coefficient float64
}

func NewLiquid(x, y, width, height, coefficient float64) *Liquid {
	return &Liquid{
		&vec.Vec2{x, y},
		&vec.Vec2{width, height},
		coefficient,
	}
}
func (l *Liquid) Display(ctx *canvas.Context) {
	ctx.BeginPath()
	ctx.Rect(l.loc[0], l.loc[1], l.size[0], l.size[1])
	ctx.SetFillStyle(color.RGBA{0, 0, 255, 100})
	ctx.Fill()
	ctx.SetFillStyle(color.RGBA{0, 0, 0, 255})
}
