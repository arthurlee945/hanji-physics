package suhag

import (
	"errors"
	"math/rand"
)

func FloatInBetween(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func StdDeviation(sd, mean, num float64) float64 {
	return sd*num + mean
}

func Map(value, currentMin, currentMax, mapMin, mapMax float64) (float64, error) {
	if currentMin >= currentMax || mapMin >= mapMax {
		return 0, errors.New("minium value is larger then max value")
	}
	if value <= currentMin {
		return mapMin, nil
	}
	if value >= currentMax {
		return mapMax, nil
	}
	return (value-currentMin)/(currentMax-currentMin)*(mapMax-mapMin) + mapMin, nil
}

func Clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
