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
		geometry.NewSphere(nil, nil),
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
