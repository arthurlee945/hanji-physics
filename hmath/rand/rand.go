package rand

import "github.com/arthurlee945/hanji-physics/hmath"

func MonteCarloRand(min, max float64) float64 {
	for {
		r1, r2 := hmath.FloatInBetween(min, max), hmath.FloatInBetween(min, max)
		if r2 < r1 {
			return r1
		}
	}
}
