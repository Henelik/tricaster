package geometry

import (
	"math"

	"github.com/Henelik/tricaster/matrix"
	"github.com/Henelik/tricaster/ray"
	"github.com/Henelik/tricaster/tuple"
)

type Sphere struct {
	M *matrix.Matrix
}

func NewSphere(m *matrix.Matrix) *Sphere {
	return &Sphere{m}
}

func (s *Sphere) Intersects(r *ray.Ray) []Intersection {
	rt := r.Transform(s.M.Inverse())
	sphereToRay := rt.Origin.Sub(tuple.NewPoint(0, 0, 0))
	a := rt.Direction.DotProd(rt.Direction)
	b := 2 * rt.Direction.DotProd(sphereToRay)
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
