package canvas

type Canvas2D struct {
	Matrix [][]uint16
}

func (c2d *Canvas2D) isCanvas() {}

func NewCanvas2D(x, y uint16) *Canvas2D {
	return &Canvas2D{make2DMatrix(x, y)}
}

func make2DMatrix(x, y uint16) [][]uint16 {
	matrix := make([][]uint16, y)
	// bytes := make([]uint16, x*y)
	for i := range matrix {
		// matrix[i] = bytes[i*int(x) : (i+1)*int(x) : (i+1)*int(x)]
		matrix[i] = make([]uint16, x)
	}
	return matrix
}
