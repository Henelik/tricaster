package scene

import (
	"github.com/Henelik/tricaster/color"
	"github.com/Henelik/tricaster/geometry"
	"github.com/Henelik/tricaster/matrix"
	"github.com/Henelik/tricaster/ray"
	"github.com/Henelik/tricaster/shading"
	"github.com/Henelik/tricaster/tuple"
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
}

type World struct {
	Geometry []geometry.Primitive
	Light    *shading.PointLight
}

func (w *World) Intersect(r *ray.Ray) []geometry.Intersection {
	var inters []geometry.Intersection
	for _, p := range w.Geometry {
		inters = append(inters, p.Intersects(r)...)
	}
	return geometry.SortI(inters)
}

func (w *World) Shade(c *geometry.Comp) *color.Color {
	return c.P.Shade(w.Light, c)
}
