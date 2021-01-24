package util

import "math"

const epsilon = .00001

func Equal(x, y float64) bool {
	if math.Abs(x-y) < epsilon {
		return true
	}
	return false
}
