package scene

import (
	"github.com/Henelik/tricaster/tuple"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPixelSize(t *testing.T) {
	// The pixel size for a horizontal canvas
	c1 := NewCamera(200, 125, 0, nil)
	assert.Equal(t, 0.01, c1.pixelSize)

	// The pixel size for a vertical canvas
	c2 := NewCamera(125, 200, 0, nil)
	assert.Equal(t, 0.01, c2.pixelSize)
}

func TestRayForPixel(t *testing.T) {
	// Constructing a ray through the center of the canvas
	c := NewCamera(201, 101, 0, nil)
	r := c.RayForPixel(100, 50)
	assert.Equal(t, tuple.Origin, r.Origin)
	assert.Equal(t, tuple.Forward, r.Direction)

	// Constructing a ray through a corner of the canvas
	r1 := c.RayForPixel(0, 0)
	assert.Equal(t, tuple.Origin, r1.Origin)
	assert.Equal(t, tuple.NewVector(0.6651864261194509, -0.6685123582500481, 0.33259321305972545), r1.Direction)
}
