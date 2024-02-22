package rand

type LCGFn func(*LCG)

// Linear Congruential Generator
type LCG struct {
	// max period
	max int
	a   int
	// co-prime of max
	coprime int
	seed    int
	z       int
}

func NewLCG(seed int) *LCG {
	lcg := &LCG{
		max:     4294967296, // 2 ** 32 || uint32
		a:       1664525,
		coprime: 1013904223,
		seed:    seed,
		z:       seed,
	}
	return lcg
}

func (lcg *LCG) GetSeed() int {
	return lcg.seed
}

func (lcg *LCG) Rand() float64 {
	lcg.z = (lcg.a*lcg.z + lcg.coprime) % lcg.max
	return float64(lcg.z) / float64(lcg.max)
}
