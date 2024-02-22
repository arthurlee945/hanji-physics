package suhag

func Lerp(t, a1, a2 float64) float64 {
	return a1 + t*(a2-a1)
}

func Fade(t float64) float64 {
	return ((6*t-15)*t + 10) * t * t * t
}
