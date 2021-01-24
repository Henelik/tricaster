package tuple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoint(t *testing.T) {
	a := New(4.3, -4.2, 3.1, 1.0)
	assert.Equal(t, 4.3, a.x)
	assert.Equal(t, -4.2, a.y)
	assert.Equal(t, 3.1, a.z)
	assert.Equal(t, 1.0, a.w)
	assert.Equal(t, a.IsPoint(), true)
	assert.Equal(t, a.IsVector(), false)
}

func TestVector(t *testing.T) {
	a := New(4.3, -4.2, 3.1, 0.0)
	assert.Equal(t, 4.3, a.x)
	assert.Equal(t, -4.2, a.y)
	assert.Equal(t, 3.1, a.z)
	assert.Equal(t, 0.0, a.w)
	assert.Equal(t, a.IsPoint(), false)
	assert.Equal(t, a.IsVector(), true)
}

func TestPointConstructor(t *testing.T) {
	p := Point(4, -4, 3)
	assert.Equal(t, &Tuple{4, -4, 3, 1}, p)
}

func TestVectorConstructor(t *testing.T) {
	v := Vector(4, -4, 3)
	assert.Equal(t, &Tuple{4, -4, 3, 0}, v)
}
