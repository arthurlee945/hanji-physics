package oscillation

import (
	"math"

	"github.com/arthurlee945/suhag/vec"
)

type Polar3 struct {
	Radius float64
	Theta  float64
	Phi    float64
}

func NewPolar3(vector vec.Vec3) *Polar3 {
	r := vector.Mag()
	theta := math.Acos(vector[0] / r)
	phi := math.Acos(vector[0] / (r * math.Sin(theta)))
	return &Polar3{
		Radius: r,
		Theta:  theta,
		Phi:    phi,
	}
}

func (p3 *Polar3) ToCartesian() *vec.Vec3 {
	return &vec.Vec3{
		p3.Radius * math.Sin(p3.Theta) * math.Cos(p3.Theta),
		p3.Radius * math.Sin(p3.Theta) * math.Sin(p3.Phi),
		p3.Radius * math.Cos(p3.Phi),
	}
}
