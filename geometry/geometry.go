package geometry

import (
	"github.com/Henelik/tricaster/color"
	"github.com/Henelik/tricaster/matrix"
	"github.com/Henelik/tricaster/shading"
	"github.com/Henelik/tricaster/tuple"
	"math"

	"github.com/Henelik/tricaster/ray"
)

var NilHit = Intersection{math.Inf(1), nil}

// Primitive geometry type which defines an intersection function
type Primitive interface {
	SetMatrix(m *matrix.Matrix)
	GetMatrix() *matrix.Matrix
	// Intersects returns an array of intersections where the ray meets the primitive
	Intersects(r *ray.Ray) []Intersection
	// NormalAt returns the normal vector at a given scene point
	NormalAt(p *tuple.Tuple) *tuple.Tuple
	Shade(light *shading.PointLight, c *Comp) *color.Color
}

// Intersection stores the t value of a ray intersection and a pointer to the intersected primitive
type Intersection struct {
	T float64
	P Primitive
}

// intersection precomputation
type Comp struct {
	T       float64
	P       Primitive
	Point   *tuple.Tuple
	EyeV    *tuple.Tuple
	NormalV *tuple.Tuple
	Inside  bool
}

func (i *Intersection) Precompute(r *ray.Ray) *Comp {
	c := &Comp{}
	c.T = i.T
	c.P = i.P
	c.Point = r.Position(i.T)
	c.EyeV = r.Direction.Neg()
	c.NormalV = i.P.NormalAt(c.Point)
	c.Inside = c.NormalV.DotProd(c.EyeV) < 0
	if c.Inside {
		c.NormalV = c.NormalV.Neg()
	}
	return c
}

// Hit returns the closest positive intersection
func Hit(inters []Intersection) Intersection {
	if len(inters) == 0 {
		return NilHit
	}
	var closest Intersection = NilHit
	for _, i := range inters {
		if i.T > 0 && i.T < closest.T {
			closest = i
		}
	}
	return closest
}

// SortI merge sorts a list of intersections in ascending order
func SortI(inters []Intersection) []Intersection {
	var result = make([]Intersection, 0, len(inters))
	if len(inters) <= 1 {
		return inters
	}
	mid := len(inters) / 2
	l := SortI(inters[:mid])
	r := SortI(inters[mid:])
	i, j := 0, 0
	for i < len(l) && j < len(r) {
		if l[i].T < r[j].T {
			result = append(result, l[i])
			i++
		} else {
			result = append(result, r[j])
			j++
		}
	}
	result = append(result, l[i:]...)
	result = append(result, r[j:]...)
	return result
}
