package scene

import (
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
