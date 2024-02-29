package main

import (
	"image/color"
	"math"
	"time"

	"github.com/arthurlee945/suhag"
	"github.com/arthurlee945/suhag/example/utility"
	"github.com/fzipp/canvas"
)

func main() {
	utility.StartCanvas(700, 400, "Oscillation", runCanvas)
}

func runCanvas(ctx *canvas.Context) {
	ctx.SetFillStyle(color.RGBA{0x08, 0x08, 0x08, 0xff})
	osc := &Osc{0, 0.1}
	for {
		select {
		case event := <-ctx.Events():
			if _, ok := event.(canvas.CloseEvent); ok {
				return
			}
		default:
			ctx.ClearRect(0, 0, float64(ctx.CanvasWidth()), float64(ctx.CanvasHeight()))
			ctx.BeginPath()
			for x := range ctx.CanvasWidth() {
				y, err := suhag.Map(math.Sin(osc.angle), -1, 1, 0, float64(ctx.CanvasHeight()))
				if err != nil {
					break
				}
				ctx.Rect(float64(x), float64(y), 1, 1)
				ctx.Stroke()
				osc.angle += osc.aAccel
			}
			ctx.Flush()
			time.Sleep(5 * time.Millisecond)
		}
	}
}

type Osc struct {
	angle  float64
	aAccel float64
}
