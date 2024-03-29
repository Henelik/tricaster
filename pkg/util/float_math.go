package util

import "math"

const Epsilon = .0000000001

func Equal(x, y float64) bool {
	if math.Abs(x-y) < Epsilon {
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

func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Lerp(start, end, factor float64) float64 {
	return start*(1-factor) + end*factor
}
