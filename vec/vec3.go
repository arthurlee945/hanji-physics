package vec

import "math"

type Vec3 [3]float64

func NewVec3(x, y, z float64) *Vec3 {
	return &Vec3{x, y, z}
}

func (v *Vec3) isVector() {}

func (v3 *Vec3) Dot(pv Vec3) float64 {
	return v3[0]*pv[0] + v3[1]*pv[1] + v3[2]*pv[2]
}

func (v3 *Vec3) Add(pv Vec3) {
	v3[0] += pv[0]
	v3[1] += pv[1]
	v3[2] += pv[2]
}

func (v3 *Vec3) Sub(pv Vec3) {
	v3[0] -= pv[0]
	v3[1] -= pv[1]
	v3[2] -= pv[2]
}

func (v3 *Vec3) Mult(v float64) {
	v3[0] *= v
	v3[1] *= v
	v3[2] *= v
}

func (v3 *Vec3) Div(v float64) {
	v3[0] /= v
	v3[1] /= v
	v3[2] /= v
}

func (v3 *Vec3) Mag() float64 {
	return math.Sqrt(v3[0]*v3[0] + v3[1]*v3[1] + v3[2]*v3[2])
}

func (v3 *Vec3) Normal() *Vec3 {
	mag := v3.Mag()
	return &Vec3{v3[0] / mag, v3[1] / mag, v3[2] / mag}
}

func (v3 *Vec3) Clone() *Vec3 {
	return &Vec3{v3[0], v3[1], v3[2]}
}
