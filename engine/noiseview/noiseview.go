package noiseview

import (
	"fmt"

	"github.com/arthurlee945/hanji-physics/hmath"
	"github.com/arthurlee945/hanji-physics/hmath/noise"
	"github.com/fzipp/canvas"
)

type NoiseView struct {
	x     int
	y     int
	start float64
	noise *noise.Noise
}

func NewNoiseView(canvasWidth, canvasHeight int) *NoiseView {
	noiseview := &NoiseView{
		x:     canvasWidth,
		y:     canvasHeight,
		start: 0,
		noise: noise.NewNoise(noise.WithSeededPermutation(8, noise.PERMUTATION_SIZE)),
	}
	return noiseview
}

func (nv *NoiseView) Draw(ctx *canvas.Context) {
	xoff := nv.start
	ctx.ClearRect(0, 0, float64(nv.x), float64(nv.y))
	ctx.BeginPath()
	for x := 0; x < nv.x; x++ {
		y, err := hmath.Map(nv.noise.Run(float64(xoff), 0, 0), 0, 1, 0, float64(nv.y)/1.5)
		if err != nil {
			fmt.Println(err.Error())
		}
		ctx.LineTo(float64(x), y)
		ctx.Stroke()
		xoff += 0.1
	}
	nv.start += 0.1
}
