package renderer

import (
	"math"

	"github.com/Henelik/tricaster/color"
	"github.com/Henelik/tricaster/matrix"
	"github.com/Henelik/tricaster/tuple"
)

var DefaultWorld = &World{
	Geometry: []Primitive{
		NewSphere(nil, &PhongMat{
			Ambient:   0.1,
			Diffuse:   0.7,
			Specular:  0.2,
			Shininess: 200,
			Color:     color.NewColor(0.8, 1, 0.6),
		}),
		NewSphere(matrix.Scaling(0.5, 0.5, 0.5), nil),
	},
	Light: &PointLight{
		tuple.NewPoint(-10, -10, 10),
		color.White,
	},
	Config: WorldConfig{
		Shadows:   false,
		MaxBounce: 3,
	},
}

type World struct {
	Geometry []Primitive
	Light    *PointLight
	Config   WorldConfig
}

// Intersect returns all the intersections where a ray encounters an object in the world, unsorted.
func (w *World) Intersect(r *Ray) []Intersection {
	var inters []Intersection
	for _, p := range w.Geometry {
		inters = append(inters, p.Intersects(r)...)
	}
	return SortI(inters)
}

// Shade finds the color of an object at a hit point
func (w *World) Shade(h *Hit, remainingBounce int) *color.Color {
	prim := h.P.(Primitive)

	if w.Config.Shadows {
		h.InShadow = w.IsShadowed(h.OverP)
	}

	surface := prim.Shade(w.Light, h).MultF(1 - prim.GetMaterial().(*PhongMat).Transparency)
	reflected := w.ReflectedColor(h, remainingBounce-1)
	refracted := w.RefractedColor(h, remainingBounce-1)

	mat := prim.GetMaterial().(*PhongMat)

	if mat.Reflectivity > 0 && mat.Transparency > 0 {
		reflectance := h.Schlick()

		return surface.
			Add(reflected.MultF(reflectance)).
			Add(refracted.MultF(1 - reflectance))
	}

	return surface.Add(reflected).Add(refracted)
}

// ColorAt finds a ray's hit and then calls shade at that hit
func (w *World) ColorAt(r *Ray, remainingBounce int) *color.Color {
	inters := w.Intersect(r)
	i := GetClosest(inters)
	if *i == *NilIntersect {
		return color.Black
	}
	return w.Shade(i.ToHit(r, inters), remainingBounce)
}

// ReflectedColor handles reflection ray culling and finds the next color on the light path
func (w *World) ReflectedColor(h *Hit, remainingBounce int) *color.Color {
	if remainingBounce <= 0 {
		return color.Black
	}
	prim := h.P.(Primitive)
	if m, ok := prim.GetMaterial().(*PhongMat); ok {
		if m.Reflectivity == 0 {
			return color.Black
		}
		return w.ColorAt(NewRay(h.OverP, h.ReflectV), remainingBounce).MultF(m.Reflectivity)
	}
	return color.Black
}

func (w *World) RefractedColor(h *Hit, remainingBounce int) *color.Color {
	if remainingBounce <= 0 {
		return color.Black
	}
	prim := h.P.(Primitive)
	if m, ok := prim.GetMaterial().(*PhongMat); ok {
		if m.Transparency == 0 {
			return color.Black
		}
		nRatio := h.N1 / h.N2
		cosI := h.EyeV.DotProd(h.NormalV)
		sin2T := nRatio * nRatio * (1 - cosI*cosI)
		// find the new ray's direction
		cosT := math.Sqrt(math.Abs(1.0 - sin2T))
		dir := h.NormalV.Mult(nRatio*cosI - cosT).Sub(h.EyeV.Mult(nRatio))
		return w.ColorAt(NewRay(h.UnderP, dir), remainingBounce).MultF(m.Transparency)
	}
	return color.Black
}

func (w *World) IsShadowed(p *tuple.Tuple) bool {
	v := w.Light.Pos.Sub(p)
	distance := v.Mag()
	direction := v.Norm()

	r := NewRay(p, direction)

	inters := w.Intersect(r)

	h := GetClosest(inters)

	if *h != *NilIntersect && h.T < distance {
		return true
	}
	return false
}
