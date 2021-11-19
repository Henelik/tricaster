package renderer

import (
	"math"

	"github.com/Henelik/tricaster/pkg/color"
	"github.com/Henelik/tricaster/pkg/geometry"
	"github.com/Henelik/tricaster/pkg/light"
	"github.com/Henelik/tricaster/pkg/material"
	"github.com/Henelik/tricaster/pkg/matrix"
	"github.com/Henelik/tricaster/pkg/ray"
	"github.com/Henelik/tricaster/pkg/tuple"
)

var DefaultWorld = &World{
	Geometry: []Primitive{
		geometry.NewSphere(nil, &material.PhongMat{
			Ambient:   0.1,
			Diffuse:   0.7,
			Specular:  0.2,
			Shininess: 200,
			Color:     color.NewColor(0.8, 1, 0.6),
		}),
		geometry.NewSphere(matrix.Scaling(0.5, 0.5, 0.5), nil),
	},
	Light: &light.PointLight{
		tuple.NewPoint(-10, -10, 10),
		color.White,
	},
	Config: &WorldConfig{
		Shadows:   false,
		MaxBounce: 3,
	},
}

type World struct {
	Geometry []Primitive
	Light    *light.PointLight
	Config   *WorldConfig
}

// Intersect returns all the intersections where a ray encounters an object in the world, sorted.
func (w *World) Intersect(r *ray.Ray) []ray.Intersection {
	// TODO: initialize inters to a decent length
	var inters []ray.Intersection

	for _, p := range w.Geometry {
		inters = append(inters, p.Intersects(r)...)
	}

	return ray.SimpleSort(inters)
}

// Shade finds the color of an object at a hit point
func (w *World) Shade(h *ray.Hit, remainingBounce int) *color.Color {
	if w.Config.Shadows {
		h.InShadow = w.IsShadowed(h.OverP)
	}

	primitive := h.Inters[h.Index].P.(Primitive)

	surface := primitive.Shade(w.Light, h)
	reflected := w.ReflectedColor(h, remainingBounce-1)
	refracted := w.RefractedColor(h, remainingBounce-1)

	mat := primitive.GetMaterial().(*material.PhongMat)

	if mat.Reflectivity > 0 && mat.Transparency > 0 {
		reflectance := h.Schlick()

		return surface.
			Add(reflected.MultF(reflectance)).
			Add(refracted.MultF(1 - reflectance))
	}

	return surface.Add(reflected).Add(refracted)
}

// ColorAt finds a ray's hit and then calls shade at that hit
func (w *World) ColorAt(r *ray.Ray, remainingBounce int) *color.Color {
	inters := w.Intersect(r)
	if len(inters) == 0 {
		return color.Black
	}

	return w.Shade(ray.NewHit(r, inters, ray.GetClosestPositiveIndex(inters)), remainingBounce)
}

// ReflectedColor handles reflection ray culling and finds the next color on the light path
func (w *World) ReflectedColor(h *ray.Hit, remainingBounce int) *color.Color {
	if remainingBounce <= 0 {
		return color.Black
	}
	if m, ok := h.Inters[h.Index].P.(Primitive).GetMaterial().(*material.PhongMat); ok {
		if m.Reflectivity == 0 {
			return color.Black
		}
		return w.ColorAt(ray.NewRay(h.OverP, h.ReflectV), remainingBounce).MultF(m.Reflectivity)
	}
	return color.Black
}

// RefractedColor handles refraction ray culling and finds the next color on the light path
func (w *World) RefractedColor(h *ray.Hit, remainingBounce int) *color.Color {
	if remainingBounce <= 0 {
		return color.Black
	}
	if m, ok := h.Inters[h.Index].P.(Primitive).GetMaterial().(*material.PhongMat); ok {
		if m.Transparency == 0 {
			return color.Black
		}
		nRatio := h.N1 / h.N2
		cosI := h.EyeV.DotProd(h.NormalV)
		sin2T := nRatio * nRatio * (1 - cosI*cosI)
		// find the new ray's direction
		cosT := math.Sqrt(math.Abs(1.0 - sin2T))
		dir := h.NormalV.Mult(nRatio*cosI - cosT).Sub(h.EyeV.Mult(nRatio))
		return w.ColorAt(ray.NewRay(h.UnderP, dir), remainingBounce).MultF(m.Transparency)
	}
	return color.Black
}

func (w *World) IsShadowed(p *tuple.Tuple) bool {
	v := w.Light.Pos.Sub(p)
	distance := v.Mag()
	direction := v.Norm()

	r := ray.NewRay(p, direction)

	inters := w.Intersect(r)

	h := ray.GetClosestPositive(inters)

	if *h != *ray.NilIntersect && h.T < distance {
		return true
	}
	return false
}
