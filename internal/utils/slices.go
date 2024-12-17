package utils

import (
	"github.com/spossner/aoc2024/internal/point"
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

func Filter[T any](slice []T, fn func(T) (bool, error)) ([]T, error) {
	return FilterI(slices.Values(slice), fn)
}
func FilterI[T any](iterable iter.Seq[T], fn func(T) (bool, error)) ([]T, error) {
	result := make([]T, 0)
	for el := range iterable {
		ok, err := fn(el)
		if err != nil {
			return nil, err
		}
		if ok {
			result = append(result, el)
		}
	}
	return result, nil
}

func Map[T, U any](slice []T, fn func(T) (U, error)) ([]U, error) {
	return MapI(slices.Values(slice), fn)
}

func MapI[T, U any](iterable iter.Seq[T], f func(T) (U, error)) ([]U, error) {
	result := make([]U, 0)
	for el := range iterable {
		elNew, err := f(el)
		if err != nil {
			return nil, err
		}
		result = append(result, elNew)
	}
	return result, nil
}

func Reduce[T any, U Number](slice []T, fn func(acc U, item T) U, initial U) U {
	return ReduceI(slices.Values(slice), fn, initial)
}

func ReduceI[T any, U Number](iterable iter.Seq[T], fn func(acc U, item T) U, initial U) U {
	acc := initial
	for item := range iterable {
		acc = fn(acc, item)
	}
	return acc
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

func IterateMatrix[T any](matrix [][]T) iter.Seq2[point.Point, T] {
	return func(yield func(point.Point, T) bool) {
	Outer:
		for y, row := range matrix {
			for x, cell := range row {
				if (!yield(point.Point{X: x, Y: y}, cell)) {
					break Outer
				}
			}
		}
	}
}

func PickFromMatrix[T any](matrix [][]T, pos point.Point) T {
	return matrix[pos.Y][pos.X]
}

func PickFirst[T any](slice []T) T {
	return slice[0]
}

func PickLast[T any](slice []T) T {
	return slice[len(slice)-1]
}

func Pick2From[T any](slice []T) (T, T) {
	return slice[0], slice[1]
}

func Pick3From[T any](slice []T) (T, T, T) {
	return slice[0], slice[1], slice[2]
}

func Pick4From[T any](slice []T) (T, T, T, T) {
	return slice[0], slice[1], slice[2], slice[3]
}

func Sum[T Number](slice []T) T {
	var zero T
	return Reduce(slice, func(acc T, item T) T {
		return acc + item
	}, zero)
}

func Product[T Number](slice []T) T {
	var one T
	one++

	return Reduce(slice, func(acc T, item T) T {
		return acc * item
	}, one)
}
