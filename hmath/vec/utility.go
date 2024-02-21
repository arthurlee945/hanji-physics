package vec

import (
	"math/rand"
)

func RandomVec2() *Vec2 {
	return &Vec2{rand.Float64(), rand.Float64()}
}

func RandomVec3() *Vec3 {
	return &Vec3{rand.Float64(), rand.Float64(), rand.Float64()}
}

func AddVec2(v1, v2 Vec2) *Vec2 {
	return &Vec2{v1[0] + v2[0], v1[1] + v2[1]}
}

func AddVec3(v1, v2 Vec3) *Vec3 {
	return &Vec3{v1[0] + v2[0], v1[1] + v2[1], v1[2] + v2[2]}
}

func SubVec2(v1, v2 Vec2) *Vec2 {
	return &Vec2{v1[0] - v2[0], v1[1] - v2[1]}
}

func SubVec3(v1, v2 Vec3) *Vec3 {
	return &Vec3{v1[0] - v2[0], v1[1] - v2[1], v1[2] - v2[2]}
}

func DivVec2(v1, v2 Vec2) *Vec2 {
	return &Vec2{v1[0] / v2[0], v1[1] / v2[1]}
}

func DivVec3(v1, v2 Vec3) *Vec3 {
	return &Vec3{v1[0] / v2[0], v1[1] / v2[1], v1[2] / v2[2]}
}

func MultVec2(v1, v2 Vec2) *Vec2 {
	return &Vec2{v1[0] * v2[0], v1[1] * v2[1]}
}

func MultVec3(v1, v2 Vec3) *Vec3 {
	return &Vec3{v1[0] * v2[0], v1[1] * v2[1], v1[2] * v2[2]}
}
