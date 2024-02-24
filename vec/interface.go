package vec

type Vector[Vec Vec2 | Vec3] interface {
	Dot(Vec) float64
	Add(Vec)
	Sub(Vec)
	Mult(float64)
	Div(float64)
	Mag() float64
	Normalize()
	Normal() *Vec
	Clone() *Vec
}
