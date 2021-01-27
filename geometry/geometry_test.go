package geometry

import (
	"testing"

	"github.com/Henelik/tricaster/matrix"
	"github.com/stretchr/testify/assert"
)

func TestHit(t *testing.T) {
	s := NewSphere(matrix.Identity)
	testCases := []struct {
		name   string
		inters []Intersection
		want   Intersection
	}{
		{
			name: "The hit, when all intersections have positive t",
			inters: []Intersection{
				{1, s},
				{2, s},
			},
			want: Intersection{1, s},
		},
		{
			name: "The hit, when some intersections have negative t",
			inters: []Intersection{
				{-1, s},
				{1, s},
			},
			want: Intersection{1, s},
		},
		{
			name: "The hit, when all intersections have negative t",
			inters: []Intersection{
				{-2, s},
				{-1, s},
			},
			want: NilIntersection,
		},
		{
			name: "The hit is always the lowest nonnegative intersection",
			inters: []Intersection{
				{5, s},
				{7, s},
				{-3, s},
				{2, s},
			},
			want: Intersection{2, s},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Hit(tc.inters))
		})
	}
}
