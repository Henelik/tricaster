package renderer

import (
	"math"
	"testing"

	"github.com/Henelik/tricaster/pkg/canvas"
	"github.com/Henelik/tricaster/pkg/color"
	geometry2 "github.com/Henelik/tricaster/pkg/geometry"
	"github.com/Henelik/tricaster/pkg/light"
	"github.com/Henelik/tricaster/pkg/material"
	"github.com/Henelik/tricaster/pkg/matrix"
	"github.com/Henelik/tricaster/pkg/tuple"

	"github.com/stretchr/testify/assert"
)

func TestPixelSize(t *testing.T) {
	// The pixel size for a horizontal canvas
	c1 := NewCamera(200, 125, 0, nil, 0)
	assert.Equal(t, 0.01, c1.pixelSize)

	// The pixel size for a vertical canvas
	c2 := NewCamera(125, 200, 0, nil, 0)
	assert.Equal(t, 0.01, c2.pixelSize)
}

func TestRayForPixel(t *testing.T) {
	// Constructing a ray through the center of the canvas
	c := NewCamera(201, 101, 0, nil, 0)
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
	floorMat := material.DefaultPhong.CopyWithColor(color.NewColor(1, 0.9, 0.9))
	floorMat.Specular = 0
	floor := geometry2.NewPlane(
		matrix.Identity,
		floorMat)

	middle := geometry2.NewSphere(
		matrix.Translation(5, 5, 2).Mult(matrix.Scaling(2, 2, 2)),
		&material.PhongMat{
			Ambient:   0.1,
			Diffuse:   0.9,
			Specular:  0.0,
			Shininess: 10,
			Color:     color.NewColor(0.1, 1, 0.5),
		})

	left := geometry2.NewSphere(
		matrix.Translation(2, -2, 1),
		&material.PhongMat{
			Ambient:   0.1,
			Diffuse:   0.9,
			Specular:  0.9,
			Shininess: 200,
			Color:     color.NewColor(1, 0.1, 0.1),
		})

	right := geometry2.NewSphere(
		matrix.Translation(-4, 3, 1.25).Mult(matrix.ScalingU(1.25)),
		&material.PhongMat{
			Ambient:   0.1,
			Diffuse:   0.9,
			Specular:  0.9,
			Shininess: 200,
			Color:     color.NewColor(0.2, 0.2, 1),
		})

	w := &World{
		Geometry: []Primitive{
			floor,
			middle,
			left,
			right,
		},
		Light: &light.PointLight{
			Pos:   tuple.NewPoint(0, -10, 10),
			Color: color.White,
		},
		Config: WorldConfig{
			Shadows: true,
		},
	}

	c := NewCamera(1000, 500, math.Pi/3,
		matrix.ViewTransform(
			tuple.NewPoint(-15, -10, 5),
			tuple.NewPoint(3, 3, 2),
			tuple.Up), 0)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		canv = c.Render(w)
	}
	assert.NotNil(b, canv)
}
