package vec

import "math"

type Vec2 [2]float64

func NewVec2(x, y float64) *Vec2 {
	return &Vec2{x, y}
}

func (v2 *Vec2) Dot(nv Vec2) float64 {
	return v2[0]*nv[0] + v2[1]*nv[1]
}

func (v2 *Vec2) Add(nv Vec2) {
	v2[0] += nv[0]
	v2[1] += nv[1]
}

func (v2 *Vec2) Sub(nv Vec2) {
	v2[0] += nv[0]
	v2[1] += nv[1]
}

func (v2 *Vec2) Mult(v float64) {
	v2[0] *= v
	v2[1] *= v
}

func (v2 *Vec2) Div(v float64) {
	if v == 0 {
		return
	}
	v2[0] /= v
	v2[1] /= v
}

func (v2 *Vec2) Normalize() {
	mag := v2.Mag()
	if mag == 0 {
		return
	}
	v2[0] /= mag
	v2[1] /= mag
}

func (v2 *Vec2) Mag() float64 {
	return math.Sqrt(v2[0]*v2[0] + v2[1]*v2[1])
}

func (v2 *Vec2) Normal() *Vec2 {
	mag := v2.Mag()
	if mag == 0 {
		return &Vec2{}
	}
	return &Vec2{v2[0] / mag, v2[1] / mag}
}

func (v2 *Vec2) Clone() *Vec2 {
	return &Vec2{v2[0], v2[1]}
}
