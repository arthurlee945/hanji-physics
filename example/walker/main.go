package main

import (
	"flag"
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/arthurlee945/suhag"
	"github.com/arthurlee945/suhag/noise"
	"github.com/arthurlee945/suhag/vec"
	"github.com/fzipp/canvas"
)

func main() {
	http := flag.String("http", ":8080", "HTTP service address (e.g.. '127.0.0.1:8080' or ':8080')")
	flag.Parse()

	fmt.Println("Listening on " + httpLink(*http))
	err := canvas.ListenAndServe(*http, runCanvas, &canvas.Options{
		Title:          "Walker",
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

	engine := NewWalker(ctx.CanvasWidth(), ctx.CanvasHeight())

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

type Walker struct {
	loc    *vec.Vec2
	pLoc   *vec.Vec2
	offset *vec.Vec2
	size   *vec.Vec2
	noise  *noise.Noise
}

func NewWalker(canvasWidth, canvasHeight int) *Walker {
	pointX, pointY := float64(canvasWidth/2), float64(canvasHeight/2)
	return &Walker{
		loc:    &vec.Vec2{pointX, pointY},
		pLoc:   &vec.Vec2{pointX, pointY},
		offset: &vec.Vec2{0, 10000},
		size:   &vec.Vec2{float64(canvasWidth), float64(canvasHeight)},
		noise:  noise.NewNoise()}
}

func (w *Walker) Draw(ctx *canvas.Context) {
	w.noiseMove()
	ctx.Rect(w.loc[0], w.loc[1], 1, 1)
	ctx.Fill()
}
func (w *Walker) Handle(evt canvas.Event) {
	e, ok := evt.(canvas.MouseMoveEvent)
	if !ok {
		return
	}
	w.pLoc[0] = float64(e.X)
	w.pLoc[1] = float64(e.Y)
}

func (w *Walker) attractionMove() {
	if w.pLoc[0] == w.loc[0] && w.pLoc[1] == w.loc[1] {
		newX := rand.Intn(3) - 1
		newY := rand.Intn(3) - 1
		w.loc[0] += float64(newX)
		w.loc[1] += float64(newY)
	} else {
		randX, randY := rand.Float32(), rand.Float32()
		newX, newY := -1+randX*2, -1+randY*2
		distX, distY := suhag.StdDeviation(0.5, 1, float64(newX)), suhag.StdDeviation(0.5, 1, float64(newY))
		if w.loc[0] < w.pLoc[0] {
			if randX < 0.2 {
				w.loc[0] -= distX
			} else if randX > 0.5 {
				w.loc[0] += distX
			}
		} else {
			if randX < 0.2 {
				w.loc[0] += distX
			} else if randX > 0.5 {
				w.loc[0] -= distX
			}
		}

		if w.loc[1] < w.pLoc[1] {
			if randY < 0.2 {
				w.loc[1] -= distY
			} else if randY > 0.5 {
				w.loc[1] += distY
			}
		} else {
			if randY < 0.2 {
				w.loc[1] += distY
			} else if randY > 0.5 {
				w.loc[1] -= distY
			}
		}
	}
}

func (w *Walker) noiseMove() {
	newX, errX := suhag.Map(w.noise.Run(w.offset[0], 0, 0), 0, 1, -2, 2)
	if errX != nil {
		fmt.Println(errX)
	}
	newY, errY := suhag.Map(w.noise.Run(w.offset[1], 0, 0), 0, 1, -2, 2)
	if errY != nil {
		fmt.Println(errY)
	}
	w.loc[0] += float64(newX)
	w.loc[1] += float64(newY)
	w.offset[0] += 0.01
	w.offset[1] += 0.01
}

func (w *Walker) move() {
	newX := rand.Intn(3) - 1
	newY := rand.Intn(3) - 1
	w.loc[0] += float64(newX)
	w.loc[1] += float64(newY)
}
