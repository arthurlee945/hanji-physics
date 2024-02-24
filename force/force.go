package force

import "github.com/arthurlee945/suhag/vec"

func Acceleration[Vec vec.Vec2 | vec.Vec3](force Vec, mass float64) *Vec {
	return vec.Div(force, mass)
}

func Force[Vec vec.Vec2 | vec.Vec3](acceleration Vec, mass float64) *Vec {
	return vec.Mult(acceleration, mass)
}
