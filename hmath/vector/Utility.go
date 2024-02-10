package vector

func Dot2(p1, p2 *Vec2) float64 {
	return p1[0]*p2[0] + p1[1]*p2[1]
}

func Dot3(p1, p2 *Vec3) float64 {
	return p1[0]*p2[0] + p1[1]*p2[1] + p1[2]*p2[2]
}
