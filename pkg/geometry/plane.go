package geometry

import (
	"math"

	"github.com/Henelik/tricaster/pkg/color"
	"github.com/Henelik/tricaster/pkg/light"
	"github.com/Henelik/tricaster/pkg/material"
	"github.com/Henelik/tricaster/pkg/matrix"
	"github.com/Henelik/tricaster/pkg/ray"
	"github.com/Henelik/tricaster/pkg/tuple"
	"github.com/Henelik/tricaster/pkg/util"
)

type Plane struct {
	// the transformation matrix
	m *matrix.Matrix
	// the inverse transformation matrix
	im *matrix.Matrix
	// the plane's normal vector
	n *tuple.Tuple
	// the material
	Mat material.Material
}

func NewPlane(m *matrix.Matrix, mat material.Material) *Plane {
	p := &Plane{
		m:   matrix.Identity,
		im:  matrix.Identity,
		n:   tuple.Up,
		Mat: material.DefaultPhong,
	}

	if m != nil {
		p.m = m
		p.im = m.Inverse()
		p.n = p.im.MultTuple(tuple.Up)
	}

	if mat != nil {
		p.Mat = mat
	}

	return p
}

func (p *Plane) SetMatrix(m *matrix.Matrix) {
	p.m = m
	p.im = m.Inverse()
}

func (p *Plane) GetMatrix() *matrix.Matrix {
	return p.m
}

func (p *Plane) Intersects(r *ray.Ray) []ray.Intersection {
	rt := r.Transform(p.im)

	if math.Abs(rt.Direction.Z) < util.Epsilon {
		return []ray.Intersection{}
	}

	t := -rt.Origin.Z / rt.Direction.Z

	return []ray.Intersection{{t, p}}
}

func (p *Plane) NormalAt(pos *tuple.Tuple) *tuple.Tuple {
	return p.n
}

func (p *Plane) Shade(light *light.PointLight, h *ray.Hit) *color.Color {
	return p.Mat.Lighting(light, h)
}

func (p *Plane) GetMaterial() material.Material {
	return p.Mat
}

func (p *Plane) GetIOR() float64 {
	return p.Mat.GetIOR()
}
