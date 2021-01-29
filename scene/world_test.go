package scene

import (
	"github.com/Henelik/tricaster/color"
	"github.com/Henelik/tricaster/geometry"
	"github.com/Henelik/tricaster/ray"
	"github.com/Henelik/tricaster/tuple"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntersect(t *testing.T) {
	// Intersect a world with a ray
	got := DefaultWorld.Intersect(
		ray.NewRay(
			tuple.NewPoint(0, 0, -5),
			tuple.NewVector(0, 0, 1)))

	assert.Equal(t, 4, len(got))
	assert.Equal(t, 4.0, got[0].T)
	assert.Equal(t, 4.5, got[1].T)
	assert.Equal(t, 5.5, got[2].T)
	assert.Equal(t, 6.0, got[3].T)
}

func TestShading(t *testing.T) {
	// Shading an intersection
	r := ray.NewRay(tuple.NewPoint(-5, 0, 0), tuple.Right)
	s := DefaultWorld.Geometry[0]
	i := &geometry.Intersection{4, s}
	col := DefaultWorld.Shade(i.Precompute(r))
	assert.Equal(t, color.NewColor(0.38066119308103435, 0.47582649135129296, 0.28549589481077575), col)

	// Shading an intersection from the inside
	w := DefaultWorld
	w.Light.Pos = tuple.NewPoint(0, 0.25, 0)
	r2 := ray.NewRay(tuple.Origin, tuple.Right)
	s2 := DefaultWorld.Geometry[1]
	i2 := &geometry.Intersection{0.5, s2}
	col2 := w.Shade(i2.Precompute(r2))
	assert.Equal(t, color.Grey(0.9049844720832575), col2)
}
