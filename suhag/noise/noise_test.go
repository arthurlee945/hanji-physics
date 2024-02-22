package noise

import (
	"testing"

	"github.com/arthurlee945/hanji-physics/suhag/vec"
)

func TestNoise(t *testing.T) {
	inputs := []vec.Vec3{
		{2, 3, 4},
		{6.5, 2, 9.3},
	}
	noise := NewNoise(WithSeededPermutation(5, 4095))
	noise2 := NewNoise(WithSeededPermutation(5, 4095))
	noise3 := NewNoise(WithSeededPermutation(8, 4095))

	for _, input := range inputs {
		result := noise.Run(input[0], input[1], input[2])
		result2 := noise2.Run(input[0], input[1], input[2])
		result3 := noise3.Run(input[0], input[1], input[2])
		if result != result2 {
			t.Errorf("two same seed result returned different result, noise-1 = %v, noise-2 = %v", result, result2)
		}
		if result3 == result || result3 == result2 {
			t.Errorf("Different seeded noise returend same value noise-1 = %v, noise-2 = %v, got = %v", result, result2, result3)
		}
	}
}
