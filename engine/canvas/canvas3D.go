package canvas

type Canvas3D struct {
	Matrix [][][]byte
}

func (c3d *Canvas3D) isCanvas() {}

func NewCanvas3D(x, y, z uint16) *Canvas3D {
	matrix := make([][][]byte, z)
	for i := 0; i < int(z); i++ {
		matrix[i] = make2DMatrix(x, y)
	}
	return &Canvas3D{matrix}
}
