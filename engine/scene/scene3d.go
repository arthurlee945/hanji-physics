package scene

type Scene3D struct {
	matrix [][][]byte
}

func NewScene3D(x, y, z uint16) *Scene3D {
	matrix := make([][][]byte, z)
	for i := 0; i < int(z); i++ {
		matrix[i] = make2DMatrix(x, y)
	}
	return &Scene3D{matrix}
}
