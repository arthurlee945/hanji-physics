package main

import (
	"flag"
	"fmt"
	"image/color"
	"log"
	"math"
	"math/rand/v2"
	"time"

	"github.com/arthurlee945/suhag"
	"github.com/arthurlee945/suhag/example/utility"
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
		Width:          800,
		Height:         800,
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
	movers    [10]*utility.Mover
	attractor *Attractor
}

func NewMoverView(w, h float64) *MoverView {
	mv := &MoverView{
		attractor: NewAttractor(w/2, h/2, 30),
	}
	for i := range len(mv.movers) {
		mv.movers[i] = utility.NewMover(float64(rand.Int32N(int32(w))), float64(rand.Int32N(int32(h))), rand.Float64()*8+4)
	}
	return mv
}

func (mv *MoverView) Draw(ctx *canvas.Context) {
	width, height := float64(ctx.CanvasWidth()), float64(ctx.CanvasHeight())
	ctx.ClearRect(0, 0, width, height)
	ctx.BeginPath()
	ctx.Rect(0, 0, width, height)
	ctx.Stroke()
	mv.attractor.Display(ctx)

	for _, m := range mv.movers {
		gravitationalPull := force.Attraction(mv.attractor.G, mv.attractor.mass, m.Mass, *m.Loc, *mv.attractor.loc)
		gravMag := suhag.Clamp(gravitationalPull.Mag(), 0.4, 3)
		gravitationalPull.Normalize()
		gravitationalPull.Mult(gravMag)
		m.ApplyForce(gravitationalPull)
		// for m2i, m2 := range mv.movers {
		// 	if mi == m2i {
		// 		continue
		// 	}
		// 	repulsion := force.Repulsion(mv.attractor.G, mv.attractor.mass, m.Mass, *m.Loc, *m2.Loc)
		// 	repulsionMag := suhag.Clamp(repulsion.Mag(), 0.1, 2)
		// 	repulsion.Normalize()
		// 	repulsion.Mult(repulsionMag)
		// 	m.ApplyForce(repulsion)
		// }
		m.CheckEdges(float64(ctx.CanvasWidth()), float64(ctx.CanvasHeight()))
		m.Move()
		m.Display(ctx)

	}
}
func (mv *MoverView) Handle(evt canvas.Event) {
	e, ok := evt.(canvas.MouseMoveEvent)
	if !ok {
		return
	}
	mv.attractor.loc[0] = float64(e.X)
	mv.attractor.loc[1] = float64(e.Y)
}

type Attractor struct {
	mass float64
	G    float64
	loc  *vec.Vec2
}

func NewAttractor(x, y, mass float64) *Attractor {
	return &Attractor{mass, 2, vec.NewVec2(x, y)}
}

func (a *Attractor) Display(ctx *canvas.Context) {
	ctx.BeginPath()
	ctx.Arc(a.loc[0], a.loc[1], a.mass, 0, math.Pi*2, false)
	ctx.SetFillStyle(color.RGBA{138, 43, 226, 150})
	ctx.Fill()
	ctx.SetFillStyle(color.RGBA{75, 0, 130, 255})
}
