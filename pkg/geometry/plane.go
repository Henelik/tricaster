package geometry

import (
	"math"

	"github.com/Henelik/tricaster/pkg/color"
	"github.com/Henelik/tricaster/pkg/light"
	material2 "github.com/Henelik/tricaster/pkg/material"
	"github.com/Henelik/tricaster/pkg/matrix"
	ray2 "github.com/Henelik/tricaster/pkg/ray"
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
	Mat material2.Material
}

func NewPlane(m *matrix.Matrix, mat material2.Material) *Plane {
	p := &Plane{
		m:   matrix.Identity,
		im:  matrix.Identity,
		n:   tuple.Up,
		Mat: material2.DefaultPhong,
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

func (p *Plane) Intersects(r *ray2.Ray) []ray2.Intersection {
	rt := r.Transform(p.im)
	if math.Abs(rt.Direction.Z) < util.Epsilon {
		return []ray2.Intersection{}
	}
	t := -rt.Origin.Z / rt.Direction.Z
	return []ray2.Intersection{{t, p}}
}

func (p *Plane) NormalAt(pos *tuple.Tuple) *tuple.Tuple {
	return p.n
}

func (p *Plane) Shade(light *light.PointLight, h *ray2.Hit) *color.Color {
	return p.Mat.Lighting(light, h)
}

func (p *Plane) GetMaterial() material2.Material {
	return p.Mat
}

func (p *Plane) GetIOR() float64 {
	return p.Mat.GetIOR()
}
