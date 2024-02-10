package noise

import (
	"math/rand"

	"github.com/arthurlee945/hanji-physics/hmath/vector"
)

const (
	ONE_DIMENSION   = "1D"
	TWO_DIMENSION   = "2D"
	THREE_DIMENSION = "3D"
)

var (
	Vec2TopRight    = &vector.Vec2{1.0, 1.0}
	Vec2TopLeft     = &vector.Vec2{-1.0, 1.0}
	Vec2BottomRight = &vector.Vec2{1.0, -1.0}
	Vec2BottomLeft  = &vector.Vec2{-1.0, -1.0}
)

type NoiseDimension string
type NoiseFn func(*Noise)

// Perlin Noise Implementation
type Noise struct {
	dimension        NoiseDimension
	permutationCount uint
	permutation      []int
}

func NewNoise(opt ...NoiseFn) *Noise {
	noise := &Noise{
		permutationCount: 256,
		permutation:      rand.Perm(256),
	}
	return noise
}

func (n *Noise) UpdatePermutations() {}

func WithDimension(dimension NoiseDimension) NoiseFn {
	return func(n *Noise) {
		switch dimension {
		case ONE_DIMENSION:
		case TWO_DIMENSION:
		case THREE_DIMENSION:
		}
	}
}

func create2DNoise() {

}

func getConstVector(value int) *vector.Vec2 {
	switch value & 3 {
	case 0:
		return Vec2TopRight
	case 1:
		return Vec2TopLeft
	case 2:
		return Vec2BottomLeft
	default:
		return Vec2BottomRight
	}
}
