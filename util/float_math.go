package util

import "git.maze.io/go/math32"

const Epsilon = .00007

func Equal(x, y float32) bool {
	if math32.Abs(x-y) < Epsilon {
		return true
	}
	return false
}

func Clamp(n, min, max float32) float32 {
	if n < min {
		return min
	}
	if n > max {
		return max
	}
	return n
}

func Min(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}

func Max(a, b float32) float32 {
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

func Lerp(start, end, factor float32) float32 {
	return start*(1-factor) + end*factor
}
