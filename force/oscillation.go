package force

import "math"

func Normal(gravity, radian float64) float64 {
	return gravity * math.Sin(radian)
}
