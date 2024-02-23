package main

import (
	"flag"
	"fmt"
	"image/color"
	"log"
	"math"
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
		Height:         600,
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

	engine := NewVectorView(ctx.CanvasWidth(), ctx.CanvasHeight())

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

type VectorView struct {
	loc          *vec.Vec2
	speed        *vec.Vec2
	accelaration *vec.Vec2
	size         *vec.Vec2
	pointer      *vec.Vec2
	topspeed     float64
}

func NewVectorView(canvasWidth, canvasHeight int) *VectorView {
	return &VectorView{
		loc:          &vec.Vec2{float64(canvasWidth) / 2, float64(canvasHeight) / 2},
		speed:        &vec.Vec2{0.5, 3.33},
		accelaration: &vec.Vec2{},
		size:         &vec.Vec2{float64(canvasWidth), float64(canvasHeight)},
		pointer:      &vec.Vec2{},
		topspeed:     10,
	}
}

func (vv *VectorView) Draw(ctx *canvas.Context) {
	vv.speed.Add(*vv.accelaration)
	if vv.speed[0] > 10 {
		vv.speed[0] = 10
	}
	if vv.speed[1] > 10 {
		vv.speed[1] = 10
	}
	vv.loc.Add(*vv.speed)
	ctx.ClearRect(0, 0, float64(ctx.CanvasWidth()), float64(ctx.CanvasHeight()))
	ctx.Rect(0, 0, vv.size[0], vv.size[1])
	ctx.Stroke()
	if vv.loc[0] > vv.size[0] && vv.speed[0] > 0 || vv.loc[0] < 0 && vv.speed[0] < 0 {
		vv.speed[0] *= -1
	}
	if vv.loc[1] > vv.size[1] && vv.speed[1] > 0 || vv.loc[1] < 0 && vv.speed[1] < 0 {
		vv.speed[1] *= -1
	}
	ctx.BeginPath()
	ctx.Arc(vv.loc[0], vv.loc[1], 10, 0, math.Pi*2, false)
	ctx.Fill()
}

func (vv *VectorView) Handle(evt canvas.Event) {
	e, ok := evt.(canvas.MouseMoveEvent)
	if !ok {
		return
	}
	vv.pointer[0] = float64(e.X)
	vv.pointer[1] = float64(e.Y)
	normalized := vec.SubV2(vec.Vec2{vv.pointer[0], vv.pointer[1]}, *vv.loc).Normal()
	normalized.Mult(0.5)
	vv.accelaration = normalized
}
