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
	newSet := NewSet[T](items...)
	for k := range s {
		newSet[k] = struct{}{}
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

// Intersect creates a new set containing only items which are included in all given sets
func Intersect[T comparable](sets ...Set[T]) Set[T] {
	switch len(sets) {
	case 0:
		return NewSet[T]()
	case 1:
		return sets[0].Clone()
	}

	set := NewSet[T]()
Outer:
	for item := range sets[0] {
		for _, that := range sets[1:] {
			if len(that) == 0 {
				return NewSet[T]()
			}
			if !that.Contains(item) {
				continue Outer
			}
		}
		set[item] = struct{}{}
	}
	return set
}

func Subtract[T comparable](set Set[T], others ...Set[T]) Set[T] {
	newSet := NewSet[T]()
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
	set := NewSet[T]()
	for _, that := range sets {
		for item := range that {
			set[item] = struct{}{}
		}
	}
	return set
}

func (s Set[T]) Contains(item T, items ...T) bool {
	if _, ok := s[item]; !ok {
		return false
	}
	for _, anotherItem := range items {
		if _, ok := s[anotherItem]; !ok {
			return false
		}
	}
	return true
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
