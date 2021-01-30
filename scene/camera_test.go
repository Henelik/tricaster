package scene

import (
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
