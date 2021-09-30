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

func TestAdd(t *testing.T) {
	c1 := NewColor(0.9, 0.6, 0.75)
	c2 := NewColor(0.7, 0.1, 0.25)

	e := NewColor(1.6, 0.7, 1.0)

	assert.True(t, e.Equal(c1.Add(c2)))
	assert.True(t, e.Equal(c2.Add(c1)))
}

func TestSub(t *testing.T) {
	c1 := NewColor(0.9, 0.6, 0.75)
	c2 := NewColor(0.7, 0.1, 0.25)

	e := NewColor(0.2, 0.5, 0.5)

	assert.True(t, e.Equal(c1.Sub(c2)))
}

func TestMultF(t *testing.T) {
	c := NewColor(0.2, 0.3, 0.4)

	e := NewColor(0.4, 0.6, 0.8)

	assert.True(t, e.Equal(c.MultF(2)))
}

func TestMultCol(t *testing.T) {
	c1 := NewColor(1, 0.2, 0.4)
	c2 := NewColor(0.9, 1, 0.1)

	e := NewColor(0.9, 0.2, 0.04)

	assert.True(t, e.Equal(c1.MultCol(c2)))
	assert.True(t, e.Equal(c2.MultCol(c1)))
}
