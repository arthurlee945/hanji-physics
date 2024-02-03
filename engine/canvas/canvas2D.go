package canvas

type Canvas2D struct {
	Matrix [][]byte
}

func (c2d *Canvas2D) isCanvas() {}

func NewCanvas2D(x, y uint16) *Canvas2D {
	return &Canvas2D{make2DMatrix(x, y)}
}

func make2DMatrix(x, y uint16) [][]byte {
	matrix := make([][]byte, y)
	bytes := make([]byte, x*y)
	for i := 0; i < int(y); i++ {
		matrix[i] = bytes[i*int(x) : (i+1)*int(x)]
	}
	return matrix
}
