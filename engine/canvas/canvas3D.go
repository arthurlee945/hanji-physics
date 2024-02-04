package canvas

type Canvas3D struct {
	Matrix [][][]uint16
}

func (c3d *Canvas3D) isCanvas() {}

func NewCanvas3D(x, y, z uint16) *Canvas3D {
	matrix := make([][][]uint16, z)
	for i := range matrix {
		matrix[i] = make2DMatrix(x, y)
	}
	return &Canvas3D{matrix}
}
