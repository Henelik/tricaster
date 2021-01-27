package geometry

import (
	"math"

	"github.com/Henelik/tricaster/matrix"
	"github.com/Henelik/tricaster/ray"
	"github.com/Henelik/tricaster/tuple"
)

type Sphere struct {
	// the transformation matrix
	M *matrix.Matrix
	// the inverse transformation matrix
	im *matrix.Matrix
}

func NewSphere(m *matrix.Matrix) *Sphere {
	return &Sphere{m, m.Inverse()}
}

func (s *Sphere) UpdateIM() {
	s.im = s.M.Inverse()
}

func (s *Sphere) Intersects(r *ray.Ray) []Intersection {
	rt := r.Transform(s.im)
	sphereToRay := rt.Origin.Sub(tuple.NewPoint(0, 0, 0))
	a := rt.Direction.DotProd(rt.Direction)
	b := -2 * rt.Direction.DotProd(sphereToRay)
	discriminant := b*b - 4*a*(sphereToRay.DotProd(sphereToRay)-1)
	if discriminant < 0 {
		return []Intersection{}
	}
	return []Intersection{
		{(b - math.Sqrt(discriminant)) / (2 * a), s},
		{(b + math.Sqrt(discriminant)) / (2 * a), s},
	}
}