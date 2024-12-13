package utils

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// GCD calculates the greatest common divisor of the given numbers a and b
func GCD[T int](a, b T) T {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM computes the lowest common multiple of the given numbers (at least 2)
func LCM(a, b int, nums ...int) int {
	result := a * b / GCD(a, b)

	for _, n := range nums {
		result = LCM(result, n)
	}

	return result
}

func Bound[T constraints.Ordered](a, low, high T) T {
	if a > high {
		return high
	}
	if a < low {
		return low
	}
	return a
}
