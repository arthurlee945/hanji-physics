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

func Add[Vec Vec2 | Vec3](v1, v2 Vec) *Vec {
	var newVec Vec
	for i := range len(newVec) {
		newVec[i] = v1[i] + v2[i]
	}
	return &newVec
}

func Sub[Vec Vec2 | Vec3](v1, v2 Vec) *Vec {
	var newVec Vec
	for i := range len(newVec) {
		newVec[i] = v1[i] - v2[i]
	}
	return &newVec
}

func Div[Vec Vec2 | Vec3](vector Vec, v float64) *Vec {
	var newVec Vec
	for i := range len(newVec) {
		newVec[i] = vector[i] / v
	}
	return &newVec
}

func Mult[Vec Vec2 | Vec3](vector Vec, v float64) *Vec {
	var newVec Vec
	for i := range len(newVec) {
		newVec[i] = vector[i] * v
	}
	return &newVec
}
