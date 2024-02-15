package noiseview

import (
	"fmt"
	"image"
	"image/color"

	"github.com/arthurlee945/hanji-physics/hmath"
	"github.com/arthurlee945/hanji-physics/hmath/noise"
	"github.com/fzipp/canvas"
)

type NoiseView struct {
	x     int
	y     int
	xoff  float64
	zoff  float64
	noise *noise.Noise
}

func NewNoiseView(canvasWidth, canvasHeight int) *NoiseView {
	noiseview := &NoiseView{
		x:     canvasWidth,
		y:     canvasHeight,
		xoff:  0,
		zoff:  0,
		noise: noise.NewNoise(noise.WithSeededPermutation(8, noise.PERMUTATION_SIZE)),
	}
	return noiseview
}

func (nv *NoiseView) Draw(ctx *canvas.Context) {

	nv.draw2D(ctx)
}

func (nv *NoiseView) draw1D(ctx *canvas.Context) {
	xoff := nv.xoff
	ctx.ClearRect(0, 0, float64(nv.x), float64(nv.y))
	ctx.BeginPath()
	for x := range nv.x {
		y, err := hmath.Map(nv.noise.Run(xoff, 0, 0), 0, 1, 0, float64(nv.y))
		if err != nil {
			fmt.Println(err.Error())
		}
		ctx.LineTo(float64(x), y)
		ctx.Stroke()
		xoff += 0.01
	}
	nv.xoff += 0.01
}

func (nv *NoiseView) draw2D(ctx *canvas.Context) {
	xoff := 0.0
	ctx.ClearRect(0, 0, float64(nv.x), float64(nv.y))
	rgbaImg := image.NewRGBA(image.Rect(0, 0, nv.x, nv.y))
	for x := range nv.x {
		yoff := 0.0
		for y := range nv.y {
			brightness, err := hmath.Map(nv.noise.Run(xoff, yoff, nv.zoff), 0, 1, 0, 255)
			if err != nil {
				fmt.Println(err.Error())
			}
			rgbaImg.Set(x, y, color.RGBA{75, 0, 130, uint8(brightness)})
			yoff += 0.01
		}
		xoff += 0.01
	}
	ctx.DrawImage(ctx.CreateImageData(rgbaImg), 0, 0)
	nv.zoff += 0.01
}

// func (nv *NoiseView) draw2D(ctx *canvas.Context) {
// 	xoff := 0.0
// 	wg := sync.WaitGroup{}
// 	ctx.ClearRect(0, 0, float64(nv.x), float64(nv.y))
// 	rgbaImg := image.NewRGBA(image.Rect(0, 0, nv.x, nv.y))
// 	for x := range nv.x {
// 		yoff := 0.0
// 		for y := range nv.y {
// 			wg.Add(1)
// 			go func() {
// 				brightness, err := hmath.Map(nv.noise.Run(xoff, yoff, 0), 0, 1, 0, 255)
// 				fmt.Println(xoff, yoff)
// 				if err != nil {
// 					fmt.Println(err.Error())
// 				}
// 				rgbaImg.Set(x, y, color.RGBA{255, 0, 0, uint8(brightness)})
// 				wg.Done()
// 			}()
// 			yoff += 0.01
// 		}
// 		xoff += 0.01
// 	}
// 	wg.Wait()
// 	ctx.DrawImage(ctx.CreateImageData(rgbaImg), 0, 0)
// }
