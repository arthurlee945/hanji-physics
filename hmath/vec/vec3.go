package vec

type Vec3 [3]float64

func (v *Vec3) isVector() {}

func (v3 *Vec3) Dot(p Vec3) float64 {
	return v3[0]*p[0] + v3[1]*p[1] + v3[2]*p[2]
}

func (v3 *Vec3) Add(p Vec3) {
	for i := range v3 {
		v3[i] += p[i]
	}
}
