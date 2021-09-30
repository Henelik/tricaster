package geometry

import (
	"testing"

	"github.com/Henelik/tricaster/pkg/ray"
	"github.com/Henelik/tricaster/pkg/tuple"

	"github.com/stretchr/testify/assert"
)

func TestCube_Intersects(t *testing.T) {
	c := NewCube(nil, nil)
	testCases := []struct {
		name      string
		origin    *tuple.Tuple
		direction *tuple.Tuple
		t1        float64
		t2        float64
	}{
		{
			name:      "+x",
			origin:    tuple.NewPoint(5, 0.5, 0),
			direction: tuple.NewVector(-1, 0, 0),
			t1:        4,
			t2:        6,
		},
		{
			name:      "-x",
			origin:    tuple.NewPoint(-5, 0.5, 0),
			direction: tuple.NewVector(1, 0, 0),
			t1:        4,
			t2:        6,
		},
		{
			name:      "+y",
			origin:    tuple.NewPoint(0.5, 5, 0),
			direction: tuple.NewVector(0, -1, 0),
			t1:        4,
			t2:        6,
		},
		{
			name:      "-y",
			origin:    tuple.NewPoint(0.5, -5, 0),
			direction: tuple.NewVector(0, 1, 0),
			t1:        4,
			t2:        6,
		},
		{
			name:      "+z",
			origin:    tuple.NewPoint(0.5, 0, 5),
			direction: tuple.NewVector(0, 0, -1),
			t1:        4,
			t2:        6,
		},
		{
			name:      "-z",
			origin:    tuple.NewPoint(0.5, 0, -5),
			direction: tuple.NewVector(0, 0, 1),
			t1:        4,
			t2:        6,
		},
		{
			name:      "inside",
			origin:    tuple.NewPoint(0, 0.5, 0),
			direction: tuple.NewVector(0, 0, 1),
			t1:        -1,
			t2:        1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := ray.NewRay(tc.origin, tc.direction)
			xs := c.Intersects(r)
			assert.Equal(t, 2, len(xs))
			assert.Equal(t, tc.t1, xs[0].T)
			assert.Equal(t, tc.t2, xs[1].T)
		})
	}
}

func TestCube_Intersects_Miss(t *testing.T) {
	c := NewCube(nil, nil)
	testCases := []struct {
		origin    *tuple.Tuple
		direction *tuple.Tuple
	}{
		{
			origin:    tuple.NewPoint(-2, 0, 0),
			direction: tuple.NewVector(0.2673, 0.5345, 0.8018),
		},
		{
			origin:    tuple.NewPoint(0, -2, 0),
			direction: tuple.NewVector(0.8018, 0.2673, 0.5345),
		},
		{
			origin:    tuple.NewPoint(0, 0, -2),
			direction: tuple.NewVector(0.5345, 0.8018, 0.2673),
		},
		{
			origin:    tuple.NewPoint(2, 0, 2),
			direction: tuple.NewVector(0, 0, -1),
		},
		{
			origin:    tuple.NewPoint(0, 2, 2),
			direction: tuple.NewVector(0, -1, 0),
		},
		{
			origin:    tuple.NewPoint(2, 2, 0),
			direction: tuple.NewVector(-1, 0, 0),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.origin.Fmt()+tc.direction.Fmt(), func(t *testing.T) {
			r := ray.NewRay(tc.origin, tc.direction)
			xs := c.Intersects(r)
			assert.Equal(t, 0, len(xs))
		})
	}
}

func TestCube_NormalAt(t *testing.T) {
	c := NewCube(nil, nil)
	testCases := []struct {
		point *tuple.Tuple
		want  *tuple.Tuple
	}{
		{
			point: tuple.NewPoint(1, 0.5, -0.8),
			want:  tuple.NewVector(1, 0, 0),
		},
		{
			point: tuple.NewPoint(-1, -0.2, 0.9),
			want:  tuple.NewVector(-1, 0, 0),
		},
		{
			point: tuple.NewPoint(-0.4, 1, -0.1),
			want:  tuple.NewVector(0, 1, 0),
		},
		{
			point: tuple.NewPoint(0.3, -1, -0.7),
			want:  tuple.NewVector(0, -1, 0),
		},
		{
			point: tuple.NewPoint(-0.6, 0.3, 1),
			want:  tuple.NewVector(0, 0, 1),
		},
		{
			point: tuple.NewPoint(0.4, 0.4, -1),
			want:  tuple.NewVector(0, 0, -1),
		},
		{
			point: tuple.NewPoint(1, 1, 1),
			want:  tuple.NewVector(1, 0, 0),
		},
		{
			point: tuple.NewPoint(-1, -1, -1),
			want:  tuple.NewVector(-1, 0, 0),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.point.Fmt(), func(t *testing.T) {
			got := c.NormalAt(tc.point)
			assert.Equal(t, tc.want, got)
		})
	}
}
