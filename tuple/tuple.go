package tuple

import "github.com/Henelik/tricaster/util"

type Tuple struct {
	x float64
	y float64
	z float64
	w float64
}

func New(x, y, z, w float64) *Tuple {
	return &Tuple{x, y, z, w}
}

func NewPoint(x, y, z float64) *Tuple {
	return &Tuple{x, y, z, 1.0}
}

func NewVector(x, y, z float64) *Tuple {
	return &Tuple{x, y, z, 0.0}
}

func (t *Tuple) IsPoint() bool {
	return t.w == 1.0
}

func (t *Tuple) IsVector() bool {
	return t.w == 0.0
}

func (t *Tuple) Equal(o *Tuple) bool {
	if util.Equal(t.x, o.x) &&
		util.Equal(t.y, o.y) &&
		util.Equal(t.z, o.z) &&
		util.Equal(t.w, o.w) {
		return true
	}
	return false
}

func (t *Tuple) Add(o *Tuple) *Tuple {
	return New(t.x+o.x, t.y+o.y, t.z+o.z, t.w+o.w)
}

func (t *Tuple) Sub(o *Tuple) *Tuple {
	return New(t.x-o.x, t.y-o.y, t.z-o.z, t.w-o.w)
}
