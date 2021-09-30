package geometry

import (
	"math"

	"github.com/Henelik/tricaster/pkg/color"
	"github.com/Henelik/tricaster/pkg/light"
	material2 "github.com/Henelik/tricaster/pkg/material"
	"github.com/Henelik/tricaster/pkg/matrix"
	ray2 "github.com/Henelik/tricaster/pkg/ray"
	"github.com/Henelik/tricaster/pkg/tuple"
)

type Sphere struct {
	// the transformation matrix
	m *matrix.Matrix
	// the inverse transformation matrix
	im *matrix.Matrix
	// the transposition of the inverse matrix
	imt *matrix.Matrix
	// the material
	Mat material2.Material
}

func NewSphere(m *matrix.Matrix, mat material2.Material) *Sphere {
	s := &Sphere{
		matrix.Identity,
		matrix.Identity,
		matrix.Identity,
		material2.DefaultPhong,
	}
	if m != nil {
		s.m = m
		s.im = m.Inverse()
		s.imt = s.im.Transpose()
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

func (s *Sphere) Intersects(r *ray2.Ray) []ray2.Intersection {
	rt := r.Transform(s.im)
	sphereToRay := rt.Origin.Sub(tuple.Origin)
	a := rt.Direction.DotProd(rt.Direction)
	b := -2 * rt.Direction.DotProd(sphereToRay)
	discriminant := b*b - 4*a*(sphereToRay.DotProd(sphereToRay)-1)
	if discriminant < 0 {
		return []ray2.Intersection{}
	}
	return []ray2.Intersection{
		{(b - math.Sqrt(discriminant)) / (2 * a), s},
		{(b + math.Sqrt(discriminant)) / (2 * a), s},
	}
}

func (s *Sphere) NormalAt(pos *tuple.Tuple) *tuple.Tuple {
	worldNormal := s.imt.MultTuple(s.im.MultTuple(pos).Sub(tuple.Origin))
	worldNormal.W = 0
	return worldNormal.Norm()
}

func (s *Sphere) Shade(light *light.PointLight, h *ray2.Hit) *color.Color {
	return s.Mat.Lighting(light, h)
}

func (s *Sphere) GetMaterial() material2.Material {
	return s.Mat
}

func (s *Sphere) GetIOR() float64 {
	return s.Mat.GetIOR()
}
