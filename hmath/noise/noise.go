package noise

import (
	"math"
	"math/rand"
	"sync"

	"github.com/arthurlee945/hanji-physics/hmath"
)

const (
	YWRAP_BITWISE    = 4
	YWRAP            = 1 << YWRAP_BITWISE //16
	ZWRAP_BITWISE    = 8
	ZWRAP            = 1 << ZWRAP_BITWISE // 256
	PERMUTATION_SIZE = 4095               // 256 * 16 - 1
	OCTAVES          = 4
	AMP_FALLOFF      = 0.5
)

type NoiseFn func(*Noise)

// * Perlin Noise Implementation
type Noise struct {
	permutSize int
	permut     []float64
	octaves    int
	ampFO      float64
	mu         sync.RWMutex
}

func NewNoise(opts ...NoiseFn) *Noise {
	noise := &Noise{
		permutSize: PERMUTATION_SIZE,
		permut:     getPermutations(PERMUTATION_SIZE),
		octaves:    OCTAVES,
		ampFO:      AMP_FALLOFF,
	}
	for _, o := range opts {
		o(noise)
	}
	return noise
}

func (n *Noise) Run(x, y, z float64) float64 {
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	if z < 0 {
		z = -z
	}

	xi, yi, zi := int(math.Floor(x)), int(math.Floor(y)), int(math.Floor(z))
	xoff, yoff, zoff := x-float64(xi), y-float64(yi), z-float64(zi)

	var (
		r    float64 = 0 //result
		ampl float64 = 0.5
	)

	for i := 0; i < n.octaves; i++ {
		of := xi + (yi << YWRAP_BITWISE) + (zi << ZWRAP_BITWISE)
		rxf := scaledCos(xoff)
		ryf := scaledCos(yoff)

		//Get Noise From X and Y
		n1 := n.permut[of&n.permutSize]
		n1 += rxf * (n.permut[(of+1)&n.permutSize] - n1)
		n2 := n.permut[(of+YWRAP)&n.permutSize]
		n2 += rxf * (n.permut[(of+YWRAP+1)&n.permutSize] - n2)
		n1 += ryf * (n2 - n1) //Lerp between n1 and n2 at ryf

		of += ZWRAP

		//Get Noise From XYNoise and Z
		n2 = n.permut[of&n.permutSize]
		n2 += rxf * (n.permut[(of+1)&n.permutSize] - n2)
		n3 := n.permut[(of+YWRAP)&n.permutSize]
		n3 += rxf * (n.permut[(of+YWRAP+1)&n.permutSize] - n3)
		n2 += ryf * (n3 - n2) //Lerp between n3 and n2 at ryf

		n1 += scaledCos(zoff) * (n2 - n1) //Lerp between prev two Lerps (n1 and n2) at ryf

		r += n1 * ampl
		ampl *= n.ampFO
		xi, yi, zi = xi<<1, yi<<1, zi<<1          // int * 2
		xoff, yoff, zoff = xoff*2, yoff*2, zoff*2 // float * 2

		if xoff >= 1 {
			xi++
			xoff--
		}
		if yoff >= 1 {
			yi++
			yoff--
		}
		if zoff >= 1 {
			zi++
			zoff--
		}
	}
	return r
}

func WithPermutationSize(size int) NoiseFn {
	return func(n *Noise) {
		n.permutSize = size
		n.permut = getPermutations(size)
	}
}

func WithDetail(oct int, ampFO float64) NoiseFn {
	return func(n *Noise) {
		n.octaves = oct
		n.ampFO = ampFO
	}
}

func WithSeededPermutation(seed int, size int) NoiseFn {
	return func(n *Noise) {
		lcg := hmath.NewLCG(seed)
		permutations := make([]float64, size+1)
		for i := range permutations {
			permutations[i] = lcg.Rand()
		}
		n.permut = permutations
	}
}

func getPermutations(size int) []float64 {
	permutations := make([]float64, size+1)
	for i := range permutations {
		permutations[i] = rand.Float64()
	}
	return permutations
}

// Smoothing Function
func scaledCos(v float64) float64 {
	return 0.5 * (1 - math.Cos(v*math.Pi))
}
