package geometry

import (
	"testing"

	"github.com/Henelik/tricaster/matrix"
	"github.com/Henelik/tricaster/ray"
	"github.com/Henelik/tricaster/tuple"
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
			want: NilHit,
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

func BenchmarkIntersection128(b *testing.B) {
	var hit bool
	for n := 0; n < b.N; n++ {
		s := NewSphere(matrix.Translation(0, 5, 0))

		w := 128
		h := 128

		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				// shoot a ray at the sphere
				xPos := (float64(x) - float64(w)/2.0) / (float64(w) * 0.4)
				yPos := -(float64(y) - float64(h)/2.0) / (float64(h) * 0.4)
				r := ray.NewRay(
					tuple.NewPoint(xPos, 0, yPos),
					tuple.NewVector(0, 1, 0),
				)
				h := Hit(s.Intersects(r))
				if h != NilHit {
					hit = true
				} else {
					hit = false
				}
			}
		}
	}
	if hit {
		return
	}
}
