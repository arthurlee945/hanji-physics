package osc

import "math"

//Takes a degree and returns radians | Oscillation
func Radian(deg float64) float64 {
	return 2 * math.Pi * (deg / 360)
}
