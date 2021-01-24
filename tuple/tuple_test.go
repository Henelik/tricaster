package tuple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoint(t *testing.T) {
	a := NewPoint(4.3, -4.2, 3.1)
	assert.Equal(t, 4.3, a.x)
	assert.Equal(t, -4.2, a.y)
	assert.Equal(t, 3.1, a.z)
	assert.Equal(t, 1.0, a.w)
	assert.Equal(t, a.IsPoint(), true)
	assert.Equal(t, a.IsVector(), false)
}

func TestVector(t *testing.T) {
	a := NewVector(4.3, -4.2, 3.1)
	assert.Equal(t, 4.3, a.x)
	assert.Equal(t, -4.2, a.y)
	assert.Equal(t, 3.1, a.z)
	assert.Equal(t, 0.0, a.w)
	assert.Equal(t, a.IsPoint(), false)
	assert.Equal(t, a.IsVector(), true)
}

func TestPointConstructor(t *testing.T) {
	p := NewPoint(4, -4, 3)
	assert.Equal(t, &Tuple{4, -4, 3, 1}, p)
}

func TestVectorConstructor(t *testing.T) {
	v := NewVector(4, -4, 3)
	assert.Equal(t, &Tuple{4, -4, 3, 0}, v)
}

func TestAdd(t *testing.T) {
	p := NewPoint(3, -2, 5)
	v := NewVector(-2, 3, 1)

	a := p.Add(v)
	b := v.Add(p)

	e := New(1, 1, 6, 1)

	assert.Equal(t, e, a)
	assert.Equal(t, e, b)
}

func TestSubPoints(t *testing.T) {
	p1 := NewPoint(3, 2, 1)
	p2 := NewPoint(5, 6, 7)

	a := p1.Sub(p2)

	e := NewVector(-2, -4, -6)

	assert.Equal(t, e, a)
}

func TestSubPointAndVector(t *testing.T) {
	p := NewPoint(3, 2, 1)
	v := NewVector(5, 6, 7)

	a := p.Sub(v)

	e := NewPoint(-2, -4, -6)

	assert.Equal(t, e, a)
}

func TestSubVectors(t *testing.T) {
	p1 := NewVector(3, 2, 1)
	p2 := NewVector(5, 6, 7)

	a := p1.Sub(p2)

	e := NewVector(-2, -4, -6)

	assert.Equal(t, e, a)
}
