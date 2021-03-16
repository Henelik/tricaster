package renderer

import (
	"log"
	"math"

	"github.com/Henelik/tricaster/util"

	"github.com/Henelik/tricaster/tuple"
)

var NilIntersect = &Intersection{math.Inf(1), nil}

// Intersection stores the t value of a ray intersection and a pointer to the intersected primitive
type Intersection struct {
	T float64
	P Primitive
}

// intersection precomputation
type Hit struct {
	T        float64
	P        Primitive
	Pos      *tuple.Tuple
	EyeV     *tuple.Tuple
	NormalV  *tuple.Tuple
	ReflectV *tuple.Tuple
	Inside   bool
	InShadow bool
	OverP    *tuple.Tuple
	UnderP   *tuple.Tuple
	N1       float64
	N2       float64
	Inters   []Intersection
}

func (i *Intersection) ToHit(r *Ray, inters []Intersection) *Hit {
	h := &Hit{}
	h.T = i.T
	h.P = i.P

	h.Pos = r.Position(i.T)

	h.EyeV = r.Direction.Neg()
	h.NormalV = i.P.NormalAt(h.Pos)
	h.Inside = h.NormalV.DotProd(h.EyeV) < 0
	if h.Inside {
		h.NormalV = h.NormalV.Neg()
	}
	h.ReflectV = r.Direction.Reflect(h.NormalV)

	h.OverP = h.Pos.Add(h.NormalV.Mult(util.Epsilon * 1000.0))
	h.UnderP = h.Pos.Sub(h.NormalV.Mult(util.Epsilon * 1000.0))

	h.Inters = inters
	ComputeRefractIOR(h)

	return h
}

// Hit returns the closest positive intersection
func GetClosest(inters []Intersection) *Intersection {
	if len(inters) == 0 {
		return NilIntersect
	}
	closest := *NilIntersect
	for _, i := range inters {
		if i.T > 0 && i.T < closest.T {
			closest = i
		}
	}
	return &closest
}

// SortI merge sorts a list of intersections in ascending order
func SortI(inters []Intersection) []Intersection {
	result := make([]Intersection, 0, len(inters))
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

func (h *Hit) Schlick() float64 {
	cos := h.EyeV.DotProd(h.NormalV)

	if h.N1 > h.N2 {
		n := h.N1 / h.N2

		sin2t := n * n * (1 - cos*cos)

		if sin2t > 1 {
			return 1
		}

		cos = math.Sqrt(1 - sin2t)
	}
	r0 := (h.N1 - h.N2) / (h.N1 + h.N2)
	r0 *= r0
	return r0 + (1-r0)*math.Pow(1-cos, 5)
}

func ComputeRefractIOR(h *Hit) {
	containers := make([]Primitive, 0, len(h.Inters))
	var removed bool
	for _, inter := range h.Inters {
		if h.T == inter.T && h.P == inter.P {
			if len(containers) == 0 {
				h.N1 = 1
			} else {
				h.N1 = containers[len(containers)-1].GetMaterial().(*PhongMat).IOR
			}
		}

		// if this object is in the containers, remove it.  Otherwise, append it.
		containers, removed = removePrimitiveFromArr(inter.P, containers)
		if !removed {
			containers = append(containers, inter.P.(Primitive))
		}

		if h.T == inter.T && h.P == inter.P {
			if len(containers) == 0 {
				h.N2 = 1
			} else {
				h.N2 = containers[len(containers)-1].GetMaterial().(*PhongMat).IOR
			}
			return
		}
	}
	log.Fatal("ComputeRefractIOR: Hit was not included in the intersections!")
}

func removePrimitiveFromArr(item Primitive, arr []Primitive) ([]Primitive, bool) {
	prim := item.(Primitive)
	for i, na := range arr {
		if na == prim {
			return append(arr[:i], arr[i+1:]...), true
		}
	}
	return arr, false
}
