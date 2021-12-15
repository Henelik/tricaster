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

type Cone struct {
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

func NewCone(min, max float64, closed bool, m *matrix.Matrix, mat material.Material) *Cone {
	c := &Cone{
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

func (cone *Cone) SetMatrix(m *matrix.Matrix) {
	cone.m = m
	cone.im = m.Inverse()
	cone.imt = cone.im.Transpose()
}

func (cone *Cone) GetMatrix() *matrix.Matrix {
	return cone.m
}

func (cone *Cone) Intersects(r *ray.Ray) []ray.Intersection {
	rt := r.Transform(cone.im)
	inters := make([]ray.Intersection, 0, 4)

	inters = append(inters, cone.intersectCaps(rt)...)

	a := rt.Direction.X*rt.Direction.X + rt.Direction.Y*rt.Direction.Y - rt.Direction.Z*rt.Direction.Z
	if math.Abs(a) < util.Epsilon {
		return inters
	}

	b := 2*rt.Origin.X*rt.Direction.X + 2*rt.Origin.Y*rt.Direction.Y - 2*rt.Origin.Z*rt.Direction.Z
	c := rt.Origin.X*rt.Origin.X + rt.Origin.Y*rt.Origin.Y - rt.Origin.Z*rt.Origin.Z

	if a == 0 {
		if b == 0 {
			return inters
		}

		return append(inters, ray.Intersection{-c / 2 * b, cone})
	}

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

	if cone.min < y0 && y0 < cone.max {
		inters = append(inters, ray.Intersection{t0, cone})
	}

	if cone.min < y1 && y1 < cone.max {
		inters = append(inters, ray.Intersection{t1, cone})
	}

	return inters
}

// TODO: move if closed statement to caller
// intersectCaps checks for intersections at the caps of the Cone
func (cone *Cone) intersectCaps(r *ray.Ray) []ray.Intersection {
	inters := make([]ray.Intersection, 0, 2)

	if !cone.closed || math.Abs(r.Direction.Z) < util.Epsilon {
		return inters
	}

	t := (cone.min - r.Origin.Z) / r.Direction.Z
	if checkConeCap(r, t) {
		inters = append(inters, ray.Intersection{t, cone})
	}

	t = (cone.max - r.Origin.Z) / r.Direction.Z
	if checkConeCap(r, t) {
		inters = append(inters, ray.Intersection{t, cone})
	}

	return inters
}

// checkCylinderCap checks for an intersection within the radius of the Cone
func checkConeCap(r *ray.Ray, t float64) bool {
	x := r.Origin.X + t*r.Direction.X
	y := r.Origin.Y + t*r.Direction.Y

	return (x*x + y*y) <= 1
}

func (cone *Cone) NormalAt(pos *tuple.Tuple) *tuple.Tuple {
	localPos := cone.im.MultTuple(pos)
	if pos.X == 0.0 && pos.Y == 0.0 && pos.Z == 0.0 {
		return tuple.NewVector(0, 0, 0)
	}

	n := cone.imt.MultTuple(cone.LocalNormalAt(localPos))
	n.W = 0

	return n.Norm()
}

func (cone *Cone) LocalNormalAt(pos *tuple.Tuple) *tuple.Tuple {
	dist := pos.X*pos.X + pos.Y*pos.Y

	if dist < 1 {
		switch {
		case pos.Z >= cone.max-util.Epsilon:
			return tuple.NewVector(0, 0, 1)
		case pos.Z <= cone.min+util.Epsilon:
			return tuple.NewVector(0, 0, -1)
		}
	}

	z := math.Sqrt(pos.X*pos.X + pos.Y*pos.Y)
	if pos.Z > 0 {
		z = -z
	}

	return tuple.NewVector(pos.X, pos.Y, z)
}

func (cone *Cone) Shade(light *light.PointLight, h *ray.Hit) *color.Color {
	return cone.Mat.Lighting(light, h)
}

func (cone *Cone) GetMaterial() material.Material {
	return cone.Mat
}

func (cone *Cone) GetIOR() float64 {
	return cone.Mat.GetIOR()
}
