package utils

import (
	"iter"
	"slices"
)

func Transpose[S ~[][]T, T any](slice S) S {
	xl := len(slice[0])
	yl := len(slice)
	result := make(S, xl)

	for i := range result {
		result[i] = make([]T, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func Cut[S ~[]T, T any](slice S, index int) S {
	newSlice := make(S, 0, len(slice)-1)
	newSlice = append(newSlice, slice[:index]...)
	newSlice = append(newSlice, slice[index+1:]...)
	return newSlice
}

func Filter[S ~[]T, T any](slice S, f func(T) (bool, error)) (S, error) {
	result := make(S, 0)
	for _, el := range slice {
		ok, err := f(el)
		if err != nil {
			return nil, err
		}
		if ok {
			result = append(result, el)
		}
	}
	return result, nil
}

func Map[T, U any](slice []T, f func(T) (U, error)) ([]U, error) {
	result := make([]U, 0)
	for _, el := range slice {
		elNew, err := f(el)
		if err != nil {
			return nil, err
		}
		result = append(result, elNew)
	}
	return result, nil
}

func Any[S ~[]E, E any](s S, f func(E) bool) bool {
	return slices.ContainsFunc(s, f)
}

func Batched[S ~[]E, E any](s S, n int) iter.Seq2[int, S] {
	return func(yield func(int, S) bool) {
		loop := 0
		for i := 0; i < len(s); i += n {
			if !yield(loop, s[i:min(i+n, len(s))]) {
				break
			}
			loop++
		}
	}
}
