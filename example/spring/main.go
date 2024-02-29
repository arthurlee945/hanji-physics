package main

import (
	"fmt"
	"image/color"

	"github.com/arthurlee945/suhag/example/utility"
	"github.com/fzipp/canvas"
)

func main() {
	utility.StartCanvas(500, 500, "Oscillation", runCanvas)
}

func runCanvas(ctx *canvas.Context) {
	width, height := float64(ctx.CanvasWidth()), float64(ctx.CanvasHeight())
	ctx.SetFillStyle(color.RGBA{0x08, 0x08, 0x08, 0xff})
	for {
		fmt.Println(width, height)

	}
}
