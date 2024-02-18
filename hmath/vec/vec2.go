package vec

import "math"

type Vec2 [2]float64

func (v2 *Vec2) isVector() {}

func (v2 *Vec2) Dot(pv Vec2) float64 {
	return v2[0]*pv[0] + v2[1]*pv[1]
}

func (v2 *Vec2) Add(pv Vec2) {
	v2[0] += pv[0]
	v2[1] += pv[1]
}

func (v2 *Vec2) Sub(pv Vec2) {
	v2[0] += pv[0]
	v2[1] += pv[1]
}

func (v2 *Vec2) Mult(v float64) {
	v2[0] *= v
	v2[1] *= v
}

func (v2 *Vec2) Div(v float64) {
	v2[0] /= v
	v2[1] /= v
}

func (v2 *Vec2) Mag() float64 {
	return math.Sqrt(v2[0]*v2[0] + v2[1]*v2[1])
}

func (v2 *Vec2) Normal() Vec2 {
	mag := v2.Mag()
	return Vec2{v2[0] / mag, v2[1] / mag}
}
