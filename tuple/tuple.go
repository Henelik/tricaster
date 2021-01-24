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

func Point(x, y, z float64) *Tuple {
	return &Tuple{x, y, z, 1.0}
}

func Vector(x, y, z float64) *Tuple {
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
