package vec

import "math"

type Vec3 [3]float64

func NewVec3(x, y, z float64) *Vec3 {
	return &Vec3{x, y, z}
}

func (v3 *Vec3) Dot(nv Vec3) float64 {
	return v3[0]*nv[0] + v3[1]*nv[1] + v3[2]*nv[2]
}

func (v3 *Vec3) Add(nv Vec3) {
	v3[0] += nv[0]
	v3[1] += nv[1]
	v3[2] += nv[2]
}

func (v3 *Vec3) Sub(nv Vec3) {
	v3[0] -= nv[0]
	v3[1] -= nv[1]
	v3[2] -= nv[2]
}

func (v3 *Vec3) Mult(v float64) {
	v3[0] *= v
	v3[1] *= v
	v3[2] *= v
}

func (v3 *Vec3) Div(v float64) {
	if v == 0 {
		return
	}
	v3[0] /= v
	v3[1] /= v
	v3[2] /= v
}

func (v3 *Vec3) Normalize() {
	mag := v3.Mag()
	if mag == 0 {
		return
	}
	v3[0] /= mag
	v3[1] /= mag
	v3[2] /= mag
}

func (v3 *Vec3) Mag() float64 {
	return math.Sqrt(v3[0]*v3[0] + v3[1]*v3[1] + v3[2]*v3[2])
}

func (v3 *Vec3) Normal() *Vec3 {
	mag := v3.Mag()
	if mag == 0 {
		return &Vec3{}
	}
	return &Vec3{v3[0] / mag, v3[1] / mag, v3[2] / mag}
}

func (v3 *Vec3) Clone() *Vec3 {
	return &Vec3{v3[0], v3[1], v3[2]}
}
