package ray

import (
	"log"
)

type IORHaver interface {
	GetIOR() float64
}

func ComputeIORs(h *Hit) {
	containers := make([]IORHaver, 0, len(h.Inters))
	var removed bool
	for _, inter := range h.Inters {
		if h.Inters[h.Index].T == inter.T && h.Inters[h.Index].P == inter.P {
			if len(containers) == 0 {
				h.N1 = 1
			} else {
				h.N1 = containers[len(containers)-1].(IORHaver).GetIOR()
			}
		}

		// if this object is in the containers, remove it.  Otherwise, append it.
		containers, removed = removeIORHaverFromArr(inter.P.(IORHaver), containers)
		if !removed {
			containers = append(containers, inter.P.(IORHaver))
		}

		if h.Inters[h.Index].T == inter.T && h.Inters[h.Index].P == inter.P {
			if len(containers) == 0 {
				h.N2 = 1
			} else {
				h.N2 = containers[len(containers)-1].(IORHaver).GetIOR()
			}
			return
		}
	}
	log.Fatal("ComputeRefractIOR: Hit was not included in the intersections!")
}

func removeIORHaverFromArr(item IORHaver, arr []IORHaver) ([]IORHaver, bool) {
	prim := item.(IORHaver)
	for i, na := range arr {
		if na == prim {
			return append(arr[:i], arr[i+1:]...), true
		}
	}
	return arr, false
}
