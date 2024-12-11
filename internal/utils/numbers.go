package utils

import (
	"golang.org/x/exp/constraints"
	"strconv"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func MustAtoi(s string) int {
	return Must(strconv.Atoi(s))
}
