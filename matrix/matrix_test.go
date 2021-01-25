package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatrix4x4(t *testing.T) {
	m, err := NewMatrix(
		1, 2, 3, 4,
		5.5, 6.5, 7.5, 8.5,
		9, 10, 11, 12,
		13.5, 14.5, 15.5, 16.5,
	)
	assert.Equal(t, nil, err)
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
	assert.Equal(t, nil, err)
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
	assert.Equal(t, nil, err)
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
