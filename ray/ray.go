package ray

import (
	"github.com/Henelik/tricaster/matrix"
	"github.com/Henelik/tricaster/tuple"
)

type Ray struct {
	Origin    *tuple.Tuple
	Direction *tuple.Tuple
}

// NewRay creates a ray.
// Origin must be a point, direction must be a vector.
// Returns nil if given invalid arguments.
func NewRay(origin, direction *tuple.Tuple) *Ray {
	if !origin.IsPoint() || !direction.IsVector() {
		return nil
	}
	return &Ray{
		Origin:    origin,
		Direction: direction,
	}
}

// Position gets a point along the ray
func (r *Ray) Position(t float32) *tuple.Tuple {
	return r.Origin.Add(r.Direction.Mult(t))
}

func (r *Ray) Transform(m *matrix.Matrix) *Ray {
	return NewRay(
		m.MultTuple(r.Origin),
		m.MultTuple(r.Direction),
	)
}
