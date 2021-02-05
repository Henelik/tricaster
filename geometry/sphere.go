package geometry

import (
	"git.maze.io/go/math32"
	"github.com/Henelik/tricaster/color"
	"github.com/Henelik/tricaster/shading"

	"github.com/Henelik/tricaster/matrix"
	"github.com/Henelik/tricaster/ray"
	"github.com/Henelik/tricaster/tuple"
)

type Sphere struct {
	// the transformation matrix
	m *matrix.Matrix
	// the inverse transformation matrix
	im *matrix.Matrix
	// the material
	Mat shading.Material
}

func NewSphere(m *matrix.Matrix, mat shading.Material) *Sphere {
	s := &Sphere{
		matrix.Identity,
		matrix.Identity,
		shading.DefaultPhong,
	}
	if m != nil {
		s.m = m
		s.im = m.Inverse()
	}
	if mat != nil {
		s.Mat = mat
	}
	return s
}

func (s *Sphere) SetMatrix(m *matrix.Matrix) {
	s.m = m
	s.im = m.Inverse()
}

func (s *Sphere) GetMatrix() *matrix.Matrix {
	return s.m
}

func (s *Sphere) Intersects(r *ray.Ray) []ray.Intersection {
	rt := r.Transform(s.im)
	sphereToRay := rt.Origin.Sub(tuple.Origin)
	a := rt.Direction.DotProd(rt.Direction)
	b := -2 * rt.Direction.DotProd(sphereToRay)
	discriminant := b*b - 4*a*(sphereToRay.DotProd(sphereToRay)-1)
	if discriminant < 0 {
		return []ray.Intersection{}
	}
	return []ray.Intersection{
		{(b - math32.Sqrt(discriminant)) / (2 * a), s},
		{(b + math32.Sqrt(discriminant)) / (2 * a), s},
	}
}

func (s *Sphere) NormalAt(pos *tuple.Tuple) *tuple.Tuple {
	objectPoint := s.im.MultTuple(pos)
	objectNormal := objectPoint.Sub(tuple.Origin)
	worldNormal := s.im.Transpose().MultTuple(objectNormal)
	worldNormal.W = 0
	return worldNormal.Norm()
}

func (s *Sphere) Shade(light *shading.PointLight, h *ray.Hit) *color.Color {
	return s.Mat.Lighting(light, h)
}

func (s *Sphere) GetMaterial() shading.Material {
	return s.Mat
}
