package main

import (
	"image/color"
	"math"
	"time"

	"github.com/arthurlee945/suhag/example/utility"
	"github.com/arthurlee945/suhag/osc"
	"github.com/arthurlee945/suhag/vec"
	"github.com/fzipp/canvas"
)

func main() {
	utility.StartCanvas(500, 500, "Oscillation", runCanvas)
}

func runCanvas(ctx *canvas.Context) {
	width, height := float64(ctx.CanvasWidth()), float64(ctx.CanvasHeight())
	ctx.SetFillStyle(color.RGBA{0x08, 0x08, 0x08, 0xff})
	pendulum := &Pendulum{&osc.Polar2{Radius: 125, Theta: math.Pi / 4}, 0, 0}
	origin := vec.NewVec2(width/2, height/3)
	for {
		pendulum.Update()

		ctx.ClearRect(0, 0, width, height)
		ctx.BeginPath()
		ctx.Rect(0, 0, width, height)
		ctx.Stroke()
		cart := pendulum.polar.ToCartesian()
		loc := vec.NewVec2(cart[1], cart[0])
		ctx.BeginPath()
		loc.Add(*origin)
		ctx.MoveTo(origin[0], origin[1])
		ctx.LineTo(loc[0], loc[1])
		ctx.Stroke()

		ctx.BeginPath()
		ctx.Arc(loc[0], loc[1], 16, 0, math.Pi*2, false)
		ctx.Fill()
		ctx.Flush()
		time.Sleep(5 * time.Millisecond)
	}
}

type Pendulum struct {
	polar  *osc.Polar2
	aVel   float64
	aAccel float64
}

func (p *Pendulum) Update() {
	p.aAccel = (-1 * 0.4 / p.polar.Radius) * math.Sin(p.polar.Theta)
	p.aVel += p.aAccel
	p.polar.Theta += p.aVel
	p.aVel *= 0.995
}
