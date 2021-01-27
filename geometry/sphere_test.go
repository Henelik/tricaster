package geometry

import (
	"testing"

	"github.com/Henelik/tricaster/ray"
	"github.com/Henelik/tricaster/tuple"
	"github.com/stretchr/testify/assert"
)

func TestIntersects(t *testing.T) {
	s := NewSphere()
	testCases := []struct {
		name string
		r    *ray.Ray
		want []Intersection
	}{
		{
			name: "A ray intersects a sphere at two points",
			r: ray.NewRay(
				tuple.NewPoint(0, 0, -5),
				tuple.NewVector(0, 0, 1),
			),
			want: []Intersection{{4, s}, {6, s}},
		},
		{
			name: "A ray intersects a sphere at a tangent",
			r: ray.NewRay(
				tuple.NewPoint(0, 1, -5),
				tuple.NewVector(0, 0, 1),
			),
			want: []Intersection{{5, s}, {5, s}},
		},
		{
			name: "A ray misses a sphere",
			r: ray.NewRay(
				tuple.NewPoint(0, 2, -5),
				tuple.NewVector(0, 0, 1),
			),
			want: []Intersection{},
		},
		{
			name: "A ray originates inside a sphere",
			r: ray.NewRay(
				tuple.NewPoint(0, 0, 0),
				tuple.NewVector(0, 0, 1),
			),
			want: []Intersection{{-1, s}, {1, s}},
		},
		{
			name: "A sphere is behind a ray",
			r: ray.NewRay(
				tuple.NewPoint(0, 0, 5),
				tuple.NewVector(0, 0, 1),
			),
			want: []Intersection{{-6, s}, {-4, s}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			xs := s.Intersects(tc.r)
			assert.Equal(t, tc.want, xs)
		})
	}
}
