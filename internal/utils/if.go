package utils

import "reflect"

func If[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b
}

func Iff(cond bool, a, b any) any {
	if cond {
		if reflect.TypeOf(a).Kind() == reflect.Func {
			return a.(func() any)()
		}
		return a
	}
	if reflect.TypeOf(b).Kind() == reflect.Func {
		return b.(func() any)()
	}
	return b
}
