package geometry

import (
	"testing"

	"github.com/Henelik/tricaster/ray"

	"github.com/Henelik/tricaster/matrix"
	"github.com/Henelik/tricaster/shading"
	"github.com/Henelik/tricaster/tuple"
	"github.com/stretchr/testify/assert"
)

func TestHit(t *testing.T) {
	s := NewSphere(matrix.Identity, shading.DefaultPhong)
	testCases := []struct {
		name   string
		inters []ray.Intersection
		want   ray.Intersection
	}{
		{
			name: "The hit, when all intersections have positive t",
			inters: []ray.Intersection{
				{1, s},
				{2, s},
			},
			want: ray.Intersection{1, s},
		},
		{
			name: "The hit, when some intersections have negative t",
			inters: []ray.Intersection{
				{-1, s},
				{1, s},
			},
			want: ray.Intersection{1, s},
		},
		{
			name: "The hit, when all intersections have negative t",
			inters: []ray.Intersection{
				{-2, s},
				{-1, s},
			},
			want: *ray.NilIntersect,
		},
		{
			name: "The hit is always the lowest nonnegative intersection",
			inters: []ray.Intersection{
				{5, s},
				{7, s},
				{-3, s},
				{2, s},
			},
			want: ray.Intersection{2, s},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, *ray.GetClosest(tc.inters))
		})
	}
}

func BenchmarkIntersection128(b *testing.B) {
	var hit bool
	for n := 0; n < b.N; n++ {
		s := NewSphere(matrix.Translation(0, 5, 0), shading.DefaultPhong)

		w := 128
		h := 128

		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				// shoot a ray at the sphere
				xPos := (float64(x) - float64(w)/2.0) / (float64(w) * 0.4)
				yPos := -(float64(y) - float64(h)/2.0) / (float64(h) * 0.4)
				r := ray.NewRay(
					tuple.NewPoint(xPos, 0, yPos),
					tuple.Backward,
				)
				i := ray.GetClosest(s.Intersects(r))
				if i != ray.NilIntersect {
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
	r := ray.NewRay(
		tuple.NewPoint(0, 0, -5),
		tuple.Up)
	s := NewSphere(nil, nil)
	i := &ray.Intersection{4, s}

	h := i.ToHit(r)

	assert.Equal(t, i.T, h.T)
	assert.Equal(t, i.P, h.P)
	assert.Equal(t, tuple.NewPoint(0, 0, -1), h.Pos)
	assert.Equal(t, tuple.Down, h.EyeV)
	assert.Equal(t, tuple.Down, h.NormalV)
	assert.Equal(t, false, h.Inside)
}

func TestToHitInside(t *testing.T) {
	r := ray.NewRay(
		tuple.Origin,
		tuple.Up)
	s := NewSphere(nil, nil)
	i := &ray.Intersection{1, s}

	h := i.ToHit(r)

	assert.Equal(t, i.T, h.T)
	assert.Equal(t, i.P, h.P)
	assert.Equal(t, tuple.NewPoint(0, 0, 1), h.Pos)
	assert.Equal(t, tuple.Down, h.EyeV)
	assert.Equal(t, tuple.Down, h.NormalV)
	assert.Equal(t, true, h.Inside)
}
