package utils

import (
	"log"
	"strconv"
)

func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func MustAtoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("cannot parse %s: %v\n", s, err)
	}
	return n
}
