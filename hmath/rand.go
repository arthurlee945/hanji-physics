package hmath

func MonteCarloRand(min, max float64) float64 {
	for {
		r1, r2 := FloatInBetween(min, max), FloatInBetween(min, max)
		if r2 < r1 {
			return r1
		}
	}
}
