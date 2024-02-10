package hmath

func Lerp(t, a1, a2 float64) float64 {
	return a1 + t*(a2-a1)
}

func Fade(t float64) float64 {
	return ((6*t-15)*t + 10) * t * t * t
}

func FractalBrownianMotion(x, y float64, octaves uint) {
	// var (
	// 	result    float64 = 0.0
	// 	amplitude float64 = 0.1
	// 	frequency float64 = 0.005
	// )
	for i := 0; i < int(octaves); i++ {

	}
}
