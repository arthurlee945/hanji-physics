package vector

type Vector interface {
	isVector()
}

type Vec []float64

func (v *Vec) isVector() {}

type Vec2 [2]float64

func (v2 *Vec2) isVector() {}

func (v2 *Vec2) Dot(p *Vec2) float64 {
	return Dot2(v2, p)
}

type Vec3 [3]float64

func (v3 *Vec3) isVector() {}

func (v3 *Vec3) Dot(p *Vec3) float64 {
	return Dot3(v3, p)
}
