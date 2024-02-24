package force

import (
	"github.com/arthurlee945/suhag/vec"
)

func Acceleration[Vec vec.Vec2 | vec.Vec3](force Vec, mass float64) *Vec {
	return vec.Div(force, mass)
}

func Force[Vec vec.Vec2 | vec.Vec3](acceleration Vec, mass float64) *Vec {
	return vec.Mult(acceleration, mass)
}

// coefficient | strenth of a friction force for a particular surface
//
// normal | the force perpendicular to the object's motion along a surface
func Friction[Vec vec.Vec2 | vec.Vec3](vel vec.Vector[Vec], coefficent, normal float64) *Vec {
	friction := any(vel.Clone()).(vec.Vector[Vec])
	friction.Mult(-1)
	friction.Normalize()
	friction.Mult(coefficent * normal)
	return any(friction).(*Vec)
}
