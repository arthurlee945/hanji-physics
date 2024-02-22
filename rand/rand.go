package rand

import "github.com/arthurlee945/suhag"

func MonteCarloRand(min, max float64) float64 {
	for {
		r1, r2 := suhag.FloatInBetween(min, max), suhag.FloatInBetween(min, max)
		if r2 < r1 {
			return r1
		}
	}
}
