package scene

import (
	"github.com/Henelik/tricaster/matrix"
	"math"
)

var DefaultCameraTransform = &matrix.Matrix{
	Order: 4,
	Data: [][]float64{
		{-1, 0, 0, 0},
		{0, 0, 1, 0},
		{0, 1, 0, 0},
		{0, 0, 0, 1}}}

var DefaultFOV = math.Pi / 2

type Camera struct {
	hSize      uint16
	vSize      uint16
	FOV        float64
	Transform  *matrix.Matrix
	halfWidth  float64
	halfHeight float64
	pixelSize  float64
}

// NewCamera creates a new camera.
// Set fov to 0 and transform to nil to use defaults.
func NewCamera(hSize, vSize uint16, fov float64, transform *matrix.Matrix) *Camera {
	c := &Camera{
		hSize:     hSize,
		vSize:     vSize,
		FOV:       DefaultFOV,
		Transform: DefaultCameraTransform,
	}

	if fov != 0 {
		c.FOV = fov
	}
	if transform != nil {
		c.Transform = transform
	}

	halfView := math.Tan(c.FOV / 2)
	aspect := float64(hSize) / float64(vSize)
	if aspect >= 1 {
		c.halfWidth = halfView
		c.halfHeight = halfView / aspect
	} else {
		c.halfWidth = halfView * aspect
		c.halfHeight = halfView
	}
	c.pixelSize = (c.halfWidth * 2) / float64(c.hSize)

	return c
}
