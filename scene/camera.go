package scene

import (
	"github.com/Henelik/tricaster/matrix"
	"github.com/Henelik/tricaster/ray"
	"github.com/Henelik/tricaster/tuple"
	"math"
)

var DefaultFOV = math.Pi / 2

type Camera struct {
	hSize      uint16
	vSize      uint16
	FOV        float64
	m          *matrix.Matrix
	im         *matrix.Matrix
	halfWidth  float64
	halfHeight float64
	pixelSize  float64
}

// NewCamera creates a new camera.
// Set fov to 0 and transform to nil to use defaults.
func NewCamera(hSize, vSize uint16, fov float64, transform *matrix.Matrix) *Camera {
	c := &Camera{
		hSize: hSize,
		vSize: vSize,
	}

	if fov == 0 {
		c.FOV = DefaultFOV
	} else {
		c.FOV = fov
	}
	if transform == nil {
		c.m = matrix.Identity
	} else {
		c.m = transform
	}
	c.im = c.m.Inverse()

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

func (c *Camera) GetMatrix() *matrix.Matrix {
	return c.m
}

func (c *Camera) SetMatrix(m *matrix.Matrix) {
	c.m = m
	c.im = m.Inverse()
}

func (c *Camera) GetTransform() *matrix.Matrix {
	return c.im
}

func (c *Camera) SetTransform(im *matrix.Matrix) {
	c.im = im
	c.m = im.Inverse()
}

func (c *Camera) RayForPixel(x, y uint16) *ray.Ray {
	// the offset from the edge of the canvas to the pixel's center
	xOffset := (float64(x) + 0.5) * c.pixelSize
	zOffset := (float64(y) + 0.5) * c.pixelSize

	// the untransformed coordinates of the pixel in world space.
	// (remember that the camera looks toward -y, so +x is to the *right*.)
	worldX := c.halfWidth - xOffset
	worldZ := c.halfHeight - zOffset

	// using the camera matrix, transform the canvas point and the origin,
	// and then compute the ray's direction vector.
	// (remember that the canvas is at y=-1)
	pixel := c.im.MultTuple(tuple.NewPoint(worldX, -1, worldZ))
	origin := c.im.MultTuple(tuple.Origin)
	direction := pixel.Sub(origin).Norm()

	return ray.NewRay(origin, direction)
}
