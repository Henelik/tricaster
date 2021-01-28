package tuple

import (
	"fmt"
	"math"

	"github.com/Henelik/tricaster/util"
)

type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
}

var Origin = NewPoint(0, 0, 0)

func New(X, Y, Z, W float64) *Tuple {
	return &Tuple{X, Y, Z, W}
}

func NewPoint(X, Y, Z float64) *Tuple {
	return &Tuple{X, Y, Z, 1.0}
}

func NewVector(X, Y, Z float64) *Tuple {
	return &Tuple{X, Y, Z, 0.0}
}

func (t *Tuple) IsPoint() bool {
	return t.W == 1.0
}

func (t *Tuple) IsVector() bool {
	return t.W == 0.0
}

func (t *Tuple) Equal(o *Tuple) bool {
	return util.Equal(t.X, o.X) &&
		util.Equal(t.Y, o.Y) &&
		util.Equal(t.Z, o.Z) &&
		util.Equal(t.W, o.W)
}

func (t *Tuple) Add(o *Tuple) *Tuple {
	return New(t.X+o.X, t.Y+o.Y, t.Z+o.Z, t.W+o.W)
}

func (t *Tuple) Sub(o *Tuple) *Tuple {
	return New(t.X-o.X, t.Y-o.Y, t.Z-o.Z, t.W-o.W)
}

func (t *Tuple) Neg() *Tuple {
	return New(-t.X, -t.Y, -t.Z, -t.W)
}

func (t *Tuple) Mult(n float64) *Tuple {
	return New(t.X*n, t.Y*n, t.Z*n, t.W*n)
}

func (t *Tuple) Div(n float64) *Tuple {
	return New(t.X/n, t.Y/n, t.Z/n, t.W/n)
}

func (t *Tuple) Mag() float64 {
	return math.Sqrt(t.X*t.X + t.Y*t.Y + t.Z*t.Z + t.W*t.W)
}

func (t *Tuple) Norm() *Tuple {
	m := t.Mag()
	return New(t.X/m, t.Y/m, t.Z/m, t.W/m)
}

func (t *Tuple) DotProd(o *Tuple) float64 {
	return t.X*o.X +
		t.Y*o.Y +
		t.Z*o.Z +
		t.W*o.W
}

func (t *Tuple) CrossProd(o *Tuple) *Tuple {
	return NewVector(t.Y*o.Z-t.Z*o.Y,
		t.Z*o.X-t.X*o.Z,
		t.X*o.Y-t.Y*o.X)
}

func (t *Tuple) Fmt() string {
	return fmt.Sprintf("X: %f, Y: %f, Z:%f, W:%f", t.X, t.Y, t.Z, t.W)
}

func (t *Tuple) Reflect(n *Tuple) *Tuple {
	return t.Sub(n.Mult(2 * t.DotProd(n)))
}
