package hmath

import (
	"math/rand"
)

func FloatInBetween(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func StdDeviation(sd, mean, num float64) float64 {
	return sd*num + mean
}

func Map(value, currentMin, currentMax, mapMin, mapMax float64) {}
