package geometry

import (
	"math"

	"github.com/Henelik/tricaster/pkg/color"
	"github.com/Henelik/tricaster/pkg/light"
	"github.com/Henelik/tricaster/pkg/material"
	"github.com/Henelik/tricaster/pkg/matrix"
	"github.com/Henelik/tricaster/pkg/ray"
	"github.com/Henelik/tricaster/pkg/tuple"
)

type Sphere struct {
	// the transformation matrix
	m *matrix.Matrix
	// the inverse transformation matrix
	im *matrix.Matrix
	// the transposition of the inverse matrix
	imt    *matrix.Matrix
	Mat    material.Material
	parent GroupInterface
}

func NewSphere(m *matrix.Matrix, mat material.Material) *Sphere {
	s := &Sphere{
		matrix.Identity,
		matrix.Identity,
		matrix.Identity,
		material.DefaultPhong,
		nil,
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

func (s *Sphere) Intersects(r *ray.Ray) []ray.Intersection {
	rt := r.Transform(s.im)
	sphereToRay := rt.Origin.Sub(tuple.Origin)
	a2 := 2 * rt.Direction.DotProd(rt.Direction)
	b := -2 * rt.Direction.DotProd(sphereToRay)

	discriminant := b*b - 2*a2*(sphereToRay.DotProd(sphereToRay)-1)
	if discriminant < 0 {
		return []ray.Intersection{}
	}

	return []ray.Intersection{
		{T: (b - math.Sqrt(discriminant)) / a2, P: s},
		{T: (b + math.Sqrt(discriminant)) / a2, P: s},
	}
}

func (s *Sphere) NormalAt(pos *tuple.Tuple) *tuple.Tuple {
	worldNormal := s.imt.MultTuple(s.im.MultTuple(pos).Sub(tuple.Origin))
	worldNormal.W = 0

	return worldNormal.Norm()
}

func (s *Sphere) SetParent(group GroupInterface) {
	s.parent = group
}

func (s *Sphere) WorldToObject(p *tuple.Tuple) *tuple.Tuple {
	if s.parent != nil {
		return s.im.MultTuple(s.parent.WorldToGroup(p))
	}

	return s.im.MultTuple(p)
}

func (s *Sphere) Shade(light *light.PointLight, h *ray.Hit) *color.Color {
	return s.Mat.Lighting(light, h)
}

func (s *Sphere) GetMaterial() material.Material {
	return s.Mat
}

func (s *Sphere) GetIOR() float64 {
	return s.Mat.GetIOR()
}
