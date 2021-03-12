package renderer

import (
	"testing"

	"github.com/Henelik/tricaster/matrix"
	"github.com/Henelik/tricaster/tuple"
	"github.com/stretchr/testify/assert"
)

func TestHit(t *testing.T) {
	s := NewSphere(matrix.Identity, DefaultPhong)
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
			want: *NilIntersect,
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
			assert.Equal(t, tc.want, *GetClosest(tc.inters))
		})
	}
}

func BenchmarkIntersection128(b *testing.B) {
	var hit bool
	for n := 0; n < b.N; n++ {
		s := NewSphere(matrix.Translation(0, 5, 0), DefaultPhong)

		w := 128
		h := 128

		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				// shoot a ray at the sphere
				xPos := (float64(x) - float64(w)/2.0) / (float64(w) * 0.4)
				yPos := -(float64(y) - float64(h)/2.0) / (float64(h) * 0.4)
				r := NewRay(
					tuple.NewPoint(xPos, 0, yPos),
					tuple.Backward,
				)
				i := GetClosest(s.Intersects(r))
				if i != NilIntersect {
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

func TestToHit(t *testing.T) {
	r := NewRay(
		tuple.NewPoint(0, 0, -5),
		tuple.Up)
	s := NewSphere(nil, nil)
	i := &Intersection{4, s}

	h := i.ToHit(r, []Intersection{*i})

	assert.Equal(t, i.T, h.T)
	assert.Equal(t, i.P, h.P)
	assert.Equal(t, tuple.NewPoint(0, 0, -1), h.Pos)
	assert.Equal(t, tuple.Down, h.EyeV)
	assert.Equal(t, tuple.Down, h.NormalV)
	assert.Equal(t, false, h.Inside)
}

func TestToHitInside(t *testing.T) {
	r := NewRay(
		tuple.Origin,
		tuple.Up)
	s := NewSphere(nil, nil)
	i := &Intersection{1, s}

	h := i.ToHit(r, []Intersection{*i})

	assert.Equal(t, i.T, h.T)
	assert.Equal(t, i.P, h.P)
	assert.Equal(t, tuple.NewPoint(0, 0, 1), h.Pos)
	assert.Equal(t, tuple.Down, h.EyeV)
	assert.Equal(t, tuple.Down, h.NormalV)
	assert.Equal(t, true, h.Inside)
}
