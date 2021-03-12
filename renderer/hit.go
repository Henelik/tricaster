package renderer

import (
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
}

func (i *Intersection) ToHit(r *Ray) *Hit {
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
	h.OverP = h.Pos.Add(h.NormalV.Mult(util.Epsilon * 1000.0))
	h.UnderP = h.Pos.Sub(h.NormalV.Mult(util.Epsilon * 1000.0))
	h.ReflectV = r.Direction.Reflect(h.NormalV)
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

func (h *Hit) Schlick(n1, n2 float64) float64 {
	cos := h.EyeV.DotProd(h.NormalV)

	if n1 > n2 {
		n := n1 / n2

		sin2t := n * n * (1 - cos*cos)

		if sin2t > 1 {
			return 1
		}

		cos = math.Sqrt(1 - sin2t)
	}
	r0 := (n1 - n2) / (n1 + n2)
	r0 *= r0
	return r0 + (1-r0)*math.Pow(1-cos, 5)
}
