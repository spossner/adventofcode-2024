package pair

import "iter"

type Pair[T any] struct {
	A, B T
}

func NewIntPair(a, b int) Pair[int] {
	return Pair[int]{a, b}
}

func (p Pair[T]) Pick() (T, T) {
	return p.A, p.B
}

func (p Pair[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		if yield(p.A) {
			yield(p.B)
		}
	}
}
