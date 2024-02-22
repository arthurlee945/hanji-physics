package rand

import "testing"

func TestLCG(t *testing.T) {
	seeds := [2]int{
		888,
		123,
	}

	lcg, lcg2 := NewLCG(seeds[0]), NewLCG(seeds[1])
	randVal1, randVal2 := lcg.Rand(), lcg.Rand()
	rand2Val1, rand2Val2 := lcg2.Rand(), lcg2.Rand()

	if randVal1 != randVal2 {
		t.Errorf("expected two same seeded value to be same but got: %v | %v", randVal1, randVal2)
	}
	if rand2Val1 != rand2Val2 {
		t.Errorf("expected two same seeded value to be same but got: %v | %v", rand2Val1, rand2Val2)
	}
	if randVal1 == rand2Val1 {
		t.Errorf("expected two different seeded value to be different but got same: %v | %v", randVal1, rand2Val2)
	}
}
