package ray

import (
	"math"

	"github.com/Henelik/tricaster/pkg/tuple"
	"github.com/Henelik/tricaster/pkg/util"
)

var NilIntersect = &Intersection{math.Inf(1), nil}

// Intersection stores the t value of a ray intersection and a pointer to the intersected primitive
type Intersection struct {
	T float64
	P Primitive
}

// TODO: replace T/P values with index of intersection
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

	h.OverP = h.Pos.Add(h.NormalV.Mult(util.Epsilon))
	h.UnderP = h.Pos.Sub(h.NormalV.Mult(util.Epsilon))

	h.Inters = inters

	ComputeIORs(h)

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

func SimpleSort(inters []Intersection) []Intersection {
	if len(inters) <= 1 {
		return inters
	}

	var sum int

	result := make([]Intersection, len(inters))

	for _, inter := range inters {
		sum = 0
		for _, other := range inters {
			if inter.T < other.T {
				sum++
			}
		}
		result[sum] = inter
	}

	return result
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
