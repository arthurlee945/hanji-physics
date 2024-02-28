package oscillation

import (
	"math"

	"github.com/arthurlee945/suhag/vec"
)

type Polar2 struct {
	Rad   float64
	Theta float64
}

func NewPolar(vector vec.Vec2) *Polar2 {
	return &Polar2{
		Rad:   vector.Mag(),
		Theta: math.Atan2(vector[1], vector[0]),
	}
}

func (p2 *Polar2) ToCartesian() *vec.Vec2 {
	return &vec.Vec2{p2.Rad * math.Cos(p2.Theta), p2.Rad * math.Sin(p2.Theta)}
}
