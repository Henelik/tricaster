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

type Cube struct {
	// the transformation matrix
	m *matrix.Matrix
	// the inverse transformation matrix
	im *matrix.Matrix
	// the transposition of the inverse matrix
	imt *matrix.Matrix
	// the material
	Mat material.Material
}

func NewCube(m *matrix.Matrix, mat material.Material) *Cube {
	c := &Cube{
		matrix.Identity,
		matrix.Identity,
		matrix.Identity,
		material.DefaultPhong,
	}

	if m != nil {
		c.SetMatrix(m)
	}

	if mat != nil {
		c.Mat = mat
	}

	return c
}

func (c *Cube) SetMatrix(m *matrix.Matrix) {
	c.m = m
	c.im = m.Inverse()
	c.imt = c.im.Transpose()
}

func (c *Cube) GetMatrix() *matrix.Matrix {
	return c.m
}

func (c *Cube) Intersects(r *ray.Ray) []ray.Intersection {
	rt := r.Transform(c.im)
	xtmin, xtmax := checkAxis(rt.Origin.X, rt.Direction.X)
	ytmin, ytmax := checkAxis(rt.Origin.Y, rt.Direction.Y)
	ztmin, ztmax := checkAxis(rt.Origin.Z, rt.Direction.Z)

	tmin := util.Max(util.Max(xtmin, ytmin), ztmin)
	tmax := util.Min(util.Min(xtmax, ytmax), ztmax)

	if tmin > tmax {
		return []ray.Intersection{}
	}

	return []ray.Intersection{{T: tmin, P: c}, {T: tmax, P: c}}
}

// checkAxis returns the min and max t-values where a ray intersects the cube on an axis
func checkAxis(origin, dir float64) (float64, float64) {
	var tMin, tMax float64

	tMinNumerator := (-1 - origin)
	tMaxNumerator := (1 - origin)

	if math.Abs(dir) >= util.Epsilon {
		tMin = tMinNumerator / dir
		tMax = tMaxNumerator / dir
	} else {
		tMin = tMinNumerator * math.Inf(1)
		tMax = tMaxNumerator * math.Inf(1)
	}

	// ensure the correct order of returns
	if tMin > tMax {
		return tMax, tMin
	}

	return tMin, tMax
}

func (c *Cube) NormalAt(pos *tuple.Tuple) *tuple.Tuple {
	n := c.imt.MultTuple(c.LocalNormalAt(c.im.MultTuple(pos)))
	n.W = 0

	return n.Norm()
}

func (c *Cube) LocalNormalAt(pos *tuple.Tuple) *tuple.Tuple {
	maxc := util.Max(util.Max(math.Abs(pos.X), math.Abs(pos.Y)), math.Abs(pos.Z))

	if maxc == math.Abs(pos.X) {
		return tuple.NewVector(pos.X, 0, 0)
	}
	if maxc == math.Abs(pos.Y) {
		return tuple.NewVector(0, pos.Y, 0)
	}
	return tuple.NewVector(0, 0, pos.Z)
}

func (c *Cube) Shade(light *light.PointLight, h *ray.Hit) *color.Color {
	return c.Mat.Lighting(light, h)
}

func (c *Cube) GetMaterial() material.Material {
	return c.Mat
}

func (c *Cube) GetIOR() float64 {
	return c.Mat.GetIOR()
}
