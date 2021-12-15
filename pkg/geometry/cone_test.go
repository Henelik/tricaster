package geometry

import (
	"fmt"
	"math"
	"testing"

	"github.com/Henelik/tricaster/pkg/ray"
	"github.com/Henelik/tricaster/pkg/tuple"
	"github.com/stretchr/testify/assert"
)

func TestCone_Intersects_Hit(t *testing.T) {
	cyl := NewCone(-10, 10, false, nil, nil)

	testCases := []struct {
		origin    *tuple.Tuple
		direction *tuple.Tuple
		t0        float64
		t1        float64
	}{
		{
			origin:    tuple.NewPoint(0, -5, 0),
			direction: tuple.NewVector(0, 1, 0),
			t0:        5,
			t1:        5,
		},
		{
			origin:    tuple.NewPoint(0, -5, 0),
			direction: tuple.NewVector(1, 1, 1),
			t0:        8.66025,
			t1:        8.66025,
		},
		{
			origin:    tuple.NewPoint(1, -5, 1),
			direction: tuple.NewVector(-0.5, -1, 1),
			t0:        4.55006,
			t1:        49.44994,
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

func TestCone_NormalAt(t *testing.T) {
	cyl := NewCone(-1, 1, false, nil, nil)

	testCases := []struct {
		point *tuple.Tuple
		want  *tuple.Tuple
	}{
		{
			point: tuple.NewPoint(0, 0, 0),
			want:  tuple.NewVector(0, 0, 0),
		},
		{
			point: tuple.NewPoint(1, 1, 1),
			want:  tuple.NewVector(1, 1, -math.Sqrt(2)),
		},
		{
			point: tuple.NewPoint(-1, 0, -1),
			want:  tuple.NewVector(-1, 0, 1),
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.point), func(t *testing.T) {
			got := cyl.NormalAt(tc.point)

			assert.Equal(t, tc.want, got)
		})
	}
}

func TestCone_Intersects_Caps(t *testing.T) {
	cyl := NewCone(-0.5, 0.5, false, nil, nil)

	testCases := []struct {
		origin    *tuple.Tuple
		direction *tuple.Tuple
		count     int
	}{
		{
			origin:    tuple.NewPoint(0, -5, 0),
			direction: tuple.NewVector(0, 0, 1),
			count:     0,
		},
		{
			origin:    tuple.NewPoint(0, -0.25, 0),
			direction: tuple.NewVector(0, 1, 1),
			count:     2,
		},
		{
			origin:    tuple.NewPoint(0, -0.25, 0),
			direction: tuple.NewVector(0, 0, 1),
			count:     4,
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
