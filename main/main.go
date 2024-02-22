package main

import (
	"flag"
	"fmt"
	"image/color"
	"log"
	"time"

	"github.com/arthurlee945/hanji-physics/example"
	"github.com/arthurlee945/hanji-physics/example/noiseview"
	"github.com/arthurlee945/hanji-physics/example/vectorview"
	"github.com/arthurlee945/hanji-physics/example/walker"
	"github.com/fzipp/canvas"
)

const (
	WALKER     = "WALKER"
	NOISEVIEW  = "NOISEVIEW"
	VECTORVIEW = "VECTORVIEW"
)

type PhysicsType string

func main() {
	http := flag.String("http", ":8080", "HTTP service address (e.g.. '127.0.0.1:8080' or ':8080')")
	flag.Parse()

	fmt.Println("Listening on " + httpLink(*http))
	err := canvas.ListenAndServe(*http, runCanvas, &canvas.Options{
		Title:          "Hanji Physics",
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

	engine := physicsToRun(NOISEVIEW, ctx)

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

func physicsToRun(typeName PhysicsType, ctx *canvas.Context) example.EnginePart {
	switch typeName {
	case WALKER:
		return walker.NewWalker(ctx.CanvasWidth(), ctx.CanvasHeight())
	case NOISEVIEW:
		return noiseview.NewNoiseView(ctx.CanvasWidth(), ctx.CanvasHeight())
	case VECTORVIEW:
		return vectorview.NewVectorView(ctx.CanvasWidth(), ctx.CanvasHeight())
	default:
		return walker.NewWalker(ctx.CanvasWidth(), ctx.CanvasHeight())
	}
}
