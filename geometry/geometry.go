package geometry

import (
	"math"

	"github.com/Henelik/tricaster/ray"
)

// Primitive geometry type which defines an intersection function
type Primitive interface {
	// Intersects returns an array of t values along the ray that intersect with the primitive
	Intersects(r *ray.Ray) []Intersection
}

// Intersection stores the t value of a ray intersection and a pointer to the intersected primitive
type Intersection struct {
	T float64
	P Primitive
}

var NilHit = Intersection{math.Inf(1), nil}

func Hit(inters []Intersection) Intersection {
	if len(inters) == 0 {
		return NilHit
	}
	var closest Intersection = NilHit
	for _, i := range inters {
		if i.T > 0 && i.T < closest.T {
			closest = i
		}
	}
	return closest
}
