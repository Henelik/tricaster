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

type Hit struct {
	Index    int
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

func NewHit(r *Ray, inters []Intersection, index int) *Hit {
	h := &Hit{
		Index: index,
		Pos:   r.Position(inters[index].T),
		EyeV:  r.Direction.Neg(),
	}

	h.NormalV = inters[index].P.NormalAt(h.Pos)
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

// GetClosestPositiveIndex returns the index of the closest positive intersection,
// assuming they are sorted.
// Returns 0 if none of the intersections are positive.
func GetClosestPositiveIndex(inters []Intersection) int {
	for i, inter := range inters {
		if inter.T > 0 {
			return i
		}
	}
	return -1
}

func GetClosestPositive(inters []Intersection) *Intersection {
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
			if inter.T > other.T {
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
