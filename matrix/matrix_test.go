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

	a := tuple.New(1, 2, 3, 4)

	assert.Equal(t, a, Identity.MultTuple(a))
}

func TestTranspose(t *testing.T) {
	m, err := NewMatrix(
		0, 9, 3, 0,
		9, 8, 0, 8,
		1, 8, 5, 3,
		0, 0, 5, 8,
	)
	assert.Nil(t, err)

	e, err := NewMatrix(
		0, 9, 1, 0,
		9, 8, 8, 0,
		3, 0, 5, 5,
		0, 8, 3, 8,
	)
	assert.Nil(t, err)

	assert.Equal(t, e, m.Transpose())
}

func TestSubmatrix3(t *testing.T) {
	a, err := NewMatrix(
		1, 5, 0,
		-3, 2, 7,
		0, 6, -3,
	)
	assert.Nil(t, err)

	e, err := NewMatrix(
		-3, 2,
		0, 6,
	)
	assert.Nil(t, err)

	assert.Equal(t, e, a.Submatrix(0, 2))
}

func TestSubmatrix4(t *testing.T) {
	a, err := NewMatrix(
		-6, 1, 1, 6,
		-8, 5, 8, 6,
		-1, 0, 8, 2,
		-7, 1, -1, 1,
	)
	assert.Nil(t, err)

	e, err := NewMatrix(
		-6, 1, 6,
		-8, 8, 6,
		-7, -1, 1,
	)
	assert.Nil(t, err)

	assert.Equal(t, e, a.Submatrix(2, 1))
}

func TestMinor(t *testing.T) {
	a, err := NewMatrix(
		3, 5, 0,
		2, -1, -7,
		6, -1, -5,
	)
	assert.Nil(t, err)

	assert.Equal(t, -25.0, a.Minor(1, 0))
}

func TestCofactor(t *testing.T) {
	m, err := NewMatrix(
		3, 5, 0,
		2, -1, -7,
		6, -1, 5,
	)
	assert.Nil(t, err)

	assert.Equal(t, -12.0, m.Minor(0, 0))
	assert.Equal(t, -12.0, m.Cofactor(0, 0))

	assert.Equal(t, 25.0, m.Minor(1, 0))
	assert.Equal(t, -25.0, m.Cofactor(1, 0))
}

func TestDeterminant2(t *testing.T) {
	m, err := NewMatrix(
		1, 5,
		-3, 2,
	)
	assert.Nil(t, err)

	assert.Equal(t, 17.0, m.Determinant())
}

func TestDeterminant3(t *testing.T) {
	m, err := NewMatrix(
		1, 2, 6,
		-5, 8, -4,
		2, 6, 4,
	)
	assert.Nil(t, err)

	assert.Equal(t, -196.0, m.Determinant())
}

func TestDeterminant4(t *testing.T) {
	m, err := NewMatrix(
		-2, -8, 3, 5,
		-3, 1, 7, 3,
		1, 2, -9, 6,
		-6, 7, 7, -9,
	)
	assert.Nil(t, err)

	assert.Equal(t, -4071.0, m.Determinant())
}
