package force

import "github.com/arthurlee945/suhag/vec"

func AccelerationV2(force *vec.Vec2, mass float64) *vec.Vec2 {
	return vec.DivV2(force, mass)
}

func AccelerationV3(force *vec.Vec3, mass float64) *vec.Vec3 {
	return vec.DivV3(force, mass)
}

func ForceV2(acceleration *vec.Vec2, mass float64) *vec.Vec2 {
	return vec.MultV2(acceleration, mass)
}

func ForceV3(acceleration *vec.Vec3, mass float64) *vec.Vec3 {
	return vec.MultV3(acceleration, mass)
}
