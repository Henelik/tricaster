package ray

import (
	"log"

	"github.com/Henelik/tricaster/pkg/matrix"
	"github.com/Henelik/tricaster/pkg/tuple"
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
		log.Fatalf("(%v, %v) is not a proper ray!", origin, direction)
		return nil
	}
	return &Ray{
		Origin:    origin,
		Direction: direction,
	}
}

// Position gets a point along the ray
func (r *Ray) Position(t float64) *tuple.Tuple {
	return r.Origin.Add(r.Direction.Mult(t))
}

func (r *Ray) Transform(m *matrix.Matrix) *Ray {
	return NewRay(
		m.MultTuple(r.Origin),
		m.MultTuple(r.Direction),
	)
}
