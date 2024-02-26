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

/*
F = -1 * μ * N * v^(velocity unit vector)

c(μ) - coefficient of friction | strenth of a friction force for a particular surface

normal(N) | the force perpendicular to the object's motion along a surface

c * normal = friction magnitude
*/
func Friction[Vec vec.Vec2 | vec.Vec3](vel vec.Vector[Vec], c, normal float64) *Vec {
	friction := any(vel.Clone()).(vec.Vector[Vec])
	friction.Mult(-1)
	friction.Normalize()
	friction.Mult(c * normal)
	return any(friction).(*Vec)
}

/*
F = -1/2 * p(rho - density of liquid) * v*2(velocity magnitude) * A(cross sectional area) * Cd (Coefficient of Drag) * v^(velocity unit vector)

c (Cd) - coefficient of drag | Coefficient of Drag)

c * speed * speed = drag magnitude
TODO: cross sectional area need to be implemented
*/
func Drag[Vec vec.Vec2 | vec.Vec3](vel vec.Vector[Vec], c float64) *Vec {
	speed := vel.Mag()
	drag := any(vel.Clone()).(vec.Vector[Vec])
	drag.Mult(-1)
	drag.Normalize()
	drag.Mult(c * speed * speed)
	return any(drag).(*Vec)
}
