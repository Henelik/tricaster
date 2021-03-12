package renderer

import (
	"math"

	"github.com/Henelik/tricaster/util"

	"github.com/Henelik/tricaster/color"
	"github.com/Henelik/tricaster/matrix"
	"github.com/Henelik/tricaster/tuple"
)

type Plane struct {
	// the transformation matrix
	m *matrix.Matrix
	// the inverse transformation matrix
	im *matrix.Matrix
	// the plane's normal vector
	n *tuple.Tuple
	// the material
	Mat Material
}

func NewPlane(m *matrix.Matrix, mat Material) *Plane {
	p := &Plane{
		m:   matrix.Identity,
		im:  matrix.Identity,
		n:   tuple.Up,
		Mat: DefaultPhong,
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

func (p *Plane) Intersects(r *Ray) []Intersection {
	rt := r.Transform(p.im)
	if math.Abs(rt.Direction.Z) < util.Epsilon {
		return []Intersection{}
	}
	t := -rt.Origin.Z / rt.Direction.Z
	return []Intersection{{t, p}}
}

func (p *Plane) NormalAt(pos *tuple.Tuple) *tuple.Tuple {
	return p.n
}

func (p *Plane) Shade(light *PointLight, h *Hit) *color.Color {
	return p.Mat.Lighting(light, h)
}

func (p *Plane) GetMaterial() Material {
	return p.Mat
}
