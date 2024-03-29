package renderer

import (
	"math"
	"testing"

	"github.com/Henelik/tricaster/pkg/color"
	"github.com/Henelik/tricaster/pkg/geometry"
	"github.com/Henelik/tricaster/pkg/light"
	"github.com/Henelik/tricaster/pkg/material"
	"github.com/Henelik/tricaster/pkg/matrix"
	"github.com/Henelik/tricaster/pkg/ray"
	"github.com/Henelik/tricaster/pkg/tuple"

	"github.com/stretchr/testify/assert"
)

func TestIntersect(t *testing.T) {
	// Intersect a world with a ray
	got := ray.SimpleSort(DefaultWorld.Intersect(
		ray.NewRay(
			tuple.NewPoint(0, 0, -5),
			tuple.NewVector(0, 0, 1))))

	assert.Equal(t, 4, len(got))
	assert.Equal(t, 4.0, got[0].T)
	assert.Equal(t, 4.5, got[1].T)
	assert.Equal(t, 5.5, got[2].T)
	assert.Equal(t, 6.0, got[3].T)
}

func TestShading(t *testing.T) {
	// Shading an intersection
	r := ray.NewRay(tuple.NewPoint(-5, 0, 0), tuple.Right)
	s := DefaultWorld.Geometry[0]
	i := ray.Intersection{4, s}
	col := DefaultWorld.Shade(ray.NewHit(r, []ray.Intersection{i}, 0), 0)
	assert.Equal(t, color.NewColor(0.38066119308103435, 0.47582649135129296, 0.28549589481077575), col)

	// Shading an intersection from the inside
	w := *DefaultWorld
	w.Light = &light.PointLight{tuple.NewPoint(0, 0.25, 0), color.White}
	r2 := ray.NewRay(tuple.Origin, tuple.Right)
	s2 := DefaultWorld.Geometry[1]
	i2 := ray.Intersection{0.5, s2}
	col2 := w.Shade(ray.NewHit(r2, []ray.Intersection{i2}, 0), 0)
	assert.Equal(t, color.Grey(0.9049844720832575), col2)
}

func TestColorAtMiss(t *testing.T) {
	// The color when a ray misses
	r := ray.NewRay(tuple.NewPoint(0, 0, -5), tuple.Forward)
	c := DefaultWorld.ColorAt(r, 0)
	assert.Equal(t, color.Black, c)
}

func TestColorAtHit(t *testing.T) {
	// The color when a ray hits
	r := ray.NewRay(tuple.NewPoint(-5, 0, 0), tuple.Right)
	col := DefaultWorld.ColorAt(r, 0)
	assert.Equal(t, color.NewColor(0.38066119308103435, 0.47582649135129296, 0.28549589481077575), col)
}

func TestColorAtHitBehind(t *testing.T) {
	// The color when a ray hits
	w := &World{
		Geometry: []Primitive{
			geometry.NewSphere(nil, &material.PhongMat{
				Ambient: 1,
				Color:   color.Red,
			}),
			geometry.NewSphere(
				matrix.Scaling(0.5, 0.5, 0.5),
				&material.PhongMat{
					Ambient: 1,
					Color:   color.Blue,
				}),
		},
		Light: &light.PointLight{
			tuple.NewPoint(-10, -10, 10),
			color.White,
		},
		Config: &WorldConfig{
			Shadows: false,
		},
	}
	r := ray.NewRay(tuple.NewPoint(0.75, 0, 0), tuple.Left)
	col := w.ColorAt(r, 0)
	assert.Equal(t, color.Blue, col)
}

func TestShadow(t *testing.T) {
	w := *DefaultWorld
	w.Config = &WorldConfig{Shadows: true}
	testCases := []struct {
		name string
		p    *tuple.Tuple
		want bool
	}{
		{
			name: "There is no shadow when nothing is collinear with point and light",
			p:    tuple.NewPoint(0, 0, 10),
			want: false,
		},
		{
			name: "The shadow when an object is between the point and the light",
			p:    tuple.NewPoint(10, 10, -10),
			want: true,
		},
		{
			name: "There is no shadow when an object is behind the light",
			p:    tuple.NewPoint(-20, -20, 20),
			want: false,
		},
		{
			name: "There is no shadow when an object is behind the point",
			p:    tuple.NewPoint(-2, -2, 2),
			want: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, w.IsShadowed(tc.p))
		})
	}
}

// The refracted color with an opaque surface
func TestWorldRefractOpaque(t *testing.T) {
	s := DefaultWorld.Geometry[0]
	r := ray.NewRay(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))
	inters := []ray.Intersection{{4, s}, {6, s}}
	h := ray.NewHit(r, inters, 0)
	col := DefaultWorld.RefractedColor(h, 3)
	assert.Equal(t, color.Black, col)
}

// The refracted color at the maximum recursive depth
func TestWorldRefractMax(t *testing.T) {
	w := *DefaultWorld
	s := w.Geometry[0].(*geometry.Sphere)
	s.Mat = &material.PhongMat{
		Ambient:      0.1,
		Diffuse:      0.5,
		Specular:     0,
		Shininess:    0,
		Reflectivity: 0,
		Transparency: 1,
		IOR:          1.5,
		Color:        color.Red,
		Pattern:      nil,
	}

	r := ray.NewRay(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))
	inters := []ray.Intersection{{4, s}, {6, s}}
	h := ray.NewHit(r, inters, 0)
	col := DefaultWorld.RefractedColor(h, 0)
	assert.Equal(t, color.Black, col)
}

// The refracted color under total internal reflection
func TestWorldRefractTotalInternalReflection(t *testing.T) {
	w := *DefaultWorld
	s := w.Geometry[0].(*geometry.Sphere)
	s.Mat = &material.PhongMat{
		Ambient:      0,
		Diffuse:      0,
		Specular:     0,
		Shininess:    0,
		Reflectivity: 0,
		Transparency: 1,
		IOR:          1.5,
		Color:        color.White,
		Pattern:      nil,
	}

	r := ray.NewRay(tuple.NewPoint(0, 0, math.Sqrt2/2), tuple.NewVector(0, 1, 0))
	inters := []ray.Intersection{{-math.Sqrt2 / 2, s}, {math.Sqrt2 / 2, s}}
	h := ray.NewHit(r, inters, 1)
	col := w.RefractedColor(h, 3)
	assert.Equal(t, color.Black, col)
}

// The refracted color with a refracted ray
func TestWorldRefractRefract(t *testing.T) {
	w := *DefaultWorld

	a := w.Geometry[0].(*geometry.Sphere)
	a.Mat = &material.PhongMat{
		Ambient:      1,
		Diffuse:      0,
		Specular:     0,
		Shininess:    0,
		Reflectivity: 0,
		Transparency: 0,
		IOR:          1,
		Pattern:      NewTestPattern(nil),
	}

	b := w.Geometry[1].(*geometry.Sphere)
	b.Mat = &material.PhongMat{
		Ambient:      0,
		Diffuse:      0,
		Specular:     0,
		Shininess:    0,
		Reflectivity: 0,
		Transparency: 1,
		IOR:          1.5,
		Color:        color.White,
	}
	r := ray.NewRay(tuple.NewPoint(0, 0, 0.1), tuple.NewVector(0, 1, 0))
	inters := []ray.Intersection{{-0.9899, a}, {-0.4899, b}, {0.4899, b}, {0.9899, a}}
	h := ray.NewHit(r, inters, 2)
	col := w.RefractedColor(h, 5)
	assert.Equal(t, color.NewColor(0, 0.998884682797801, 0.04721642163417859), col)
}
