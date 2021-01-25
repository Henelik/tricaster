package util

import "math"

const epsilon = .0000000000000001

func Equal(x, y float64) bool {
	if math.Abs(x-y) < epsilon {
		return true
	}
	return false
}

func Clamp(n, min, max float64) float64 {
	if n < min {
		return min
	}
	if n > max {
		return max
	}
	return n
}
