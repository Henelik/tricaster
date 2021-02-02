package geometry

import (
	"github.com/Henelik/tricaster/matrix"
	"github.com/Henelik/tricaster/shading"
)

type Plane struct {
	// the transformation matrix
	m *matrix.Matrix
	// the inverse transformation matrix
	im *matrix.Matrix
	// the material
	Mat *shading.PhongMat
}
