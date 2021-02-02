package scene

import (
	"github.com/Henelik/tricaster/color"
	"github.com/Henelik/tricaster/geometry"
	"github.com/Henelik/tricaster/matrix"
	"github.com/Henelik/tricaster/ray"
	"github.com/Henelik/tricaster/shading"
	"github.com/Henelik/tricaster/tuple"
	"github.com/Henelik/tricaster/util"
)

var DefaultWorld = &World{
	Geometry: []geometry.Primitive{
		geometry.NewSphere(nil, &shading.PhongMat{
			Ambient:   0.1,
			Diffuse:   0.7,
			Specular:  0.2,
			Shininess: 200,
			Color:     color.NewColor(0.8, 1, 0.6),
		}),
		geometry.NewSphere(matrix.Scaling(0.5, 0.5, 0.5), nil),
	},
	Light: &shading.PointLight{
		tuple.NewPoint(-10, -10, 10),
		color.White,
	},
	Config: WorldConfig{
		Shadows: false,
	},
}

type World struct {
	Geometry []geometry.Primitive
	Light    *shading.PointLight
	Config   WorldConfig
}

func (w *World) Intersect(r *ray.Ray) []geometry.Intersection {
	var inters []geometry.Intersection
	for _, p := range w.Geometry {
		inters = append(inters, p.Intersects(r)...)
	}
	return geometry.SortI(inters)
}

func (w *World) IntersectNoSort(r *ray.Ray) []geometry.Intersection {
	var inters []geometry.Intersection
	for _, p := range w.Geometry {
		inters = append(inters, p.Intersects(r)...)
	}
	return inters
}

func (w *World) Shade(c *geometry.Comp) *color.Color {
	if !w.Config.Shadows {
		return c.P.Shade(w.Light, c, false)
	}
	overP := c.Point.Add(c.NormalV.Mult(util.Epsilon))
	inShadow := w.IsShadowed(overP)
	return c.P.Shade(w.Light, c, inShadow)
}

func (w *World) ColorAt(r *ray.Ray) *color.Color {
	h := geometry.Hit(w.IntersectNoSort(r))
	if *h == *geometry.NilHit {
		return color.Black
	}
	return w.Shade(h.Precompute(r))
}

func (w *World) IsShadowed(p *tuple.Tuple) bool {
	v := w.Light.Pos.Sub(p)
	distance := v.Mag()
	direction := v.Norm()

	r := ray.NewRay(p, direction)

	inters := w.IntersectNoSort(r)

	h := geometry.Hit(inters)

	if *h != *geometry.NilHit && h.T < distance {
		return true
	}
	return false
}
