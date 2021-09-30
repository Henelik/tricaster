package ray

import (
	"testing"

	"github.com/Henelik/tricaster/pkg/matrix"
	"github.com/Henelik/tricaster/pkg/tuple"

	"github.com/stretchr/testify/assert"
)

func TestPosition(t *testing.T) {
	r := NewRay(
		tuple.NewPoint(2, 3, 4),
		tuple.NewVector(1, 0, 0),
	)
	assert.Equal(t,
		tuple.NewPoint(2, 3, 4),
		r.Position(0),
	)
	assert.Equal(t,
		tuple.NewPoint(3, 3, 4),
		r.Position(1),
	)
	assert.Equal(t,
		tuple.NewPoint(1, 3, 4),
		r.Position(-1),
	)
	assert.Equal(t,
		tuple.NewPoint(4.5, 3, 4),
		r.Position(2.5),
	)
}

func TestTransform(t *testing.T) {
	testCases := []struct {
		name string
		r    *Ray
		m    *matrix.Matrix
		want *Ray
	}{
		{
			name: "The hit, when all intersections have positive t",
			r: NewRay(
				tuple.NewPoint(1, 2, 3),
				tuple.NewVector(0, 1, 0),
			),
			m: matrix.Translation(3, 4, 5),
			want: NewRay(
				tuple.NewPoint(4, 6, 8),
				tuple.NewVector(0, 1, 0),
			),
		},
		{
			name: "Scaling a ray",
			r: NewRay(
				tuple.NewPoint(1, 2, 3),
				tuple.NewVector(0, 1, 0),
			),
			m: matrix.Scaling(2, 3, 4),
			want: NewRay(
				tuple.NewPoint(2, 6, 12),
				tuple.NewVector(0, 3, 0),
			),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.r.Transform(tc.m))
		})
	}
}
