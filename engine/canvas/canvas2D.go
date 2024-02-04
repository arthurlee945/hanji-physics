package canvas

type Canvas2D struct {
	X      uint16
	Y      uint16
	Matrix [][]uint16
}

func (c2d *Canvas2D) isCanvas() {}

func NewCanvas2D(x, y uint16) *Canvas2D {
	return &Canvas2D{x, y, make2DMatrix(x, y)}
}

func make2DMatrix(x, y uint16) [][]uint16 {
	matrix := make([][]uint16, y)
	for i := range matrix {
		matrix[i] = make([]uint16, x)
	}
	return matrix
}
