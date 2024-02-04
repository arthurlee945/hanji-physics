package canvas

type CanvasType interface {
	isCanvas()
}

type Canvas struct {
	X      int
	y      int
	Matrix []uint16
}

func NewCanvas(x, y int) *Canvas {
	return &Canvas{x, y, makeMatrix(x, y)}
}

func makeMatrix(x, y int) []uint16 {
	matrix := make([]uint16, x*y*4)
	for i := range matrix {
		matrix[i] = 255
	}
	return matrix
}
