package force

import "github.com/arthurlee945/suhag/vec"

/*
F = -1 * k * (currentLength - restLength)

k | constants

restLength | length during rest position

anchor | origin point where object is connected to

obj | object hanging from the anchor
*/
func Spring[Vec vec.Vec2 | vec.Vec3](k, restLength float64, anchor, obj Vec) *Vec {
	force := any(vec.Sub(obj, anchor)).(vec.Vector[Vec])
	currLength := force.Mag()
	force.Normalize()
	force.Mult(-1 * k * (restLength - currLength))
	return any(force).(*Vec)
}
