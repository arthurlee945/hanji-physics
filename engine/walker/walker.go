package walker

type Walker struct {
	x int32
	y int32
}

func NewWalker(x int32, y int32) *Walker {
	return &Walker{x, y}
}
