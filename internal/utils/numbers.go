package utils

import (
	"golang.org/x/exp/constraints"
)

// AbsInt give the absolute value of the given number
func AbsInt[T constraints.Integer | constraints.Float](n T) T {
	if n < 0 {
		return -n
	}
	return n
}

// MinMax returns the minimum and maximum number (in that order) from the given values
func MinMax[T constraints.Ordered](a, b T, nums ...T) (lower T, upper T) {
	lower, upper = a, b
	if a > b {
		return b, a
	}

	for _, m := range nums {
		if m < lower {
			lower = m
		}
		if m > upper {
			upper = m
		}
	}

	return lower, upper
}

// GCD calculates the greatest common divisor of the given numbers a and b
func GCD[T constraints.Integer](a, b T) T {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM computes the lowest common multiple of the given numbers (at least 2)
func LCM[T constraints.Integer](a, b T, nums ...T) T {
	result := a * b / GCD(a, b)

	for _, n := range nums {
		result = LCM[T](result, n)
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

// Mod returns the (euclidean; python like; unlike go) modulus of x and y.
// The result will always be positive.
func Mod[T constraints.Integer](x, y T) T {
	return (x%y + y) % y
}
