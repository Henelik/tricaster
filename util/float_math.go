package util

import "math"

const epsilon = .000000000000001

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

func Min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func Max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func MinInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
