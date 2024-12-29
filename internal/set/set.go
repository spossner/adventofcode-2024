package set

import (
	"iter"
	"slices"
)

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](items ...T) Set[T] {
	set := make(Set[T])
	for _, item := range items {
		set[item] = struct{}{}
	}
	return set
}

// Clone clones the set, adds any optional additional items specified and returns the new set.
func (s Set[T]) Clone(items ...T) Set[T] {
	newSet := make(Set[T])
	for k := range s {
		newSet[k] = struct{}{}
	}
	for _, item := range items {
		newSet[item] = struct{}{}
	}
	return newSet
}

func (s Set[T]) All() iter.Seq[T] {
	return func(yield func(value T) bool) {
		for k := range s {
			if !yield(k) {
				break
			}
		}
	}
}

func (s Set[T]) List() []T {
	return slices.Collect(s.All())
}

func FromSlice[T comparable](sets ...[]T) Set[T] {
	set := make(Set[T])
	for _, slice := range sets {
		for _, item := range slice {
			set[item] = struct{}{}
		}
	}
	return set
}

func Intersect[T comparable](sets ...Set[T]) Set[T] {
	set := make(Set[T])
	switch len(sets) {
	case 0:
		return set
	case 1:
		return sets[0]
	}

Outer:
	for item := range sets[0] {
		for _, that := range sets[1:] {
			if that.Contains(item) {
				set[item] = struct{}{}
				continue Outer
			}
		}
	}
	return set
}

func Subtract[T comparable](set Set[T], others ...Set[T]) Set[T] {
	newSet := make(Set[T])
Outer:
	for item := range set {
		for _, that := range others {
			if that.Contains(item) {
				continue Outer
			}
		}
		newSet[item] = struct{}{}
	}
	return newSet
}

func Union[T comparable](sets ...Set[T]) Set[T] {
	set := make(Set[T])
	for _, that := range sets {
		for item := range that {
			set[item] = struct{}{}
		}
	}
	return set
}

func (s Set[T]) Contains(item T) bool {
	_, ok := s[item]
	return ok
}

func (s Set[T]) Add(items ...T) {
	for _, item := range items {
		s[item] = struct{}{}
	}
}

func (s Set[T]) Remove(items ...T) {
	for _, item := range items {
		delete(s, item)
	}
}

func (s Set[T]) Extend(items iter.Seq[T]) {
	for item := range items {
		s[item] = struct{}{}
	}
}

func (s Set[T]) Pop() T {
	var zero T
	for item, _ := range s {
		delete(s, item)
		return item
	}
	return zero
}
