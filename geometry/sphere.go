package geometry

import (
	"math"

	"github.com/Henelik/tricaster/ray"
	"github.com/Henelik/tricaster/tuple"
)

type Sphere struct {
}

func NewSphere() *Sphere {
	return &Sphere{}
}

func (s *Sphere) Intersects(r *ray.Ray) []Intersection {
	sphereToRay := r.Origin.Sub(tuple.NewPoint(0, 0, 0))
	a := r.Direction.DotProd(r.Direction)
	b := 2 * r.Direction.DotProd(sphereToRay)
	c := sphereToRay.DotProd(sphereToRay) - 1
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return []Intersection{}
	}
	return []Intersection{
		{(-b - math.Sqrt(discriminant)) / (2 * a), s},
		{(-b + math.Sqrt(discriminant)) / (2 * a), s},
	}
}
