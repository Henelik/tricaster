package geometry

import (
	"math"
	"testing"

	"github.com/Henelik/tricaster/pkg/material"
	"github.com/Henelik/tricaster/pkg/matrix"
	"github.com/Henelik/tricaster/pkg/ray"
	"github.com/Henelik/tricaster/pkg/tuple"

	"github.com/stretchr/testify/assert"
)

func TestSphere_Intersects(t *testing.T) {
	s := NewSphere(matrix.Identity, material.DefaultPhong)
	testCases := []struct {
		name string
		r    *ray.Ray
		want []ray.Intersection
	}{
		{
			name: "A ray intersects a sphere at two points",
			r: ray.NewRay(
				tuple.NewPoint(0, 0, -5),
				tuple.Up,
			),
			want: []ray.Intersection{{4, s}, {6, s}},
		},
		{
			name: "A ray intersects a sphere at a tangent",
			r: ray.NewRay(
				tuple.NewPoint(0, 1, -5),
				tuple.Up,
			),
			want: []ray.Intersection{{5, s}, {5, s}},
		},
		{
			name: "A ray misses a sphere",
			r: ray.NewRay(
				tuple.NewPoint(0, 2, -5),
				tuple.Up,
			),
			want: []ray.Intersection{},
		},
		{
			name: "A ray originates inside a sphere",
			r: ray.NewRay(
				tuple.Origin,
				tuple.Up,
			),
			want: []ray.Intersection{{-1, s}, {1, s}},
		},
		{
			name: "A sphere is behind a ray",
			r: ray.NewRay(
				tuple.NewPoint(0, 0, 5),
				tuple.Up,
			),
			want: []ray.Intersection{{-6, s}, {-4, s}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			xs := s.Intersects(tc.r)
			assert.Equal(t, tc.want, xs)
		})
	}
}

func TestSphere_IntersectsTransformed(t *testing.T) {
	r := ray.NewRay(
		tuple.NewPoint(0, 0, -5),
		tuple.Up,
	)

	// Intersecting a scaled sphere with a ray
	s := NewSphere(matrix.Scaling(2, 2, 2), material.DefaultPhong)
	want := []ray.Intersection{
		{3, s},
		{7, s},
	}

	assert.Equal(t, want, s.Intersects(r))

	// Intersecting a translated sphere with a ray
	s2 := NewSphere(matrix.Translation(5, 0, 0), material.DefaultPhong)
	want2 := []ray.Intersection{}

	assert.Equal(t, want2, s2.Intersects(r))
}

func TestSphere_NormalAt(t *testing.T) {
	testCases := []struct {
		name string
		s    *Sphere
		p    *tuple.Tuple
		want *tuple.Tuple
	}{
		{
			name: "The normal on a sphere at a point on the x axis",
			s:    NewSphere(matrix.Identity, material.DefaultPhong),
			p:    tuple.NewPoint(1, 0, 0),
			want: tuple.Right,
		},
		{
			name: "The normal on a sphere at a point on the y axis",
			s:    NewSphere(matrix.Identity, material.DefaultPhong),
			p:    tuple.NewPoint(0, 1, 0),
			want: tuple.Backward,
		},
		{
			name: "The normal on a sphere at a point on the z axis",
			s:    NewSphere(matrix.Identity, material.DefaultPhong),
			p:    tuple.NewPoint(0, 0, 1),
			want: tuple.Up,
		},
		{
			name: "The normal on a sphere at a nonaxial point",
			s:    NewSphere(matrix.Identity, material.DefaultPhong),
			p:    tuple.NewPoint(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3),
			want: tuple.NewVector(1, 1, 1).Norm(),
		},
		{
			name: "Computing the normal on a translated sphere",
			s:    NewSphere(matrix.Translation(0, 1, 0), material.DefaultPhong),
			p:    tuple.NewPoint(0, 1.70711, -0.70711),
			want: tuple.NewVector(0, 1, -1).Norm(),
		},
		{
			name: "Computing the normal on a transformed sphere",
			s: NewSphere(
				matrix.Scaling(1, 0.5, 1).Mult(matrix.RotationZ(math.Pi/5)),
				material.DefaultPhong),
			p:    tuple.NewPoint(0, math.Sqrt(2)/2, -math.Sqrt(2)/2).Norm(),
			want: tuple.NewVector(0, 0.970160000001, -0.24254).Norm(),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.True(t, tc.want.Equal(tc.s.NormalAt(tc.p)))
		})
	}
}
