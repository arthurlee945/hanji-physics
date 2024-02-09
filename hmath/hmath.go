package hmath

import "math/rand"

func FloatInBetween(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func MonteCarloRand(min, max float64) float64 {
	for {
		r1, r2 := FloatInBetween(min, max), FloatInBetween(min, max)
		if r2 < r1 {
			return r1
		}
	}
}

func StdDeviation(sd, mean, num float64) float64 {
	return sd*num + mean
}

func Map(value, currentMin, currentMax, mapMin, mapMax float64) {}
