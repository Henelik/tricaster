package geometry

import (
	"fmt"
	"testing"

	"github.com/Henelik/tricaster/pkg/ray"
	"github.com/Henelik/tricaster/pkg/tuple"
	"github.com/stretchr/testify/assert"
)

func TestCylinder_Intersects_Miss(t *testing.T) {
	cyl := NewCylinder(-1, 1, false, nil, nil)

	testCases := []struct {
		origin    *tuple.Tuple
		direction *tuple.Tuple
	}{
		{
			origin:    tuple.NewPoint(1, 0, 0),
			direction: tuple.NewVector(0, 0, 1),
		},
		{
			origin:    tuple.NewPoint(0, 0, 0),
			direction: tuple.NewVector(0, 0, 1),
		},
		{
			origin:    tuple.NewPoint(0, -5, 0),
			direction: tuple.NewVector(1, 1, 1),
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v, %v", tc.origin, tc.direction), func(t *testing.T) {
			inters := cyl.Intersects(ray.NewRay(
				tc.origin,
				tc.direction.Norm()))

			assert.Equal(t, 0, len(inters))
		})
	}
}

func TestCylinder_Intersects_Hit(t *testing.T) {
	cyl := NewCylinder(-10, 10, false, nil, nil)

	testCases := []struct {
		origin    *tuple.Tuple
		direction *tuple.Tuple
		t0        float64
		t1        float64
	}{
		{
			origin:    tuple.NewPoint(1, -5, 0),
			direction: tuple.NewVector(0, 1, 0),
			t0:        5,
			t1:        5,
		},
		{
			origin:    tuple.NewPoint(0, -5, 0),
			direction: tuple.NewVector(0, 1, 0),
			t0:        4,
			t1:        6,
		},
		{
			origin:    tuple.NewPoint(0.5, -5, 0),
			direction: tuple.NewVector(0.1, 1, 1),
			t0:        6.80798191702732,
			t1:        7.088723439378861,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v, %v", tc.origin, tc.direction), func(t *testing.T) {
			inters := cyl.Intersects(ray.NewRay(
				tc.origin,
				tc.direction.Norm()))

			if assert.Equal(t, 2, len(inters)) {
				assert.Equal(t, tc.t0, inters[0].T)

				assert.Equal(t, tc.t1, inters[1].T)
			}
		})
	}
}

func TestCylinder_NormalAt(t *testing.T) {
	cyl := NewCylinder(-1, 1, false, nil, nil)

	testCases := []struct {
		point *tuple.Tuple
		want  *tuple.Tuple
	}{
		{
			point: tuple.NewPoint(1, 0, 0),
			want:  tuple.NewVector(1, 0, 0),
		},
		{
			point: tuple.NewPoint(0, -1, 5),
			want:  tuple.NewVector(0, -1, 0),
		},
		{
			point: tuple.NewPoint(0, 1, -2),
			want:  tuple.NewVector(0, 1, 0),
		},
		{
			point: tuple.NewPoint(-1, 0, 1),
			want:  tuple.NewVector(-1, 0, 0),
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.point), func(t *testing.T) {
			got := cyl.NormalAt(tc.point)

			assert.Equal(t, tc.want, got)
		})
	}
}

func TestCylinder_NormalAt_Caps(t *testing.T) {
	cyl := NewCylinder(1, 2, false, nil, nil)

	testCases := []struct {
		point *tuple.Tuple
		want  *tuple.Tuple
	}{
		{
			point: tuple.NewPoint(0, 0, 1),
			want:  tuple.NewVector(0, 0, -1),
		},
		{
			point: tuple.NewPoint(0.5, 0, 1),
			want:  tuple.NewVector(0, 0, -1),
		},
		{
			point: tuple.NewPoint(0, 0.5, 1),
			want:  tuple.NewVector(0, 0, -1),
		},
		{
			point: tuple.NewPoint(0, 0, 2),
			want:  tuple.NewVector(0, 0, 1),
		},
		{
			point: tuple.NewPoint(0.5, 0, 2),
			want:  tuple.NewVector(0, 0, 1),
		},
		{
			point: tuple.NewPoint(0, 0.5, 2),
			want:  tuple.NewVector(0, 0, 1),
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.point), func(t *testing.T) {
			got := cyl.NormalAt(tc.point)

			assert.Equal(t, tc.want, got)
		})
	}
}

func TestCylinder_Intersects_Capped(t *testing.T) {
	cyl := NewCylinder(1, 2, false, nil, nil)

	testCases := []struct {
		origin    *tuple.Tuple
		direction *tuple.Tuple
		count     int
	}{
		{
			origin:    tuple.NewPoint(0, 0, 1.5),
			direction: tuple.NewVector(0.1, 0, 1),
			count:     0,
		},
		{
			origin:    tuple.NewPoint(0, -5, 3),
			direction: tuple.NewVector(0, 1, 0),
			count:     0,
		},
		{
			origin:    tuple.NewPoint(0, -5, 0),
			direction: tuple.NewVector(0, 1, 0),
			count:     0,
		},
		{
			origin:    tuple.NewPoint(0, -5, 2),
			direction: tuple.NewVector(0, 1, 0),
			count:     0,
		},
		{
			origin:    tuple.NewPoint(0, -5, 1),
			direction: tuple.NewVector(0, 1, 0),
			count:     0,
		},
		{
			origin:    tuple.NewPoint(0, -5, 1.5),
			direction: tuple.NewVector(0, 1, 0),
			count:     2,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v, %v", tc.origin, tc.direction), func(t *testing.T) {
			inters := cyl.Intersects(ray.NewRay(
				tc.origin,
				tc.direction.Norm()))

			assert.Equal(t, tc.count, len(inters))
		})
	}
}

func TestCylinder_Intersects_Cap(t *testing.T) {
	cyl := NewCylinder(1, 2, true, nil, nil)

	testCases := []struct {
		point     *tuple.Tuple
		direction *tuple.Tuple
		count     int
	}{
		{
			point:     tuple.NewPoint(0, 0, 3),
			direction: tuple.NewVector(0, 0, -1),
			count:     2,
		},
		{
			point:     tuple.NewPoint(0, -2, 3),
			direction: tuple.NewVector(0, 2, -1),
			count:     2,
		},
		{
			point:     tuple.NewPoint(0, -2, 4),
			direction: tuple.NewVector(0, 1, -1),
			count:     2,
		},
		{
			point:     tuple.NewPoint(0, -2, 0),
			direction: tuple.NewVector(0, 2, 1),
			count:     2,
		},
		{
			point:     tuple.NewPoint(0, -2, -1),
			direction: tuple.NewVector(0, 1, 1),
			count:     2,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v, %v", tc.point, tc.direction), func(t *testing.T) {
			inters := cyl.Intersects(ray.NewRay(
				tc.point,
				tc.direction.Norm()))

			assert.Equal(t, tc.count, len(inters))
		})
	}
}
