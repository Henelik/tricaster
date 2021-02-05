package scene

import (
	"sync"

	"git.maze.io/go/math32"
	"github.com/Henelik/tricaster/color"

	"github.com/Henelik/tricaster/canvas"
	"github.com/Henelik/tricaster/matrix"
	"github.com/Henelik/tricaster/ray"
	"github.com/Henelik/tricaster/tuple"
)

var DefaultFOV float32 = math32.Pi / 2

type Camera struct {
	hSize      int
	vSize      int
	FOV        float32
	m          *matrix.Matrix
	im         *matrix.Matrix
	halfWidth  float32
	halfHeight float32
	pixelSize  float32
	AALevel    int
}

// NewCamera creates a new camera.
// Set fov to 0 and transform to nil to use defaults.
func NewCamera(hSize, vSize int, fov float32, transform *matrix.Matrix, aaLevel int) *Camera {
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

	switch aaLevel {
	case 2:
		c.AALevel = 2
	case 4:
		c.AALevel = 4
	case 8:
		c.AALevel = 8
	case 16:
		c.AALevel = 16
	default:
		c.AALevel = 1
	}

	halfView := math32.Tan(c.FOV / 2)
	aspect := float32(hSize) / float32(vSize)
	if aspect >= 1 {
		c.halfWidth = halfView
		c.halfHeight = halfView / aspect
	} else {
		c.halfWidth = halfView * aspect
		c.halfHeight = halfView
	}
	c.pixelSize = (c.halfWidth * 2) / float32(c.hSize)

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

func (c *Camera) RayForPixel(x, y int) *ray.Ray {
	// the offset from the edge of the canvas to the pixel's center
	xOffset := (float32(x) + 0.5) * c.pixelSize
	yOffset := (float32(y) + 0.5) * c.pixelSize

	// the untransformed coordinates of the pixel in world space.
	// (remember that the camera looks toward -z, so +x is to the *right*.)
	worldX := c.halfWidth - xOffset
	worldY := c.halfHeight - yOffset

	// using the camera matrix, transform the canvas point and the origin,
	// and then compute the ray's direction vector.
	// (remember that the canvas is at y=-1)
	pixel := c.im.MultTuple(tuple.NewPoint(worldX, worldY, -1))
	origin := c.im.MultTuple(tuple.Origin)
	direction := pixel.Sub(origin).Norm()

	return ray.NewRay(origin, direction)
}

func (c *Camera) AARaysForPixel(x, y int) []*ray.Ray {
	if c.AALevel == 0 {
		return []*ray.Ray{c.RayForPixel(x, y)}
	}
	rs := make([]*ray.Ray, 0, c.AALevel*c.AALevel)

	// the offset from the edge of the canvas to the pixel's center
	xOffset := (float32(x) + 0.5) * c.pixelSize
	yOffset := (float32(y) + 0.5) * c.pixelSize
	// the distance between sampled sub-pixel points on the canvas
	aaOffset := c.pixelSize / float32(c.AALevel)
	origin := c.im.MultTuple(tuple.Origin)

	for aax := 0; aax < c.AALevel; aax++ {
		for aay := 0; aay < c.AALevel; aay++ {
			// the untransformed coordinates of the pixel in world space.
			// (remember that the camera looks toward -z, so +x is to the *right*.)
			worldX := c.halfWidth - xOffset + float32(aax)*aaOffset
			worldY := c.halfHeight - yOffset + float32(aay)*aaOffset

			// using the camera matrix, transform the canvas point and the origin,
			// and then compute the ray's direction vector.
			// (remember that the canvas is at y=-1)
			pixel := c.im.MultTuple(tuple.NewPoint(worldX, worldY, -1))
			direction := pixel.Sub(origin).Norm()
			rs = append(rs, ray.NewRay(origin, direction))
		}
	}

	return rs
}

func (c *Camera) Render(w *World) *canvas.Canvas {
	canv := canvas.NewCanvas(c.hSize, c.vSize)
	for x := 0; x < c.hSize; x++ {
		for y := 0; y < c.vSize; y++ {
			r := c.RayForPixel(x, y)
			col := w.ColorAt(r, w.Config.MaxBounce, nil)
			canv.Set(x, y, col)
		}
	}
	return canv
}

// GoRender divides the image into an n*n grid and renders each cell in a goroutine
func (c *Camera) GoRender(w *World, gridNum int) *canvas.Canvas {
	canv := canvas.NewCanvas(c.hSize, c.vSize)

	// set up a wait group for the number of subdivisions
	var wg sync.WaitGroup
	wg.Add(gridNum * gridNum)

	subH := c.hSize / gridNum
	subV := c.vSize / gridNum

	worker := func(xs, ys int) {
		defer wg.Done()
		for x := xs; x < subH+xs; x++ {
			for y := ys; y < subV+ys; y++ {
				rs := c.AARaysForPixel(x, y)
				cols := make([]*color.Color, len(rs))
				for i, r := range rs {
					cols[i] = w.ColorAt(r, w.Config.MaxBounce, nil)
				}
				canv.Set(x, y, color.Avg(cols))
			}
		}
	}

	for sh := 0; sh < gridNum; sh++ {
		for sv := 0; sv < gridNum; sv++ {
			go worker(sh*subH, sv*subV)
		}
	}

	wg.Wait()
	return canv
}
