package color

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewColor(t *testing.T) {
	c := NewColor(-0.5, 0.4, 1.7)
	assert.Equal(t, -0.5, c.R)
	assert.Equal(t, 0.4, c.G)
	assert.Equal(t, 1.7, c.B)
}
