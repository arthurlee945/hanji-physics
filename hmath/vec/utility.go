package vec

import "math"

func dot[VecT Vec2 | Vec3](ov VecT, pv VecT) float64 {
	var dotVal float64
	for i := range len(ov) {
		dotVal += ov[i] * pv[i]
	}
	return dotVal
}

func add[VecT *Vec2 | *Vec3](ov VecT, pv VecT) {
	for i := range len(ov) {
		ov[i] += pv[i]
	}
}

func sub[VecT *Vec2 | *Vec3](ov VecT, pv VecT) {
	for i := range len(ov) {
		ov[i] -= pv[i]
	}
}

func mult[VecT *Vec2 | *Vec3](ov VecT, v float64) {
	for i := range len(ov) {
		ov[i] *= v
	}
}

func div[VecT *Vec2 | *Vec3](ov VecT, v float64) {
	for i := range len(ov) {
		ov[i] /= v
	}
}

func mag[VecT *Vec2 | *Vec3](ov VecT) float64 {
	var magSq float64
	for i := range len(ov) {
		magSq += ov[i] * ov[i]
	}
	return math.Sqrt(magSq)
}
