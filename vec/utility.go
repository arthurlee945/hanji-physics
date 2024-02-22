package vec

import (
	"math/rand"
)

func RandomV2() *Vec2 {
	return &Vec2{rand.Float64(), rand.Float64()}
}

func RandomV3() *Vec3 {
	return &Vec3{rand.Float64(), rand.Float64(), rand.Float64()}
}

func AddV2(v1, v2 Vec2) *Vec2 {
	return &Vec2{v1[0] + v2[0], v1[1] + v2[1]}
}

func AddV3(v1, v2 Vec3) *Vec3 {
	return &Vec3{v1[0] + v2[0], v1[1] + v2[1], v1[2] + v2[2]}
}

func SubV2(v1, v2 Vec2) *Vec2 {
	return &Vec2{v1[0] - v2[0], v1[1] - v2[1]}
}

func SubV3(v1, v2 Vec3) *Vec3 {
	return &Vec3{v1[0] - v2[0], v1[1] - v2[1], v1[2] - v2[2]}
}

func DivV2(vector *Vec2, v float64) *Vec2 {
	return &Vec2{vector[0] / v, vector[1] / v}
}

func DivV3(vector *Vec3, v float64) *Vec3 {
	return &Vec3{vector[0] / v, vector[1] / v, vector[2] / v}
}

func MultV2(vector *Vec2, v float64) *Vec2 {
	return &Vec2{vector[0] * v, vector[1] * v}
}

func MultV3(vector *Vec3, v float64) *Vec3 {
	return &Vec3{vector[0] * v, vector[1] * v, vector[2] * v}
}
