package vectorview

import (
	"math"

	"github.com/arthurlee945/hanji-physics/hmath/vec"
	"github.com/fzipp/canvas"
)

type VectorView struct {
	loc   *vec.Vec2
	speed *vec.Vec2
	size  *vec.Vec2
}

func NewVectorView(canvasWidth, canvasHeight int) *VectorView {
	return &VectorView{
		loc:   &vec.Vec2{float64(canvasWidth) / 2, float64(canvasHeight) / 2},
		speed: &vec.Vec2{0.5, 3.33},
		size:  &vec.Vec2{float64(canvasWidth), float64(canvasHeight)},
	}
}

func (vv *VectorView) Draw(ctx *canvas.Context) {
	vv.loc[0] += vv.speed[0]
	vv.loc[1] += vv.speed[1]

	if vv.loc[0] > vv.size[0] && vv.speed[0] > 0 || vv.loc[0] < 0 && vv.speed[0] < 0 {
		vv.speed[0] *= -1
	}
	if vv.loc[0] > vv.size[1] && vv.speed[1] > 0 || vv.loc[0] < 0 && vv.speed[1] < 0 {
		vv.speed[1] *= -1
	}
	ctx.Arc(vv.loc[0], vv.loc[0], 16, 0, math.Pi*2, false)
	ctx.Fill()
}
func (vv *VectorView) Handle(evt canvas.Event) {}
