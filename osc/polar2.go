package osc

//Oscillation

import (
	"math"

	"github.com/arthurlee945/suhag/vec"
)

type Polar2 struct {
	Radius float64
	Theta  float64
}

func NewPolar2(vector vec.Vec2) *Polar2 {
	return &Polar2{
		Radius: vector.Mag(),
		Theta:  math.Atan2(vector[1], vector[0]),
	}
}

func (p2 *Polar2) ToCartesian() *vec.Vec2 {
	return &vec.Vec2{p2.Radius * math.Cos(p2.Theta), p2.Radius * math.Sin(p2.Theta)}
}
