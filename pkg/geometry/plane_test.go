package geometry

import (
	"testing"

	"github.com/Henelik/tricaster/pkg/material"
	"github.com/Henelik/tricaster/pkg/matrix"
	"github.com/Henelik/tricaster/pkg/ray"
	"github.com/Henelik/tricaster/pkg/tuple"

	"github.com/stretchr/testify/assert"
)

func TestPlane_Intersects(t *testing.T) {
	p := NewPlane(matrix.Identity, material.DefaultPhong)
	testCases := []struct {
		name string
		r    *ray.Ray
		want []ray.Intersection
	}{
		{
			name: "Intersect with a ray parallel to the plane",
			r: ray.NewRay(
				tuple.NewPoint(0, 0, 10),
				tuple.Forward,
			),
			want: []ray.Intersection{},
		},
		{
			name: "Intersect with a coplanar ray",
			r: ray.NewRay(
				tuple.NewPoint(0, 0, 0),
				tuple.Forward,
			),
			want: []ray.Intersection{},
		},
		{
			name: "A ray intersecting a plane from above",
			r: ray.NewRay(
				tuple.NewPoint(0, 0, 1),
				tuple.Down,
			),
			want: []ray.Intersection{{1, p}},
		},
		{
			name: "A ray intersecting a plane from below",
			r: ray.NewRay(
				tuple.NewPoint(0, 0, -1),
				tuple.Up,
			),
			want: []ray.Intersection{{1, p}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := p.Intersects(tc.r)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestPlane_NormalAt(t *testing.T) {
	testCases := []struct {
		name  string
		p     *Plane
		point *tuple.Tuple
		want  *tuple.Tuple
	}{
		{
			name:  "The normal of a plane is constant everywhere 1",
			p:     NewPlane(matrix.Identity, material.DefaultPhong),
			point: tuple.NewPoint(1, 0, 0),
			want:  tuple.Up,
		},
		{
			name:  "The normal of a plane is constant everywhere 2",
			p:     NewPlane(matrix.Identity, material.DefaultPhong),
			point: tuple.NewPoint(10, 0, -10),
			want:  tuple.Up,
		},
		{
			name:  "The normal of a plane is constant everywhere 3",
			p:     NewPlane(matrix.Identity, material.DefaultPhong),
			point: tuple.NewPoint(-5, 0, 150),
			want:  tuple.Up,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.True(t, tc.want.Equal(tc.p.NormalAt(tc.point)))
		})
	}
}
