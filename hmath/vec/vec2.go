package vec

type Vec2 [2]float64

func (v2 *Vec2) isVector() {}

func (v2 *Vec2) Dot(p Vec2) float64 {
	return v2[0]*p[0] + v2[1]*p[1]
}

func (v2 *Vec2) Add(p Vec2) {
	for i := range v2 {
		v2[i] += p[i]
	}
}
