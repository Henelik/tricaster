package matrix

import (
	"testing"

	"github.com/Henelik/tricaster/tuple"
	"github.com/stretchr/testify/assert"
)

func TestMatrix4x4(t *testing.T) {
	m, err := NewMatrix(
		1, 2, 3, 4,
		5.5, 6.5, 7.5, 8.5,
		9, 10, 11, 12,
		13.5, 14.5, 15.5, 16.5,
	)
	assert.Nil(t, err)
	assert.Equal(t, 1.0, m.Data[0][0])
	assert.Equal(t, 4.0, m.Data[0][3])
	assert.Equal(t, 5.5, m.Data[1][0])
	assert.Equal(t, 7.5, m.Data[1][2])
	assert.Equal(t, 11.0, m.Data[2][2])
	assert.Equal(t, 13.5, m.Data[3][0])
	assert.Equal(t, 15.5, m.Data[3][2])
}

func TestMatrix2x2(t *testing.T) {
	m, err := NewMatrix(
		-3, 5,
		1, -2,
	)
	assert.Nil(t, err)
	assert.Equal(t, -3.0, m.Data[0][0])
	assert.Equal(t, 5.0, m.Data[0][1])
	assert.Equal(t, 1.0, m.Data[1][0])
	assert.Equal(t, -2.0, m.Data[1][1])
}

func TestMatrix3x3(t *testing.T) {
	m, err := NewMatrix(
		-3, 5, 0,
		1, -2, -7,
		0, 1, 1,
	)
	assert.Nil(t, err)
	assert.Equal(t, -3.0, m.Data[0][0])
	assert.Equal(t, 5.0, m.Data[0][1])
	assert.Equal(t, 1.0, m.Data[1][0])
	assert.Equal(t, -2.0, m.Data[1][1])
	assert.Equal(t, 1.0, m.Data[2][2])
}

func TestEqual(t *testing.T) {
	m, err := NewMatrix(
		-3, 5,
		1, -2,
	)
	assert.Nil(t, err)

	o, err := NewMatrix(
		-3, 5,
		1, -2,
	)
	assert.Nil(t, err)

	assert.True(t, m.Equal(o))

	o, err = NewMatrix(
		-3, 5,
		1, -1,
	)
	assert.Nil(t, err)

	assert.False(t, m.Equal(o))
}

func TestMatrixImproperSize(t *testing.T) {
	m, err := NewMatrix(
		1, 2, 3, 4,
		5.5, 6.5, 7.5, 8.5,
		9, 10, 11, 12,
		13.5, 14.5, 15.5,
	)
	assert.EqualError(t, err, "matrix must be a square")
	assert.Nil(t, m)
}

func TestMult(t *testing.T) {
	a, err := NewMatrix(
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 8, 7, 6,
		5, 4, 3, 2,
	)
	assert.Nil(t, err)

	b, err := NewMatrix(
		-2, 1, 2, 3,
		3, 2, 1, -1,
		4, 3, 6, 5,
		1, 2, 7, 8,
	)
	assert.Nil(t, err)

	e, err := NewMatrix(
		20, 22, 50, 48,
		44, 54, 114, 108,
		40, 58, 110, 102,
		16, 26, 46, 42,
	)
	assert.Nil(t, err)

	assert.Equal(t, e, a.Mult(b))
}

func TestMultTuple(t *testing.T) {
	a, err := NewMatrix(
		1, 2, 3, 4,
		2, 4, 4, 2,
		8, 6, 4, 1,
		0, 0, 0, 1,
	)
	assert.Nil(t, err)

	tup := tuple.New(1, 2, 3, 1)

	e := tuple.New(18, 24, 33, 1)

	assert.Equal(t, e, a.MultTuple(tup))
}

func TestIdentity(t *testing.T) {
	m, err := NewMatrix(
		0, 1, 2, 4,
		1, 2, 4, 8,
		2, 4, 8, 16,
		4, 8, 16, 32,
	)
	assert.Nil(t, err)

	assert.Equal(t, m, m.Mult(Identity))
}
