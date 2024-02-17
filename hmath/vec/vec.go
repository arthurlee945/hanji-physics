package vec

type Vec [1]float64

func (v *Vec) isVector() {}

func (v *Vec) Dot(p Vec) float64 {
	return v[0] - p[0]
}
