package scene

import (
	"log"
	"math"

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
	Config: WorldConfig{
		Shadows:   false,
		MaxBounce: 3,
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

// Shade finds the color of an object at a hit point
func (w *World) Shade(h *ray.Hit, remainingBounce int, inters []ray.Intersection) *color.Color {
	prim := h.P.(geometry.Primitive)

	if w.Config.Shadows {
		h.InShadow = w.IsShadowed(h.OverP)
	}

	n1, n2 := ComputeRefractIOR(h, inters)

	surface := prim.Shade(w.Light, h).MultF(1 - prim.GetMaterial().(*shading.PhongMat).Transparency)
	reflected := w.ReflectedColor(h, remainingBounce-1)
	refracted := w.RefractedColor(h, remainingBounce-1, inters, n1, n2)

	mat := prim.GetMaterial().(*shading.PhongMat)

	if mat.Reflectivity > 0 && mat.Transparency > 0 {
		reflectance := h.Schlick(n1, n2)

		return surface.
			Add(reflected.MultF(reflectance)).
			Add(refracted.MultF(1 - reflectance))
	}

	return surface.Add(reflected).Add(refracted)
}

// ColorAt finds a ray's hit and then calls shade at that hit
func (w *World) ColorAt(r *ray.Ray, remainingBounce int) *color.Color {
	inters := w.Intersect(r)
	i := ray.GetClosest(inters)
	if *i == *ray.NilIntersect {
		return color.Black
	}
	return w.Shade(i.ToHit(r), remainingBounce, inters)
}

// ReflectedColor handles reflection ray culling and finds the next color on the light path
func (w *World) ReflectedColor(h *ray.Hit, remainingBounce int) *color.Color {
	if remainingBounce <= 0 {
		return color.Black
	}
	prim := h.P.(geometry.Primitive)
	if m, ok := prim.GetMaterial().(*shading.PhongMat); ok {
		if m.Reflectivity == 0 {
			return color.Black
		}
		return w.ColorAt(ray.NewRay(h.OverP, h.ReflectV), remainingBounce).MultF(m.Reflectivity)
	}
	return color.Black
}

func (w *World) RefractedColor(h *ray.Hit, remainingBounce int, inters []ray.Intersection, n1, n2 float64) *color.Color {
	if remainingBounce <= 0 {
		return color.Black
	}
	prim := h.P.(geometry.Primitive)
	if m, ok := prim.GetMaterial().(*shading.PhongMat); ok {
		if m.Transparency == 0 {
			return color.Black
		}
		nRatio := n1 / n2
		cosI := h.EyeV.DotProd(h.NormalV)
		sin2T := nRatio * nRatio * (1 - cosI*cosI)
		// find the new ray's direction
		cosT := math.Sqrt(1.0 - sin2T)
		dir := h.NormalV.Mult(nRatio*cosI - cosT).Sub(h.EyeV.Mult(nRatio))
		return w.ColorAt(ray.NewRay(h.UnderP, dir), remainingBounce).MultF(m.Transparency)
	}
	return color.Black
}

func ComputeRefractIOR(h *ray.Hit, inters []ray.Intersection) (float64, float64) {
	inters = ray.SortI(inters)
	containers := make([]geometry.Primitive, 0, len(inters))
	var n1, n2 float64
	for _, inter := range inters {
		if h.T == inter.T && h.P == inter.P {
			if len(containers) == 0 {
				n1 = 1
			} else {
				m := inters[len(containers)-1].P.(geometry.Primitive).GetMaterial().(*shading.PhongMat)
				n1 = m.IOR
			}
		}
		var removed bool
		containers, removed = removeNormalAter(inter.P, containers)
		if !removed {
			containers = append(containers, inter.P.(geometry.Primitive))
		}
		if h.T == inter.T && h.P == inter.P {
			if len(containers) == 0 {
				n2 = 1
			} else {
				m := inters[len(containers)-1].P.(geometry.Primitive).GetMaterial().(*shading.PhongMat)
				n2 = m.IOR
			}
			return n1, n2
		}
	}
	log.Fatal("ComputeRefractIOR: Ray was not included in the intersections!")
	return 0, 0
}

func removeNormalAter(item interface{}, arr []geometry.Primitive) ([]geometry.Primitive, bool) {
	prim := item.(geometry.Primitive)
	for i, na := range arr {
		if na == prim {
			return append(arr[:i], arr[i+1:]...), true
		}
	}
	return arr, false
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
