package ray

import (
	"math"
)

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
