package main

import (
	"flag"
	"fmt"
	"image/color"
	"log"
	"math"
	"math/rand/v2"
	"time"

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

	engine := NewMoverView()

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

type MoverView [100]*Mover

func NewMoverView() *MoverView {
	movers := &MoverView{}
	for i := range 100 {
		movers[i] = NewMover(0, 0, rand.Float64()*8+2)
	}
	return movers
}

func (mv *MoverView) Draw(ctx *canvas.Context) {
	width, height := float64(ctx.CanvasWidth()), float64(ctx.CanvasHeight())
	ctx.ClearRect(0, 0, width, height)
	ctx.BeginPath()
	ctx.Rect(0, 0, width, height)
	ctx.Stroke()

	wind := vec.NewVec2(0.01, 0)

	for _, m := range mv {
		gravity := vec.NewVec2(0, 0.1*m.mass)
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
	m.acceleration.Add(*vec.DivV2(force, m.mass))
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
