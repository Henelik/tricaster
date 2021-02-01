package scene

import (
	"math"
	"testing"

	"github.com/Henelik/tricaster/canvas"

	"github.com/Henelik/tricaster/color"
	"github.com/Henelik/tricaster/geometry"
	"github.com/Henelik/tricaster/shading"

	"github.com/Henelik/tricaster/matrix"
	"github.com/Henelik/tricaster/tuple"
	"github.com/stretchr/testify/assert"
)

func TestPixelSize(t *testing.T) {
	// The pixel size for a horizontal canvas
	c1 := NewCamera(200, 125, 0, nil)
	assert.Equal(t, 0.01, c1.pixelSize)

	// The pixel size for a vertical canvas
	c2 := NewCamera(125, 200, 0, nil)
	assert.Equal(t, 0.01, c2.pixelSize)
}

func TestRayForPixel(t *testing.T) {
	// Constructing a ray through the center of the canvas
	c := NewCamera(201, 101, 0, nil)
	r := c.RayForPixel(100, 50)
	assert.Equal(t, tuple.Origin, r.Origin)
	assert.Equal(t, tuple.Down, r.Direction)

	// Constructing a ray through a corner of the canvas
	r1 := c.RayForPixel(0, 0)
	assert.Equal(t, tuple.Origin, r1.Origin)
	assert.Equal(t, tuple.NewVector(0.6651864261194509, 0.33259321305972545, -0.6685123582500481), r1.Direction)

	// Constructing a ray when the camera is transformed
	c.SetMatrix(matrix.RotationX(math.Pi / 4).Mult(matrix.Translation(0, -2, 5)))
	r2 := c.RayForPixel(100, 50)
	assert.True(t, tuple.NewPoint(0, 2, -5).Equal(r2.Origin))
	assert.True(t, tuple.NewVector(0, -math.Sqrt2/2, -math.Sqrt2/2).Equal(r2.Direction))
}

func BenchmarkRender(b *testing.B) {
	var canv *canvas.Canvas
	floor := geometry.NewSphere(
		matrix.Scaling(15, 15, 0.01),
		shading.DefaultPhong.CopyWithColor(color.NewColor(1, 0.9, 0.9)))
	floor.Mat.Specular = 0

	leftWall := geometry.NewSphere(
		matrix.Translation(0, 10, 0).Mult(
			matrix.RotationY(-math.Pi/4).Mult(
				matrix.RotationX(math.Pi/2).Mult(
					matrix.Scaling(15, 15, 0.01)))),
		floor.Mat)

	rightWall := geometry.NewSphere(
		matrix.Translation(10, 0, 0).Mult(
			matrix.RotationZ(math.Pi/2).Mult(
				matrix.RotationX(math.Pi/2).Mult(
					matrix.Scaling(15, 15, 0.01)))),
		floor.Mat)

	middle := geometry.NewSphere(
		matrix.Translation(5, 5, 2).Mult(matrix.Scaling(2, 2, 2)),
		&shading.PhongMat{
			Ambient:   0.1,
			Diffuse:   0.9,
			Specular:  0.0,
			Shininess: 10,
			Color:     color.NewColor(0.1, 1, 0.5),
		})

	left := geometry.NewSphere(
		matrix.Translation(2, -2, 1),
		&shading.PhongMat{
			Ambient:   0.1,
			Diffuse:   0.9,
			Specular:  0.9,
			Shininess: 200,
			Color:     color.NewColor(1, 0.1, 0.1),
		})

	right := geometry.NewSphere(
		matrix.Translation(-4, 3, 1.25).Mult(matrix.ScalingU(1.25)),
		&shading.PhongMat{
			Ambient:   0.1,
			Diffuse:   0.9,
			Specular:  0.9,
			Shininess: 200,
			Color:     color.NewColor(0.2, 0.2, 1),
		})

	w := &World{
		Geometry: []geometry.Primitive{
			floor,
			leftWall,
			rightWall,
			middle,
			left,
			right,
		},
		Light: &shading.PointLight{
			Pos:   tuple.NewPoint(0, -10, 10),
			Color: color.White,
		},
	}

	c := NewCamera(1000, 500, math.Pi/3,
		matrix.ViewTransform(
			tuple.NewPoint(-15, -10, 5),
			tuple.NewPoint(3, 3, 2),
			tuple.Up))

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		canv = c.GoRender(2, w)
	}
	assert.NotNil(b, canv)
}
