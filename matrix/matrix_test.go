package matrix

import (
	"math"
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

func TestIsInvertible(t *testing.T) {
	a, err := NewMatrix(
		-2, -8, 3, 5,
		-3, 1, 7, 3,
		1, 2, -9, 6,
		-6, 7, 7, -9,
	)
	assert.Nil(t, err)

	assert.Equal(t, true, a.IsInvertible())

	b, err := NewMatrix(
		-4, 2, -2, -3,
		9, 6, 2, 6,
		0, -5, 1, -5,
		0, 0, 0, 0,
	)
	assert.Nil(t, err)

	assert.Equal(t, false, b.IsInvertible())
}

func TestInverse(t *testing.T) {
	a, err := NewMatrix(
		-5, 2, 6, -8,
		1, -5, 1, 8,
		7, 7, -6, -7,
		1, -3, 7, 4,
	)
	assert.Nil(t, err)

	e := &Matrix{Order: 4, Data: [][]float64{[]float64{0.21804511278195488, 0.45112781954887216, 0.24060150375939848, -0.045112781954887216}, []float64{-0.8082706766917294, -1.4567669172932332, -0.44360902255639095, 0.5206766917293233}, []float64{-0.07894736842105263, -0.2236842105263158, -0.05263157894736842, 0.19736842105263158}, []float64{-0.5225563909774437, -0.8139097744360902, -0.3007518796992481, 0.30639097744360905}}}

	assert.True(t, a.IsInvertible())
	assert.Equal(t, 532.0, a.Determinant())
	assert.Equal(t, e, a.Inverse())

	b, err := NewMatrix(
		8, -5, 9, 2,
		7, 5, 6, 1,
		-6, 0, 9, 6,
		-3, 0, -9, -4,
	)
	assert.Nil(t, err)

	e = &Matrix{Order: 4, Data: [][]float64{[]float64{-0.15384615384615385, -0.15384615384615385, -0.28205128205128205, -0.5384615384615384}, []float64{-0.07692307692307693, 0.12307692307692308, 0.02564102564102564, 0.03076923076923077}, []float64{0.358974358974359, 0.358974358974359, 0.4358974358974359, 0.9230769230769231}, []float64{-0.6923076923076923, -0.6923076923076923, -0.7692307692307693, -1.9230769230769231}}}

	assert.True(t, b.IsInvertible())
	assert.Equal(t, -585.0, b.Determinant())
	assert.Equal(t, e, b.Inverse())

	c, err := NewMatrix(
		9, 3, 0, 9,
		-5, -2, -6, -3,
		-4, 9, 6, 4,
		-7, 6, 6, 2,
	)
	assert.Nil(t, err)

	e = &Matrix{Order: 4, Data: [][]float64{[]float64{-0.040740740740740744, -0.07777777777777778, 0.14444444444444443, -0.2222222222222222}, []float64{-0.07777777777777778, 0.03333333333333333, 0.36666666666666664, -0.3333333333333333}, []float64{-0.029012345679012345, -0.14629629629629629, -0.10925925925925926, 0.12962962962962962}, []float64{0.17777777777777778, 0.06666666666666667, -0.26666666666666666, 0.3333333333333333}}}

	assert.True(t, c.IsInvertible())
	assert.Equal(t, 1620.0, c.Determinant())
	assert.Equal(t, e, c.Inverse())
}

func TestMultInverse(t *testing.T) {
	a, err := NewMatrix(
		3, -9, 7, 3,
		3, -8, 2, -9,
		-4, 4, 4, 1,
		-6, 5, -1, 1,
	)
	assert.Nil(t, err)

	b, err := NewMatrix(
		8, 2, 2, 2,
		3, -1, 7, 0,
		7, 0, 5, 4,
		6, -2, 0, 5,
	)
	assert.Nil(t, err)

	e, err := NewMatrix(
		64, 9, -22, 49,
		-40, 32, -40, -31,
		14, -14, 40, 13,
		-34, -19, 18, -11,
	)
	assert.Nil(t, err)

	c := a.Mult(b)

	assert.Equal(t, e, c)

	assert.True(t, a.Equal(c.Mult(b.Inverse())))
}

func TestTranslation(t *testing.T) {
	trans := Translation(5, -3, 2)

	p := tuple.NewPoint(-3, 4, 5)

	e := tuple.NewPoint(2, 1, 7)

	assert.Equal(t, e, trans.MultTuple(p))
}

func TestTranslationInverse(t *testing.T) {
	trans := Translation(5, -3, 2)

	p := tuple.NewPoint(-3, 4, 5)

	e := tuple.NewPoint(-8, 7, 3)

	assert.Equal(t, e, trans.Inverse().MultTuple(p))
}

func TestTranslationVector(t *testing.T) {
	trans := Translation(5, -3, 2)

	v := tuple.NewVector(-3, 4, 5)

	assert.Equal(t, v, trans.MultTuple(v))
}

func TestScalingPoint(t *testing.T) {
	s := Scaling(2, 3, 4)

	p := tuple.NewPoint(-4, 6, 8)

	e := tuple.NewPoint(-8, 18, 32)

	assert.Equal(t, e, s.MultTuple(p))
}

func TestScalingVector(t *testing.T) {
	s := Scaling(2, 3, 4)

	v := tuple.NewVector(-4, 6, 8)

	e := tuple.NewVector(-8, 18, 32)

	assert.Equal(t, e, s.MultTuple(v))
}

func TestScalingInverse(t *testing.T) {
	s := Scaling(2, 3, 4)

	v := tuple.NewVector(-4, 6, 8)

	e := tuple.NewVector(-2, 2, 2)

	assert.Equal(t, e, s.Inverse().MultTuple(v))
}

func TestScalingReflection(t *testing.T) {
	s := Scaling(-1, 1, 1)

	p := tuple.NewPoint(2, 3, 4)

	e := tuple.NewPoint(-2, 3, 4)

	assert.Equal(t, e, s.MultTuple(p))
}

func TestRotationX(t *testing.T) {
	p := tuple.NewPoint(0, 1, 0)

	eighth := RotationX(math.Pi / 4)
	eighthPoint := tuple.NewPoint(0, math.Sqrt(2)/2, math.Sqrt(2)/2)
	assert.True(t, eighthPoint.Equal(eighth.MultTuple(p)))

	quarter := RotationX(math.Pi / 2)
	quarterPoint := tuple.NewPoint(0, 0, 1)
	assert.True(t, quarterPoint.Equal(quarter.MultTuple(p)))
}

func TestRotationXInverse(t *testing.T) {
	p := tuple.NewPoint(0, 1, 0)

	eighth := RotationX(math.Pi / 4)
	e := tuple.NewPoint(0, math.Sqrt(2)/2, -math.Sqrt(2)/2)
	assert.True(t, e.Equal(eighth.Inverse().MultTuple(p)))
}

func TestRotationY(t *testing.T) {
	p := tuple.NewPoint(0, 0, 1)

	eighth := RotationY(math.Pi / 4)
	eighthPoint := tuple.NewPoint(math.Sqrt(2)/2, 0, math.Sqrt(2)/2)
	assert.True(t, eighthPoint.Equal(eighth.MultTuple(p)))

	quarter := RotationY(math.Pi / 2)
	quarterPoint := tuple.NewPoint(1, 0, 0)
	assert.True(t, quarterPoint.Equal(quarter.MultTuple(p)))
}

func TestRotationZ(t *testing.T) {
	p := tuple.NewPoint(0, 1, 0)

	eighth := RotationZ(math.Pi / 4)
	eighthPoint := tuple.NewPoint(-math.Sqrt(2)/2, math.Sqrt(2)/2, 0)
	assert.True(t, eighthPoint.Equal(eighth.MultTuple(p)))

	quarter := RotationZ(math.Pi / 2)
	quarterPoint := tuple.NewPoint(-1, 0, 0)
	assert.True(t, quarterPoint.Equal(quarter.MultTuple(p)))
}
