package tuple

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoint(t *testing.T) {
	a := NewPoint(4.3, -4.2, 3.1)
	assert.Equal(t, 4.3, a.X)
	assert.Equal(t, -4.2, a.Y)
	assert.Equal(t, 3.1, a.Z)
	assert.Equal(t, 1.0, a.W)
	assert.Equal(t, a.IsPoint(), true)
	assert.Equal(t, a.IsVector(), false)
}

func TestVector(t *testing.T) {
	a := NewVector(4.3, -4.2, 3.1)
	assert.Equal(t, 4.3, a.X)
	assert.Equal(t, -4.2, a.Y)
	assert.Equal(t, 3.1, a.Z)
	assert.Equal(t, 0.0, a.W)
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
	v1 := NewVector(3, 2, 1)
	v2 := NewVector(5, 6, 7)

	a := v1.Sub(v2)

	e := NewVector(-2, -4, -6)

	assert.Equal(t, e, a)
}

func TestNeg(t *testing.T) {
	a := New(1, -2, 3, -4)

	b := a.Neg()

	e := New(-1, 2, -3, 4)

	assert.Equal(t, e, b)
}

func TestMult(t *testing.T) {
	a := New(1, -2, 3, -4)

	b := a.Mult(0.5)

	e := New(0.5, -1, 1.5, -2)

	assert.Equal(t, e, b)
}

func TestDiv(t *testing.T) {
	a := New(1, -2, 3, -4)

	b := a.Div(2)

	e := New(0.5, -1, 1.5, -2)

	assert.Equal(t, e, b)
}

func TestMag(t *testing.T) {
	testCases := []struct {
		name string
		v    *Tuple
		want float64
	}{
		{
			name: "X unit vector",
			v:    Right,
			want: 1,
		},
		{
			name: "Y unit vector",
			v:    Backward,
			want: 1,
		},
		{
			name: "Z unit vector",
			v:    Up,
			want: 1,
		},
		{
			name: "all positive vector",
			v:    NewVector(1, 2, 3),
			want: 3.7416573867739413,
		},
		{
			name: "all negative vector",
			v:    NewVector(-1, -2, -3),
			want: 3.7416573867739413,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.v.Mag())
		})
	}
}

func TestNorm(t *testing.T) {
	testCases := []struct {
		name string
		v    *Tuple
		want *Tuple
	}{
		{
			name: "unit vector multiple",
			v:    NewVector(4, 0, 0),
			want: Right,
		},
		{
			name: "all positive vector",
			v:    NewVector(1, 2, 3),
			want: NewVector(0.2672612419124244, 0.5345224838248488, 0.8017837257372732),
		},
		{
			name: "all negative vector",
			v:    NewVector(-1, -2, -3),
			want: NewVector(-0.2672612419124244, -0.5345224838248488, -0.8017837257372732),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.v.Norm())
		})
	}
}

func TestDotProd(t *testing.T) {
	v := NewVector(1, 2, 3)
	o := NewVector(2, 3, 4)

	assert.Equal(t, 20.0, v.DotProd(o))
}

func TestCrossProd(t *testing.T) {
	a := NewVector(1, 2, 3)
	b := NewVector(2, 3, 4)

	axb := NewVector(-1, 2, -1)
	bxa := NewVector(1, -2, 1)

	assert.Equal(t, axb, a.CrossProd(b))
	assert.Equal(t, bxa, b.CrossProd(a))
}

func TestReflect(t *testing.T) {
	// Reflecting a vector off a slanted surface
	v1 := NewVector(1, -1, 0)
	n1 := Backward
	r1 := v1.Reflect(n1)
	e1 := NewVector(1, 1, 0)
	assert.True(t, e1.Equal(r1))

	// Reflecting a vector off a slanted surface
	v2 := Forward
	n2 := NewVector(math.Sqrt2/2, math.Sqrt2/2, 0)
	r2 := v2.Reflect(n2)
	e2 := Right
	assert.True(t, e2.Equal(r2))
}

func TestIsPoint(t *testing.T) {
	assert.True(t, (&Tuple{1, 2, 3, 1}).IsPoint())
	assert.False(t, (&Tuple{1, 2, 3, 0}).IsPoint())
}

func TestIsVector(t *testing.T) {
	assert.True(t, (&Tuple{1, 2, 3, 0}).IsVector())
	assert.False(t, (&Tuple{1, 2, 3, 1}).IsVector())
}

func TestFmt(t *testing.T) {
	assert.Equal(t, "X: 1.000000, Y: 2.000000, Z:3.000000, W:1.000000", (&Tuple{1, 2, 3, 1}).Fmt())
}
