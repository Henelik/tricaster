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

// Intersect returns all the intersections where a ray encounters an object in the world, unsorted.
func (w *World) Intersect(r *ray.Ray) []ray.Intersection {
	var inters []ray.Intersection
	for _, p := range w.Geometry {
		inters = append(inters, p.Intersects(r)...)
	}
	return inters
}

func (w *World) Shade(h *ray.Hit) *color.Color {
	prim := h.P.(geometry.Primitive)
	if !w.Config.Shadows {
		return prim.Shade(w.Light, h)
	}
	overP := h.Pos.Add(h.NormalV.Mult(util.Epsilon))
	h.InShadow = w.IsShadowed(overP)
	return prim.Shade(w.Light, h)
}

func (w *World) ColorAt(r *ray.Ray) *color.Color {
	i := ray.GetClosest(w.Intersect(r))
	if *i == *ray.NilIntersect {
		return color.Black
	}
	return w.Shade(i.ToHit(r))
}

func (w *World) IsShadowed(p *tuple.Tuple) bool {
	v := w.Light.Pos.Sub(p)
	distance := v.Mag()
	direction := v.Norm()

	r := ray.NewRay(p, direction)

	inters := w.Intersect(r)

	h := ray.GetClosest(inters)

	if *h != *ray.NilIntersect && h.T < distance {
		return true
	}
	return false
}
