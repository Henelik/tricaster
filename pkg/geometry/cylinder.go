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

type Cylinder struct {
	// the transformation matrix
	m *matrix.Matrix
	// the inverse transformation matrix
	im *matrix.Matrix
	// the transposition of the inverse matrix
	imt    *matrix.Matrix
	min    float64
	max    float64
	closed bool
	// the material
	Mat material.Material
}

func NewCylinder(min, max float64, closed bool, m *matrix.Matrix, mat material.Material) *Cylinder {
	c := &Cylinder{
		m:      matrix.Identity,
		im:     matrix.Identity,
		imt:    matrix.Identity,
		min:    min,
		max:    max,
		closed: closed,
		Mat:    material.DefaultPhong,
	}

	if m != nil {
		c.SetMatrix(m)
	}

	if mat != nil {
		c.Mat = mat
	}

	return c
}

func (cyl *Cylinder) SetMatrix(m *matrix.Matrix) {
	cyl.m = m
	cyl.im = m.Inverse()
	cyl.imt = cyl.im.Transpose()
}

func (cyl *Cylinder) GetMatrix() *matrix.Matrix {
	return cyl.m
}

func (cyl *Cylinder) Intersects(r *ray.Ray) []ray.Intersection {
	rt := r.Transform(cyl.im)
	inters := make([]ray.Intersection, 0, 2)

	inters = append(inters, cyl.intersectCaps(rt)...)

	a := rt.Direction.X*rt.Direction.X + rt.Direction.Y*rt.Direction.Y
	if a < util.Epsilon {
		return inters
	}

	b := 2*rt.Origin.X*rt.Direction.X + 2*rt.Origin.Y*rt.Direction.Y
	c := rt.Origin.X*rt.Origin.X + rt.Origin.Y*rt.Origin.Y - 1

	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return inters
	}

	t0 := (-b - math.Sqrt(discriminant)) / (2 * a)
	t1 := (-b + math.Sqrt(discriminant)) / (2 * a)

	if t0 > t1 {
		t0, t1 = t1, t0
	}

	y0 := rt.Origin.Z + t0*rt.Direction.Z
	y1 := rt.Origin.Z + t1*rt.Direction.Z

	if cyl.min < y0 && y0 < cyl.max {
		inters = append(inters, ray.Intersection{t0, cyl})
	}

	if cyl.min < y1 && y1 < cyl.max {
		inters = append(inters, ray.Intersection{t1, cyl})
	}

	return inters
}

// TODO: move if closed statement to caller
// intersectCaps checks for intersections at the caps of the cylinder
func (cyl *Cylinder) intersectCaps(r *ray.Ray) []ray.Intersection {
	inters := make([]ray.Intersection, 0, 2)

	if !cyl.closed || math.Abs(r.Direction.Z) < util.Epsilon {
		return inters
	}

	t := (cyl.min - r.Origin.Z) / r.Direction.Z
	if checkCylinderCap(r, t) {
		inters = append(inters, ray.Intersection{t, cyl})
	}

	t = (cyl.max - r.Origin.Z) / r.Direction.Z
	if checkCylinderCap(r, t) {
		inters = append(inters, ray.Intersection{t, cyl})
	}

	return inters
}

// checkCylinderCap checks for an intersection within the radius of the cylinder
func checkCylinderCap(r *ray.Ray, t float64) bool {
	x := r.Origin.X + t*r.Direction.X
	y := r.Origin.Y + t*r.Direction.Y

	return (x*x + y*y) <= 1
}

func (cyl *Cylinder) NormalAt(pos *tuple.Tuple) *tuple.Tuple {
	n := cyl.imt.MultTuple(cyl.LocalNormalAt(cyl.im.MultTuple(pos)))
	n.W = 0

	return n.Norm()
}

func (cyl *Cylinder) LocalNormalAt(pos *tuple.Tuple) *tuple.Tuple {
	dist := pos.X*pos.X + pos.Y*pos.Y

	if dist < 1 {
		switch {
		case pos.Z >= cyl.max-util.Epsilon:
			return tuple.NewVector(0, 0, 1)
		case pos.Z <= cyl.min+util.Epsilon:
			return tuple.NewVector(0, 0, -1)
		}
	}

	return tuple.NewVector(pos.X, pos.Y, 0)
}

func (cyl *Cylinder) Shade(light *light.PointLight, h *ray.Hit) *color.Color {
	return cyl.Mat.Lighting(light, h)
}

func (cyl *Cylinder) GetMaterial() material.Material {
	return cyl.Mat
}

func (cyl *Cylinder) GetIOR() float64 {
	return cyl.Mat.GetIOR()
}
